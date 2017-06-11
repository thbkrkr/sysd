package main

import (
	"encoding/json"
	"fmt"
)

type GlobalMetrics struct {
	System SystemMetrics
	Docker DockerContainerMetrics
}

func main() {
	g := GlobalMetrics{
		System: GetSystemMetrics(),
		Docker: GetDockerMetrics(),
	}

	print(g)
}

func print(i interface{}) {
	bytes, _ := json.Marshal(i)
	fmt.Println(string(bytes))
}
