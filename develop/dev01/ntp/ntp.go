package ntp

import (
	"fmt"
	"github.com/beevik/ntp"
)

// // Hi returns a friendly greeting
func TimeNow() error {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return err
	} else {
		fmt.Println(time)
	}
	return nil
}
