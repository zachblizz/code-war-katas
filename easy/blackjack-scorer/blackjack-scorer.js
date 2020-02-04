const handMap = {
  "J": 10,
  "Q": 10,
  "K": 10
}

function scoreHand_(cards) {
  let result = 0
  let As = 0

  for (const c of cards) {
    if (handMap[c] && c !== "A") {
      result += handMap[c]
    } else if (c === "A") {
      As++
    } else {
      result += parseInt(c, 10)
    }
  }

  tmpResult = result
  for (i = 0; i < As; i++) {
    if (tmpResult + 11 <= 21) {
      tmpResult += 11
    } else {
      tmpResult += 1
    }
  }

  if (tmpResult > 21) {
    for (i = 0; i < As; i++) {
      result += 1
    }
  } else {
    result = tmpResult
  }

  return result
}

const cardsMap = {1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9, 10: 10, J: 10, Q: 10, K: 10, A: 1}

function scoreHand(cards){
  const sum = cards.reduce((acc, card) => acc + cardsMap[card], 0)
  return sum < 12 && cards.includes('A') ? sum + 10 : sum
}

module.exports = scoreHand
