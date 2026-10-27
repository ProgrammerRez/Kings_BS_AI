package main

import(
	log "dfs/logger"
)

func main(){

	log.Init()
	log.SetLevel(log.InfoLevel)

	log.Info("Started Logging")
	log.Info("Started Distributed File System")
}