package domain

type Employee struct {
	ID      int                    `json:"id"`
	Name    string                 `json:"name"`
	PhotoID int                    `json:"photo_id"`
	Meta    map[string]interface{} `json:"meta"`
}
