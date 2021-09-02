package main

import (
	"FOFNotifier/handler"
	"time"
)

func main() {
	handler.MainRoutine()

	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		handler.MainRoutine()
	}
}
