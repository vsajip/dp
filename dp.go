package dp

import "fmt"

func Print(a interface{}) (int, error) {
	if a == nil {
		return fmt.Println("[nil]")
	}
	return fmt.Printf("%T(%v)", a, a)
}