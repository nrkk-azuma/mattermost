// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package main

import (
	"net/http"
	"time"

	"github.com/mattermost/mattermost/server/public/plugin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Plugin struct {
	plugin.MattermostPlugin
}

func (p *Plugin) ServeMetrics(_ *plugin.Context, w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/subpath" {
		w.Write([]byte("METRICS SUBPATH"))
		return
	}

	w.Write([]byte("METRICS"))
}

func main() {
	NewRelicAgent, err := newrelic.NewApplication(newrelic.ConfigFromEnvironment())
	if err != nil {
		panic(err)
	}

	plugin.ClientMain(&Plugin{})

	NewRelicAgent.Shutdown(5 * time.Second)
}
