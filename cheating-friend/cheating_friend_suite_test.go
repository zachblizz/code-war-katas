package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCheatingFriend(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CheatingFriend Suite")
}
