# Library and Command Line Spec

## High Level
This project will include the following...
1) A set of libraries that make common tasks easier
2) A command line tool that allows you to use all of these convenience methods directly from your command line
3) A testing framework allowing you to try things locally before using real ETH

## Interacting with ETH Networks
We will use the go-ethereum client which will allow for rpc or ipc connections

#### Testing
We can use [Ganache](https://github.com/trufflesuite/ganache) for testing locally

## Library Interface

### Method Index

+ erc20.getTx.byTxID([options], txId)
  + options
  ```
  retryTimes
  retryInterval
  ```
+ erc20.getTx.byNonce
+ erc20.getTxs(searchDataObject)
  + searchDataObject
  ```
  fromAddress
  toAddress
  amount
  ```
+ erc20.safeResend

## Command Line Spec

### Config
~/.ethutils.conf
```
ethHosts: [
  {
    hostname: "host1.io",
    port: 8443,
    account: "0x0000000",
    passwordFile: "~/blah.pw"
  },
  {
    hostname: "host1.io",
    port: 8443,
    account: "0x0000000",
    passwordFile: "~/blah.pw"
  }
]
```
#### Notes
+ passwordFile would be encrypted and the command line tool could prompt for passwords, and encrypt them with a password set through the command line tool

### Examples
```
$ ethutils erc20 getTx byTxID 0x0000000000000000...
$ ethutils erc20 safeResend
