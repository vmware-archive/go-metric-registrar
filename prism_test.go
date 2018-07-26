package prism_test

import (
    . "github.com/onsi/ginkgo"
    "fmt"
    . "github.com/onsi/gomega"
    "github.com/pivotal-cf/go-prism"
)

var _ = Describe("PrismClient", func() {
    It("logs events", func() {
        p := newMockPrinter()
        prismLogger := prism.New(prism.WithDefaultTags(map[string]string{
            "globalTag": "globalValue",
        }), prism.WithPrinter(p))

        prismLogger.LogEvent("title", "body", map[string]string{"tag": "tag value"})

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
        prismLogger := prism.New(prism.WithDefaultTags(map[string]string{
            "globalTag": "globalValue",
        }), prism.WithPrinter(p))

        prismLogger.LogGauge("name", 1.5, map[string]string{"tag": "tag value"})

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
        prismLogger := prism.New(prism.WithDefaultTags(map[string]string{
            "globalTag": "globalValue",
        }), prism.WithPrinter(p))

        prismLogger.LogCounter("name", 1, map[string]string{"tag": "tag value"})

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
