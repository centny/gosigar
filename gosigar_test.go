package gosigar

import (
	"fmt"
	"os"
	"testing"
)

func TestSigar(t *testing.T) {
	fmt.Println(Version())
	//
	//
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
	//
	//
	fmt.Println("\n\nQueryProcStat...")
	stat, err := sg.QueryProcStat()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(stat)
	//
	pid := int64(os.Getpid())
	//
	//
	fmt.Println("\n\nQueryProcMem...")
	pmem, err := sg.QueryProcMem(pid)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(pmem)
	//
	//
	// fmt.Println("\n\nQueryProcDiskIO...")
	// pdisk, err := sg.QueryProcDiskIO(pid)
	// if err != nil {
	// 	t.Error(err.Error())
	// 	return
	// }
	// fmt.Println(pdisk)
	//
	//
	// fmt.Println("\n\nQueryProcCumulativeDiskIO...")
	// pcdisk, err := sg.QueryProcCumulativeDiskIO(pid)
	// if err != nil {
	// 	t.Error(err.Error())
	// 	return
	// }
	// fmt.Println(pcdisk)
	//
	//
	fmt.Println("\n\nQueryDumpCache...")
	dump, err := sg.QueryDumpCache()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(dump)
	//
	//
	fmt.Println("\n\nQueryProcCred...")
	pcred, err := sg.QueryProcCred(pid)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(pcred)
	//
	//
	fmt.Println("\n\nQueryProcCredName..")
	pcredn, err := sg.QueryProcCredName(pid)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(pcredn)
	//
	//
	fmt.Println("\n\nQueryProcTime..")
	pt, err := sg.QueryProcTime(pid)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(pt)
	//
	//
	fmt.Println("\n\nQueryProcCPU..")
	pcpu, err := sg.QueryProcCPU(pid)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(pcpu)
	//
	//
	fmt.Println("\n\nQueryProcState..")
	pst, err := sg.QueryProcState(pid)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(pst)
	//
	//
	fmt.Println("\n\nQueryProcArgs..")
	pargs, err := sg.QueryProcArgs(pid)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(pargs)
	//
	//
	// fmt.Println("\n\nQueryProcFd..")
	// pfd, err := sg.QueryProcFD(pid)
	// if err != nil {
	// 	t.Error(err.Error())
	// 	return
	// }
	// fmt.Println(pfd)
	//
	//
	fmt.Println("\n\nQueryProcExe..")
	pex, err := sg.QueryProcExe(pid)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(pex)
	//
	//
	fmt.Println("\n\nQueryThreadCPU..")
	tcpu, err := sg.QueryThreadCPU(0)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(tcpu)
	//
	//
	fmt.Println("\n\nQueryFileSystem..")
	fss, err := sg.QueryFileSystems()
	if err != nil {
		t.Error(err.Error())
		return
	}
	for _, fs := range fss {
		fmt.Println()
		fmt.Println(sg.PingFS(fs))
		fmt.Println("-->", fs.String())
		fmt.Println(sg.QueryDiskUsage(fs.DevName))
		fmt.Println(sg.QueryFileSystemUsage(fs.DirName))
		fmt.Println()
	}
	fmt.Println(fss)
	//
	//
	fmt.Println("\n\nQueryNetInfo..")
	nis, err := sg.QueryNetInfo()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(nis)
	//
	//
	fmt.Println("\n\nQueryNetRoutes..")
	nrs, err := sg.QueryNetRoutes()
	if err != nil {
		t.Error(err.Error())
		return
	}
	// for _, nr := range nrs {
	// 	fmt.Println(nr.String())
	// 	fmt.Println("\n")
	// }
	fmt.Println(nrs[0].String())
	fmt.Println(len(nrs))
	//
	//
	fmt.Println("\n\nQueryNetConfig..")
	ncfg, err := sg.QueryNetConfig("")
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(ncfg)
	sg.QueryNetConfig("en0")
	//
	//
	fmt.Println("\n\nQueryNetNames..")
	nns, err := sg.QueryNetNames()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(nns)
	//
	//
	fmt.Println("\n\nQueryNetStat..")
	for _, nn := range nns {
		//
		//
		nst, err := sg.QueryNetIfStat(nn)
		if err != nil {
			t.Error(err.Error())
			return
		}
		fmt.Println(nst)
	}
	//
	//
	ncf := []int{
		SIGAR_NETCONN_TCP,
		SIGAR_NETCONN_UDP,
		SIGAR_NETCONN_RAW,
		SIGAR_NETCONN_UNIX,
	}
	fmt.Println("\n\nQueryNetConnections..")
	for _, f := range ncf {
		ncs, err := sg.QueryNetConnections(f)
		if err != nil {
			t.Error(err.Error())
			return
		}
		fmt.Println(len(ncs))
	}
	//
	//
	fmt.Println("\n\nQueryNetStat..")
	nss, err := sg.QueryNetStat(SIGAR_NETCONN_TCP)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(nss)
	//
	//
	fmt.Println("\n\nQueryTCP..")
	tcp, err := sg.QueryTCP()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(tcp)
	//
	//
	fmt.Println("\n\nQueryWhoes..")
	ws, err := sg.QueryWhoes()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(ws)
	//
	//
	fmt.Println("\n\nQuerySysInfo..")
	sys, err := sg.QuerySysInfo()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(sys)
}

func TestSigarErr(t *testing.T) {
	sg := NewSigar()
	sg.Open()
	defer sg.Close()
	sg.IsOk = func(int) bool {
		return false
	}
	pid := int64(os.Getpid())
	sg.QueryMem()
	sg.QuerySwap()
	sg.QueryCpu()
	sg.QueryCpus()
	sg.QueryCpuInfoes()
	sg.QueryUptime()
	sg.QueryLoadAvg()
	sg.QueryProcs()
	sg.QueryResLimit()
	sg.QueryProcStat()
	sg.QueryProcMem(pid)
	sg.QueryDumpCache()
	sg.QueryProcCred(pid)
	sg.QueryProcCredName(pid)
	sg.QueryProcTime(pid)
	sg.QueryProcCPU(pid)
	sg.QueryProcState(pid)
	sg.QueryProcArgs(pid)
	sg.QueryProcFD(pid)
	sg.QueryProcExe(pid)
	sg.QueryProcDiskIO(pid)
	sg.QueryProcCumulativeDiskIO(pid)
	sg.QueryThreadCPU(0)
	sg.QueryFileSystems()
	sg.QueryDiskUsage("/")
	sg.QueryFileSystemUsage("/")
	sg.QueryNetInfo()
	sg.QueryNetRoutes()
	sg.QueryNetConfig("")
	sg.QueryNetIfStat("en0")
	sg.QueryNetNames()
	sg.QueryNetStat(0)
	sg.QueryTCP()
	sg.QueryWhoes()
	sg.QuerySysInfo()
}
