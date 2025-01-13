#!/usr/bin/env bash

set -eu

ipc_dir=${1:-../ipc}
hoku_exporter_dir=$(cd `dirname $0`; pwd)

function compile_abi {
  local sol_file_and_contract=$1
  local go_package=$2
  local go_file=${3:-$go_package.go}

  local abi_file=$hoku_exporter_dir/abi/$(echo $sol_file_and_contract | sed -e 's/.*://').abi
  (
    cd $ipc_dir/$contract_dir
    forge inspect $sol_file_and_contract abi > $abi_file
  )
  mkdir -p ./contracts/$go_package
  abigen --abi=$abi_file --pkg=$go_package --out=contracts/$go_package/$go_file
}

contract_dir=contracts
compile_abi ./contracts/gateway/GatewayGetterFacet.sol:GatewayGetterFacet gateway
compile_abi ./contracts/subnet/SubnetActorGetterFacet.sol:SubnetActorGetterFacet subnet subnet_actor_getter.go
compile_abi ./lib/murky/lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol:ERC20 erc20

contract_dir=hoku-contracts
compile_abi ./src/wrappers/BlobManager.sol:BlobManager blobs
