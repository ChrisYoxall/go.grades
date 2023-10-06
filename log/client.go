package log

import (
	"bytes"
	"fmt"
	"grades/registry"
	stlog "log"
	"net/http"
)

func SetClientLogger(serviceURL string, clientService registry.ServiceName) {
	stlog.SetPrefix(fmt.Sprintf("[%v] - ", clientService))
	stlog.SetFlags(0)
	logger = clientLogger{url: serviceURL}
	stlog.SetOutput(&logger)
}

type clientLogger struct {
	url string
}

var logger clientLogger

func (cl clientLogger) Write(data []byte) (int, error) {
	b := bytes.NewBuffer(data)
	res, err := http.Post(cl.url+"/log", "text/plain", b)
	if err != nil {
		return 0, err
	}
	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to send log message. Service responded with %v - %v", res.StatusCode, res.Status)
	}
	return len(data), nil
}

func Debug(m string) {
	_, err := logger.Write([]byte("[DEBUG] " + m))
	if err != nil {
		stlog.Printf("Could not write log message: %v", err)
	}
}
