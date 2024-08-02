package handlers

import (
	"fmt"

	"GpioController/pkg/utils"
	"GpioController/types"
)

func GetUserInput(c *types.Controller) error {
	var (
		userPin        int
		userTickerTime int
		err            error
	)

	// Loop until valid GPIO pin number is entered
	for {
		fmt.Print("===> Enter GPIO pin number: ")
		_, err = fmt.Scan(&userPin)
		if err != nil {
			fmt.Println("Error reading input:", err)
			utils.ClearBuffer()
			continue
		}

		if !(utils.IsValidPin(userPin)) {
			fmt.Println("Invalid GPIO pin number. Please enter a valid pin.")
			continue
		}

		// If GPIO pin is valid, break out of the loop
		break
	}

	// Loop until valid ticker time is entered
	for {
		fmt.Print("===> Enter ticker time in seconds: ")
		_, err = fmt.Scan(&userTickerTime)
		if err != nil {
			fmt.Println("Error reading input:", err)
			utils.ClearBuffer()
			continue
		}

		if !utils.IsValidTicker(userTickerTime) {
			fmt.Println("Invalid ticker time. Must be between 1 and 600 seconds.")
			continue
		}

		// If ticker time is valid, break out of the loop
		break
	}

	utils.InitControllerFields(c, userPin, userTickerTime)

	return nil
}
