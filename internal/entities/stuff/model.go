package stuff

type Stuff struct {
	ID      int                    `json:"id,omitempty"`
	Name    string                 `json:"name,omitempty"`
	PhotoID int                    `json:"photo_id,omitempty"`
	Meta    map[string]interface{} `json:"meta,omitempty"`
}
