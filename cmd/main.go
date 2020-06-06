package main

import (
	"github.com/edgexfoundry/device-random"
	"github.com/edgexfoundry/device-random/driver"
	"github.com/edgexfoundry/device-sdk-go/pkg/startup"
)

const (
	serviceName string = "device-random"
)

func main() {
	d := driver.NewProtocolDriver()
	startup.Bootstrap(serviceName, device_random.Version, d)

}
