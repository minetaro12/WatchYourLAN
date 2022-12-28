package scan

import (
	"fmt"
	"log"

	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

func hostInDB(path string, host models.Host, dbHosts []models.Host) bool { // Check if host is already in DB
	for _, oneHost := range dbHosts {
		if host.IP == oneHost.IP && host.Mac == oneHost.Mac && host.Hw == oneHost.Hw {
			oneHost.Date = host.Date
			oneHost.Now = 1
			db.Update(path, oneHost)
			return true
		}
	}
	return false
}

func hostsCompare(path string, foundHosts, dbHosts []models.Host) {
	for _, oneHost := range foundHosts {
		if !(hostInDB(path, oneHost, dbHosts)) {
			oneHost.Now = 1 // Mark host online
			msg := fmt.Sprintf("UNKNOWN HOST IP: '%s', MAC: '%s', Hw: '%s'", oneHost.IP, oneHost.Mac, oneHost.Hw)
			log.Println("WARN:", msg)
			// shoutr_notify(msg) // Notify through Shoutrrr
			db.Insert(path, oneHost)
		}
	}
}