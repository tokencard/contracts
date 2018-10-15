package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gobwas/glob"
	"github.com/logrusorgru/aurora"
	"github.com/tokencard/ethertest"
)

// TODO(tav): Figure out the base cost at runtime.
const baseCost = 21400

var (
	evmFalse = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	evmTrue  = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
)

var solcArgs = []string{
	"--combined-json",
	"abi,bin,metadata",
	"--optimize",
}

type contract struct {
	abi      abi.ABI
	binding  *bind.BoundContract
	events   map[common.Hash]string
	filename string
	network  *network
	tally    map[string]uint64
}

func (c *contract) call(fn string) {
	tx, err := c.binding.Transact(c.network.auth, fn)
	if err != nil {
		logError("Could not run %s", fn)
		return
	}
	c.network.backend.Commit()
	// TODO(tav): The gas cost from the transaction receipt does not seem to be
	// accurate. Figure out a more accurate way.
	receipt, err := c.network.backend.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		logFatal("Could not get receipt for %s: %s", fn, err)
	}
	if len(receipt.Logs) > 0 {
		fmt.Printf("   .. %s\n\n", fn)
		for _, entry := range receipt.Logs {
			val := formatEventData(entry.Data)
			for _, event := range entry.Topics {
				fmt.Printf("      -> %s(%s)\n", c.events[event], val)
			}
		}
		fmt.Println("")
	}
	c.tally[fn] = receipt.GasUsed - baseCost
}

func (c *contract) printTally(fns []string, width int) {
	logProgress("Gas usage for %s", c.filename)
	format := fmt.Sprintf("\t%%%ds    %%d\n", width)
	for _, group := range groupFunctions(fns) {
		for _, fn := range group {
			gas, ok := c.tally[fn]
			if !ok {
				continue
			}
			fmt.Printf(format, fn, gas)
		}
		fmt.Println("")
	}
}

func (c *contract) test(pattern glob.Glob) {
	logProgress("Testing %s", c.filename)
	fns, width := getFunctions(c.abi, pattern)
	c.tally = map[string]uint64{}
	for _, fn := range fns {
		c.call(fn)
	}
	c.printTally(fns, width)
}

type network struct {
	auth    *bind.TransactOpts
	backend ethertest.TestBackend
}

func (n *network) deploy(source *source) (c *contract) {
	logProgress("Deploying %s", stripExt(source.filename))
	abi := abi.ABI{}
	if err := json.Unmarshal([]byte(source.abi), &abi); err != nil {
		logError("Could not decode ABI: %s", err)
		return nil
	}
	defer func() {
		if r := recover(); r != nil {
			logError("Failed to deploy: %s", r)
			c = nil
		}
	}()
	_, _, bound, err := bind.DeployContract(n.auth, abi, common.FromHex(source.bin), n.backend)
	if err != nil {
		logError("Failed to deploy: %s", err)
		return nil
	}
	events := map[common.Hash]string{}
	for _, event := range abi.Events {
		events[event.Id()] = event.Name
	}
	return &contract{
		abi:      abi,
		binding:  bound,
		events:   events,
		filename: stripExt(source.filename),
		network:  n,
	}
}

type source struct {
	abi      string
	bin      string
	filename string
}

func compile(filenames []string) map[string]*source {
	var (
		stderr, stdout bytes.Buffer
	)
	logProgress("Compiling contracts")
	cmd := exec.Command("solc", append(solcArgs, filenames...)...)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		fmt.Println(stderr.String())
		logFatal("Failed to compile contracts: %s", err)
	}
	sources, err := parse(stdout.Bytes())
	if err != nil {
		logFatal("Failed to process the compiled JSON output: %s", err)
	}
	return sources
}

func ensure(filenames []string, sources map[string]*source) {
	for _, filename := range filenames {
		source := sources[filename]
		if source == nil {
			logFatal("Could not find a contract named 'Test' within %s", filename)
		}
	}
}

func eth(amount int64) *big.Int {
	v := big.NewInt(1000000000000000000)
	return v.Mul(v, big.NewInt(amount))
}

func formatEventData(data []byte) string {
	val := ""
	if len(data) != 32 {
		val = "?"
	} else if bytes.Equal(data, evmTrue) {
		val = "true"
	} else if bytes.Equal(data, evmFalse) {
		val = "false"
	} else {
		val = "?"
	}
	return val
}

func getFunctions(abi abi.ABI, pattern glob.Glob) ([]string, int) {
	var fns []string
	width := 0
	for _, fn := range abi.Methods {
		if len(fn.Inputs) > 0 || len(fn.Outputs) > 0 {
			continue
		}
		name := fn.Name
		if !pattern.Match(name) {
			continue
		}
		if len(name) > width {
			width = len(name)
		}
		fns = append(fns, name)
	}
	sort.Strings(fns)
	return fns, width
}

func groupFunctions(fns []string) [][]string {
	var groups [][]string
	var group []string
	prefix := ""
	prev := ""
	for idx, fn := range fns {
		if idx == 0 {
			group = append(group, fn)
			prev = fn
			continue
		}
		if prefix == "" {
			prefix = longestCommonPrefix(fn, prev)
			if prefix == "" {
				groups = append(groups, group)
				group = []string{fn}
				prev = fn
				continue
			}
			group = append(group, fn)
			continue
		}
		if strings.HasPrefix(fn, prefix) {
			group = append(group, fn)
			continue
		} else {
			groups = append(groups, group)
			group = []string{fn}
			prefix = ""
			prev = fn
			continue
		}
	}
	if len(group) > 0 {
		groups = append(groups, group)
	}
	return groups
}

func initNetwork() *network {
	logProgress("Initialising testnet")
	root := ethertest.NewAccount()
	rig := ethertest.NewTestRig()
	rig.AddGenesisAccountAllocation(root.Address(), eth(1000000000))
	backend := rig.NewTestBackend()
	return &network{
		auth:    root.TransactOpts(),
		backend: backend,
	}
}

func logError(format string, args ...interface{}) {
	fmt.Println(aurora.Red(fmt.Sprintf("!! ERROR: "+format+"\n", args...)))
}

func logFatal(format string, args ...interface{}) {
	fmt.Println(aurora.Red(fmt.Sprintf("!! ERROR: "+format, args...)))
	os.Exit(1)
}

func logProgress(format string, args ...interface{}) {
	fmt.Println(aurora.Blue(fmt.Sprintf(">> "+format+"\n", args...)))
}

func longestCommonPrefix(s1, s2 string) string {
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] != s2[i] {
			if i == 0 {
				return ""
			}
			return s1[:i]
		}
	}
	return s1
}

func parse(combined []byte) (map[string]*source, error) {
	var toplevel struct {
		Contracts map[string]struct {
			Abi string
			Bin string
		}
	}
	if err := json.Unmarshal(combined, &toplevel); err != nil {
		return nil, err
	}
	sources := map[string]*source{}
	for name, info := range toplevel.Contracts {
		if !strings.HasSuffix(name, ":Test") {
			continue
		}
		filename := name[:len(name)-5]
		sources[filename] = &source{
			abi:      info.Abi,
			bin:      info.Bin,
			filename: filename,
		}
	}
	return sources, nil
}

func parseArgs(args []string) ([]string, map[string]glob.Glob) {
	var filenames []string
	patterns := map[string]glob.Glob{}
	for _, arg := range args {
		split := strings.SplitN(arg, ":", 2)
		filename := split[0]
		pattern := "*"
		if len(split) == 2 {
			pattern = split[1]
		}
		filenames = append(filenames, filename)
		matcher, err := glob.Compile(pattern)
		if err != nil {
			logFatal("Invalid glob pattern %q: %s", pattern, err)
		}
		patterns[filename] = matcher
	}
	return filenames, patterns
}

func process(args []string) {
	filenames, patterns := parseArgs(args)
	sources := compile(filenames)
	ensure(filenames, sources)
	testnet := initNetwork()
	for _, filename := range filenames {
		contract := testnet.deploy(sources[filename])
		if contract == nil {
			continue
		}
		contract.test(patterns[filename])
	}
}

func stripExt(filename string) string {
	if strings.HasSuffix(filename, ".sol") {
		return filename[:len(filename)-4]
	}
	return filename
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: gasmeter FILE_PATH[:PATTERN] ...")
		os.Exit(1)
	}
	process(os.Args[1:])
}
