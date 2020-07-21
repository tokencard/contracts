#!/bin/bash

if [ ! -d "contracts" ]; then 
	echo "error: script needs to be run from project root './tools/manticore/manticore.sh'"
	exit 1
fi

docker run --rm -v "$PWD":/contracts -it --ulimit stack=100000000:100000000 --workdir=/contracts trailofbits/eth-security-toolbox@sha256:e5e2ebbffcc4c3a334063ba3871ec25a4e1dc4915d77166aef1aa265e0a5978f -c 'solc-select 0.6.11 &&
manticore contracts/controller.sol --contract=Controller --config=tools/manticore/manticore.yaml --quick-mode &&
manticore contracts/wallet.sol --contract=Wallet --config=tools/manticore/manticore.yaml --quick-mode &&
manticore contracts/oracle.sol --contract=Oracle --config=tools/manticore/manticore.yaml --quick-mode &&
manticore contracts/tokenWhitelist.sol --contract=TokenWhitelist --config=tools/manticore/manticore.yaml --quick-mode &&
manticore contracts/licence.sol --contract=Licence --config=tools/manticore/manticore.yaml --quick-mode &&
manticore contracts/holder.sol --contract=Holder --config=tools/manticore/manticore.yaml --quick-mode &&
manticore contracts/walletDeployer.sol --contract=WalletDeployer --config=tools/manticore/manticore.yaml --quick-mode'
