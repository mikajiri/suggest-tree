package crawler

import (
	"bytes"
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"io"
	"time"

	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const userAgent = "Mozilla/5.0 (Windows NT 6.1; Trident/7.0; rv:11.0) like Gecko"

type GoogleSuggestions struct {
	CompleteSuggestions []struct {
		Suggestion struct {
			Data string `xml:"data,attr"`
		} `xml:"suggestion"`
	} `xml:"CompleteSuggestion"`
}

func GetSuggestions(kw string) []string {
	ret := []string{}
	xmlSrc := fetch(kw)
	if xmlSrc == nil {
		return ret
	}
	doc := parse(xmlSrc)
	if doc == nil {
		return ret
	}
	for _, s := range doc.CompleteSuggestions {
		ret = append(ret, s.Suggestion.Data)
	}
	return ret
}

func parse(src []byte) *GoogleSuggestions {
	suggest := new(GoogleSuggestions)
	if err := xml.Unmarshal(src, suggest); err != nil {
		fmt.Println(err)
		return nil
	}
	return suggest
}

func fetch(kw string) []byte {
	endpoint := "https://www.google.com/complete/search"
	req, _ := http.NewRequest("GET", endpoint, nil)
	h := getDefaultHeader()
	for k, v := range h {
		req.Header.Add(k, v)
	}
	req.Header.Add("User-Agent", userAgent)

	values := url.Values{}
	values.Add("hl", "ja")
	values.Add("output", "toolbar")
	values.Add("q", kw)

	req.URL.RawQuery = values.Encode()

	client := newClient()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(extract(resp.Body))
	return buf.Bytes()
}

func newClient() *http.Client {
	jar, _ := cookiejar.New(&cookiejar.Options{})
	client := &http.Client{
		Jar:     jar,
		Timeout: 30 * time.Second,
	}
	return client
}

func getDefaultHeader() map[string]string {
	h := make(map[string]string, 0)
	h["Accept-Language"] = "ja"
	h["Pragma"] = "no-cache"
	h["Cache-Control"] = "no-cache"
	h["Accept-Encoding"] = "gzip,deflate"
	h["Accept"] = "*/*"
	return h
}

func extract(zr io.Reader) io.Reader {
	ret, err := gzip.NewReader(zr)
	//cfg := &brotli.ReaderConfig{}
	//return brotli.NewReader(zr, cfg)
	if err != nil {
		return zr
	}
	return ret
}
