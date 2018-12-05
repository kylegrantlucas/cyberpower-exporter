package main

import (
	"log"
	"net/http"
	"os"

	cyberpower "github.com/kylegrantlucas/cyberpower-exporter/services"

	"github.com/kylegrantlucas/cyberpower-exporter/collectors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	foo := collectors.NewCyberpowerCollector(cyberpower.Client{Host: os.Getenv("CYBERPOWER_HOST"), Port: os.Getenv("CYBERPOWER_PORT")})
	prometheus.MustRegister(foo)

	http.Handle("/metrics", promhttp.Handler())
	log.Print("Beginning to serve on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
