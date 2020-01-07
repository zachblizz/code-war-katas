const lookAndSay = require("./look-and-say")

test("should return 1", () => expect(lookAndSay("1", 1)).toEqual("1"))
test("should return 21", () => expect(lookAndSay("1", 3)).toEqual("21"))
test("should return 111221", () => expect(lookAndSay("1", 5)).toEqual("111221"))
test("should return 22", () => expect(lookAndSay("22", 10)).toEqual("22"))
test("should return 1114", () => expect(lookAndSay("14", 2)).toEqual("1114"))