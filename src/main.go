package main

import (
	mc "mathcomputation"
	config "configuration"
	"fmt"
)

func main() {
	fmt.Println("Initializing Math Computation....")

	config := &config.Configuration{}
	config.Init()

	ops := &mc.MathOps{}
	monitor := &mc.APIMonitor{}

	mc.SetupRoutes(ops, monitor)
}
