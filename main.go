package main

import (
	agent_services "iptv-monitor-service/monitor-services"
	"log"
	"github.com/headend/iptv-monitor-service/monitor-services"
)


func main()  {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Monitor service")
	monitor_services.StartServer()
}
