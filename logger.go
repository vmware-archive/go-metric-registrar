package metric_registrar

import (
    "encoding/json"
    "log"
    "os"
)

type printer interface {
    Printf(format string, v ...interface{})
}

type StructuredLogger struct {
    defaultTags map[string]string
    printer
}

func NewLogger(options ...LoggerOption) *StructuredLogger {
    logger := &StructuredLogger{
        printer: log.New(os.Stdout, "", 0),
    }

    for _, option := range options {
        option(logger)
    }

    return logger
}

type LoggerOption func(logger *StructuredLogger)

func WithDefaultTags(defaultTags map[string]string) LoggerOption {
    return func(logger *StructuredLogger) {
        logger.defaultTags = defaultTags
    }
}

func WithPrinter(loggerPrinter printer) LoggerOption {
    return func(logger *StructuredLogger) {
        logger.printer = loggerPrinter
    }
}


type event struct {
    Type  string            `json:"type"`
    Title string            `json:"title"`
    Body  string            `json:"body"`
    Tags  map[string]string `json:"tags"`
}

func (l *StructuredLogger) LogEvent(title, body string, tags map[string]string) {
    bytes, err := json.Marshal(&event{
        Type:  "event",
        Title: title,
        Body:  body,
        Tags:  l.addDefaultTags(tags),
    })
    if err != nil {
        l.Printf("unable to marshal event json: %s\n", err)
        return
    }

    l.Printf("%s\n", bytes)
}

type gauge struct {
    Type  string            `json:"type"`
    Name  string            `json:"name"`
    Value float64           `json:"value"`
    Tags  map[string]string `json:"tags"`
}

func (l *StructuredLogger) LogGauge(name string, value float64, tags map[string]string) {
    bytes, err := json.Marshal(&gauge{
        Type:  "gauge",
        Name:  name,
        Value: value,
        Tags:  l.addDefaultTags(tags),
    })
    if err != nil {
        l.Printf("unable to marshal gauge json: %s\n", err)
        return
    }

    l.Printf("%s\n", bytes)
}

type counter struct {
    Type  string            `json:"type"`
    Name  string            `json:"name"`
    Delta uint              `json:"delta"`
    Tags  map[string]string `json:"tags"`
}

func (l *StructuredLogger) LogCounter(name string, delta uint, tags map[string]string) {
    bytes, err := json.Marshal(&counter{
        Type:  "counter",
        Name:  name,
        Delta: delta,
        Tags:  l.addDefaultTags(tags),
    })
    if err != nil {
        l.Printf("unable to marshal counter json: %s\n", err)
        return
    }

    l.Printf("%s\n", bytes)
}

func (l *StructuredLogger) addDefaultTags(tags map[string]string) map[string]string {
    if tags == nil {
        tags = map[string]string{}
    }

    for tag, value := range l.defaultTags {
        tags[tag] = value
    }

    return tags
}