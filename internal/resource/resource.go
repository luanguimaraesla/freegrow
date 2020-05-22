package resource

type Tags map[string]string

type Meta struct {
	Name string `yaml:"name"`
	Tags Tags   `yaml:"tags"`
}
