package vers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestVers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Vers Suite")
}
