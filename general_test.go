package fakku

import "testing"

func TestFakkuGeneralInformation_1(t *testing.T) {
	// Currently, this will fail as this api function doesn't return what it says it is supposed to :(
	_, err := GetGeneralInformation()
	if err != nil {
		t.Error(err)
	}
}
