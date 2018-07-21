#!/bin/bash

echo
echo " ____    _____      _      ____    _____ "
echo "/ ___|  |_   _|    / \    |  _ \  |_   _|"
echo "\___ \    | |     / _ \   | |_) |   | |  "
echo " ___) |   | |    / ___ \  |  _ <    | |  "
echo "|____/    |_|   /_/   \_\ |_| \_\   |_|  "
echo
echo "Build your first network (BYFN) end-to-end test"
echo
CHANNEL_NAME="$1"
DELAY="$2"
LANGUAGE="$3"
TIMEOUT="$4"
VERBOSE="$5"
: ${CHANNEL_NAME:="mychannel"}
: ${DELAY:="500"}
: ${LANGUAGE:="golang"}
: ${TIMEOUT:="10"}
: ${VERBOSE:="false"}
LANGUAGE=`echo "$LANGUAGE" | tr [:upper:] [:lower:]`
COUNTER=1
MAX_RETRY=5

echo "Channel name : "$CHANNEL_NAME

# import utils
. scripts/utils.sh

createChannel() {

	peer channel create -o orderer.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls --cafile $ORDERER_CA >&log.txt
	
	cat log.txt
	sleep $DELAY
}

joinChannel () {
	for org in 1 2; do
	    for peer in 0 1; do
		joinChannelWithRetry $peer $org
		echo "===================== peer${peer}.org${org} joined channel '$CHANNEL_NAME' ===================== "
		sleep $DELAY
		echo
	    done
	done
}

## Create channel
echo "Creating channel..."
createChannel

## Join all the peers to the channel
echo "Having all peers join the channel..."
joinChannel

## Set the anchor peers for each org in the channel
echo "Updating anchor peers for org1..."
updateAnchorPeers 0 1
sleep $DELAY
echo "Updating anchor peers for org2..."
updateAnchorPeers 0 2
sleep $DELAY

## Install chaincode on peer0.org1 and peer0.org2

echo "Installing chaincode on peer0.org1..."
installChaincode 0 1
sleep $DELAY

echo "Installing chaincode on peer1.org1..."
installChaincode 1 1
sleep $DELAY

echo "Install chaincode on peer0.org2..."
installChaincode 0 2
sleep $DELAY

echo "Install chaincode on peer1.org2..."
installChaincode 1 2
sleep $DELAY

echo "Instantiating chaincode on peer0.org1..."
instantiateChaincode
sleep $DELAY
echo
echo
echo
