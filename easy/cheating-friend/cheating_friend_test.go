package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/zach-blizz/code-war-katas/cheating-friend"
)

func dotest(n uint64, exp [][2]uint64) {
	var ans = RemovNb(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("CheatingFriend", func() {
	It("should handle basic cases howMuch", func() {
		dotest(26, [][2]uint64{{15, 21}, {21, 15}})
		dotest(100, nil)
		dotest(101, [][2]uint64{{55, 91}, {91, 55}})
		dotest(102, [][2]uint64{{70, 73}, {73, 70}})
		dotest(110, [][2]uint64{{70, 85}, {85, 70}})
	})
})
