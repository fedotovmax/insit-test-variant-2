package redis

import "fmt"

func Operation(id string) string {
	return fmt.Sprintf("Operation:%s", id)
}
