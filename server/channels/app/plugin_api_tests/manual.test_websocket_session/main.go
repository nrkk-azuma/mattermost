// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package main

import (
	"time"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Plugin struct {
	plugin.MattermostPlugin
	sessionCh chan string
}

func (p *Plugin) MessageWillBePosted(_ *plugin.Context, _ *model.Post) (*model.Post, string) {
	return nil, <-p.sessionCh
}

func (p *Plugin) WebSocketMessageHasBeenPosted(connID, userID string, req *model.WebSocketRequest) {
	p.sessionCh <- req.Session.Id
}

func main() {
	NewRelicAgent, err := newrelic.NewApplication(newrelic.ConfigFromEnvironment())
	if err != nil {
		panic(err)
	}

	plugin.ClientMain(&Plugin{
		sessionCh: make(chan string, 1),
	})

	NewRelicAgent.Shutdown(5 * time.Second)
}
