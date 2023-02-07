package main

import (
	"fmt"
	"strconv"

	"github.com/shirou/gopsutil/process"
)

func main() {
	processes, _ := process.Processes()
	processMap := make(map[int32]*process.Process)
	for _, p := range processes {
		processMap[p.Pid] = p
	}

	for _, p := range processes {
		pid := p.Pid
		_, err := p.Ppid()
		if err != nil {
			fmt.Println("Error while getting parent process ID:", err)
			return
		}

		for i := 0; i < getDepth(processMap, pid); i++ {
			fmt.Print("    ")
		}

		name, err := p.Name()
		if err != nil {
			fmt.Println("Error while getting process name:", err)
			return
		}

		fmt.Print("└─" + name + "───")
		children, _ := p.Children()
		if len(children) == 0 {
			fmt.Println("1*[{" + name + "}]")
		} else {
			fmt.Println(strconv.Itoa(len(children)) + "*[{" + name + "}]")
			for _, child := range children {
				childPID := child.Pid
				displayProcessTree(processMap, childPID, getDepth(processMap, pid)+1)
			}
		}
	}
}

func getDepth(processMap map[int32]*process.Process, pid int32) int {
	depth := 0
	for {
		ppid, _ := processMap[pid].Ppid()
		if ppid == 0 {
			break
		}
		depth++
		pid = ppid
	}
	return depth
}

func displayProcessTree(processMap map[int32]*process.Process, pid int32, depth int) {
	p := processMap[pid]
	for i := 0; i < depth; i++ {
		fmt.Print("    ")
	}

	name, err := p.Name()
	if err != nil {
		fmt.Println("Error while getting process name:", err)
		return
	}

	fmt.Print("└─" + name + "───")
	children, _ := p.Children()
	if len(children) == 0 {
		fmt.Println("1*[{" + name + "}]")
	} else {
		fmt.Println(strconv.Itoa(len(children)) + "*[{" + name + "}]")
		for _, child := range children {
			childPID := child.Pid
			displayProcessTree(processMap, childPID, depth+1)
		}
	}
}
