package hardwareaddr

import (
	"testing"
)

func TestGenerateEUI48(test *testing.T) {
	hardwareAddress, err := GenerateEUI48()

	if err != nil {
		test.Error("Error Generating MAC:" + err.Error())
	}

	test.Logf("Generated MAC:%v\n", hardwareAddress)
}
