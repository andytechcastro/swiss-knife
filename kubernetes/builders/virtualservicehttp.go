package builders

// VirtualServiceHTTP struct for http information for virtualService
type VirtualServiceHTTP struct {
	Name  string        `json:"name"`
	Match []interface{} `json:"match,omitempty"`
}
