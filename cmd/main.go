// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018-2019 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/edgexfoundry/device-random"
	"github.com/edgexfoundry/device-random/driver"
	"github.com/edgexfoundry/device-sdk-go/pkg/startup"
)

const (
	serviceName string = "eer-edgex"
)

func main() {
	d := driver.NewProtocolDriver()
	startup.Bootstrap(serviceName, eer_edgex.Version, d)
}
