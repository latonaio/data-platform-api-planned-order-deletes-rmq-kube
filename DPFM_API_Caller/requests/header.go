package requests

type Header struct {
	PlannedOrder     	int   `json:"PlannedOrder"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}
