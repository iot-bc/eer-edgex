package main

import (
	device_random "eer-edgex"
	"eer-edgex/config"
	"eer-edgex/driver"
	"github.com/edgexfoundry/device-sdk-go/pkg/startup"
)

const (
	serviceName string = "device-random"
)

func main() {

	// 先进行基础配置
	config.Config()

	d := driver.NewProtocolDriver()
	startup.Bootstrap(serviceName, device_random.Version, d)

}
