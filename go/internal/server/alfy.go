package server

type Alfy struct {
	Items []item `json:"items"`
}

type item struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
}
