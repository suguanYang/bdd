package main

import (
	"os"

	"github.com/suguanyang/commandlist/branch"
	"github.com/suguanyang/promptui"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	branchs := branch.GetBranchs()
	chosedIndex := []int{}
	prompt := promptui.Select{
		Label:       "Branchs: ",
		Checkbox:    true,
		ChosedIcon:  promptui.IconGood,
		ChosenIndex: &chosedIndex,
		Items:       branchs,
	}
	app := cli.NewApp()
	app.Name = "branchdeletor"
	app.Action = func(c *cli.Context) error {
		_, _, err := prompt.Run()
		if err != nil {
			return err
		}

		chosedBranchs := []string{}
		for chosed := range chosedIndex {
			chosedBranchs = append(chosedBranchs, branchs[chosed])
		}

		branch.DeleteBranches(chosedBranchs)
		return err
	}
	app.Run(os.Args)
}
