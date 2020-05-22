// +build unit

package machine

import (
	"testing"

	tt "github.com/luanguimaraesla/freegrow/test"
)

func TestMain(m *testing.M) {
	configurations := []tt.ConfigurationFunc{
		tt.ConfigureFlags,
	}

	setup := tt.GetSetup(configurations)
	setup(m)
}

//func TestGadgetsUnmarshal(t *testing.T) {
//	gadgetsData := []byte(`---
//- class: irrigator
//  spec:
//    name: dafault
//    port: 14
//    states:
//    - name: "on"
//      schedule: "* * */1 * *"
//    - name: "off"
//      schedule: "* * */1 * *"`)
//
//	expectedGadgets := []*Gadget{
//		&Gadget{
//			Class: "irrigator",
//			Spec:  nil,
//		},
//	}
//
//	gadgets := []*Gadget{}
//
//	err := yaml.Unmarshal(gadgetsData, &gadgets)
//	assert.NilError(t, err)
//
//	for _, g := range gadgets {
//		g.Spec = nil
//	}
//
//	assert.DeepEqual(t, gadgets, expectedGadgets)
//}
