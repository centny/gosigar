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

func (s *Sigar) QueryUptime() (float64, error) {
	var uptime C.sigar_uptime_t
	status := C.sigar_uptime_get(s.sigar, &uptime)
	if !s.IsOk(int(status)) {
		return 0, s.terror("QueryUptime", status)
	}
	return float64(uptime.uptime), nil
}

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
