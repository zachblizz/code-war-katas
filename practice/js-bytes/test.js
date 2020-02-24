const sha1 = require("sha1");


function testSha1ByteStuff(input) {
  const sinput = sha1(input)
  const ret = Buffer.from(sinput, "hex");

  debugger;
  console.log(sinput);
  console.log(ret.readInt8(0));
  console.log(ret.readInt8(1));
  return ret;
}

console.log(testSha1ByteStuff("hello world"));
