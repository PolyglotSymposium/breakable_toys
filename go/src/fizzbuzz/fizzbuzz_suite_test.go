package fizzbuzz

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGolangSandbox(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FizzBuzz Suite")
}
