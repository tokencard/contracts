# Mythril

Mythril security analysis can be executed for all contracts using the provided script:

	./tools/mythril/mythril.sh

Additional contracts can be added by appending them to the command string inside `mythril.sh`.

## Development

To start the mythril environment manually during development, use the provided Docker image:

	docker run --rm -v "$PWD":/contracts -it --workdir=/contracts --entrypoint=bash mythril/myth@sha256:<hash>

To test a contract use the `myth` command inside the container:

	myth analyze --solv=<version> <contract>.sol --execution-timeout=<number>
