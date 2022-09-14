package main

import (
	"fmt"
	"ntp/ntp"
)

func main() {

	err := ntp.TimeNow()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(err)

}
