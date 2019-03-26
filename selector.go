package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/suguanyang/bdd/branch"
	"github.com/suguanyang/promptui"
	"gopkg.in/urfave/cli.v1"
)

func getChosedBranchs(chosedIndex []int, branchs []string) []string {
	chosedBranchs := []string{}
	for _, chosed := range chosedIndex {
		chosedBranchs = append(chosedBranchs, branchs[chosed])
	}
	return chosedBranchs
}

func main() {
	branchs := branch.GetBranchs()
	chosedIndex := []int{}
	searcher := func(input string, index int) bool {
		branch := branchs[index]
		name := strings.Replace(strings.ToLower(branch), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}
	prompt := promptui.Select{
		Label:       "Branchs: ",
		Checkbox:    true,
		ChosedIcon:  promptui.IconGood,
		ChosenIndex: &chosedIndex,
		Items:       branchs,
	}
	size := "5"
	app := cli.NewApp()
	app.Name = "bdd"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "size, s",
			Value:       "5",
			Usage:       "chose how many branchs `per` a page",
			Destination: &size,
		},
	}
	app.Action = func(c *cli.Context) error {
		convertedSize, err := strconv.Atoi(size)

		if err != nil {
			convertedSize = 5
		}
		prompt.Size = convertedSize
		if c.NArg() > 0 && c.Args().First() == "checkout" {
			prompt.Searcher = searcher
			prompt.Checkbox = false
		}
		_, chosedBranch, promptErr := prompt.Run()
		if prompt.Checkbox {
			branch.Delete(getChosedBranchs(chosedIndex, branchs))
		} else {
			branch.Checkout(chosedBranch)
		}
		return promptErr
	}
	app.Run(os.Args)
}
