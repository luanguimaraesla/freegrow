package resource

type Tags map[string]string

type Metadata struct {
	Name string `yaml:"name"`
	Tags Tags   `yaml:"tags"`
}

type List struct {
	Resources interface{} `json:"resources"`
}

func NewList(i interface{}) *List {
	return &List{
		Resources: i,
	}
}
