package nap

import (
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

func (c *Client) ProcessRequest(baseURL string, res *RestResource, params map[string]string) error {
	endpoint := strings.TrimLeft(res.RenderEndpoint(params), "/")
	trimmedBaseURL := strings.TrimRight(baseURL, "/")
	url := trimmedBaseURL + "/" + endpoint
	req, err := http.NewRequest(res.Method, url, nil)
	if err != nil {
		return err
	}
	if c.AuthInfo != nil {
		req.Header.Add("Authorization", c.AuthInfo.AuthorizationHeader())
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if err := res.Router.CallFunc(resp, nil); err != nil {
		return err
	}
	return nil
}
