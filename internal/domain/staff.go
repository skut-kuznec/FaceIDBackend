package domain

type Employee struct {
	ID      uint64
	Name    string
	PhotoID uint64
	Meta    map[string]interface{}
}
