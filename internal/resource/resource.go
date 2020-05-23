package resource

type Tags map[string]string

type Metadata struct {
	Name string `yaml:"name"`
	Tags Tags   `yaml:"tags"`
}
