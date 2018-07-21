
var fs = require('fs');
let peerPem = fs.readFileSync('/home/firedragon/Documents/hung/Blockchain/test/first-network/crypto-config/peerOrganizations/org1.com/peers/peer0.org1.com/tls/ca.crt');
let ordererPem = fs.readFileSync('/home/firedragon/Documents/hung/Blockchain/test/first-network/crypto-config/ordererOrganizations/com/orderers/orderer.com/tls/ca.crt');
module.exports = {
    PEER_PEM: peerPem,
    ORDERER_PEM: ordererPem,
    ORDERER_DOMAIN: "orderer.com",
    PEER_DOMAIN: "peer0.org1.com",
    TLS_ENABLED: "true",
    MSPID: "Org1MSP",
    CA_SERVER_NAME: "ca-org1"
};
