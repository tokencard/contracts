const assert = require('assert')
const ganache = require('ganache-cli')
const Web3 = require('web3')
const provider = ganache.provider()
const web3 = new Web3(provider)
const { card } = require('../compile')
const { expect } = require('chai')

let accounts
let inbox

let controller
let owner
let newOwner
let neitherOwnerNorController

beforeEach(async () => {
  accounts = await web3.eth.getAccounts()
  controller = accounts[0]
  owner = accounts[1]
  newOwner = accounts[2]
  neitherOwnerNorController = accounts[3]
  inbox = await new web3.eth.Contract(JSON.parse(card.interface))
    .deploy({ data: card.bytecode, arguments: [controller]})
    .send({ from: owner, gas: '1000000'})
})

describe('Card', () => {
  it('deploys contract', () => {
    assert.ok(inbox.options.address)
  })

  describe('owner()', () => {
    it("should return address of the owner", async () => {
      expect(await inbox.methods.owner().call()).to.be.equal(owner)
    })
  })

  describe('controller()', () => {
    it('should return address of the controller', async() => {
      expect(await inbox.methods.controller().call()).to.be.equal(controller)
    })
  })

  describe('changeOwner', () => {
    context('when the caller is the owner', () => {
      it('should change the owner', async () => {
        await inbox.methods.changeOwner(newOwner).send({from: owner})
        expect(await inbox.methods.owner().call()).to.be.equal(newOwner)
      })
    })
    context('when the caller is the controller', () => {
      it('should change the owner', async () => {
        await inbox.methods.changeOwner(newOwner).send({from: controller})
        expect(await inbox.methods.owner().call()).to.be.equal(newOwner)
      })
    })
    context('when the caller is neither owner nor controller', () => {
      it('should reject the transaction', async () => {
        await expect(inbox.methods.changeOwner(newOwner).send({from: neitherOwnerNorController}))
          .to.eventually.be.rejectedWith('VM Exception while processing transaction: revert')
      })
    })
  })

  describe('unlock', () => {
    context('when the caller is the owner', () => {
      it('should unlock the contract', async () => {
        await inbox.methods.unlock().send({from: owner})
        expect(await inbox.methods.unlockAt().call()).not.to.be.equal(0)
      })
    })
    context('when the caller is the controller', () => {
      it('should reject the transaction', async () => {
        await expect(inbox.methods.unlock().send({from: controller}))
          .to.eventually.be.rejectedWith('VM Exception while processing transaction: revert')
      })
    })
    context('when the caller is neither owner nor controller', () => {
      it('should reject the transaction', async () => {
        await expect(inbox.methods.unlock().send({from: neitherOwnerNorController}))
          .to.eventually.be.rejectedWith('VM Exception while processing transaction: revert')
      })
    })
  })

})
