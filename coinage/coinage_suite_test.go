package coinage_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCoinage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Coinage Suite")
}
