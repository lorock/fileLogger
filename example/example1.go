// Package: example
// File: example1.go
// Created by: mint(mint.zhao.chiu@gmail.com)_aiwuTech
// Useage: example
// DATE: 14-8-24 14:08
package main

import (
	"github.com/aiwuTech/fileLogger"
	"sync"
)

var (
	TRACE  *fileLogger.FileLogger
	INFO   *fileLogger.FileLogger
	WARN   *fileLogger.FileLogger
	ERROR  *fileLogger.FileLogger
)

func init() {

	TRACE = fileLogger.NewDefaultLogger("/usr/local/bbsvr/log", "trace.log")
	INFO = fileLogger.NewDefaultLogger("/usr/local/bbsvr/log", "info.log")
	WARN = fileLogger.NewDefaultLogger("/usr/local/bbsvr/log", "warn.log")
	ERROR = fileLogger.NewDefaultLogger("/usr/local/bbsvr/log", "error.log")

	TRACE.SetPrefix("[TRACE] ")
	INFO.SetPrefix("[INFO] ")
	WARN.SetPrefix("[WARN] ")
	ERROR.SetPrefix("[ERROR] ")
}

func main() {

	wg := new(sync.WaitGroup)

	wg.Add(1)
	go traceTest(wg)

	wg.Add(1)
	go infoTest(wg)

	wg.Add(1)
	go warnTest(wg)

	wg.Add(1)
	go errorTest(wg)

	wg.Wait()
}

func traceTest(wg *sync.WaitGroup) {
	for i := 1; i <= 1000; i++ {
		TRACE.Printf("This is the No[%v] TRACE log using fileLogger that written by aiwuTech.", i)
	}

	wg.Done()
}

func infoTest(wg *sync.WaitGroup) {
	for i := 1; i <= 1000; i++ {
		INFO.Printf("This is the No[%v] INFO log using fileLogger that written by aiwuTech.", i)
	}

	wg.Done()
}

func warnTest(wg *sync.WaitGroup) {
	for i := 1; i <= 1000; i++ {
		WARN.Printf("This is the No[%v] WARN log using fileLogger that written by aiwuTech.", i)
	}

	wg.Done()
}

func errorTest(wg *sync.WaitGroup) {
	for i := 1; i <= 1000; i++ {
		ERROR.Printf("This is the No[%v] ERROR log using fileLogger that written by aiwuTech.", i)
	}

	wg.Done()
}


