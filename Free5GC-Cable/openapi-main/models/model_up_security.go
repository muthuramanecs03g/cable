/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type UpSecurity struct {
	UpIntegr UpIntegrity       `json:"upIntegr" yaml:"upIntegr" bson:"upIntegr" mapstructure:"UpIntegr"`
	UpConfid UpConfidentiality `json:"upConfid" yaml:"upConfid" bson:"upConfid" mapstructure:"UpConfid"`
}