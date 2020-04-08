#!/bin/bash

# Development script which runs all validation tools in order.

set -e

if [ ! -d "tools" ]; then 
	echo "error: script needs to be run from project root './tools/run-all.sh'"
	exit 1
fi

echo "========================"
echo "Formatting contracts (prettier) ..."
echo "========================"

./tools/prettier/format.sh

echo "========================"
echo "Flattening contracts (slither) ..."
echo "========================"

./tools/slither/flatten.sh

echo "========================================"
echo "Performing static analysis (slither) ..."
echo "========================================"

./tools/slither/slither.sh

echo "============================================"
echo "Performing symbolic analysis (manticore) ..."
echo "============================================"

./tools/manticore/manticore.sh

echo "================================"
echo "Running fuzz tests (echidna) ..."
echo "================================"

./tools/echidna/echidna.sh
