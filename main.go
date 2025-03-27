package main

import (
	"github.com/Wareload/service-apisix/internal/oidc"
	"github.com/apache/apisix-go-plugin-runner/pkg/log"
	"github.com/apache/apisix-go-plugin-runner/pkg/plugin"
	"github.com/apache/apisix-go-plugin-runner/pkg/runner"
)

var plugins = map[string]plugin.Plugin{
	"oidc": &oidc.Oidc{},
}

func main() {
	for name, p := range plugins {
		err := plugin.RegisterPlugin(p)
		if err != nil {
			log.Fatalf("failed to register plugin '%s': %s", name, err)
		} else {
			log.Infof("plugin '%s' registered", name)
		}
	}
	runner.Run(runner.RunnerConfig{})
}
