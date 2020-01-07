const decode = require('./reversing-a-process')

function testDecode(r, expected) {
  const actual = decode(r)

  expect(actual).toEqual(expected)
}

test("basic tests", () => {
  testDecode("6015mer", "ekx")
  testDecode("1544749cdcizljymhdmvvypyjamowl", "mfmwhbpoudfujjozopaugcb")
  testDecode("105860ymmgegeeiwaigsqkcaeguicc", "Impossible to decode")
  testDecode("1122305vvkhrrcsyfkvejxjfvafzwpsdqgp", "rrsxppowmjsrclfljrajtybwviqb")
  testDecode("1273409kuqhkoynvvknsdwljantzkpnmfgf", "uogbucwnddunktsjfanzlurnyxmx")
})
