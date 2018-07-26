package prism

import (
    "encoding/json"
    "log"
    "os"
)

type printer interface {
    Printf(format string, v ...interface{})
}

type prismLogger struct {
    defaultTags map[string]string
    printer
}

func New(options ...LoggerOption) *prismLogger {
    logger := &prismLogger{
        printer: log.New(os.Stdout, "", 0),
    }

    for _, option := range options {
        option(logger)
    }

    return logger
}

type LoggerOption func(logger *prismLogger)

func WithDefaultTags(defaultTags map[string]string) LoggerOption {
    return func(logger *prismLogger) {
        logger.defaultTags = defaultTags
    }
}

func WithPrinter(loggerPrinter printer) LoggerOption {
    return func(logger *prismLogger) {
        logger.printer = loggerPrinter
    }
}


type event struct {
    Type  string            `json:"type"`
    Title string            `json:"title"`
    Body  string            `json:"body"`
    Tags  map[string]string `json:"tags"`
}

func (l *prismLogger) LogEvent(title, body string, tags map[string]string) {
    for tag, value := range l.defaultTags {
        tags[tag] = value
    }

    bytes, err := json.Marshal(&event{
        Type:  "event",
        Title: title,
        Body:  body,
        Tags:  tags,
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

func (l *prismLogger) LogGauge(name string, value float64, tags map[string]string) {
    for tag, value := range l.defaultTags {
        tags[tag] = value
    }

    bytes, err := json.Marshal(&gauge{
        Type:  "gauge",
        Name:  name,
        Value: value,
        Tags:  tags,
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

func (l *prismLogger) LogCounter(name string, delta uint, tags map[string]string) {
    for tag, value := range l.defaultTags {
        tags[tag] = value
    }

    bytes, err := json.Marshal(&counter{
        Type:  "counter",
        Name:  name,
        Delta: delta,
        Tags:  tags,
    })
    if err != nil {
        l.Printf("unable to marshal counter json: %s\n", err)
        return
    }

    l.Printf("%s\n", bytes)
}
