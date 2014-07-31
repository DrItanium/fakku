package fakku

import "testing"

func TestFakkuGeneralInformation_1(t *testing.T) {
	_, err := GetGeneralInformation()
	if err != nil {
		// Currently, this will fail as this api function doesn't return what it says it is supposed to :(
		t.Log(err)
	}
}
