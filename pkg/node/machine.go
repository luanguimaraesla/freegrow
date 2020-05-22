package node

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Machine struct {
	Host string `yaml:"url"`
	Port int    `yaml:"port"`
}

func (m *Machine) Register(node *Node) error {
	body, err := json.Marshal(node)
	if err != nil {
		return err
	}

	url := m.url("nodes/register")
	buf := bytes.NewBuffer(body)

	req, err := http.NewRequest("POST", url, buf)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated &&
		resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("error registering node: %s", string(body))
	}

	return nil
}

func (m *Machine) url(path string) string {
	u := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%d", m.Host, m.Port),
		Path:   path,
	}

	return u.String()
}
