# Library and Command Line Spec

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
