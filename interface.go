package prohermes

type IPromelven interface {
	NewPrometheusHook() (*PrometheusHook, error)
}
