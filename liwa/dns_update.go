package liwa

import (
	"context"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/rdegges/go-ipify"
)

func UpdateDNSRecord(ctx context.Context, api *cloudflare.API, record cloudflare.DNSRecord) {
	record.Content = GetPublicIp()
	api.UpdateDNSRecord(ctx, record.ZoneID, record.ID, record)
}
func GetPublicIp() string {
	ip, err := ipify.GetIp()
	if err != nil {
		log.Fatal("Error getting the Public IP")
	}
	return ip
}
