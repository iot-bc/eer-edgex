package driver

import (
	"fmt"
	"sync"
	"time"

	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

var once sync.Once
var driver *MyDriver

type MyDriver struct {
	lc        logger.LoggingClient
	asyncCh   chan<- *dsModels.AsyncValues
	myDevices sync.Map
}

func NewProtocolDriver() dsModels.ProtocolDriver {
	once.Do(func() {
		driver = new(MyDriver)
	})
	return driver
}

func (d *MyDriver) DisconnectDevice(deviceName string, protocols map[string]models.ProtocolProperties) error {
	d.lc.Info(fmt.Sprintf("MyDriver.DisconnectDevice: driver is disconnecting to %s", deviceName))
	return nil
}

func (d *MyDriver) Initialize(lc logger.LoggingClient, asyncCh chan<- *dsModels.AsyncValues, deviceCh chan<- []dsModels.DiscoveredDevice) error {
	d.lc = lc
	d.asyncCh = asyncCh
	return nil
}

func (d *MyDriver) HandleReadCommands(deviceName string, protocols map[string]models.ProtocolProperties,
	reqs []dsModels.CommandRequest) (res []*dsModels.CommandValue, err error) {
	rd := d.retrieveDevice(deviceName)

	res = make([]*dsModels.CommandValue, len(reqs))
	now := time.Now().UnixNano()

	for i, req := range reqs {
		v, err := rd.value()
		handle()
		dataCollect()
		if err != nil {
			return nil, err
		}
		var cv *dsModels.CommandValue

		cv = dsModels.NewStringValue(req.DeviceResourceName, now, v)
		res[i] = cv
	}

	return res, nil
}

func (d *MyDriver) retrieveDevice(deviceName string) (rdv *myDevice) {
	rd, ok := d.myDevices.LoadOrStore(deviceName, newDevice())
	if rdv, ok = rd.(*myDevice); !ok {
		panic("The value in myDevices has to be a reference of myDevice")
	}
	return rdv
}

func (d *MyDriver) HandleWriteCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []dsModels.CommandRequest,
	params []*dsModels.CommandValue) error {
	rd := d.retrieveDevice(deviceName)

	for _, param := range params {
		switch param.DeviceResourceName {
		case "eer_data":
			v, err := param.StringValue()
			if err != nil {
				return fmt.Errorf("MyDriver.HandleWriteCommands: %v", err)
			}
			rd.data = v
		default:
			return fmt.Errorf("MyDriver.HandleWriteCommands: there is no matched device resource for %s", param.String())
		}
	}
	return nil
}

func (d *MyDriver) Stop(force bool) error {
	d.lc.Info("Driver.Stop: device driver is stopping...")
	return nil
}

func (d *MyDriver) AddDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
	d.lc.Debug(fmt.Sprintf("a new Device is added: %s", deviceName))
	return nil
}

func (d *MyDriver) UpdateDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
	d.lc.Debug(fmt.Sprintf("Device %s is updated", deviceName))
	return nil
}

func (d *MyDriver) RemoveDevice(deviceName string, protocols map[string]models.ProtocolProperties) error {
	d.lc.Debug(fmt.Sprintf("Device %s is removed", deviceName))
	return nil
}
