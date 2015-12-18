package gosigar

import (
	"fmt"
	"testing"
)

func TestSigar(t *testing.T) {
	sg := NewSigar()
	sg.Open()
	defer sg.Close()
	//
	//
	fmt.Println("\n\nQueryMem...")
	mem, err := sg.QueryMem()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(mem)
	//
	//
	fmt.Println("\n\nQuerySwap...")
	swap, err := sg.QuerySwap()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(swap)
	//
	//
	fmt.Println("\n\nQueryCpu...")
	cpu, err := sg.QueryCpu()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(cpu)
	//
	//
	fmt.Println("\n\nQueryCpus...")
	cpus, err := sg.QueryCpus()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(cpus)
	//
	//
	fmt.Println("\n\nQueryCpuInfoes...")
	cpuis, err := sg.QueryCpuInfoes()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(cpuis)
	//
	//
	fmt.Println("\n\nQueryUptime...")
	upt, err := sg.QueryUptime()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(upt)
	//
	//
	fmt.Println("\n\nQueryLoadAvg...")
	avgs, err := sg.QueryLoadAvg()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(avgs)
	//
	//
	fmt.Println("\n\nQueryProcs...")
	procs, err := sg.QueryProcs()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(procs)
	//
	//
	fmt.Println("\n\nQueryResLimit...")
	limit, err := sg.QueryResLimit()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(limit)
}

func TestSigarErr(t *testing.T) {
	sg := NewSigar()
	sg.Open()
	defer sg.Close()
	sg.IsOk = func(int) bool {
		return false
	}
	sg.QueryMem()
	sg.QuerySwap()
	sg.QueryCpu()
	sg.QueryCpus()
	sg.QueryCpuInfoes()
	sg.QueryUptime()
	sg.QueryLoadAvg()
	sg.QueryProcs()
	sg.QueryResLimit()
}
