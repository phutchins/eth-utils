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
