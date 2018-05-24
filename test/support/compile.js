const path = require('path')
const fs = require('fs')
const solc = require('solc')
const instrumentSolidity = require('solidity-coverage/lib/instrumentSolidity')



function compile(fileName, key) {
  const contractPath = path.resolve(__dirname, fileName)
  const source = fs.readFileSync(contractPath, 'utf8')

  const instrumented = instrumentSolidity(source, contractPath)

  const output = solc.compile(instrumented.contract, 1)
  if (!output.contracts[key]) {
    throw new Error(`could not compile ${pth}: ${output}`)
  }
  return { contractPath, instrumented, ...output.contracts[key]}

}



module.exports = {
  card: compile('../../v1/card.sol',':Card'),
  controller: compile('../../v1/controller.sol',':Controller'),
  wallet: compile('../../v1/wallet.sol',':Wallet')
}
