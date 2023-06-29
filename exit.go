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
)

// Function to exit if an error occured
func ExitIfError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: "+fmt.Sprint(err))
		log.Printf("-< %s >-\n", fmt.Sprint(err))
		os.Exit(1)
	}
}

// Function to exit if pointer is nill
func ExitIfNill(ptr interface{}) {
	if ptr == nil {
		fmt.Fprintln(os.Stderr, "Error: got a nil pointer.")
		log.Printf("-< Error: got a nil pointer.  >-\n")
		os.Exit(1)
	}
}

// Function to exit with the nagios standard exit code and word (WARNING, CRITICAL, UNKNOWN)
// if error was not nil
func ExitWithNagiosCode(exitValue int, err error) {
	if err != nil {
		var nagiosCode string
		switch exitValue {
		case 1:
			nagiosCode = "WARNING"
		case 2:
			nagiosCode = "CRITICAL"
		default:
			nagiosCode = "UNKNOWN"
		}
		fmt.Fprintln(os.Stderr, nagiosCode+" Error: "+fmt.Sprint(err))
		log.Printf("%s -< %s >-\n", nagiosCode, fmt.Sprint(err))
		os.Exit(exitValue)
	}
}
