package metric_registrar_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"

    "testing"
)

func TestMetricRegistrarClient(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Metric Registrar Suite")
}
