package main

import (
	"fmt"
	"log"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("time.google.com")
	if err != nil {
		log.Println("Ошибка получения времени:", err)
		os.Exit(1)
	}
	fmt.Println("Точное время:", time)
}
