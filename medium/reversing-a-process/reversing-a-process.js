const alpha = {
  'a': 0, 'b': 1, 'c': 2,
  'd': 3, 'e': 4, 'f': 5,
  'g': 6, 'h': 7, 'i': 8,
  'j': 9, 'k': 10, 'l': 11,
  'm': 12, 'n': 13, 'o': 14,
  'p': 15, 'q': 16, 'r': 17,
  's': 18, 't': 19, 'u': 20,
  'v': 21, 'w': 22, 'x': 23,
  'y': 24, 'z': 25
}

// x * num % 26
function decode(r) {
  const alphaArr = Object.keys(alpha)
  const strArr = []
  const strNum = []
  
  for (const s of r) {
    if (/[0-9]/.test(s)) {
      strNum.push(s)
    } else {
      strArr.push(s)
    }
  }
  
  const origR = strArr.join("")
  const num = parseInt(strNum.join(""), 10)

  console.log({num, origR, len: origR.length})
  const newR = strArr.map(s => alphaArr[alpha[s] * 45 % 26]).join("")

  return newR !== origR ? newR : "Impossible to decode"
}

// 15
// console.log(decode("1273409kuqhkoynvvknsdwljantzkpnmfgf"))

// 45
console.log(decode("1544749cdcizljymhdmvvypyjamowl"))

module.exports = decode
