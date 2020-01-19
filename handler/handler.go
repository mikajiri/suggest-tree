package handler

import (
	"net/http"
	"strconv"
	"suggest-tree/crawler"
	"suggest-tree/reponse"

	"github.com/labstack/echo"
)

func Suggest(c echo.Context) error {
	kw := c.Param("q")
	depth, err := strconv.Atoi(c.Param("depth"))
	if err != nil {
		depth = 1
	}
	if depth <= 0 {
		depth = 1
	}
	if depth > 3 {
		depth = 3
	}
	res := fetchSuggest(kw, 1, depth)

	return c.JSON(http.StatusOK, res.Nodes)
}

func fetchSuggest(kw string, depth int, maxdepth int) reponse.SuggestNode {
	res := crawler.GetSuggestions(kw)
	nodes := []reponse.SuggestNode{}
	for _, k := range res {
		var node reponse.SuggestNode
		if depth >= maxdepth {
			node = reponse.SuggestNode{Text: k}
		} else {
			node = fetchSuggest(k, depth+1, maxdepth)
		}
		nodes = append(nodes, node)
	}
	return reponse.SuggestNode{
		Text:  kw,
		Nodes: nodes,
		Tags:  []string{strconv.Itoa(len(nodes))},
	}
}
