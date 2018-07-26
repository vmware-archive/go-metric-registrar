package prism_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"

    "testing"
)

func TestPrismClient(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "PrismClient Suite")
}
