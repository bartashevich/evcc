package server

type product struct {
	Name     string `json:"name"`
	Template string `json:"template"`
	Group    string `json:"group"`
}

type products []product
