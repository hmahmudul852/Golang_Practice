package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string
	var mxRecord *net.MX

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range mxRecords {
		mxRecord = record
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasSuffix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Printf("ErrorL%v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("domain = %v\nhasMX = %v\nmxRecord = %v\nhasSPF = %v\nsprRecord = %v\nhasDMARC = %v\ndmarcRecord = %v\n\n", domain, hasMX, mxRecord, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input: %v\n", err)
	}
}
