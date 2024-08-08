// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package main

import (
	"time"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/pkg/errors"
)

type PluginUsingLogAPI struct {
	plugin.MattermostPlugin
}

type Foo struct {
	bar float64
}

func main() {
	NewRelicAgent, err := newrelic.NewApplication(newrelic.ConfigFromEnvironment())
	if err != nil {
		panic(err)
	}

	plugin.ClientMain(&PluginUsingLogAPI{})

	NewRelicAgent.Shutdown(5 * time.Second)
}

func (p *PluginUsingLogAPI) MessageWillBePosted(_ *plugin.Context, _ *model.Post) (*model.Post, string) {
	p.API.LogDebug("LogDebug", "one", 1, "two", "two", "foo", Foo{bar: 3.1416})
	p.API.LogInfo("LogInfo", "one", 1, "two", "two", "foo", Foo{bar: 3.1416})
	p.API.LogWarn("LogWarn", "one", 1, "two", "two", "foo", Foo{bar: 3.1416})
	p.API.LogError("LogError", "error", errors.WithStack(errors.New("boom!")))
	return nil, "OK"
}
