/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Represents a custom set of attributes applied for either a domain; a domain and project; or domain, project and workflow name. These are used to override system level defaults for kubernetes cluster resource management, default execution values, and more all across different levels of specificity.
type AdminMatchableAttributesConfiguration struct {
	Attributes *AdminMatchingAttributes `json:"attributes,omitempty"`
	Domain     string                   `json:"domain,omitempty"`
	Project    string                   `json:"project,omitempty"`
	Workflow   string                   `json:"workflow,omitempty"`
	LaunchPlan string                   `json:"launch_plan,omitempty"`
}
