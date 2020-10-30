package main

import (
	"github.com/maxmind/mmdbwriter"
	"github.com/maxmind/mmdbwriter/mmdbtype"
	"log"
	"net"
	"os"
)

// 1.0.16.0/24,2519,"ARTERIA Networks Corporation"
// 1.0.64.0/18,18144,"Energia Communications,Inc."

func main() {
	writer, err := mmdbwriter.New(
		mmdbwriter.Options{
			DatabaseType: "My-ASN-DB",
			RecordSize:   24,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	ip1 := "1.0.16.0/24"
	ip2 := "1.0.16.0/12"
	str1 := "fdj"
	str2 := "afdj"

	_, network, err := net.ParseCIDR(ip1)
	if err != nil {
		log.Fatal(err)
	}
	record := mmdbtype.Map{}
	record["flag"] = mmdbtype.String(str1)

	err = writer.Insert(network, record)
	if err != nil {
		log.Fatal(err)
	}

	_, network, err = net.ParseCIDR(ip2)
	if err != nil {
		log.Fatal(err)
	}
	record = mmdbtype.Map{}
	record["flag"] = mmdbtype.String(str2)

	err = writer.Insert(network, record)
	if err != nil {
		log.Fatal(err)
	}

	fh, err := os.Create("out.mmdb")
	if err != nil {
		log.Fatal(err)
	}

	_, err = writer.WriteTo(fh)
	if err != nil {
		log.Fatal(err)
	}
}
