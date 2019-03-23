package commandlist

import (
	"fmt"
	"log"
	"os/exec"
)

// GetBranchList return a repo's branchs
func GetBranchList() []byte {
	cmd := exec.Command("git", "branch")
	branchs, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(branchs))
	return branchs
}

// func BranchDeletor() []byte {
// 	branchs := getBranchList()
//
// 	return branchs
// }
