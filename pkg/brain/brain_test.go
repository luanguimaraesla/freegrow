// +build unit

package brain

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
