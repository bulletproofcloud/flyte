/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

type FlyteidladminTaskNodeMetadata struct {
	// Captures the status of caching for this execution.
	CacheStatus   *CoreCatalogCacheStatus `json:"cache_status,omitempty"`
	CatalogKey    *CoreCatalogMetadata    `json:"catalog_key,omitempty"`
	CheckpointUri string                  `json:"checkpoint_uri,omitempty"`
}
