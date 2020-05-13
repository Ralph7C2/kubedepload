package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PodData struct {
	Name string `json:"name"`
	Resources []Usage
}

type Usage struct {
	CPU string `json:"cpu"`
	Mem string `json:"memory"`
}

var usageByPod = make(map[string]Usage)

func main() {
	podUsage := strings.Split(os.Args[1], "\n")
	for _, pod := range podUsage {
		pparts := strings.Split(pod, " ")
		usageByPod[pparts[0]] = Usage{pparts[1], pparts[2]}
	}
	var pods []PodData
	err := json.Unmarshal([]byte(os.Args[2]), &pods)
	if err != nil {
		panic(err)
	}
	for _, pod := range pods {
		name := pod.Name
		var cpu, mem int
		for _, res := range pod.Resources {
			cpu += parseCPU(res.CPU)
			mem += parseMem(res.Mem)
		}
		usedCPU := parseCPU(usageByPod[name].CPU)
		usedMem := parseMem(usageByPod[name].Mem)
		fmt.Printf("%s CPU: %.2f%% MEM: %.2f%%\n", pod.Name, (float64(usedCPU)/float64(cpu))*100, (float64(usedMem)/float64(mem))*100)
	}
}

func parseCPU(cpuStr string) int {
	if strings.HasSuffix(cpuStr, "m") {
		c := strings.TrimSuffix(cpuStr, "m")
		cpu, err := strconv.Atoi(c)
		if err != nil {
			panic(err)
		}
		return cpu
	}
	cpu, err := strconv.Atoi(cpuStr)
	if err != nil {
		panic(err)
	}
	return cpu * 1000
}

func parseMem(memStr string) int {
	if strings.HasSuffix(memStr,  "Mi") {
		m := strings.TrimSuffix(memStr, "Mi")
		mem, err := strconv.Atoi(m)
		if err != nil {
			panic(err)
		}
		return mem
	}
	if strings.HasSuffix(memStr, "Gi") {
		m := strings.TrimSuffix(memStr, "Gi")
		mem, err := strconv.Atoi(m)
		if err != nil {
			panic(err)
		}
		return mem * 1024
	}
	panic(memStr)
}
