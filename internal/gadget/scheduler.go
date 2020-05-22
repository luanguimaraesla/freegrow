package gadget

type State struct {
	Name     string `yaml:"name"`
	Schedule string `yaml:"schedule"`
}

type Scheduler struct {
	States []*State `yaml:"states"`
}
