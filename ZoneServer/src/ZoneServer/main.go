package main
import (
       	"jznet"
       )
       
func main() {
	zoneserver := NewZoneServer("zoneserver")
	jznet.StartService(zoneserver,"10.1.100.39:31892")
	return
}
