package resource

type Tags map[string]string

type Metadata struct {
	Name string `yaml:"name" json:"name"`
	Tags Tags   `yaml:"tags" json:"tags"`
}

type List struct {
	Resources interface{} `json:"resources"`
}

func NewList(i interface{}) *List {
	return &List{
		Resources: i,
	}
}
