// Copyright (c) 2025 The bel2 developers

package mempool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type API struct {
	config Config
}

func NewAPI(config Config) *API {
	switch strings.ToLower(config.Network) {
	case "mainnet":
		config.ApiBaseUrl = Mainnet_ApiBaseUrl
	case "testnet":
		config.ApiBaseUrl = Testnet_ApiBaseUrl
	}
	return &API{config: config}
}

func (api *API) GetRawTransaction(txid string) (*GetRawTransactionResponse, error) {
	buf, err := api.doRequst(txid)
	if err != nil {
		return nil, err
	}
	body := buf.String()
	fmt.Println(body)

	response := GetRawTransactionResponse{}
	err = json.Unmarshal(buf.Bytes(), &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (api *API) doRequst(url string) (*bytes.Buffer, error) {
	headersParams := api.getHeaders()
	req := api.getRequest(headersParams, url)

	client := api.getClient()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(resp.Body)

	return buf, nil
}

func (api *API) getRequest(headersParams map[string]string, url string) *http.Request {
	req, err := http.NewRequest("GET", getRequestUrl(api.config.ApiBaseUrl, url), bytes.NewBuffer([]byte{}))
	if err != nil {
		panic(err)
	}

	for key, value := range headersParams {
		req.Header.Set(key, value)
	}

	return req
}

func (api *API) getClient() *http.Client {
	client := &http.Client{}

	if api.config.Proxy != "" {
		proxyUrl, _ := url.Parse(api.config.Proxy)
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
		client.Transport = transport
	}

	return client
}

func (api *API) getHeaders() map[string]string {
	headersParams := map[string]string{
		"Content-Type": "application/json",
	}
	return headersParams
}

func getRequestUrl(apiBaseUrl string, api string) string {
	u, _ := url.Parse(apiBaseUrl + api)
	return u.String()
}

func (api *API) getUrl(url string, request map[string]interface{}) string {
	return url + "?" + api.getQueryString(request)
}

func (api *API) getQueryString(request map[string]interface{}) string {
	queryString := ""
	count := 0
	for key, value := range request {
		queryString += key + "=" + value.(string)
		count++
		if count < len(request) {
			queryString += "&"
		}
	}
	return queryString
}
