package gosigar

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "sigar.h"
#cgo darwin CPPFLAGS: -I/usr/local/include
#cgo darwin LDFLAGS: -L/usr/local/lib -lsigar
#cgo linux CPPFLAGS: -I/usr/local/include
#cgo linux LDFLAGS: -L/usr/local/lib -lsigar
#cgo win LDFLAGS: -LC:\LibreOffice4\sdk\lib -LC:\GOPATH\src\github.com\Centny\oogo -loogo -licppu -licppuhelper -lipurpenvhelper -lisal -lisalhelper

char* gs_char_t_(char **data, int idx) {
	return data[idx];
}

sigar_cpu_t gs_cpu_t_(sigar_cpu_list_t* cpus, int idx) {
	return cpus->data[idx];
}

sigar_cpu_info_t gs_cpu_info_t_(sigar_cpu_info_list_t* cpus, int idx) {
	return cpus->data[idx];
}

sigar_pid_t gs_pid_t_(sigar_proc_list_t* proc, int idx){
	return proc->data[idx];
}

sigar_file_system_t gs_fs_t_(sigar_file_system_list_t* fs, int idx) {
	return fs->data[idx];
}

void gs_nr_addr_set_t_(sigar_net_address_t* addr, int* family,
		sigar_uint32_t* in, sigar_uint32_t* in6, void* mac) {
	*family = addr->family;
	*in = addr->addr.in;
	memcpy(in6, addr->addr.in6, 4);
	memcpy(mac, addr->addr.mac, 8);
}

sigar_net_route_t gs_nr_t_(sigar_net_route_list_t* nrs, int idx) {
	return nrs->data[idx];
}

sigar_net_connection_t gs_net_con_t_(sigar_net_connection_list_t* cons, int idx) {
	return cons->data[idx];
}

*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

func cstring_(cs *C.char) string {
	clen := C.strlen(cs)
	if clen < 1 {
		return ""
	}
	buf := make([]byte, clen+1)
	C.strcpy((*C.char)(unsafe.Pointer(&buf[0])), cs)
	return string(buf[:clen])
}

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
		tcpus = append(tcpus, &CpuInfo{
			Vendor:         cstring_(&cpu.vendor[0]),
			Model:          cstring_(&cpu.model[0]),
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
	return &ProcCredName{
		User:  cstring_(&cred.user[0]),
		Group: cstring_(&cred.group[0]),
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
	return &ProcState{
		Name:      cstring_(&state.name[0]),
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
		targs = append(targs, cstring_(C.gs_char_t_(args.data, C.int(i))))
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
	return &ProcExe{
		Name: cstring_(&exe.name[0]),
		Cwd:  cstring_(&exe.cwd[0]),
		Root: cstring_(&exe.root[0]),
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

//query file system
func (s *Sigar) QueryFileSystems() ([]*FileSystem, error) {
	var fss C.sigar_file_system_list_t
	status := C.sigar_file_system_list_get(s.sigar, &fss)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryFileSystems", status)
	}
	defer C.sigar_file_system_list_destroy(s.sigar, &fss)
	clen := int(fss.number)
	tfss := []*FileSystem{}
	for i := 0; i < clen; i++ {
		fs := C.gs_fs_t_(&fss, C.int(i))
		ttyp := SIGAR_FSTYPE_UNKNOWN
		switch fs._type {
		case C.SIGAR_FSTYPE_UNKNOWN:
			ttyp = SIGAR_FSTYPE_UNKNOWN
		case C.SIGAR_FSTYPE_NONE:
			ttyp = SIGAR_FSTYPE_NONE
		case C.SIGAR_FSTYPE_LOCAL_DISK:
			ttyp = SIGAR_FSTYPE_LOCAL_DISK
		case C.SIGAR_FSTYPE_NETWORK:
			ttyp = SIGAR_FSTYPE_NETWORK
		case C.SIGAR_FSTYPE_RAM_DISK:
			ttyp = SIGAR_FSTYPE_RAM_DISK
		case C.SIGAR_FSTYPE_CDROM:
			ttyp = SIGAR_FSTYPE_CDROM
		case C.SIGAR_FSTYPE_SWAP:
			ttyp = SIGAR_FSTYPE_SWAP
		case C.SIGAR_FSTYPE_MAX:
			ttyp = SIGAR_FSTYPE_MAX
		}
		tfss = append(tfss, &FileSystem{
			DirName:     cstring_(&fs.dir_name[0]),
			DevName:     cstring_(&fs.dev_name[0]),
			TypeName:    cstring_(&fs.type_name[0]),
			SysTypeName: cstring_(&fs.sys_type_name[0]),
			Options:     cstring_(&fs.options[0]),
			Type:        ttyp,
			Flags:       uint64(fs.flags),
		})
	}
	return tfss, nil
}

//query disk usage
func (s *Sigar) QueryDiskUsage(name string) (*DiskUsage, error) {
	name_ := C.CString(name)
	defer C.free(unsafe.Pointer(name_))
	var disk C.sigar_disk_usage_t
	status := C.sigar_disk_usage_get(s.sigar, name_, &disk)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryDiskUsage", status)
	}
	return &DiskUsage{
		Reads:       uint64(disk.reads),
		Writes:      uint64(disk.writes),
		WriteBytes:  uint64(disk.write_bytes),
		ReadBytes:   uint64(disk.read_bytes),
		RTime:       uint64(disk.rtime),
		WTime:       uint64(disk.wtime),
		QTime:       uint64(disk.qtime),
		Time:        uint64(disk.time),
		Snaptime:    uint64(disk.snaptime),
		ServiceTime: float64(disk.service_time),
		Queue:       float64(disk.queue),
	}, nil
}

//query file system usage
func (s *Sigar) QueryFileSystemUsage(dirname string) (*FileSystemUsage, error) {
	dirname_ := C.CString(dirname)
	defer C.free(unsafe.Pointer(dirname_))
	var fsu C.sigar_file_system_usage_t
	status := C.sigar_file_system_usage_get(s.sigar, dirname_, &fsu)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryFileSystemUsage", status)
	}
	disk := fsu.disk
	return &FileSystemUsage{
		Disk: DiskUsage{
			Reads:       uint64(disk.reads),
			Writes:      uint64(disk.writes),
			WriteBytes:  uint64(disk.write_bytes),
			ReadBytes:   uint64(disk.read_bytes),
			RTime:       uint64(disk.rtime),
			WTime:       uint64(disk.wtime),
			QTime:       uint64(disk.qtime),
			Time:        uint64(disk.time),
			Snaptime:    uint64(disk.snaptime),
			ServiceTime: float64(disk.service_time),
			Queue:       float64(disk.queue),
		},
		UsePercent: float64(fsu.use_percent),
		Total:      uint64(fsu.total),
		Free:       uint64(fsu.free),
		Used:       uint64(fsu.used),
		Avail:      uint64(fsu.avail),
		Files:      uint64(fsu.files),
		FreeFiles:  uint64(fsu.free_files),
	}, nil
}

//ping the file system
func (s *Sigar) PingFS(fs *FileSystem) int {
	var cfs C.sigar_file_system_t
	dir, dev := []byte(fs.DirName), []byte(fs.DevName)
	tn, sys, opt := []byte(fs.TypeName), []byte(fs.SysTypeName), []byte(fs.Options)
	C.strcpy(&cfs.dir_name[0], (*C.char)(unsafe.Pointer(&dir[0])))
	C.strcpy(&cfs.dev_name[0], (*C.char)(unsafe.Pointer(&dev[0])))
	C.strcpy(&cfs.type_name[0], (*C.char)(unsafe.Pointer(&tn[0])))
	C.strcpy(&cfs.sys_type_name[0], (*C.char)(unsafe.Pointer(&sys[0])))
	C.strcpy(&cfs.options[0], (*C.char)(unsafe.Pointer(&opt[0])))
	switch fs.Type {
	case SIGAR_FSTYPE_NONE:
		cfs._type = C.SIGAR_FSTYPE_NONE
	case SIGAR_FSTYPE_LOCAL_DISK:
		cfs._type = C.SIGAR_FSTYPE_LOCAL_DISK
	case SIGAR_FSTYPE_NETWORK:
		cfs._type = C.SIGAR_FSTYPE_NETWORK
	case SIGAR_FSTYPE_RAM_DISK:
		cfs._type = C.SIGAR_FSTYPE_RAM_DISK
	case SIGAR_FSTYPE_CDROM:
		cfs._type = C.SIGAR_FSTYPE_RAM_DISK
	case SIGAR_FSTYPE_SWAP:
		cfs._type = C.SIGAR_FSTYPE_SWAP
	case SIGAR_FSTYPE_MAX:
		cfs._type = C.SIGAR_FSTYPE_MAX
	default:
		cfs._type = C.SIGAR_FSTYPE_UNKNOWN
	}
	cfs.flags = C.ulong(fs.Flags)
	return int(C.sigar_file_system_ping(s.sigar, &cfs))
}

//query net info
func (s *Sigar) QueryNetInfo() (*NetInfo, error) {
	var net C.sigar_net_info_t
	status := C.sigar_net_info_get(s.sigar, &net)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryDiskUsage", status)
	}
	return &NetInfo{
		DefaultGateway:          cstring_(&net.default_gateway[0]),
		DefaultGatewayInterface: cstring_(&net.default_gateway_interface[0]),
		HostName:                cstring_(&net.host_name[0]),
		DomainName:              cstring_(&net.domain_name[0]),
		PrimaryDns:              cstring_(&net.primary_dns[0]),
		SecondaryDns:            cstring_(&net.secondary_dns[0]),
	}, nil
}

func (s *Sigar) net_addr(addr *C.sigar_net_address_t) NetAddr {
	dest := NetAddr{}
	C.gs_nr_addr_set_t_(addr,
		(*C.int)(unsafe.Pointer(&dest.Family)),
		(*C.sigar_uint32_t)(unsafe.Pointer(&dest.In)),
		(*C.sigar_uint32_t)(unsafe.Pointer(&dest.In6[0])),
		unsafe.Pointer(&dest.Mac[0]),
	)
	switch dest.Family {
	case C.SIGAR_AF_UNSPEC:
		dest.Family = SIGAR_AF_UNSPEC
	case C.SIGAR_AF_INET:
		dest.Family = SIGAR_AF_INET
	case C.SIGAR_AF_INET6:
		dest.Family = SIGAR_AF_INET6
	case C.SIGAR_AF_LINK:
		dest.Family = SIGAR_AF_LINK
	}
	return dest
}

//query net route
func (s *Sigar) QueryNetRoutes() ([]*NetRoute, error) {
	var nrs C.sigar_net_route_list_t
	status := C.sigar_net_route_list_get(s.sigar, &nrs)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryFileSystems", status)
	}
	defer C.sigar_net_route_list_destroy(s.sigar, &nrs)
	clen := int(nrs.number)
	tnrs := []*NetRoute{}
	for i := 0; i < clen; i++ {
		nr := C.gs_nr_t_(&nrs, C.int(i))
		tnrs = append(tnrs, &NetRoute{
			Destination: s.net_addr(&nr.destination),
			Gateway:     s.net_addr(&nr.gateway),
			Mask:        s.net_addr(&nr.mask),
			Flags:       uint64(nr.flags),
			Refcnt:      uint64(nr.refcnt),
			Use:         uint64(nr.use),
			Metric:      uint64(nr.metric),
			Mtu:         uint64(nr.mtu),
			Window:      uint64(nr.window),
			Irtt:        uint64(nr.irtt),
			IfName:      cstring_(&nr.ifname[0]),
		})
	}
	return tnrs, nil
}

//query net config
func (s *Sigar) QueryNetConfig(name string) (*NetConfig, error) {
	var net C.sigar_net_interface_config_t
	var status C.int
	if len(name) < 1 {
		status = C.sigar_net_interface_config_primary_get(s.sigar, &net)
	} else {
		name_ := C.CString(name)
		defer C.free(unsafe.Pointer(name_))
		status = C.sigar_net_interface_config_get(s.sigar, name_, &net)
	}
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryNetConfig", status)
	}
	return &NetConfig{
		Name:          cstring_(&net.name[0]),
		Type:          cstring_(&net._type[0]),
		Description:   cstring_(&net.description[0]),
		Hwaddr:        s.net_addr(&net.hwaddr),
		Address:       s.net_addr(&net.address),
		Destination:   s.net_addr(&net.destination),
		Broadcast:     s.net_addr(&net.broadcast),
		Netmask:       s.net_addr(&net.netmask),
		Address6:      s.net_addr(&net.address6),
		Prefix6Length: int(net.prefix6_length),
		Scope6:        int(net.scope6),
		Flags:         uint64(net.flags),
		Mtu:           uint64(net.mtu),
		Metric:        uint64(net.metric),
		TxQueueLen:    int(net.tx_queue_len),
	}, nil
}

//query net stat
func (s *Sigar) QueryNetStat(name string) (*NetStat, error) {
	var net C.sigar_net_interface_stat_t
	name_ := C.CString(name)
	defer C.free(unsafe.Pointer(name_))
	status := C.sigar_net_interface_stat_get(s.sigar, name_, &net)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryNetStat", status)
	}
	return &NetStat{
		RxPackets:    uint64(net.rx_packets),
		RxBytes:      uint64(net.rx_bytes),
		RxErrors:     uint64(net.rx_errors),
		RxDropped:    uint64(net.rx_dropped),
		RxOverruns:   uint64(net.rx_overruns),
		RxFrame:      uint64(net.rx_frame),
		TxPackets:    uint64(net.tx_packets),
		TxBytes:      uint64(net.tx_bytes),
		TxErrors:     uint64(net.tx_errors),
		TxDropped:    uint64(net.tx_dropped),
		TxOverruns:   uint64(net.tx_overruns),
		TxCollisions: uint64(net.tx_collisions),
		TxCarrier:    uint64(net.tx_carrier),
		Speed:        uint64(net.speed),
	}, nil
}

//query net name
func (s *Sigar) QueryNetNames() ([]string, error) {
	var names C.sigar_net_interface_list_t
	status := C.sigar_net_interface_list_get(s.sigar, &names)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryNetNames", status)
	}
	defer C.sigar_net_interface_list_destroy(s.sigar, &names)
	clen := int(names.number)
	tnames := []string{}
	for i := 0; i < clen; i++ {
		tnames = append(tnames, cstring_(C.gs_char_t_(names.data, C.int(i))))
	}
	return tnames, nil
}

func (s *Sigar) QueryNetConnections(flags int) ([]*NetConnection, error) {
	var flags_ C.int = C.int(flags)
	var cons C.sigar_net_connection_list_t
	status := C.sigar_net_connection_list_get(s.sigar, &cons, flags_)
	if !s.IsOk(int(status)) {
		return nil, s.terror("QueryNetConnections", status)
	}
	defer C.sigar_net_connection_list_destroy(s.sigar, &cons)
	clen := int(cons.number)
	tcons := []*NetConnection{}
	for i := 0; i < clen; i++ {
		con := C.gs_net_con_t_(&cons, C.int(i))
		tcons = append(tcons, &NetConnection{
			LocalPort:     uint64(con.local_port),
			LocalAddress:  s.net_addr(&con.local_address),
			RemotePort:    uint64(con.remote_port),
			RemoteAddress: s.net_addr(&con.local_address),
			Uid:           uint64(con.uid),
			Inode:         uint64(con.inode),
			Type:          int(con._type),
			State:         int(con.state),
			SendQueue:     uint64(con.send_queue),
			ReceiveQueue:  uint64(con.receive_queue),
		})
	}
	return tcons, nil
}
