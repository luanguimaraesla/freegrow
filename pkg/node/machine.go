package node

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/luanguimaraesla/freegrow/internal/global"
	"go.uber.org/zap"
)

type Machine struct {
	Host   string `yaml:"host" json:"host"`
	Port   int    `yaml:"port" json:"port"`
	logger *zap.Logger
}

func (m *Machine) Register(node *Node) error {
	body, err := json.Marshal(node)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	resp, err := m.post("nodes", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated &&
		resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf("error registering node: %d %s", resp.StatusCode, body)
	}

	return nil
}

func (m *Machine) post(path string, body []byte) (*http.Response, error) {
	url := m.url(path)
	buf := bytes.NewBuffer(body)

	req, err := http.NewRequest("POST", url, buf)
	req.Header.Set("Content-Type", "application/json")

	m.Logger().With(
		zap.String("url", url),
	).Debug("posting to machine API")

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 100

	client := retryClient.StandardClient()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *Machine) url(path string) string {
	u := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%d", m.Host, m.Port),
		Path:   path,
	}

	return u.String()
}

func (m *Machine) Logger() *zap.Logger {
	if m.logger == nil {
		m.logger = global.Logger.With(
			zap.String("entity", "machine"),
			zap.String("host", m.Host),
			zap.Int("entity", m.Port),
		)
	}

	return m.logger
}
