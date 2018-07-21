byfn.sh -m down
../bin/cryptogen generate --config=./crypto-config.yaml
export FABRIC_CFG_PATH=$PWD
../bin/configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block
export CHANNEL_NAME=mychannel  && ../bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME
../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP
../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org2MSP
docker-compose -f docker-compose-cli.yaml -f docker-compose-couch.yaml up -d
docker exec -it cli bash
export CHANNEL_NAME=mychannel
peer channel create -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
peer channel join -b mychannel.block

export  CORE_PEER_ADDRESS=peer0.org2.example.com:7051
export CORE_PEER_LOCALMSPID="Org2MSP"
export  CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt

export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt


peer channel update -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/Org1MSPanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

peer chaincode install -n users -v 1.0 -p github.com/chaincode/user01/go

export CHANNEL_NAME=mychannel
peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n users -v 1.0 -c '{"Args":["init"]}' -P "OR ('Org0MSP.peer','Org1MSP.peer')"

peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n users_information -v 1.0 -c '{"Args":["init"]}' -P "OR ('Org1MSP.peer','Org2MSP.peer')"

peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n aaa -c '{"Args":["initUser","aaa1","3012","Truong Van Luat","30-12-1996","Nam","Ha Tinh"]}'
peer chaincode invoke -o orderer.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/com/orderers/orderer.com/msp/tlscacerts/tlsca.com-cert.pem -C $CHANNEL_NAME -n aaa -c '{"Args":["updateUser","aaa1","3012","Truong Van Luat","30-12-1996","Nam","Ha Nam"]}'

peer chaincode invoke -o orderer.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/com/orderers/orderer.com/msp/tlscacerts/tlsca.com-cert.pem -C $CHANNEL_NAME -n aaa -c '{"Args":["initProfile","aaa2","3012","10A1,Ha Huy Tap,2017-2018,Kante,Mbappe,Toan#9.3&Ly#9.5,Kha,Hoc sinh gioi tinh mon hoa hoc","Tot nghiep cap 2"]}'
peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n aaa -c '{"Args":["updateProfile","aaa2","3012","10A1,Ha Huy Tap,2017-2018,Kante,Mbappe,Toan#9.3&Ly#9.5,Kha,Hoc sinh gioi tinh mon hoa hoc#Hoc sinh gioi tinh mon van","Tot nghiep cap 1#Tot nghiep cap 2","11"]}'

peer chaincode invoke -o orderer.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/com/orderers/orderer.com/msp/tlscacerts/tlsca.com-cert.pem -C $CHANNEL_NAME -n aaa -c '{"Args":["deleteUser","aaa1","aaa2","3012"]}'
peer chaincode query -o orderer.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/com/orderers/orderer.com/msp/tlscacerts/tlsca.com-cert.pem -C $CHANNEL_NAME -n aaa -c '{"Args":["getUserByID","aaa1","3012"]}'
peer chaincode query -o orderer.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/com/orderers/orderer.com/msp/tlscacerts/tlsca.com-cert.pem -C $CHANNEL_NAME -n aaa -c '{"Args":["getProfileByID","aaa2","3012"]}'
peer chaincode query -o orderer.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/com/orderers/orderer.com/msp/tlscacerts/tlsca.com-cert.pem -C $CHANNEL_NAME -n aaa -c '{"Args":["getListProfileOfClass","aaa2","class_10","2017-2018","10A1"]}'

peer chaincode query -o orderer.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/com/orderers/orderer.com/msp/tlscacerts/tlsca.com-cert.pem -C $CHANNEL_NAME -n aaa -c '{"Args":["initScore","aaa3","Toan#9.3&Ly#9.5"]}'


peer chaincode install -n aaa -v 14.3 -p github.com/chaincode/user01/go/main_chaincode
peer chaincode upgrade -o orderer.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/com/orderers/orderer.com/msp/tlscacerts/tlsca.com-cert.pem -C $CHANNEL_NAME -n aaa -v 14.3 -c '{"Args":["init"]}' -P "OR ('Org1MSP.peer','Org2MSP.peer')"