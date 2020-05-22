package gadget

type State struct {
	Name     string `yaml:"name"`
	Schedule string `yaml:"schedule"`
}

type Scheduler struct {
	Board  string   `yaml:"name"`
	States []*State `yaml:"states"`
}
