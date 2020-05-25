package node

import (
	"io/ioutil"

	"github.com/luanguimaraesla/freegrow/internal/controller"
	"github.com/luanguimaraesla/freegrow/internal/global"
	"github.com/luanguimaraesla/freegrow/internal/resource"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Node struct {
	Kind     string             `yaml:"kind" json:"kind"`
	Metadata *resource.Metadata `yaml:"metadata" json:"metadata"`
	Spec     *NodeSpec          `yaml:"spec" json:"spec"`
	Status   *NodeStatus        `yaml:"status,omitempty" json:"status,omitempty"`
	logger   *zap.Logger
}

type NodeSpec struct {
	Board   string   `yaml:"board" json:"board"`
	Machine *Machine `yaml:"machine" json:"machine"`
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

	if err := controller.DefineController(n.Spec.Board); err != nil {
		return err
	}

	return nil
}

func (n *Node) Run() error {
	n.Logger().Info("running")

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
