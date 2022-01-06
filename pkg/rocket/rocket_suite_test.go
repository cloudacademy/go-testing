package rocket_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRocket(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rocket Suite")
}
