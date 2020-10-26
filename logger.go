package main

import (
	"log"
	"os"
)

var (
	// WarningLogger provides warn level
	WarningLogger *log.Logger
	// InfoLogger provides info level
	InfoLogger *log.Logger
	// ErrorLogger provides err level
	ErrorLogger *log.Logger
)

func init() {
	InfoLogger = log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	WarningLogger = log.New(os.Stdout, "WARNING ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime)
}
