package devices

import (
	"fmt"
	"sync"
	"time"

	"github.com/vapor-ware/synse-sdk/sdk"
)

const (
	// Valid statuses of a lock device.
	statusLock   = "locked"
	statusUnlock = "unlocked_electrically"

	// Valid actions of a lock device.
	actionLock        = "lock"
	actionUnlock      = "unlock"
	actionPulseUnlock = "pulseUnlock"
)

var (
	// mux provides mutual exclusion for reading/writing to lock status.
	mux sync.Mutex
)

// Lock is the handler for the emulated Lock device(s).
var Lock = sdk.DeviceHandler{
	Name:  "lock",
	Read:  lockRead,
	Write: lockWrite,
}

// lockRead is the read handler for the emulated Lock device(s). It
// returns the status values for the device. If no status has previously
// been set, this will set the status to 'locked'.
func lockRead(device *sdk.Device) ([]*sdk.Reading, error) {
	mux.Lock()
	defer mux.Unlock()

	var lockStatus string

	dState, ok := deviceState[device.GUID()]

	if ok {
		if _, ok := dState["lockStatus"]; ok {
			lockStatus = dState["lockStatus"].(string)
		}
	}

	// if we have no stored device lock status, default to "locked"
	if lockStatus == "" {
		lockStatus = statusLock
	}

	statusReading, err := device.GetOutput("lock.status").MakeReading(lockStatus)
	if err != nil {
		return nil, err
	}

	return []*sdk.Reading{
		statusReading,
	}, nil
}

// lockWrite is the write handler for the emulated Lock device(s). It
// sets the status values for the device.
func lockWrite(device *sdk.Device, data *sdk.WriteData) error {
	mux.Lock()
	defer mux.Unlock()

	dState, ok := deviceState[device.GUID()]

	switch action := data.Action; action {
	case actionLock:
		if !ok {
			deviceState[device.GUID()] = map[string]interface{}{"lockStatus": statusLock}
		} else {
			dState["lockStatus"] = statusLock
		}
	case actionUnlock:
		if !ok {
			deviceState[device.GUID()] = map[string]interface{}{"lockStatus": statusUnlock}
		} else {
			dState["lockStatus"] = statusUnlock
		}
	case actionPulseUnlock:
		// Unlock the device for 5 seconds then lock it.
		if !ok {
			deviceState[device.GUID()] = map[string]interface{}{"lockStatus": statusUnlock}
		} else {
			dState["lockStatus"] = statusUnlock
		}

		go func() {
			time.Sleep(5 * time.Second)

			mux.Lock()
			defer mux.Unlock()

			if !ok {
				deviceState[device.GUID()] = map[string]interface{}{"lockStatus": statusLock}
			} else {
				dState["lockStatus"] = statusLock
			}
		}()
	default:
		return fmt.Errorf("unsupported command for status action: %v", action)
	}

	return nil
}