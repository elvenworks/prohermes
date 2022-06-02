package promelven

type IPromelven interface {
	NewPrometheusHook() (*PrometheusHook, error)
}
