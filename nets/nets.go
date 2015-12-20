package main

import (
	"fmt"
	"github.com/Centny/gosigar"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: nets <net interface name>")
		return
	}
	sigar := gosigar.NewSigar()
	sigar.Open()
	defer sigar.Close()
	ltx, lrx := uint64(0), uint64(0)
	for {
		nss, err := sigar.QueryNetIfStat(os.Args[1])
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
