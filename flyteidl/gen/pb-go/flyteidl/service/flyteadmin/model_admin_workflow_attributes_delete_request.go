/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

type AdminWorkflowAttributesDeleteRequest struct {
	Project      string                  `json:"project,omitempty"`
	Domain       string                  `json:"domain,omitempty"`
	Workflow     string                  `json:"workflow,omitempty"`
	ResourceType *AdminMatchableResource `json:"resource_type,omitempty"`
}
