# Elven Alertmanager golang library
Use Warning logs to send metrics to Prometheus.
Logrus hook to expose the number of log messages as Prometheus metrics:

```
log_messages{level="warning", "message": ""}
log_messages{level="error", "message": ""}
```

To use the lib, you need to invoke logrus and send a parameter in the final of the string "[status_code]:" inside the string and the status code you want.

## Usage

Sample code:
```go
package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	promelven "github.com/elvenworks/prohermes"
)

func main() {
	hook := promelven.MustNewPrometheusHook()

	logrus.AddHook(hook)

	go http.ListenAndServe(":8080", promhttp.Handler())

  logrus.Warning("Error Message [status_code]: 55")
}
```