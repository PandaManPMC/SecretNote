package psutil

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

// SysInfo saves the basic system information
type SysInfo struct {
	Hostname string `bson:hostname`
	Platform string `bson:platform`
	CPU      string `bson:cpu`
	RAM      uint64 `bson:ram`
	Disk     uint64 `bson:disk`
}

func TestPCInfo(t *testing.T) {
	hostStat, _ := host.Info()
	cpuStat, _ := cpu.Info()
	vmStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("\\") // If you're in Unix change this "\\" for "/"

	info := new(SysInfo)
	info.Hostname = hostStat.Hostname
	info.Platform = hostStat.Platform
	info.CPU = cpuStat[0].ModelName
	info.RAM = vmStat.Total / 1024 / 1024
	info.Disk = diskStat.Total / 1024 / 1024

	fmt.Printf("%+v\n", info)

	fmt.Printf("Hostname: %s\n", hostStat.Hostname)
	fmt.Printf("Platform: %s\n", hostStat.Platform)
	fmt.Printf("CPU: %s\n", cpuStat[0].ModelName)
	fmt.Printf("RAM: %sgb\n", strconv.FormatUint(vmStat.Total/1024/1024, 10)[0:2])
	//fmt.Printf("Disk：%dm\n",info.Disk)

	// gets GPU info
	Info := exec.Command("cmd", "/C", "wmic path win32_VideoController get name")
	History, _ := Info.Output()
	replace := strings.Replace(string(History), "Name", "", -1)
	replace2 := strings.Replace(replace, "LuminonCore IDDCX Adapter", "", -1)

	// gets BOARD info
	Infos := exec.Command("cmd", "/C", "wmic path win32_BaseBoard get Product")
	Historys, _ := Infos.Output()
	replaces := strings.Replace(string(Historys), "Product", "", -1)
	fmt.Println("GPU: " + replace2)
	fmt.Println("Mainboard: " + replaces)

}
