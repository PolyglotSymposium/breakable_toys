package gosoon_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGosoon(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gosoon Suite")
}
