package base

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	addr   string
	client *http.Client
}

func NewClient(addr string) *Client {
	s := &Client{
		addr: addr,
	}
	s.client = &http.Client{
		Timeout: time.Second * 10,
	}
	return s
}

func (s *Client) DoGet(uri string, obj interface{}) error {
	req, _ := http.NewRequest("GET", uri, nil)
	return DoHttp(s.client, req, obj)
}

func (s *Client) DoPost(uri string, contentType string, body io.Reader) error {
	req, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", contentType)
	return DoHttp(s.client, req, nil)
}

func (s *Client) DoPut(uri string, contentType string, body io.Reader) error {
	req, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", contentType)
	return DoHttp(s.client, req, nil)
}

func (s *Client) DoDelete(uri string, contentType string, body io.Reader) error {
	req, err := http.NewRequest("DELETE", uri, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", contentType)
	return DoHttp(s.client, req, nil)
}

func DoHttp(client *http.Client, req *http.Request, obj interface{}) error {
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%s %q bad status: %q", req.Method, req.URL, res.Status)
	}

	if obj != nil {
		if err = json.NewDecoder(res.Body).Decode(obj); err != nil {
			return err
		}
	}
	return nil
}
