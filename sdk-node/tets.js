
function timeout(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

// async function sleep() {
//     for (var i = 0; i < 2 * 8; i ++) {
//         var start = Date.now();
//         console.log("starting timer: ", i , "-", start );
//         await timeout(125);
//         getTimeInvoke(i);
//     }
//     // await timeout(3000);
//     // return fn(...args);
// }

// function sleep (getTimeInvoke,i) { 
//     return new Promise((resolve) => {
//         // wait 3s before calling fn(par)
//         setTimeout(() => resolve(getTimeInvoke(i)), 125)
//     })
// }

// async function test() {
//     for (var i = 0; i < 2 * 8; i ++) {
//         var start = Date.now();
//         console.log("starting timer: ", i , "-", start );
//         await sleep(getTimeInvoke,i);
//     }
// } 

// const sleep = require('util').promisify(setTimeout)

// async function main() {
//     for (var i = 0; i  < 5; i ++) {
//         console.log("Slept for1", i);
//         await timeout(1000);
//         console.log("Slept for2", i);
//     }
// }

// main()


async function getTimer() {
    for (var i = 0; i < 2 * 8; i ++) {
        var start = Date.now();
        console.log("starting timer: ", i , "-", start );
        await setTimeout(function() {
        },125);
        getTimeInvoke(i);
    }
}

async function getTimeInvoke(i) {
    if (i == 2) {    
        await timeout(300);
        var end = Date.now();
        console.log("ending timer: ", i , "-", end );
    } else {
        await timeout(400);
        var end = Date.now();
        console.log("ending timer: ", i , "-", end );
    }
}

getTimer();

