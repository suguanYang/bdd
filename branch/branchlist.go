package branch

import (
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func filter(s string) string {
	rule := `(\*)|(\ )`
	reg, err := regexp.Compile(rule)
	if err != nil {
		log.Fatal(err)
	}

	return reg.ReplaceAllString(s, "")
}

// GetBranchs returns a repo all branchs
func GetBranchs() []string {
	branchText := filter(getGitBranch())
	branchs := strings.Split(branchText, "\n")
	return branchs[:len(branchs)-1]
}

func getGitBranch() string {
	cmd := exec.Command("git", "branch")
	branchs, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(branchs)
}

// DeleteBranches delete passed branch
func DeleteBranches(branchs []string) {
	cmdList := append([]string{"branch", "-D"}, branchs...)
	cmd := exec.Command("git", cmdList...)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
