#!/usr/bin/env bash

# Fail the build if Go bindings have not been regenerated after making changes to a contract.

function hasChanged {
    return $(! git diff --quiet $(git merge-base --fork-point master) $1)
}

for contract in *.sol; do
    if hasChanged ${contract} && ! hasChanged pkg/bindings/${contract%.sol}.go
    then
        echo "${contract} bindings have not been regenerated"
        exit 1
    fi
done


