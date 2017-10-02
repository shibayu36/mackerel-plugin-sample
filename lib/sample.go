package mpsample

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"

	mp "github.com/mackerelio/go-mackerel-plugin"
)

const version = "v0.0.1"

// SamplePlugin mackerel plugin
type SamplePlugin struct {
	Prefix string
}

// MetricKeyPrefix interface for PluginWithPrefix
func (p *SamplePlugin) MetricKeyPrefix() string {
	if p.Prefix == "" {
		p.Prefix = "sample"
	}
	return p.Prefix
}

// GraphDefinition interface for mackerelplugin
func (p *SamplePlugin) GraphDefinition() map[string]mp.Graphs {
	labelPrefix := strings.Title(p.Prefix)
	return map[string]mp.Graphs{
		"dice": {
			Label: labelPrefix + " Dice Value",
			Unit:  "integer",
			Metrics: []mp.Metrics{
				{Name: "d6", Label: "Dice(d6)"},
				{Name: "d20", Label: "Dice(d20)"},
			},
		},
	}
}

// FetchMetrics interface for mackerelplugin
func (p *SamplePlugin) FetchMetrics() (map[string]float64, error) {
	rand.Seed(time.Now().UnixNano())
	metrics := map[string]float64{
		"d6":  float64(rand.Intn(6) + 1),
		"d20": float64(rand.Intn(20) + 1),
	}
	return metrics, nil
}

// Do the plugin
func Do() {
	optVersion := flag.Bool("version", false, "Show version")
	optPrefix := flag.String("metric-key-prefix", "", "Metric key prefix")
	flag.Parse()

	if *optVersion {
		fmt.Println(version)
		return
	}

	plugin := mp.NewMackerelPlugin(&SamplePlugin{
		Prefix: *optPrefix,
	})
	plugin.Run()
}
