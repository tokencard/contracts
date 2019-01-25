package holder_test

	import (
		// "errors"
		"math/big"
		// "context"

	  // "github.com/ethereum/go-ethereum/core/types"

		. "github.com/onsi/ginkgo"
		. "github.com/onsi/gomega"
		. "github.com/tokencard/contracts/test/shared"
		"github.com/tokencard/ethertest"
		// "github.com/ethereum/go-ethereum/common"

	)

	var _ = Describe("TokenHolder", func() {

		When("burn is NOT called by the burner contract", func() {

      It("Should revert", func(){
        tx, err := TokenHolder.Burn(RandomAccount.TransactOpts(ethertest.WithGasLimit(100000)),RandomAccount.Address(), big.NewInt(301))
        Expect(err).ToNot(HaveOccurred())
        Backend.Commit()
        Expect(isGasExhausted(tx, 100000)).To(BeFalse())
        Expect(isSuccessful(tx)).To(BeFalse())
        Expect(err).ToNot(HaveOccurred())
      })

    })
})
