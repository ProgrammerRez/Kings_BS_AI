package main 

import (
	log "github.com/ProgrammerRez/rezstore/logger"
)

func main(){
	log.Init()
	log.SetLevel(log.InfoLevel)

	log.Info("Distributed File System Started")
}