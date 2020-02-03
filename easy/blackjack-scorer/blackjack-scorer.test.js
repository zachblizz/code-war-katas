const scoreHand = require("./blackjack-scorer")

// Calculate scores with cards 2-10
test("should equal 5", () => expect(scoreHand(["2", "3"])).toEqual(5))
test("should equal 22 (7,7,8)", () => expect(scoreHand(["7", "7", "8"])).toEqual(22))
test("should equal 19", () => expect(scoreHand(["4", "7", "8"])).toEqual(19))

// should score J, Q and K as 10
test("should equal 30", () => expect(scoreHand(["K", "J", "Q"])).toEqual(30))

// should core hands with Aces correctly
test("should equal 14", () => expect(scoreHand(["A", "3"])).toEqual(14))
test("should equal 22", () => expect(scoreHand(["A", "2", "A", "9", "9"])).toEqual(22))
test("should equal 12", () => expect(scoreHand(["A", "A"])).toEqual(12))
test("should equal 12", () => expect(scoreHand(["A", "J", "A"])).toEqual(12))
test("should equal 12", () => expect(scoreHand(["A", "J"])).toEqual(21))