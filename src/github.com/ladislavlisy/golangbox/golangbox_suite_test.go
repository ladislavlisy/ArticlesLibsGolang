package golangbox_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGolangbox(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Golangbox Suite")
}
