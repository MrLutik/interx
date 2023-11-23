#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/Transactions/query-block-by-height-or-hash.sh
set -e
set -x
. /etc/profile

TEST_NAME="TX-BLOCK-HEIGHTORHASH-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

VALIDATOR_ADDRESS=$(showAddress validator)
addAccount testuser2
TESTUSER_ADDRESS=$(showAddress testuser2)

TXRESULT=$(sekaid tx bank send validator $TESTUSER_ADDRESS 5ukex --keyring-backend=test --chain-id=$NETWORK_NAME --fees 100ukex --broadcast-mode=async --output=json --yes --home=$SEKAID_HOME 2> /dev/null || exit 1)
BLOCK_HASH=$(echo $TXRESULT | jsonQuickParse "txhash")
TXQUERYRESULT=$(sekaid query tx $TX_ID 2> /dev/null || exit 1)
BLOCK_HEIGHT=$(echo $TXQUERYRESULT | jsonQuickParse "height")

INTERX_GATEWAY="127.0.0.1:11000"
RESULT_HASH_FROM_INTERX=$(curl --fail $INTERX_GATEWAY/api/blocks/$BLOCK_HEIGHT | jq '.block_id.hash' | tr -d '"' || exit 1)
RESULT_HEIGHT_FROM_INTERX=$(curl --fail $INTERX_GATEWAY/api/blocks/0x$BLOCK_HASH | jq '.block.header.height' | tr -d '"' || exit 1)

[[ $BLOCK_HEIGHT !=  $RESULT_HEIGHT_FROM_INTERX ]] && echoErr "ERROR: Expected tx block height to be '$BLOCK_HEIGHT', but got '$RESULT_HEIGHT_FROM_INTERX'" && exit 1
[[ $BLOCK_HASH != $RESULT_HASH_FROM_INTERX ]] &&  echoErr "ERROR: Expected tx block hash to be '$BLOCK_HASH', but got '$RESULT_HASH_FROM_INTERX'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"