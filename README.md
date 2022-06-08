# Elven Alertmanager golang library
Use Warning logs to send metrics to Prometheus.
Logrus hook to expose the number of log messages as Prometheus metrics:

```
log_messages{level="warning", "message": ""}
log_messages{level="error", "message": ""}
```

## Usage

Sample code:
```go
package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	promelven "github.com/elvenworks/elven-alertmanager-golang"
)

func main() {
	hook := promelven.MustNewPrometheusHook()

	logrus.AddHook(hook)

	go http.ListenAndServe(":8080", promhttp.Handler())

  logrus.Warning("Error Message")
}
```