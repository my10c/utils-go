// Copyright (c) 2017 Badassops
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//	* Redistributions of source code must retain the above copyright
//	notice, this list of conditions and the following disclaimer.
//	* Redistributions in binary form must reproduce the above copyright
//	notice, this list of conditions and the following disclaimer in the
//	documentation and/or other materials provided with the distribution.
//	* Neither the name of the <organization> nor the
//	names of its contributors may be used to endorse or promote products
//	derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSEcw
// ARE DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// Author		:	Luc Suryo <luc@badassops.com>
//
// Version		:	0.1
//
// Date			:	Jul 11, 2017
//
// History	:
// 	Date:			Author:		Info:
//	Jul 11, 2017	LIS			Release as separate package
//
// TODO:

package utils

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path"
	"runtime"
)

const (
	// nagios exit values
	OK       = 0
	WARNING  = 1
	CRITICAL = 2
	UNKNOWN  = 3
)

var (
	// syslog need to this so configuration can use string instead of int
	SyslogPriority = map[string]int{
		"LOG_EMERG":   0,
		"LOG_ALERT":   1,
		"LOG_CRIT":    2,
		"LOG_ERR":     3,
		"LOG_WARNING": 4,
		"LOG_NOTICE":  5,
		"LOG_INFO":    6,
		"LOG_DEBUG":   7,
	}

	SyslogFacility = map[string]int{
		"LOG_MAIL":     0,
		"LOG_DAEMON":   1,
		"LOG_AUTH":     2,
		"LOG_SYSLOG":   3,
		"LOG_LPR":      4,
		"LOG_NEWS":     5,
		"LOG_UUCP":     6,
		"LOG_CRON":     7,
		"LOG_AUTHPRIV": 8,
		"LOG_FTP":      9,
		"LOG_LOCAL0":   10,
		"LOG_LOCAL1":   11,
		"LOG_LOCAL2":   12,
		"LOG_LOCAL3":   13,
		"LOG_LOCAL4":   14,
		"LOG_LOCAL5":   15,
		"LOG_LOCAL6":   16,
		"LOG_LOCAL7":   18,
	}
)

// Function to print the given message to stdout and log file
func Log(message string) {
	fmt.Printf("-< %s >-\n", message)
	log.Printf("-< %s >-\n", message)
	return
}

// Function to check if the user that runs the app is root
func IsRoot() {
	if os.Geteuid() != 0 {
		Log(fmt.Sprintf("%s must be run as root.", path.Base(os.Args[0])))
		os.Exit(1)
	}
}

// Function to check if the system is the given OS
func IsOS(osName string) (string, bool) {
	if runtime.GOOS == osName {
		return runtime.GOOS, true
	}
	return runtime.GOOS, false
}

// Function to check if we are on a Linux system
func IsLinux() {
	if osName, ok := IsOS("linux"); !ok {
		fmt.Printf("OS (%s) not supported, this check can only be run on a Linux system.\n", osName)
		os.Exit(1)
	}
	return
}

// Function to log any reveived signal
func SignalHandler() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt)
	go func() {
		sigId := <-interrupt
		Log(fmt.Sprintf("received %v", sigId))
		os.Exit(0)
	}()
}
