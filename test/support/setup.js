"use strict";
const chai = require("chai")
const chaiAsPromised = require("chai-as-promised")
const Web3 = require('web3')
var ganache = require("ganache-core");
const memdown = require('memdown')
const web3 = new Web3(ganache.provider({db: memdown()}))
const { expect } = require('chai')

chai.use(chaiAsPromised)
global.expect = expect
global.web3 = web3
