package gosigar

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "sigar.h"
#cgo darwin CPPFLAGS: -I/usr/local/include
#cgo darwin LDFLAGS: -L/usr/local/lib -lsigar
#cgo win LDFLAGS: -LC:\LibreOffice4\sdk\lib -LC:\GOPATH\src\github.com\Centny\oogo -loogo -licppu -licppuhelper -lipurpenvhelper -lisal -lisalhelper

sigar_cpu_t gs_cpu_t_(sigar_cpu_list_t* cpus, int idx) {
	return cpus->data[idx];
}

sigar_cpu_info_t gs_cpu_info_t_(sigar_cpu_info_list_t* cpus, int idx) {
	return cpus->data[idx];
}

sigar_pid_t gs_pid_t_(sigar_proc_list_t* proc, int idx){
	return proc->data[idx];
}

void gs_proc_args_cpy(sigar_proc_args_t* args, char* dest, int idx) {
	strcpy(dest, args->data[idx]);
}

*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type Sigar struct {
	sigar *C.sigar_t
	IsOk  func(int) bool
}

func NewSigar() *Sigar {
	return &Sigar{
		IsOk: func(s int) bool {
			return s == int(C.SIGAR_OK)
		},
	}
}

func (s *Sigar) Open() {
	C.sigar_open(&s.sigar)
}

func (s *Sigar) Close() {
	C.sigar_close(s.sigar)
}

func (s *Sigar) terror(m string, status C.int) error {
	estr := C.sigar_strerror(s.sigar, status)
	buf := make([]byte, int(C.strlen(estr)))
	C.strcpy((*C.char)(unsafe.Pointer(&buf[0])), estr)
	return errors.New(fmt.Sprintf("%v error(%v)->%v", m, status, string(buf)))
}

//query memory info
func (s *Sigar) QueryMem() (*Mem, error) {
	var mem C.sigar_mem_t
	status := C.sigar_mem_get(s.sigar, &mem)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryMem", status)
	}
	return &Mem{
		Ram:         uint64(mem.ram),
		Total:       uint64(mem.total),
		Used:        uint64(mem.used),
		Free:        uint64(mem.free),
		ActualUsed:  uint64(mem.actual_used),
		ActualFree:  uint64(mem.actual_free),
		UsedPercent: float64(mem.used_percent),
		FreePercent: float64(mem.free_percent),
	}, nil
}

//query swap info
func (s *Sigar) QuerySwap() (*Swap, error) {
	var swap C.sigar_swap_t
	status := C.sigar_swap_get(s.sigar, &swap)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QuerySwap", status)
	}
	return &Swap{
		Total:   uint64(swap.total),
		Used:    uint64(swap.total),
		Free:    uint64(swap.total),
		PageIn:  uint64(swap.total),
		PageOut: uint64(swap.total),
	}, nil
}

//query all cpu info
func (s *Sigar) QueryCpu() (*Cpu, error) {
	var cpu C.sigar_cpu_t
	status := C.sigar_cpu_get(s.sigar, &cpu)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryCpu", status)
	}
	return &Cpu{
		User:    uint64(cpu.user),
		Sys:     uint64(cpu.sys),
		Nice:    uint64(cpu.nice),
		Idle:    uint64(cpu.idle),
		Wait:    uint64(cpu.wait),
		Irq:     uint64(cpu.irq),
		SoftIrq: uint64(cpu.soft_irq),
		Stolen:  uint64(cpu.stolen),
		Total:   uint64(cpu.total),
	}, nil
}

//query per cpu info
func (s *Sigar) QueryCpus() ([]*Cpu, error) {
	var cpus C.sigar_cpu_list_t
	status := C.sigar_cpu_list_get(s.sigar, &cpus)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryCpus", status)
	}
	defer C.sigar_cpu_list_destroy(s.sigar, &cpus)
	clen := int(cpus.number)
	tcpus := []*Cpu{}
	for i := 0; i < clen; i++ {
		cpu := C.gs_cpu_t_(&cpus, C.int(i))
		tcpus = append(tcpus, &Cpu{
			User:    uint64(cpu.user),
			Sys:     uint64(cpu.sys),
			Nice:    uint64(cpu.nice),
			Idle:    uint64(cpu.idle),
			Wait:    uint64(cpu.wait),
			Irq:     uint64(cpu.irq),
			SoftIrq: uint64(cpu.soft_irq),
			Stolen:  uint64(cpu.stolen),
			Total:   uint64(cpu.total),
		})
	}
	return tcpus, nil
}

//query cpu physical info
func (s *Sigar) QueryCpuInfoes() ([]*CpuInfo, error) {
	var cpus C.sigar_cpu_info_list_t
	status := C.sigar_cpu_info_list_get(s.sigar, &cpus)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryCpuInfoes", status)
	}
	defer C.sigar_cpu_info_list_destroy(s.sigar, &cpus)
	clen := int(cpus.number)
	tcpus := []*CpuInfo{}
	for i := 0; i < clen; i++ {
		cpu := C.gs_cpu_info_t_(&cpus, C.int(i))
		vendor, model := make([]byte, 128), make([]byte, 128)
		C.strcpy((*C.char)(unsafe.Pointer(&vendor[0])), &cpu.vendor[0])
		C.strcpy((*C.char)(unsafe.Pointer(&model[0])), &cpu.model[0])
		tcpus = append(tcpus, &CpuInfo{
			Vendor:         string(vendor),
			Model:          string(model),
			Mhz:            int(cpu.mhz),
			MhzMax:         int(cpu.mhz_max),
			MhzMin:         int(cpu.mhz_min),
			CacheSize:      uint64(cpu.cache_size),
			TotalSockets:   int(cpu.total_sockets),
			TotalCores:     int(cpu.total_cores),
			CoresPerSocket: int(cpu.cores_per_socket),
		})
	}
	return tcpus, nil
}

//query uptime info
func (s *Sigar) QueryUptime() (float64, error) {
	var uptime C.sigar_uptime_t
	status := C.sigar_uptime_get(s.sigar, &uptime)
	if !s.IsOk(int(status)) {
		return 0, s.terror("QueryUptime", status)
	}
	return float64(uptime.uptime), nil
}

//query loadavg
func (s *Sigar) QueryLoadAvg() ([]float64, error) {
	var avg C.sigar_loadavg_t
	status := C.sigar_loadavg_get(s.sigar, &avg)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryLoadAvg", status)
	}
	return []float64{
		float64(avg.loadavg[0]),
		float64(avg.loadavg[1]),
		float64(avg.loadavg[2]),
	}, nil
}

//query proc id list
func (s *Sigar) QueryProcs() ([]int64, error) {
	var proc C.sigar_proc_list_t
	status := C.sigar_proc_list_get(s.sigar, &proc)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryProc", status)
	}
	defer C.sigar_proc_list_destroy(s.sigar, &proc)
	clen := int(proc.number)
	tproc := []int64{}
	for i := 0; i < clen; i++ {
		tproc = append(tproc, int64(C.gs_pid_t_(&proc, C.int(i))))
	}
	return tproc, nil
}

//query system limit info
func (s *Sigar) QueryResLimit() (*ResLimit, error) {
	var limit C.sigar_resource_limit_t
	status := C.sigar_resource_limit_get(s.sigar, &limit)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryResLimit", status)
	}
	return &ResLimit{
		CpuCur:           uint64(limit.cpu_cur),
		CpuMax:           uint64(limit.cpu_max),
		FileSizeCur:      uint64(limit.file_size_cur),
		FileSizeMax:      uint64(limit.file_size_max),
		PipeSizeCur:      uint64(limit.pipe_size_cur),
		PipeSizeMax:      uint64(limit.pipe_size_max),
		DataCur:          uint64(limit.data_cur),
		DataMax:          uint64(limit.data_max),
		StackCur:         uint64(limit.stack_cur),
		StackMax:         uint64(limit.stack_max),
		CoreCur:          uint64(limit.core_cur),
		CoreMax:          uint64(limit.core_max),
		MemoryCur:        uint64(limit.memory_cur),
		MemoryMax:        uint64(limit.memory_max),
		ProcessesCur:     uint64(limit.processes_cur),
		ProcessesMax:     uint64(limit.processes_max),
		OpenFilesCur:     uint64(limit.open_files_cur),
		OpenFilesMax:     uint64(limit.open_files_max),
		VirtualMemoryCur: uint64(limit.virtual_memory_cur),
		VirtualMemoryMax: uint64(limit.virtual_memory_max),
	}, nil
}

//query proc stat
func (s *Sigar) QueryProcStat() (*ProcStat, error) {
	var stat C.sigar_proc_stat_t
	status := C.sigar_proc_stat_get(s.sigar, &stat)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryProcStat", status)
	}
	return &ProcStat{
		Total:    uint64(stat.total),
		Sleeping: uint64(stat.sleeping),
		Running:  uint64(stat.running),
		Zombie:   uint64(stat.zombie),
		Stopped:  uint64(stat.stopped),
		Idle:     uint64(stat.idle),
		Threads:  uint64(stat.threads),
	}, nil
}

//query proc memory
func (s *Sigar) QueryProcMem(pid int64) (*ProcMem, error) {
	var mem C.sigar_proc_mem_t
	status := C.sigar_proc_mem_get(s.sigar, C.sigar_pid_t(pid), &mem)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryProcMem", status)
	}
	return &ProcMem{
		Size:        uint64(mem.size),
		Resident:    uint64(mem.resident),
		Share:       uint64(mem.share),
		MinorFaults: uint64(mem.minor_faults),
		MajorFaults: uint64(mem.major_faults),
		PageFaults:  uint64(mem.page_faults),
	}, nil
}

//query proc disk io
func (s *Sigar) QueryProcDiskIO(pid int64) (*ProcDiskIO, error) {
	var disk C.sigar_proc_disk_io_t
	status := C.sigar_proc_disk_io_get(s.sigar, C.sigar_pid_t(pid), &disk)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryProcDiskIO", status)
	}
	return &ProcDiskIO{
		BytesRead:    uint64(disk.bytes_read),
		BytesWritten: uint64(disk.bytes_written),
		BytesTotal:   uint64(disk.bytes_total),
	}, nil
}

//query proc cumulative disk io
func (s *Sigar) QueryProcCumulativeDiskIO(pid int64) (*ProcCumulativeDiskIO, error) {
	var disk C.sigar_proc_cumulative_disk_io_t
	status := C.sigar_proc_cumulative_disk_io_get(s.sigar, C.sigar_pid_t(pid), &disk)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryProcCumulativeDiskIO", status)
	}
	return &ProcCumulativeDiskIO{
		BytesRead:    uint64(disk.bytes_read),
		BytesWritten: uint64(disk.bytes_written),
		BytesTotal:   uint64(disk.bytes_total),
	}, nil
}

//query dump cache
func (s *Sigar) QueryDumpCache() (uint64, error) {
	var dump C.sigar_dump_pid_cache_t
	status := C.sigar_dump_pid_cache_get(s.sigar, &dump)
	if !s.IsOk(int(status)) {
		return 0, s.terror("QueryDumpCache", status)
	}
	return uint64(dump.dummy), nil
}

//query proc cred
func (s *Sigar) QueryProcCred(pid int64) (*ProcCred, error) {
	var cred C.sigar_proc_cred_t
	status := C.sigar_proc_cred_get(s.sigar, C.sigar_pid_t(pid), &cred)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryProcCred", status)
	}
	return &ProcCred{
		Uid:  uint64(cred.uid),
		Gid:  uint64(cred.gid),
		EUid: uint64(cred.euid),
		EGid: uint64(cred.egid),
	}, nil
}

//query proc cred name
func (s *Sigar) QueryProcCredName(pid int64) (*ProcCredName, error) {
	var cred C.sigar_proc_cred_name_t
	status := C.sigar_proc_cred_name_get(s.sigar, C.sigar_pid_t(pid), &cred)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryProcCredName", status)
	}
	user, group := make([]byte, 512), make([]byte, 512)
	C.strcpy((*C.char)(unsafe.Pointer(&user[0])), &cred.user[0])
	C.strcpy((*C.char)(unsafe.Pointer(&group[0])), &cred.group[0])
	return &ProcCredName{
		User:  string(user),
		Group: string(group),
	}, nil
}

//query proc time
func (s *Sigar) QueryProcTime(pid int64) (*ProcTime, error) {
	var pt C.sigar_proc_time_t
	status := C.sigar_proc_time_get(s.sigar, C.sigar_pid_t(pid), &pt)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryProcTime", status)
	}
	return &ProcTime{
		StartTime: uint64(pt.start_time),
		User:      uint64(pt.user),
		Sys:       uint64(pt.sys),
		Total:     uint64(pt.total),
	}, nil
}

//query proc cpu
func (s *Sigar) QueryProcCPU(pid int64) (*ProcCPU, error) {
	var cpu C.sigar_proc_cpu_t
	status := C.sigar_proc_cpu_get(s.sigar, C.sigar_pid_t(pid), &cpu)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryProcCPU", status)
	}
	return &ProcCPU{
		StartTime: uint64(cpu.start_time),
		User:      uint64(cpu.user),
		Sys:       uint64(cpu.sys),
		Total:     uint64(cpu.total),
		LastTime:  uint64(cpu.last_time),
		Percent:   float64(cpu.percent),
	}, nil
}

//query proc state
func (s *Sigar) QueryProcState(pid int64) (*ProcState, error) {
	var state C.sigar_proc_state_t
	status := C.sigar_proc_state_get(s.sigar, C.sigar_pid_t(pid), &state)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryProcCPU", status)
	}
	name := make([]byte, 128)
	C.strcpy((*C.char)(unsafe.Pointer(&name[0])), &state.name[0])
	return &ProcState{
		Name:      string(name),
		State:     byte(state.state),
		Ppid:      uint64(state.ppid),
		Tty:       int(state.tty),
		Priority:  int(state.priority),
		Nice:      int(state.nice),
		Processor: int(state.processor),
		Threads:   uint64(state.threads),
	}, nil
}

//query proc args
func (s *Sigar) QueryProcArgs(pid int64) ([]string, error) {
	var args C.sigar_proc_args_t
	status := C.sigar_proc_args_get(s.sigar, C.sigar_pid_t(pid), &args)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryProcArgs", status)
	}
	defer C.sigar_proc_args_destroy(s.sigar, &args)
	clen := int(args.number)
	targs := []string{}
	for i := 0; i < clen; i++ {
		arg := make([]byte, 1024)
		C.gs_proc_args_cpy(&args, (*C.char)(unsafe.Pointer(&arg[0])), C.int(i))
		arg_ := string(arg)
		targs = append(targs, arg_)
	}
	return targs, nil
}

//query proc fd
func (s *Sigar) QueryProcFD(pid int64) (uint64, error) {
	var fd C.sigar_proc_fd_t
	status := C.sigar_proc_fd_get(s.sigar, C.sigar_pid_t(pid), &fd)
	if !s.IsOk(int(status)) {
		return 0, s.terror("QueryProcFD", status)
	}
	return uint64(fd.total), nil
}

//query exe
func (s *Sigar) QueryProcExe(pid int64) (*ProcExe, error) {
	var exe C.sigar_proc_exe_t
	status := C.sigar_proc_exe_get(s.sigar, C.sigar_pid_t(pid), &exe)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryExe", status)
	}
	name, cwd, root := make([]byte, 4097), make([]byte, 4097), make([]byte, 4097)
	C.strcpy((*C.char)(unsafe.Pointer(&name[0])), &exe.name[0])
	C.strcpy((*C.char)(unsafe.Pointer(&cwd[0])), &exe.cwd[0])
	C.strcpy((*C.char)(unsafe.Pointer(&root[0])), &exe.root[0])
	return &ProcExe{
		Name: string(name),
		Cwd:  string(cwd),
		Root: string(root),
	}, nil
}

//query thread cpu.
func (s *Sigar) QueryThreadCPU(id int64) (*ThreadCPU, error) {
	var cpu C.sigar_thread_cpu_t
	status := C.sigar_thread_cpu_get(s.sigar, C.sigar_uint64_t(id), &cpu)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryThreadCPU", status)
	}
	return &ThreadCPU{
		User:  uint64(cpu.user),
		Sys:   uint64(cpu.sys),
		Total: uint64(cpu.total),
	}, nil
}
