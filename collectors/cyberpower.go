package collectors

import (
	"log"

	cyberpower "github.com/kylegrantlucas/cyberpower-exporter/services"
	"github.com/prometheus/client_golang/prometheus"
)

var client cyberpower.Client

//Define a struct for you collector that contains pointers
//to prometheus descriptors for each metric you wish to expose.
//Note you can also include fields of other types if they provide utility
//but we just won't be exposing them as metrics.
type cyberpowerCollector struct {
	Wattage *prometheus.Desc
	Load    *prometheus.Desc
}

//You must create a constructor for you collector that
//initializes every descriptor and returns a pointer to the collector
func NewCyberpowerCollector(c cyberpower.Client) *cyberpowerCollector {
	client = c
	return &cyberpowerCollector{
		Wattage: prometheus.NewDesc("wattage",
			"The wattage on the UPS",
			nil, nil,
		),
		Load: prometheus.NewDesc("load",
			"The load on the UPS",
			nil, nil,
		),
	}
}

//Each and every collector must implement the Describe function.
//It essentially writes all descriptors to the prometheus desc channel.
func (collector *cyberpowerCollector) Describe(ch chan<- *prometheus.Desc) {

	//Update this section with the each metric you create for a given collector
	ch <- collector.Wattage
	ch <- collector.Load
}

//Collect implements required collect function for all promehteus collectors
func (collector *cyberpowerCollector) Collect(ch chan<- prometheus.Metric) {
	resp, err := client.Poll()
	if err != nil {
		log.Print(err)
	}

	//Write latest value for each metric in the prometheus metric channel.
	//Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.
	ch <- prometheus.MustNewConstMetric(collector.Wattage, prometheus.GaugeValue, float64(resp.Status.Output.Watt))
	ch <- prometheus.MustNewConstMetric(collector.Load, prometheus.GaugeValue, float64(resp.Status.Output.Load))

}
