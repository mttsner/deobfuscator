package ironbrew

import (
	
)

// InitIronbrew sets up everything needed for it to work
func InitIronbrew() error {
	if err := initVmdata(); err != nil {
		return err
	}
	if _, err := initMapping(); err != nil {
		return err
	}
	return nil
}