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

func (p *Plugin) ServeHTTP(_ *plugin.Context, w http.ResponseWriter, _ *http.Request) {
	hj, ok := w.(http.Hijacker)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	conn, brw, err := hj.Hijack()
	if conn == nil || brw == nil || err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	conn.Write([]byte("HTTP/1.1 200\n\nOK"))
	conn.Close()
}

func main() {
	NewRelicAgent, err := newrelic.NewApplication(newrelic.ConfigFromEnvironment())
	if err != nil {
		panic(err)
	}

	plugin.ClientMain(&Plugin{})

	NewRelicAgent.Shutdown(5 * time.Second)
}
