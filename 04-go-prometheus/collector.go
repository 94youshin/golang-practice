package main

// import (
// 	"net/http"

// 	"github.com/prometheus/client_golang/prometheus"
// 	"github.com/prometheus/client_golang/prometheus/collectors"
// 	"github.com/prometheus/client_golang/prometheus/promhttp"
// )

// func main() {

// 	reg := prometheus.NewRegistry()
// 	reg.MustRegister(collectors.NewBuildInfoCollector())
// 	reg.MustRegister(collectors.NewGoCollector(collectors.WithGoCollections(collectors.GoRuntimeMemStatsCollection | collectors.GoRuntimeMetricsCollection)))

// 	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{EnableOpenMetrics: true}))
// 	http.ListenAndServe(":8080", nil)
// }
