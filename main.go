package main

import (
	"fmt"
	"sort"

	"github.com/shirou/gopsutil/process"
)

var processTree = make(map[int32]*process.Process)

func main() {
	processes, _ := process.Processes()

	var processMap = make(map[int32][]int32)

	for _, proc := range processes {
		pid := proc.Pid
		ppid, _ := proc.Ppid()
		processMap[ppid] = append(processMap[ppid], pid)
		processTree[pid] = proc
	}

	for pid, proc := range processTree {
		if _, ok := processMap[pid]; !ok {
			printProcess(proc, 0, processMap)
		}
	}
}

func printProcess(proc *process.Process, depth int, processMap map[int32][]int32) {
	name, _ := proc.Name()
	fmt.Printf("%s%d %s\n", indent(depth), proc.Pid, name)
	if children, ok := processMap[proc.Pid]; ok {
		sort.Slice(children, func(i, j int) bool {
			return children[i] < children[j]
		})
		for _, childPid := range children {
			childProc := processTree[childPid]
			printProcess(childProc, depth+1, processMap)
		}
	}
}

func indent(depth int) string {
	return fmt.Sprintf("%s", fmt.Sprintf("%s", "  ")[:depth*2])
}
