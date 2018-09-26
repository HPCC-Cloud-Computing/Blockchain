
var a = "aaa";
for (var i = 0; i < 10; i++) {
    a = a + i.toString();
    console.log("//");
    testThroughput(a);
}

async function testThroughput(a) {
  await wait(5000);
  test(a);
}
function wait(ms){
     return new Promise(r => setTimeout(r, ms));
}
function test(a) {
    console.log(a);
}