package node

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/luanguimaraesla/freegrow/internal/global"
	"github.com/luanguimaraesla/freegrow/internal/resource"
	"github.com/luanguimaraesla/freegrow/pkg/gadget/irrigator"
	"go.uber.org/zap"
)

var availableResources = []string{
	"irrigator",
}

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

func (m *Machine) GetResources(kind string) (*resource.ResourceList, error) {
	resp, err := m.get(filepath.Join("resources", kind))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("error getting %s resources: %d %s", kind, resp.StatusCode, body)
	}

	switch kind {
	case "irrigator":
		list := irrigator.NewIrrigatorList()
		if err := json.NewDecoder(resp.Body).Decode(list); err != nil {
			return nil, err
		}

		return resource.NewResourceList(list.Resources), nil
	default:
		return nil, fmt.Errorf("kind %s is not allowed", kind)
	}
}

func (m *Machine) get(path string) (*http.Response, error) {
	url := m.url(path)
	req, err := http.NewRequest("GET", url, nil)

	m.Logger().Debug("getting from machine API", zap.String("url", url))

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10

	client := retryClient.StandardClient()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
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
	retryClient.RetryMax = 10

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
