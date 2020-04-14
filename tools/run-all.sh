#!/bin/bash

if [ ! -d "tools" ]; then 
	echo "error: script needs to be run from project root './tools/run-all.sh'"
	exit 1
fi

read -p "Would you like to format the contracts using 'prettier' (< 10 sec) [Y/n] " answer
case $answer in 
	[Nn]) echo "Skipping ...";;
	*) ./tools/prettier/format.sh;;
esac

read -p "Would you like to flatten the contracts using 'slither-flat' (< 1 min) [Y/n] " answer
case $answer in 
	[Nn]) echo "Skipping ...";;
	*) ./tools/slither/flatten.sh;;
esac

read -p "Would you like to perform static analysis using 'slither' (< 5 min) [Y/n] " answer
case $answer in 
	[Nn]) echo "Skipping ...";;
	*) ./tools/slither/slither.sh;;
esac

read -p "Would you like to perform symbolic execution using 'manticore' (< 10 min) [Y/n] " answer
case $answer in 
	[Nn]) echo "Skipping ...";;
	*) ./tools/manticore/manticore.sh;;
esac

read -p "Would you like to perform security analysis using 'mythril' (< 30 min) [Y/n] " answer
case $answer in 
	[Nn]) echo "Skipping ...";;
	*) ./tools/mythril/mythril.sh;;
esac

read -p "Would you like to run the fuzz tests using 'echidna' (< 30 min) [Y/n] " answer
case $answer in 
	[Nn]) echo "Skipping ...";;
	*) ./tools/echidna/echidna.sh;;
esac


