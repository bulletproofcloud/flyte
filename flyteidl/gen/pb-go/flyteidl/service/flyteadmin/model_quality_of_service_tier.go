/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// QualityOfServiceTier :  - UNDEFINED: Default: no quality of service specified.
type QualityOfServiceTier string

// List of QualityOfServiceTier
const (
	QualityOfServiceTierUNDEFINED QualityOfServiceTier = "UNDEFINED"
	QualityOfServiceTierHIGH      QualityOfServiceTier = "HIGH"
	QualityOfServiceTierMEDIUM    QualityOfServiceTier = "MEDIUM"
	QualityOfServiceTierLOW       QualityOfServiceTier = "LOW"
)
