package main

import (
	"fmt"
)

func main() {
	sessions := getSessions()

	for _, session := range sessions {
		fmt.Printf("Session %s, %d windows.\n", session.name, session.windows)
	}
}
