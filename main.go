package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	tableContent := [][]string{}

	// Host Infomation
	hostInfo, err := host.Info()
	if err != nil {
		log.Fatal(err)
	}
	tableContent = append(tableContent,
		[]string{"Hostname", hostInfo.Hostname},
		[]string{"OS", hostInfo.Platform + " " + hostInfo.PlatformVersion},
		[]string{"BootTime", time.Unix(int64(hostInfo.BootTime), 0).Format("2006-01-02 15:04:05")},
		[]string{"Uptime", formatBootTime(hostInfo.Uptime)},
	)

	// CPU Information
	tableContent = append(tableContent,
		[]string{"CPU Core", fmt.Sprintf("%d", runtime.NumCPU())},
	)

	// Disk Information
	diskInfo, err := disk.Usage("/")
	if err != nil {
		log.Fatal(err)
	}
	tableContent = append(tableContent,
		[]string{"Disk Total", makeSizeReadable(diskInfo.Total)},
	)

	// Memory Infomation
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}
	tableContent = append(tableContent,
		[]string{"Memory Total", makeSizeReadable(memInfo.Total)},
	)

	table := tablewriter.NewWriter(os.Stdout)
	table.AppendBulk(tableContent)
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.Render()
}

func makeSizeReadable(size uint64) string {
	return fmt.Sprintf("%d GB", size/1024/1024/1024)
}

// The time value
const (
	Second uint64 = 1
	Minute        = Second * 60
	Hour          = Minute * 60
	Day           = Hour * 24
)

func formatBootTime(bootTime uint64) string {
	dayVal := bootTime / Day
	hourVal := (bootTime - dayVal*Day) / Hour
	minVal := (bootTime - dayVal*Day - hourVal*Hour) / Minute
	secVal := bootTime - dayVal*Day - hourVal*Hour - minVal*Minute

	return fmt.Sprintf("%d Day %d Hour %d Minute %d Second", dayVal, hourVal, minVal, secVal)
}
