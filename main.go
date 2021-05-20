package main

import (
	"context"
	"log"
	"os"

	"github.com/adfolks/liwa-nw/liwa"
	"github.com/cloudflare/cloudflare-go"
)

var localVariable string

func main() {
	records := os.Args
	zoneName := os.Getenv("ZONE_NAME")
	ctx := context.Background()
	api, err := cloudflare.NewWithAPIToken(os.Getenv("API_TOKEN"))
	if err != nil {
		log.Fatal("Error getting the authentications")
	}
	zoneId, err := api.ZoneIDByName(zoneName)
	if err != nil {
		log.Fatal("Error getting zone information")
	}
	recs, err := api.DNSRecords(ctx, zoneId, cloudflare.DNSRecord{})

	for _, dns_records := range recs {
		for _, lr := range records {
			if dns_records.Type == "A" && dns_records.Name == lr {
				liwa.UpdateDNSRecord(ctx, api, dns_records)
			}
		}

	}

}
