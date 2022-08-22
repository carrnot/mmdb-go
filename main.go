package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"os"
	"strings"

	"github.com/maxmind/mmdbwriter"
	"github.com/maxmind/mmdbwriter/mmdbtype"
)

var (
	output  string
	include string
)

var (
	cnRecord = mmdbtype.Map{
		"country": mmdbtype.Map{
			"geoname_id":           mmdbtype.Uint32(1814991),
			"is_in_european_union": mmdbtype.Bool(false),
			"iso_code":             mmdbtype.String("CN"),
			"names": mmdbtype.Map{
				"en":    mmdbtype.String("China"),
				"zh-CN": mmdbtype.String("中国"),
			},
		},
	}
)

func initFlag() {
	flag.StringVar(&include, "i", "", "text containing IP CIDR. like ip.txt")
	flag.StringVar(&output, "o", "Country.mmdb", "mmdb file name")
	flag.Parse()
}

func main() {
	initFlag()

	if include == "" {
		log.Fatal("please use the -i parameter to specify the ip file")
		return
	}

	writer, err := mmdbwriter.New(mmdbwriter.Options{
		IncludeReservedNetworks: true,
		DatabaseType:            "GeoIP2-Country",
		RecordSize:              24,
	})

	if err != nil {
		log.Panic(err)
	}

	ipList, err := os.Open(include)
	if err != nil {
		log.Panic(err)
	}

	ipBuf := bufio.NewScanner(ipList)
	for ipBuf.Scan() {
		txt := ipBuf.Text()
		if strings.Contains(txt, "#") {
			continue
		}
		var ipNet *net.IPNet
		if _, ipNet, err = net.ParseCIDR(ipBuf.Text()); err != nil {
			log.Println("ParseCIDR:", err, "->", txt)
			continue
		}
		if err = writer.Insert(ipNet, cnRecord); err != nil {
			log.Panic(err)
		}
	}

	mmdb, err := os.Create(output)
	if err != nil {
		log.Panic(err)
	}

	_, err = writer.WriteTo(mmdb)
	if err != nil {
		log.Panic(err)
	}
}
