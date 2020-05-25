package node

import (
	"io/ioutil"

	"github.com/luanguimaraesla/freegrow/internal/controller"
	"github.com/luanguimaraesla/freegrow/internal/global"
	"github.com/luanguimaraesla/freegrow/internal/resource"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Runner interface {
	Init() error
	Run() error
	Name() string
}

type Node struct {
	Kind     string             `yaml:"kind" json:"kind"`
	Metadata *resource.Metadata `yaml:"metadata" json:"metadata"`
	Spec     *NodeSpec          `yaml:"spec" json:"spec"`
	Status   *NodeStatus        `yaml:"status,omitempty" json:"status,omitempty"`
	logger   *zap.Logger
}

type NodeSpec struct {
	Board     string   `yaml:"board" json:"board"`
	Machine   *Machine `yaml:"machine" json:"machine"`
	Resources []string `yaml:"resources" json:"resources"`
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

	for _, kind := range n.Spec.Resources {
		list, err := n.Spec.Machine.GetResources(kind)
		if err != nil {
			return err
		}

		for _, rsc := range list.Resources {
			if runner, ok := rsc.(Runner); ok {
				if err := runner.Init(); err != nil {
					n.Logger().Error(
						"failed initializing runner",
						zap.String("kind", kind),
						zap.String("name", runner.Name()),
					)
				} else {
					n.Logger().Info(
						"starting runner",
						zap.String("kind", kind),
						zap.String("name", runner.Name()),
					)

					go runner.Run()
				}
			} else {
				n.Logger().Error(
					"resource is not runnable",
					zap.String("kind", kind),
				)
			}
		}
	}

	for {
	}
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
