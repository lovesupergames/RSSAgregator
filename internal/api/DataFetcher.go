package api

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
)

func (cfg *ApiConfig) DataFetcher(url string) Root {
	var root Root
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	XMLdata, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	xml.Unmarshal(XMLdata, &root)
	return root
}
