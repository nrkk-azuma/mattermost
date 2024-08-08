// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package main

import (
	"os"
	"time"

	_ "github.com/golang/mock/mockgen/model"
	"github.com/mattermost/mattermost/server/v8/cmd/mmctl/commands"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	NewRelicAgent, err := newrelic.NewApplication(newrelic.ConfigFromEnvironment())
	if err != nil {
		panic(err)
	}

	if err := commands.Run(os.Args[1:]); err != nil {
		os.Exit(1)
	}

	NewRelicAgent.Shutdown(5 * time.Second)
}
