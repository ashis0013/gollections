package gollections_test

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var (
	// Declarations for Ginkgo/GoMega DSL
	BeforeEach     = ginkgo.BeforeEach
	BeNil          = gomega.BeNil
    BeFalse        = gomega.BeFalse
    BeTrue         = gomega.BeTrue
	Context        = ginkgo.Context
	Describe       = ginkgo.Describe
	Equal          = gomega.Equal
	Expect         = gomega.Expect
	It             = ginkgo.It
)

func TestGollections(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Gollections Test Suite")
}
