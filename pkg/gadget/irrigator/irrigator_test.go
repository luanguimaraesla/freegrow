// +build unit

package irrigator

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/luanguimaraesla/freegrow/internal/gadget"
	tt "github.com/luanguimaraesla/freegrow/test"
	"gopkg.in/yaml.v2"
	"gotest.tools/assert"
)

func TestMain(m *testing.M) {
	configurations := []tt.ConfigurationFunc{
		tt.ConfigureFlags,
	}

	setup := tt.GetSetup(configurations)
	setup(m)
}

func TestIrrigatorUnmarshal(t *testing.T) {
	irrigatorData := []byte(`---
name: default
enabled: true
port: 14
states:
- name: "on"
  schedule: "* * */1 * *"
- name: "off"
  schedule: "* * */1 * *"`)

	expectedIrrigator := Irrigator{
		Gadget: gadget.Gadget{
			Name:    "default",
			Enabled: true,
		},
		Scheduler: gadget.Scheduler{
			States: []*gadget.State{
				&gadget.State{
					Name:     "on",
					Schedule: "* * */1 * *",
				},
				&gadget.State{
					Name:     "off",
					Schedule: "* * */1 * *",
				},
			},
		},
		Port: 14,
	}

	i := Irrigator{}

	err := yaml.Unmarshal(irrigatorData, &i)
	assert.NilError(t, err)

	assert.DeepEqual(t, &i, &expectedIrrigator, cmp.AllowUnexported(Irrigator{}, gadget.Gadget{}))
}
