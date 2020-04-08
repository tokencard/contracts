#!/bin/bash

if [ ! -d "contracts" ]; then 
	echo "error: script needs to be run from project root './tools/manticore/manticore.sh'"
	exit 1
fi
