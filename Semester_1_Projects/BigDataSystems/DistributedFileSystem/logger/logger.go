package logger 

import(
	"errors"
	"log"
	"os"
	"path/filepath"
)

// Creating Log Levels

const (
	InfoLevel = iota
	WarnLevel
	ErrorLevel
)

type Logger struct{
	infoLogger 		*log.Logger
	warnLogger		*log.Logger
	errorLogger		*log.Logger
	Level 			int
}


// Creating a variable pointing to the struct
var logger *Logger

const LOG_PATH = "./logs.txt"

// Init Function
func Init(){

	// Creating the Log Output File

	f, err := CreateLoggingFile(LOG_PATH)
	if err != nil{
		log.Fatalf(err.Error())
	}

	log.SetOutput(f)
	flags := log.Ldate | log.Lmicroseconds | log.Lshortfile

	logger = &Logger{
		infoLogger: log.New(f, "INFO: ", flags),
		warnLogger: log.New(f, "WARN: ", flags),
		errorLogger: log.New(f, "ERROR: ", flags),
	}
}


// This function makes sure the file is already created and ready to write
func CreateLoggingFile(path string) (*os.File, error){

	dir := filepath.Dir(path)

	if dir != "."{
		if err := os.MkdirAll(path, os.ModePerm); err != nil{
			return nil, errors.New(err.Error())
		}
	}

	f, err :=  os.OpenFile(path, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil{
		return nil, errors.New(err.Error())
	}

	return f, nil
}	

// Logging Functions

// Set the Level for the logger
func SetLevel(level int){
	if logger != nil{
		logger.Level = level
	}
}

// Info Logging
func Info(message string){
	if logger.Level <= InfoLevel{
		logger.infoLogger.Println(message)
	}
}

// Warn Logging
func Warn(message string){
	if logger.Level <= WarnLevel{
		logger.warnLogger.Println(message)
	}
}

// Error Logging
func Error(message string){
	if logger.Level <= ErrorLevel{
		logger.errorLogger.Println(message)
	}
}