package gosigar_test

import (
	"fmt"
	"github.com/Centny/gosigar"
	"time"
)

func ExampleNetSpeed() {
	//query net speed
	sigar := gosigar.NewSigar()
	sigar.Open()
	defer sigar.Close()
	ltx, lrx := uint64(0), uint64(0)
	for {
		nss, err := sigar.QueryNetIfStat("en0")
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		if ltx > 0 && lrx > 0 {
			fmt.Printf("T:%vKB/s\nR:%vKB/s\n\n", (nss.TxBytes-ltx)/1024, (nss.RxBytes-lrx)/1024)
		}
		ltx, lrx = nss.TxBytes, nss.RxBytes
		time.Sleep(time.Second)
	}
}

func ExampleMemory() {
	//query memory
	sigar := gosigar.NewSigar()
	sigar.Open()
	defer sigar.Close()
	for {
		mem, err := sigar.QueryMem()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println(mem.String())
		time.Sleep(time.Second)
	}
}

func ExampleCpu() {
	//query cpu
	sigar := gosigar.NewSigar()
	sigar.Open()
	defer sigar.Close()
	for {
		cpu, err := sigar.QueryCpu()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println(cpu.String())
		time.Sleep(time.Second)
	}
}
