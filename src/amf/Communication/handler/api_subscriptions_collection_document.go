/*
 * Namf_Communication
 *
 * AMF Communication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Communication

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// AMFStatusChangeSubscribe - Namf_Communication AMF Status Change Subscribe service Operation
func AMFStatusChangeSubscribe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}