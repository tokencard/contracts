const path = require('path')
const fs = require('fs')
const solc = require('solc')


function compile(fileName, key) {
  const pth = path.resolve(__dirname, fileName)
  const source = fs.readFileSync(pth, 'utf8')
  const output = solc.compile(source, 1)
  if (!output.contracts[key]) {
    console.log(output)
    throw new Error(`could not compile ${pth}: ${output}`)
  }
  return output.contracts[key]

}



module.exports = {
  card: compile('card.sol',':Card'),
  controller: compile('controller.sol',':Controller'),
  wallet: compile('wallet.sol',':Wallet')
}
