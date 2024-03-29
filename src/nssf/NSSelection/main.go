/*
 * NSSF NS Selection
 *
 * NSSF Network Slice Selection Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.0-alpha.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	sw "gogc/src/nssf/NSSelection/handler"
	"gogc/src/nssf/NSSelection/scenario"
	"log"
	"os"
)

func main() {
	log.Printf("Server started")

	// Init
	scenario.Init()
	router := sw.NewRouter()

	port := os.Getenv("ENV_NSSF_NSSELECTION_PORT")
	if port == "" {
		port = "8702"
	}
	port = ":" + port

	log.Fatal(router.Run(port))
}
