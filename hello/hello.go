package hello

import "fmt"

func Hello(firstname string, lastname string) string {
	return fmt.Sprintf("Hello %v %v", firstname, lastname)
}
