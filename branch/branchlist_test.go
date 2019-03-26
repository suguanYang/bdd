package branch

import (
	"fmt"
	"strings"
	"testing"

	"github.com/suguanyang/bdd/branch"
)

func TestGetBranchs(t *testing.T) {
	branchs := branch.GetBranchs()
	fmt.Printf("BranchDeletor() =%s", strings.Join(branchs, " "))
	if len(branchs) == 0 {
		t.Failed()
	}
}
