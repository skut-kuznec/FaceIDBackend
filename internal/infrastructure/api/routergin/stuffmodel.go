package routergin

type StuffAnswer struct {
	ID      int                    `json:"id" example:"1"`
	Name    string                 `json:"name" example:"Сергей"`
	PhotoID int                    `json:"photo_id" example:"0"`
	Meta    map[string]interface{} `json:"meta"`
}
