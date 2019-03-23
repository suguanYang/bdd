package commandlist

import (
	"fmt"
	"testing"
)

func TestGetBranchList(t *testing.T) {
	branchs := GetBranchList()
	fmt.Printf("BranchDeletor() = %s waha", string(branchs[:]))
	if len(branchs) == 0 {
		t.Failed()
	}
}
