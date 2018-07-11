package perm

//Policy holds permissions grouped by resource group
type Policy struct {
	Groups map[string]map[string]map[string]map[string][]*Permission //client-id -> platform -> resource group -> resource name -> Resource
}
