package interaction_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestInteraction(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Interaction Suite")
}
