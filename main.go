package main

import (
	"encoding/json"
	"fmt"
)

var (
	tagsFlag string

	tags map[string]string
)

type GlobalMetrics struct {
	System SystemMetrics
	Docker []DockerMetrics
}

func main() {
	/*g := GlobalMetrics{
		System: GetSystemMetrics(),
		Docker: GetDockerMetrics(),
	}
	print(g)
	*/

	fmt.Println(GetCPUMetrics())
	//fmt.Println(GetNetMetrics())
}

func print(i interface{}) {
	bytes, _ := json.Marshal(i)
	fmt.Println(string(bytes))
}
