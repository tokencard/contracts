# Echidna tests for TokenCard contracts

This directory contains some properties for testing TokenCard contracts where three users are defined:

* `crytic_owner`: the owner of the contract only used to execute the contract constructor
* `crytic_attack`: an unprivileged user, used to execute the randomly generated transactions
* `crytic_user`: another unprivileged user, also used to execute randomly generated transactions in cases where multiple users are necessary

We test a variety of properties, such as:
* Whitelist addition and removal operations properly modify the whitelist
* The owner of a contract cannot be changed.
* Spending limits are properly bounded by the spend limit values
* Spending limits do in fact reset daily
* Differential testing of the different interger parsing functions produce identical results

# Instructions

1. Uncompress this file in the root of [TokenCard's contract repository](https://github.com/tokencard/contracts). We tested this example with version [2.0.0](https://github.com/tokencard/contracts/commit/b99b7d1670f9ad7b90335e8391fe63fd7e20de9b).
2. Install the latest revision of [Echidna](https://github.com/crytic/echidna/#installation). Note that `cryticECRecover.sol` requires the `dev-hevm-0.29` branch of Echidna to run.
3. Execute `echidna-test` with one of the `*.sol` files and its corresponding `*.yaml` config file. For example, an invocation with `cryticSpendLimit.sol` would look like:

``` 
$ echidna-test cryticSpendLimit.sol --config cryticSpendLimit.yaml TEST
...
Analyzing contract: cryticSpendLimit.sol:TEST
crytic_fixed_spendLimitValue: passed! 🎉
crytic_bounded_spendLimitAvailable: passed! 🎉
crytic_fixed_owner: passed! 🎉
crytic_daily_spendLimitAvailable: passed! 🎉

Unique instructions: 900
Unique codehashes: 1
```
