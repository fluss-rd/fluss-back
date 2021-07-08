package httputils

// PatchOperation represents the request body of a patch request
type PatchOperation struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

// PatchRequest defines a general patch request
type PatchRequest []PatchOperation
