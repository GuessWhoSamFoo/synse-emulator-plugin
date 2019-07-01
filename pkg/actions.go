package pkg

import (
	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/sdk"
)

// ActionAirflowValueEmitterSetup initializes a ValueEmitter for each "airflow" type device.
var ActionAirflowValueEmitterSetup = sdk.DeviceAction{
	Name: "airflow value emitter setup",
	Filter: map[string][]string{
		"type": {"airflow"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		emitter := utils.NewValueEmitter(utils.RandomWalk).WithLowerBound(-100).WithUpperBound(100)
		return utils.SetEmitter(d.GetID(), emitter)
	},
}

// ActionEnergyValueEmitterSetup initializes a ValueEmitter for each "energy" type device.
var ActionEnergyValueEmitterSetup = sdk.DeviceAction{
	Name: "energy value emitter setup",
	Filter: map[string][]string{
		"type": {"energy"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		emitter := utils.NewValueEmitter(utils.Accumulate).WithLowerBound(0).WithUpperBound(100000)
		return utils.SetEmitter(d.GetID(), emitter)
	},
}

// ActionFanValueEmitterSetup initializes a ValueEmitter for each "fan" type device.
var ActionFanValueEmitterSetup = sdk.DeviceAction{
	Name: "fan value emitter setup",
	Filter: map[string][]string{
		"type": {"fan"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		emitter := utils.NewValueEmitter(utils.Store).WithSeed(0)
		return utils.SetEmitter(d.GetID(), emitter)
	},
}

// ActionHumidityValueEmitterSetup initializes a ValueEmitter for each "humidity" type device.
var ActionHumidityValueEmitterSetup = sdk.DeviceAction{
	Name: "humidity value emitter setup",
	Filter: map[string][]string{
		"type": {"humidity"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		emitter := utils.NewValueEmitter(utils.RandomWalk).WithLowerBound(0).WithUpperBound(100)
		return utils.SetEmitter(d.GetID(), emitter)
	},
}

// ActionLEDValueEmitterSetup initializes a ValueEmitter for each "led" type device.
var ActionLEDValueEmitterSetup = sdk.DeviceAction{
	Name: "LED value emitter setup",
	Filter: map[string][]string{
		"type": {"led"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		emitter := utils.NewValueEmitter(utils.Store).WithSeed(map[string]string{
			"state": "off",
			"color": "000000",
		})
		return utils.SetEmitter(d.GetID(), emitter)
	},
}

// ActionLockValueEmitterSetup initializes a ValueEmitter for each "lock" type device.
var ActionLockValueEmitterSetup = sdk.DeviceAction{
	Name: "lock value emitter setup",
	Filter: map[string][]string{
		"type": {"lock"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		emitter := utils.NewValueEmitter(utils.Store).WithSeed("locked")
		return utils.SetEmitter(d.GetID(), emitter)
	},
}

// ActionPowerValueEmitterSetup initializes a ValueEmitter for each "power" type device.
var ActionPowerValueEmitterSetup = sdk.DeviceAction{
	Name: "power value emitter setup",
	Filter: map[string][]string{
		"type": {"power"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		emitter := utils.NewValueEmitter(utils.RandomWalk).WithLowerBound(1000).WithUpperBound(3000)
		return utils.SetEmitter(d.GetID(), emitter)
	},
}

// ActionPressureValueEmitterSetup initializes a ValueEmitter for each "pressure" type device.
var ActionPressureValueEmitterSetup = sdk.DeviceAction{
	Name: "pressure value emitter setup",
	Filter: map[string][]string{
		"type": {"pressure"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		emitter := utils.NewValueEmitter(utils.RandomWalk).WithLowerBound(-5).WithUpperBound(5)
		return utils.SetEmitter(d.GetID(), emitter)
	},
}

// ActionTemperatureValueEmitterSetup initializes a ValueEmitter for each "temperature" type device.
var ActionTemperatureValueEmitterSetup = sdk.DeviceAction{
	Name: "temperature value emitter setup",
	Filter: map[string][]string{
		"type": {"temperature"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		emitter := utils.NewValueEmitter(utils.RandomWalk).WithLowerBound(0).WithUpperBound(100)
		return utils.SetEmitter(d.GetID(), emitter)
	},
}

// ActionVoltageValueEmitterSetup initializes a ValueEmitter for each "voltage" type device.
var ActionVoltageValueEmitterSetup = sdk.DeviceAction{
	Name: "voltage value emitter setup",
	Filter: map[string][]string{
		"type": {"voltage"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		emitter := utils.NewValueEmitter(utils.RandomWalk).WithLowerBound(100).WithUpperBound(500)
		return utils.SetEmitter(d.GetID(), emitter)
	},
}
