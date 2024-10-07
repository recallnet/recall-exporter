# hoku-exporter

`hoku-exporter` is a service that exports metrics about
* hoku validator node,
* hoku network.

## ABI Wrapper Code Generation

1. Install
   * `ethereum`: `brew install ethereum`
   * foundry tools (`forge`)
2. Compile `GatewayGetterFacet.sol`
```sh
cd ipc/contracts
forge inspect contracts/gateway/GatewayGetterFacet.sol:GatewayGetterFacet abi > ../../hoku-exporter/abi/GatewayGetterFacet.abi
```
3. Generate code
```sh
abigen --abi=./abi/GatewayGetterFacet.abi --pkg=gateway --out=contracts/gateway/gateway.go
```
For details see https://goethereumbook.org/smart-contract-compile/
