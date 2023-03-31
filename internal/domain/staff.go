package domain

type Employee struct {
	ID      uint64                 `json:"id"`
	Name    string                 `json:"name"`
	PhotoID int                    `json:"photo_id"`
	Meta    map[string]interface{} `json:"meta"`
}
