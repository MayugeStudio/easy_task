package code

import "fmt"

func PrintErrorMessages(messages []string) {
	if len(messages) == 0 {
		return
	}
	for _, message := range messages {
		fmt.Print(message)
	}
	fmt.Printf("\n")
}
