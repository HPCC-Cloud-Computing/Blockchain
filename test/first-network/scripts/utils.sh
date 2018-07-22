
# This is a collection of bash functions used by different scripts

ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/com/orderers/orderer.com/msp/tlscacerts/tlsca.com-cert.pem
PEER0_ORG1_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.com/peers/peer0.org1.com/tls/ca.crt
PEER0_ORG2_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.com/peers/peer0.org2.com/tls/ca.crt
PEER0_ORG3_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.com/peers/peer0.org3.com/tls/ca.crt

# verify the result of the end-to-end test
verifyResult() {
  if [ $1 -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to execute End-2-End Scenario ==========="
    echo
    exit 1
  fi
}

setGlobals() {
  PEER=$1
  ORG=$2
  if [ $ORG -eq 1 ]; then
    CORE_PEER_LOCALMSPID="Org1MSP"
    CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA
    CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.com/users/Admin@org1.com/msp
    if [ $PEER -eq 0 ]; then
      CORE_PEER_ADDRESS=peer0.org1.com:7051
    else
      CORE_PEER_ADDRESS=peer1.org1.com:7051
    fi
  elif [ $ORG -eq 2 ]; then
    CORE_PEER_LOCALMSPID="Org2MSP"
    CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG2_CA
    CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.com/users/Admin@org2.com/msp
    if [ $PEER -eq 0 ]; then
      CORE_PEER_ADDRESS=peer0.org2.com:7051
    else
      CORE_PEER_ADDRESS=peer1.org2.com:7051
    fi
  else
    echo "================== ERROR !!! ORG Unknown =================="
  fi

  if [ "$VERBOSE" == "true" ]; then
    env | grep CORE
  fi
}

updateAnchorPeers() {
  PEER=$1
  ORG=$2
  setGlobals $PEER $ORG

  peer channel update -o orderer.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/${CORE_PEER_LOCALMSPID}anchors.tx --tls --cafile $ORDERER_CA >&log.txt
   
  cat log.txt
  
  sleep $DELAY
}

## Sometimes Join takes time hence RETRY at least 5 times
joinChannelWithRetry() {
  PEER=$1
  ORG=$2
  setGlobals $PEER $ORG

  set -x
  peer channel join -b $CHANNEL_NAME.block >&log.txt
  res=$?
  set +x
  cat log.txt
  if [ $res -ne 0 -a $COUNTER -lt $MAX_RETRY ]; then
    COUNTER=$(expr $COUNTER + 1)
    echo "peer${PEER}.org${ORG} failed to join the channel, Retry after $DELAY seconds"
    sleep $DELAY
    joinChannelWithRetry $PEER $ORG
  else
    COUNTER=1
  fi
  verifyResult $res "After $MAX_RETRY attempts, peer${PEER}.org${ORG} has failed to join channel '$CHANNEL_NAME' "
}

installChaincode() {
  PEER=$1
  ORG=$2
  setGlobals $PEER $ORG

  peer chaincode install -n aaa -v 15.3 -p github.com/chaincode/user01/go/main_chaincode >&log1.txt
  cat log1.txt

  peer chaincode install -n aaa1 -v 15.3 -p github.com/chaincode/user01/go/chaincode_information >&log2.txt
  cat log2.txt

  peer chaincode install -n aaa2 -v 15.3 -p github.com/chaincode/user01/go/chaincode_school_profile >&log3.txt
  cat log3.txt

  peer chaincode install -n aaa3 -v 15.3 -p github.com/chaincode/user01/go/chaincode_score >&log4.txt
  cat log4.txt
}

instantiateChaincode() {

  export CHANNEL_NAME=mychannel

  peer chaincode instantiate -o orderer.com:7050 --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n aaa -v 15.3 -c '{"Args":["init"]}' -P "OR ('Org1MSP.peer','Org2MSP.peer')" >&log1.txt
  cat log1.txt
  
  peer chaincode instantiate -o orderer.com:7050 --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n aaa1 -v 15.3 -c '{"Args":["init"]}' -P "OR ('Org1MSP.peer','Org2MSP.peer')" >&log2.txt
  cat log2.txt

  peer chaincode instantiate -o orderer.com:7050 --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n aaa2 -v 15.3 -c '{"Args":["init"]}' -P "OR ('Org1MSP.peer','Org2MSP.peer')" >&log3.txt
  cat log3.txt

  peer chaincode instantiate -o orderer.com:7050 --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n aaa3 -v 15.3 -c '{"Args":["init"]}' -P "OR ('Org1MSP.peer','Org2MSP.peer')" >&log4.txt
  cat log4.txt
}
