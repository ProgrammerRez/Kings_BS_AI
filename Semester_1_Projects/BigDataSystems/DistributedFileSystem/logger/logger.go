package logger

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)


const(
	InfoLevel = iota
	WarnLevel
	ErrorLevel
)

// Creating the type strcut for the logger
type Logger struct{
	infoLogger 	*log.Logger
	warnLogger	*log.Logger
	errorLogger *log.Logger
	Level		int
}

// Creating an instance
var logger *Logger

const LOG_PATH = "./logs.txt"

func Init(){

	f, err := createLogFile(LOG_PATH)
	if err != nil{
		log.Fatal(err.Error())
	}

	log.SetOutput(f)
	flags := log.Ldate | log.Lmicroseconds | log.Lshortfile 
	logger = &Logger{
		Level: InfoLevel,
		infoLogger: log.New(f, "INFO: ", flags),
		warnLogger: log.New(f, "WARN: ", flags),
		errorLogger: log.New(f, "ERROR: ", flags),
	}
	
}



// Creating a function to make sure the file exists
func createLogFile(path string) (*os.File ,error){

	dir := filepath.Dir(path)

	if dir != "."{
		if err := os.MkdirAll(dir, os.ModePerm); err!= nil{
			return nil, errors.New(err.Error())
		}
	}

	f, err := os.OpenFile(path, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)

	if err != nil{
		return nil, errors.New(err.Error())
	}

	return f, nil
}

func SetLevel(lvl int){
	if logger != nil{
		logger.Level = lvl
	}
}

func Info(message string){
	if logger.Level <= InfoLevel{
		logger.infoLogger.Println(message)
	}
}

func Warn(message string){
	if logger.Level <= WarnLevel{
		logger.warnLogger.Println(message)
	}
}

func Error(message string){
	if logger.Level <= ErrorLevel{
		logger.errorLogger.Println(message)
	}
}
