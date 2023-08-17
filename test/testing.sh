#!/bin/bash
truncate -s 0 ../air.log
airsettled tx airsettle add-execution-layer "basic chain info" verificationkey.json --from alice --gas 3000000 -y 
sleep 1; 
chainid=`cat ./chainid.test.air`
creator_address="air1p3qmpeddacn8znrhlvhrkjv0xts6d9u662qlpj"

# Query Chains
# sleep 1; airsettled query airsettle chain-list $creator_address
# sleep 1; airsettled query airsettle chain-list-detailed $creator_address
# sleep 1; airsettled query airsettle list-execution-layers
# sleep 1; airsettled query airsettle verification-key $chainid
# sleep 1; airsettled query airsettle show-execution-layer $chainid

# Create Batch
# batchnumber=100 # wrong batch number error.
# sleep 1; airsettled tx airsettle add-batch $chainid $batchnumber "0xMerkleRootHash" "0xPrevMerkleRootHash" zkproof.json --from bob -y
# batchnumber=1
# sleep 1; airsettled tx airsettle add-batch $chainid $batchnumber "0xMerkleRootHash" "0xPrevMerkleRootHash" zkproof.json --from bob -y
# sleep 2; airsettled tx airsettle add-batch $chainid $batchnumber "0xMerkleRootHash" "0xPrevMerkleRootHash" zkproof.json --from $creator_address -y
# sleep 2; airsettled query airsettle show-batch $chainid $batchnumber
# batchnumber=2
# sleep 2; airsettled tx airsettle add-batch $chainid $batchnumber "0xMerkleRootHash" "0xPrevMerkleRootHash" zkproof.json --from $creator_address -y
# sleep 2; airsettled query airsettle show-batch $chainid $batchnumber
# batchnumber=3
# sleep 2; airsettled tx airsettle add-batch $chainid $batchnumber "0xMerkleRootHash" "0xPrevMerkleRootHash" zkproof.json --from $creator_address -y
# sleep 2; airsettled query airsettle show-batch $chainid $batchnumber

# airsettled query airsettle show-batch $chainid $batchnumber
# Verify Batch
# sleep 1; airsettled query airsettle verify $chainid $batchnumber inputs.json
sleep 1; airsettled tx airsettle add-validator "345789459783452978" $chainid --from $creator_address -y
sleep 1; airsettled query airsettle list-add-validators-polls