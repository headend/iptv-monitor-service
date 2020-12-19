package main

import (
	monitor_services "github.com/headend/iptv-monitor-service/monitor-services"
	"log"
)


func main()  {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Monitor service")
	monitor_services.StartServer()
}
