package gosigar

import (
	"fmt"
)

//memory info
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

//swap info
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

//cpu info
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

//cpu physical info
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

//resource limit
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

//proc stat
type ProcStat struct {
	Total    uint64
	Sleeping uint64
	Running  uint64
	Zombie   uint64
	Stopped  uint64
	Idle     uint64
	Threads  uint64
}

func (p *ProcStat) String() string {
	return fmt.Sprintf("total(%v),sleeping(%v),running(%v),zombie(%v),stopped(%v),idle(%v),threads(%v)",
		p.Total, p.Sleeping, p.Running, p.Zombie, p.Stopped, p.Idle, p.Threads)
}

//proc memory
type ProcMem struct {
	Size,
	Resident,
	Share,
	MinorFaults,
	MajorFaults,
	PageFaults uint64
}

func (p *ProcMem) String() string {
	return fmt.Sprintf("size(%v),resident(%v),share(%v),minor_faults(%v),major_faults(%v),page_faults(%v)",
		p.Size, p.Resident, p.Share, p.MinorFaults, p.MajorFaults, p.PageFaults)
}

//proc disk io
type ProcDiskIO struct {
	BytesRead,
	BytesWritten,
	BytesTotal uint64
}

func (p *ProcDiskIO) String() string {
	return fmt.Sprintf("bytes_read(%v),bytes_written(%v),bytes_total(%v)",
		p.BytesRead, p.BytesWritten, p.BytesTotal)
}

//cached proc disk io
type CachedProcDiskIO struct {
	BytesRead,
	BytesWritten,
	BytesTotal,
	LastTime,
	BytesReadDiff,
	BytesWrittenDiff,
	BytesTotalDiff uint64
}

func (c *CachedProcDiskIO) String() string {
	return fmt.Sprintf("bytes_read(%v),bytes_written(%v),bytes_total(%v),last_time(%v),bytes_read_diff(%v),bytes_written_diff(%v),bytes_total_diff(%v)",
		c.BytesRead, c.BytesWritten, c.BytesTotal, c.LastTime, c.BytesReadDiff, c.BytesWrittenDiff, c.BytesTotalDiff)
}

//proc cumulative disk io
type ProcCumulativeDiskIO struct {
	BytesRead,
	BytesWritten,
	BytesTotal uint64
}

func (p *ProcCumulativeDiskIO) String() string {
	return fmt.Sprintf("bytes_read(%v),bytes_written(%v),bytes_total(%v)",
		p.BytesRead, p.BytesWritten, p.BytesTotal)
}

//proc cred
type ProcCred struct {
	Uid  uint64
	Gid  uint64
	EUid uint64
	EGid uint64
}

func (p *ProcCred) String() string {
	return fmt.Sprintf("uid(%v),gid(%v),euid(%v),egid(%v)",
		p.Uid, p.Gid, p.EUid, p.EGid)
}

//proc cred name
type ProcCredName struct {
	User  string
	Group string
}

func (p *ProcCredName) String() string {
	return fmt.Sprintf("user(%v),group(%v)",
		p.User, p.Group)
}

//proc time
type ProcTime struct {
	StartTime,
	User,
	Sys,
	Total uint64
}

func (p *ProcTime) String() string {
	return fmt.Sprintf("start_time(%v),user(%v),sys(%v),total(%v)",
		p.StartTime, p.User, p.Sys, p.Total)
}

//proc cpu
type ProcCPU struct {
	/* must match ProcTime fields */
	StartTime,
	User,
	Sys,
	Total uint64
	LastTime uint64
	Percent  float64
}

func (p *ProcCPU) String() string {
	return fmt.Sprintf("start_time(%v),user(%v),sys(%v),total(%v),last_time(%v),percent(%v)",
		p.StartTime, p.User, p.Sys, p.Total, p.LastTime, p.Percent)
}

//proc state
type ProcState struct {
	Name      string
	State     byte
	Ppid      uint64
	Tty       int
	Priority  int
	Nice      int
	Processor int
	Threads   uint64
}

func (p *ProcState) String() string {
	return fmt.Sprintf("name(%v),state(%v),ppid(%v),tty(%v),priority(%v),nice(%v),processor(%v),threads(%v)",
		p.Name, p.State, p.Ppid, p.Tty, p.Priority, p.Nice, p.Processor, p.Threads)
}

//proc exe
type ProcExe struct {
	Name string
	Cwd  string
	Root string
}

func (p *ProcExe) String() string {
	return fmt.Sprintf("name(%v),cwd(%v),root(%v)",
		p.Name, p.Cwd, p.Root)
}

//thread cpu
type ThreadCPU struct {
	User  uint64
	Sys   uint64
	Total uint64
}

func (t *ThreadCPU) String() string {
	return fmt.Sprintf("user(%v),sys(%v),total(%v)",
		t.User, t.Sys, t.Total)
}

const (
	SIGAR_FSTYPE_UNKNOWN = iota
	SIGAR_FSTYPE_NONE
	SIGAR_FSTYPE_LOCAL_DISK
	SIGAR_FSTYPE_NETWORK
	SIGAR_FSTYPE_RAM_DISK
	SIGAR_FSTYPE_CDROM
	SIGAR_FSTYPE_SWAP
	SIGAR_FSTYPE_MAX
)

//file system
type FileSystem struct {
	DirName     string
	DevName     string
	TypeName    string
	SysTypeName string
	Options     string
	Type        int
	Flags       uint64
}

func (f *FileSystem) String() string {
	return fmt.Sprintf("dir_name(%v),dev_name(%v),type_name(%v),sys_type_name(%v),options(%v),type(%v),flags(%v)",
		f.DirName, f.DevName, f.TypeName, f.SysTypeName, f.Options, f.Type, f.Flags)
}

//disk usage
type DiskUsage struct {
	Reads,
	Writes,
	WriteBytes,
	ReadBytes,
	RTime,
	WTime,
	QTime,
	Time,
	Snaptime uint64
	ServiceTime,
	Queue float64
}

func (d *DiskUsage) String() string {
	return fmt.Sprintf("reads(%v),writes(%v),write_bytes(%v),read_bytes(%v),rtime(%v),wtime(%v),qtime(%v),time(%v),snaptime(%v),service_time(%v),queue(%v)",
		d.Reads, d.Writes, d.WriteBytes, d.ReadBytes, d.RTime, d.WTime, d.QTime, d.Time, d.Snaptime, d.ServiceTime, d.Queue)
}

//file system usage
type FileSystemUsage struct {
	Disk       DiskUsage
	UsePercent float64
	Total,
	Free,
	Used,
	Avail,
	Files,
	FreeFiles uint64
}

func (f *FileSystemUsage) String() string {
	return fmt.Sprintf("total(%v),free(%v),used(%v),avail(%v),files(%v),free_files(%v)->disk:%v",
		f.Total, f.Free, f.Used, f.Avail, f.Files, f.FreeFiles, f.Disk)
}

const (
	SIGAR_AF_UNSPEC = iota
	SIGAR_AF_INET
	SIGAR_AF_INET6
	SIGAR_AF_LINK
)

type NetAddr struct {
	Family int
	In     uint32
	In6    [4]uint32
	Mac    [8]uint8
}

func (n *NetAddr) String() string {
	return fmt.Sprintf("family(%v),in(%v),in6(%v),mac(%v)",
		n.Family, n.In, n.In6, n.Mac)
}

//net info
type NetInfo struct {
	DefaultGateway          string
	DefaultGatewayInterface string
	HostName                string
	DomainName              string
	PrimaryDns              string
	SecondaryDns            string
}

func (n *NetInfo) String() string {
	return fmt.Sprintf("default_gateway(%v),default_gateway_interface(%v),host_name(%v),domain_name(%v),primary_dns(%v),secondary_dns(%v)",
		n.DefaultGateway, n.DefaultGatewayInterface, n.HostName, n.DomainName, n.PrimaryDns, n.SecondaryDns)
}

//net route
type NetRoute struct {
	Destination NetAddr
	Gateway     NetAddr
	Mask        NetAddr
	Flags,
	Refcnt,
	Use,
	Metric,
	Mtu,
	Window,
	Irtt uint64
	IfName string
}

func (n *NetRoute) String() string {
	return fmt.Sprintf("flags(%v),refcnt(%v),use(%v),metric(%v),mtu(%v),window(%v),irtt(%v),ifname(%v)\ndestination:%v\ngateway:%v\nmask:%v",
		n.Flags, n.Refcnt, n.Use, n.Metric, n.Mtu, n.Window, n.Irtt, n.IfName, n.Destination.String(), n.Gateway.String(), n.Mask.String())
}
