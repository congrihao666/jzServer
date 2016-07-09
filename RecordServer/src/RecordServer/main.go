package main

import (
       "jznet"
       )
func main() {
	recordServer := NewRecordServer("recordserver")
	jznet.StartService(recordServer,"10.1.100.39:10035")
}