package domain

type Employee struct {
	ID      uint64
	Name    string
	PhotoID int
	Meta    map[string]interface{}
}
