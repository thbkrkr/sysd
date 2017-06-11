package main

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/docker"
)

type DockerMetrics struct {
	Info *docker.CgroupDockerStat `json:"sys.docker.info"`
	CPU  *cpu.TimesStat           `json:"sys.docker.cpu"`
	Mem  *docker.CgroupMemStat    `json:"sys.docker.mem"`
}

func GetDockerMetrics() []DockerMetrics {
	dockerStat, _ := docker.GetDockerStat()
	m := make([]DockerContainerMetrics, len(dockerStat))

	for i, c := range dockerStat {
		cpu, _ := docker.CgroupCPUDocker(c.ContainerID)
		mem, _ := docker.CgroupMemDocker(c.ContainerID)
		m[i] = DockerContainerMetrics{
			Info: &c,
			CPU:  cpu,
			Mem:  mem,
		}
	}

	return m
}
