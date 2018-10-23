package metric_registrar_test

import (
    "fmt"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "github.com/pivotal-cf/go-metric-registrar"
)

var _ = Describe("Metric Registrar Logger", func() {
    It("logs events", func() {
        p := newMockPrinter()
        registrarLogger := metric_registrar.NewLogger(metric_registrar.WithDefaultTags(map[string]string{
            "globalTag": "globalValue",
        }), metric_registrar.WithPrinter(p))

        registrarLogger.LogEvent("title", "body", map[string]string{"tag": "tag value"})

        Expect(p.printed).To(Receive(MatchJSON(
            `{
				"type": "event",
				"title": "title",
	     		"body": "body",
				"tags": {
                    "globalTag": "globalValue",
					"tag": "tag value"
				}
			}`,
        )))
    })

    It("logs gauges", func() {
        p := newMockPrinter()
        registrarLogger := metric_registrar.NewLogger(metric_registrar.WithDefaultTags(map[string]string{
            "globalTag": "globalValue",
        }), metric_registrar.WithPrinter(p))

        registrarLogger.LogGauge("name", 1.5, map[string]string{"tag": "tag value"})

        Expect(p.printed).To(Receive(MatchJSON(
            `{
				"type": "gauge",
				"name": "name",
	     		"value": 1.5,
				"tags": {
                    "globalTag": "globalValue",
					"tag": "tag value"
				}
			}`,
        )))
    })

    It("logs counters", func() {
        p := newMockPrinter()
        registrarLogger := metric_registrar.NewLogger(metric_registrar.WithDefaultTags(map[string]string{
            "globalTag": "globalValue",
        }), metric_registrar.WithPrinter(p))

        registrarLogger.LogCounter("name", 1, map[string]string{"tag": "tag value"})

        Expect(p.printed).To(Receive(MatchJSON(
            `{
				"type": "counter",
				"name": "name",
	     		"delta": 1,
				"tags": {
                    "globalTag": "globalValue",
					"tag": "tag value"
				}
			}`,
        )))
    })

    Context("nil tags", func() {
        It("logs events", func() {
            p := newMockPrinter()
            registrarLogger := metric_registrar.NewLogger(metric_registrar.WithDefaultTags(map[string]string{
                "globalTag": "globalValue",
            }), metric_registrar.WithPrinter(p))

            registrarLogger.LogEvent("title", "body", nil)

            Expect(p.printed).To(Receive())
        })

        It("logs gauges", func() {
            p := newMockPrinter()
            registrarLogger := metric_registrar.NewLogger(metric_registrar.WithDefaultTags(map[string]string{
                "globalTag": "globalValue",
            }), metric_registrar.WithPrinter(p))

            registrarLogger.LogGauge("name", 1.5, nil)

            Expect(p.printed).To(Receive())
        })

        It("logs counters", func() {
            p := newMockPrinter()
            registrarLogger := metric_registrar.NewLogger(metric_registrar.WithDefaultTags(map[string]string{
                "globalTag": "globalValue",
            }), metric_registrar.WithPrinter(p))

            registrarLogger.LogCounter("name", 1, nil)

            Expect(p.printed).To(Receive())
        })
    })
})

func newMockPrinter() *mockPrinter {
    return &mockPrinter{
        printed: make(chan string, 10),
    }
}

type mockPrinter struct {
    printed chan string
}

func (p *mockPrinter) Printf(format string, v ...interface{}) {
    p.printed <- fmt.Sprintf(format, v...)
}
