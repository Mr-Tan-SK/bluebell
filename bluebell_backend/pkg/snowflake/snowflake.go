package snowflake

import (
	"fmt"
	"github.com/sony/sonyflake"
	"time"
)

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

// Init 需传⼊当前的机器ID
func Init(startTime string, machineId uint16) (err error) {
	sonyMachineID = machineId
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return err
	}
	settings := sonyflake.Settings{
		StartTime: st,
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}

// GenID ⽣成id
func GenID() (int64, error) {
	if sonyFlake == nil {
		err := fmt.Errorf("snoy flake not inited")
		return 0, err
	}
	id, err := sonyFlake.NextID()
	return int64(id), err
}
