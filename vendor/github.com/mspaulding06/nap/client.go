package nap

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

type Client struct {
	Client   *http.Client
	AuthInfo Authentication
}

func NewClient() *Client {
	return &Client{
		Client: http.DefaultClient,
	}
}

func (c *Client) SetAuth(auth Authentication) {
	c.AuthInfo = auth
}

func (c *Client) ProcessRequest(baseURL string, res *RestResource, params map[string]string, payload interface{}) error {
	endpoint := strings.TrimLeft(res.RenderEndpoint(params), "/")
	trimmedBaseURL := strings.TrimRight(baseURL, "/")
	url := trimmedBaseURL + "/" + endpoint
	req := buildClientRequest(res.Method, url, payload)
	if c.AuthInfo != nil {
		req.Header.Add("Authorization", c.AuthInfo.AuthorizationHeader())
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	return res.Router.CallFunc(resp)
}

func buildClientRequest(method, url string, payload interface{}) *http.Request {
	if payload != nil {
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			return nil
		}
		payloadBuffer := bytes.NewBuffer(payloadBytes)
		req, err := http.NewRequest(method, url, payloadBuffer)
		return req
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil
	}
	return req
}
