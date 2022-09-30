package main

import (
	"fmt"

	"github.com/adamnasrudin03/student-service/configs"
)

func main() {
	config := configs.GetInstance()
	fmt.Printf("Running in port %s", config.Appconfig.Port)
}
