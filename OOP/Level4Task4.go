package main

import (
	"fmt"
)

type Logger interface {
	Log(msg string)
}

type StdoutLogger struct{}

func (l StdoutLogger) Log(msg string) {
	fmt.Println("Message:", msg)
}

type MockLogger struct {
	Message []string
}

func (m *MockLogger) Log(msg string) {
	m.Message = append(m.Message, msg)
}
