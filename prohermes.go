package prohermes

import (
	"fmt"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type PrometheusHook struct {
	counterVec *prometheus.CounterVec
}

var supportedLevels = []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel}

func NewPrometheusHook() (*PrometheusHook, error) {
	counterVec := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "log_messages_total",
		Help: "Total number of log messages.",
	}, []string{"level", "message", "status_code"})

	for _, level := range supportedLevels {
		counterVec.WithLabelValues(level.String(), "", "")
	}

	prometheus.Unregister(counterVec)

	if err := prometheus.Register(counterVec); err != nil {
		return nil, err
	}

	return &PrometheusHook{
		counterVec: counterVec,
	}, nil
}

func MustNewPrometheusHook() *PrometheusHook {
	hook, err := NewPrometheusHook()
	if err != nil {
		panic(err)
	}
	return hook
}

func (hook *PrometheusHook) Fire(entry *logrus.Entry) error {
	message := strings.Split(entry.Message, "[status_code]: ")

	if len(message) != 1 {
		entry.Message = strings.ReplaceAll(entry.Message, fmt.Sprintf("[status_code]: %v", message[1]), "")

		if len(message[1]) != 0 {
			hook.counterVec.WithLabelValues(entry.Level.String(), message[0], message[1]).Inc()
		}
	}

	return nil
}

func (hook *PrometheusHook) Levels() []logrus.Level {
	return supportedLevels
}
