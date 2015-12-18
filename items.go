package gosigar

import (
	"fmt"
)

type Mem struct {
	Ram,
	Total,
	Used,
	Free,
	ActualUsed,
	ActualFree uint64
	UsedPercent,
	FreePercent float64
}

func (m *Mem) String() string {
	return fmt.Sprintf("ram(%v),total(%v),used(%v),free(%v),actual_used(%v),actual_free(%v),used_percent(%v),free_percent(%v)",
		m.Ram, m.Total, m.Used, m.Free, m.ActualUsed, m.ActualFree, m.UsedPercent, m.FreePercent)
}

type Swap struct {
	Total,
	Used,
	Free,
	PageIn,
	PageOut uint64
}

func (s *Swap) String() string {
	return fmt.Sprintf("total(%v),used(%v),free(%v),page_in(%v),page_out(%v)", s.Total, s.Used, s.Free, s.PageIn, s.PageOut)
}

type Cpu struct {
	User,
	Sys,
	Nice,
	Idle,
	Wait,
	Irq,
	SoftIrq,
	Stolen,
	Total uint64
}

func (c *Cpu) String() string {
	return fmt.Sprintf("user(%v),sys(%v),nice(%v),idle(%v),wait(%v),irq(%v),soft_irq(%v),stolen(%v),total(%v)",
		c.User, c.Sys, c.Nice, c.Idle, c.Wait, c.Irq, c.SoftIrq, c.Stolen, c.Total)
}

type CpuInfo struct {
	Vendor         string
	Model          string
	Mhz            int
	MhzMax         int
	MhzMin         int
	CacheSize      uint64
	TotalSockets   int
	TotalCores     int
	CoresPerSocket int
}

func (c *CpuInfo) String() string {
	return fmt.Sprintf("vendor(%v),model(%v),mhz(%v),mhz_max(%v),mhz_min(%v),cache_size(%v),total_sockets(%v),total_cores(%v),cores_per_socket(%v)",
		c.Vendor, c.Model, c.Mhz, c.MhzMax, c.MhzMin, c.CacheSize, c.TotalSockets, c.TotalCores, c.CoresPerSocket)
}

type ResLimit struct {
	/* RLIMIT_CPU */
	CpuCur, CpuMax uint64
	/* RLIMIT_FSIZE */
	FileSizeCur, FileSizeMax uint64
	/* PIPE_BUF */
	PipeSizeCur, PipeSizeMax uint64
	/* RLIMIT_DATA */
	DataCur, DataMax uint64
	/* RLIMIT_STACK */
	StackCur, StackMax uint64
	/* RLIMIT_CORE */
	CoreCur, CoreMax uint64
	/* RLIMIT_RSS */
	MemoryCur, MemoryMax uint64
	/* RLIMIT_NPROC */
	ProcessesCur, ProcessesMax uint64
	/* RLIMIT_NOFILE */
	OpenFilesCur, OpenFilesMax uint64
	/* RLIMIT_AS */
	VirtualMemoryCur, VirtualMemoryMax uint64
}

func (r *ResLimit) String() string {
	return fmt.Sprintf("cpu_cur(%v),cpu_max(%v),file_size_cur(%v),file_size_max(%v),pipe_size_cur(%v),pipe_size_max(%v),data_cur(%v),data_max(%v),stack_cur(%v),stack_max(%v),core_cur(%v),core_max(%v),memory_cur(%v),memory_max(%v),processes_cur(%v),processes_max(%v),open_files_cur(%v),open_files_max(%v),virtual_memory_cur(%v),virtual_memory_max(%v)",
		r.CpuCur, r.CpuMax, r.FileSizeCur, r.FileSizeMax, r.PipeSizeCur, r.PipeSizeMax, r.DataCur, r.DataMax, r.StackCur, r.StackMax, r.CoreCur, r.CoreMax, r.MemoryCur, r.MemoryMax, r.ProcessesCur, r.ProcessesMax, r.OpenFilesCur, r.OpenFilesMax, r.VirtualMemoryCur, r.VirtualMemoryMax)
}
