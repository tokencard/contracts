"use strict";
const chai = require("chai")
const chaiAsPromised = require("chai-as-promised")
const Web3 = require('web3')
var ganache = require("ganache-core");
const memdown = require('memdown')
const web3 = new Web3(ganache.provider({db: memdown()}))
const { expect } = require('chai')

const CoverageMap = require('solidity-coverage/lib/coverageMap')
const _ = require('lodash')
const istanbul = require('istanbul')
const contracts = require('./compile')



chai.use(chaiAsPromised)
global.expect = expect
global.web3 = web3



// capture all events fired while tests are running
const txLog = []
web3.eth.subscribe('logs', {}, (err, data) => {
  if (data) {
    txLog.push(data)
  }
})

// generate coverage reports
process.on('exit', function() {
  const raw = _.map(txLog, (e) => {
    return JSON.stringify({
      data: e.data,
      topics: _.map(e.topics, (a) => a.replace('0x', ''))
    })
  })

  const coverageMap = new CoverageMap()

  _.forEach(contracts, (contract) => {
    coverageMap.addContract(contract.instrumented, contract.contractPath)
  })

  _.forEach(contracts, (contract) => {
    coverageMap.generate(raw, contract.contractPath)
  })


  const collector = new istanbul.Collector()
  const reporter = new istanbul.Reporter()

  collector.add(coverageMap.coverage)
  reporter.add('text')
  reporter.add('html')
  reporter.write(collector, true, () => {
    console.log('Istanbul coverage reports generated');
  })

})
