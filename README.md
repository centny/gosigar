gosigar
======
binding sigar api to golang

### Install

<br/>
`linux/unix`

* download `libsigar` source code from <https://github.com/hyperic/sigar>
* install by command 

```
./autogen.sh
./configure --prefix=/usr/local/
make
make install
```
* install gosigar by command

```
export CGO_CPPFLAGS="-I/usr/local/include"
export CGO_LDFLAGS="-L/usr/local/lib -lsigar"
go get github.com/Centny/gosigar
```

* test gosigar

```
go test -v github.com/Centny/gosigar
```

note: adding libsigar install path to LD_LIBRARY_PATH

<br/>
`window`


### Example

`NetSpeed`

```
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
```

`Memory`

```
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
```

`Cpu`

```
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
```