package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/adfolks/cloudflare-liwanetwork/liwa"
	"github.com/adfolks/cloudflare-liwanetwork/tracing"
	"github.com/cloudflare/cloudflare-go"
	"go.opentelemetry.io/otel"
)

func main() {
	tp, err := tracing.TracerProvider()
	if err != nil {
		log.Fatal("This is an error")
	}
	otel.SetTracerProvider(tp)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}(ctx)
	tr := tp.Tracer("cloudflare main")
	ctx, span := tr.Start(ctx, "dns-update")
	defer span.End()
	dnsUpdate(ctx)
}

func dnsUpdate(ctx context.Context) {
	records := os.Args
	zoneName := os.Getenv("ZONE_NAME")
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
	fmt.Print(zoneName + " Record updated")
}
