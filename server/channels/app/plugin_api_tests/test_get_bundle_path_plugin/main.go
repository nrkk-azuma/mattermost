// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"
	"github.com/mattermost/mattermost/server/v8/channels/app/plugin_api_tests"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type MyPlugin struct {
	plugin.MattermostPlugin
	configuration plugin_api_tests.BasicConfig
}

func (p *MyPlugin) OnConfigurationChange() error {
	if err := p.API.LoadPluginConfiguration(&p.configuration); err != nil {
		return err
	}
	return nil
}

func (p *MyPlugin) MessageWillBePosted(_ *plugin.Context, _ *model.Post) (*model.Post, string) {
	bundlePath, err := p.API.GetBundlePath()
	if err != nil {
		return nil, err.Error() + "failed get bundle path"
	} else if bundlePathFromConfig, _ := filepath.Abs(filepath.Join(*p.API.GetConfig().PluginSettings.Directory, "test_get_bundle_path_plugin")); bundlePathFromConfig != bundlePath {
		return nil, fmt.Sprintf("Invalid bundle path returned: %v vs %v", bundlePathFromConfig, bundlePath)
	}

	return nil, "OK"
}

func main() {
	NewRelicAgent, err := newrelic.NewApplication(newrelic.ConfigFromEnvironment())
	if err != nil {
		panic(err)
	}

	plugin.ClientMain(&MyPlugin{})

	NewRelicAgent.Shutdown(5 * time.Second)
}
