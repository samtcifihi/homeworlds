package userio

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// Int gets an int within a specified range from the user
func Int(min int, max int, cin *bufio.Reader) int {
	userIntA := ""
	userInt := 0

	validUserChoice := false
	for validUserChoice == false {
		userIntA, _ = cin.ReadString('\n')
		userIntA = strings.Replace(userIntA, "\n", "", -1)

		// Convert to an int
		userIntL, err := strconv.Atoi(userIntA)
		if err != nil {
			fmt.Println(err) // Debugging

			fmt.Println("Invalid Input")
		} else {
			userInt = userIntL
			validUserChoice = true // Exit Loop
		}

		// If it passed the above test, ensure the number is within bounds
		if validUserChoice == true {
			if userInt < min {
				fmt.Println("That number is too low.")
				validUserChoice = false // Do not Exit Loop
			} else if userInt > max {

				fmt.Println("That number is too high.") // Do not Exit Loop
				validUserChoice = false
			}

		}

	} // User has entered a valid choice

	fmt.Println()

	return userInt
}

// String gets a string from the user
func String(cin *bufio.Reader) string {

	userReply, _ := cin.ReadString('\n')
	userReply = strings.Replace(userReply, "\n", "", -1)

	fmt.Println()

	return userReply
}
