// +build unit

package machine

import (
	"testing"

	tt "github.com/luanguimaraesla/freegrow/test"

	"gopkg.in/yaml.v3"
	"gotest.tools/assert"
)

func TestMain(m *testing.M) {
	configurations := []tt.ConfigurationFunc{
		tt.ConfigureFlags,
	}

	setup := tt.GetSetup(configurations)
	setup(m)
}

func TestGadgetsUnmarshal(t *testing.T) {
	gadgetsData := []byte(`---
- class: irrigator
  name: dafault
  port: 14
  states:
  - name: "on"
    schedule: "* * */1 * *"
  - name: "off"
    schedule: "* * */1 * *"`)

	expectedGadgets := &Gadgets{
		Items: []interface{}{},
	}

	gadgets := &Gadgets{}

	err := yaml.Unmarshal(gadgetsData, &gadgets)
	assert.NilError(t, err)

	assert.DeepEqual(t, gadgets, expectedGadgets)
}
