
var loop = 8;
var sleep = require('system-sleep');
async function getTimer() {
    for (var i = 0; i < 4 * loop; i ++) {
        await setTimeout(function() {
        },1000/loop);
        aa(i);
    }
}
function aa(i) {
    console.log("Peer: ", i);
    sleep(200);
}

getTimer();