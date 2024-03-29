/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	sw "gogc/src/nssf/NSSAIAvailability/handler"
	"gogc/src/nssf/NSSAIAvailability/scenario"
	"log"
	"os"
)

func main() {
	log.Printf("Server started")

	// Init
	scenario.Init()
	router := sw.NewRouter()

	port := os.Getenv("ENV_NSSF_NSSAIAVAILABILITY_PORT")
	if port == "" {
		port = "8701"
	}
	port = ":" + port

	log.Fatal(router.Run(port))
}
