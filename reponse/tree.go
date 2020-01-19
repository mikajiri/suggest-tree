package reponse

type SuggestNode struct {
	Text  string        `json:"text,omitempty"`
	Nodes []SuggestNode `json:"nodes,omitempty"`
	Tags  []string      `json:"tags,omitempty"`
}
