package node

import (
	"io/ioutil"

	"github.com/luanguimaraesla/freegrow/internal/global"
	"github.com/luanguimaraesla/freegrow/internal/resource"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Node struct {
	Metadata *resource.Metadata `yaml:"metadata" json:"metadata"`
	Spec     *NodeSpec          `yaml:"spec" json:"spec"`
	Status   *NodeStatus        `yaml:"status,omitempty" json:"status,omitempty"`
	logger   *zap.Logger
}

type NodeSpec struct {
	Board   nodeBoard `yaml:"board"`
	Machine *Machine  `yaml:"machine"`
}

func New() *Node {
	return &Node{}
}

func (n *Node) Load(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(content, n); err != nil {
		return err
	}

	return nil
}

func (n *Node) Init() error {
	n.Logger().With(
		zap.String("host", n.Spec.Machine.Host),
		zap.Int("port", n.Spec.Machine.Port),
	).Info("registering")

	if err := n.Spec.Machine.Register(n); err != nil {
		return err
	}

	return nil
}

func (n *Node) Run() error {
	return nil
}

func (n *Node) Logger() *zap.Logger {
	if n.logger == nil {
		n.logger = global.Logger.With(
			zap.String("entity", "node"),
			zap.String("name", n.Metadata.Name),
		)
	}

	return n.logger
}
