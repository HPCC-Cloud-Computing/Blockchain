'use strict';
/*
* Copyright IBM Corp All Rights Reserved
*
* SPDX-License-Identifier: Apache-2.0
*/
/*
 * Chaincode Invoke
 */
var program = require("commander");
var defaultConfig = require("./config");
var path = require('path');

program
    .version("0.1.0")
    .option("-u, --user []", "User id", "user1")
    .option("--name, --channel []", "A channel", "mychannel")
    .option("--chaincode, --chaincode []", "A chaincode name", "chaincode_example02")
    .option("--host, --host []", "Host", defaultConfig.peerHost)
    .option("--ehost, --event-host []", "Host", defaultConfig.eventHost)
    .option("--ohost, --orderer-host []", "Host", defaultConfig.ordererHost)
    .option("-m, --method []", "A method", "invoke")
    .option(
        "-a, --arguments [value]",
        "A repeatable value",
        (val, memo) => memo.push(val) && memo,
        []
    )
    .option("-l, --loop []", "Loop", "8")
    .parse(process.argv);

// node invoke.js -u user9 --channel mychannel --chaincode mycc -m invoke -a a -a b -a 10
var store_path = path.join(__dirname, 'hfc-key-store')
const config = Object.assign({}, defaultConfig, {
    channelName: program.channel,
    user: program.user,
	storePath: store_path
});

var controller = require("./controller")(config);

var request = {
    //targets: let default to the peer assigned to the client
    chaincodeId: program.chaincode,
    fcn: program.method,
    args: program.arguments
};
var numLoop = program.loop;

getTimer();

async function getTimer() {
    for (var i = 0; i < 2*numLoop; i ++) {
        var start = Date.now();
        console.log("starting timer: ", i + "-", start );
        request.args[0] = request.args[0] + i.toString();
        await setTimeout(function() {
            // result => getTimeInvoke(request, start, i)
            // .then(result)
        },125);
        await getTimeInvoke(request, start, i);
    }
}

// each method require different certificate of user
async function getTimeInvoke(request, start, i) {
    controller
    .invoke(program.user, request, start, i)
    .then(results => {
        console.log(
            "Send transaction promise and event listener promise have completed",
            results
        );
    })
    .catch(err => {
        console.error(err);
    });
}
