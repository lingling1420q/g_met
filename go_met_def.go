// GMet is golang client of XMet API.
// For more, see also: https://github.com/chaocai2001/g_met
// Created on 2018.5
package g_met

// Metric Item
type MetricItem struct {
	Key   string      // metric item name
	Value interface{} // metric value
}

type GMet interface {
	// Send metric
	Send(metrics ...MetricItem) error
	// flush out the data from cache
	Flush()
	Close() error
}


type MetAggregator interface {
	// aggregate each metrics, implement by system module
	Aggregate(metrics []MetricItem) error
	GetMetrics() (metrics []MetricItem)
}

type MetWriter interface {
	// write the formatted metrics
	Write(msg string)
	// flush out the data from cache
	Flush()
	Close() error
}

type MetFormatter interface {
	Format(metrics []MetricItem) (string, error)
}

// Create a MetricItem. Metric method helps you write the concise code
func Metric(key string, value interface{}) MetricItem {
	return MetricItem{key, value}
}
