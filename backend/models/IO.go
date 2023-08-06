package models

type InputGraph struct {
	Input [][]string `json:"input"`
	Type  string     `json:"type"`
}

type OutputGraph struct {
	Graph  [][]string `json:"graph"`
	Status string     `json:"status"`
}
