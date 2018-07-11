package perm

//CheckResponse holds the response of checking resource
type CheckResponse struct {
	Allow ResourceXpression `json:"allow"`
	Deny  ResourceXpression `json:"deny"`
}
