package main

import (
	"os"
	"strconv"

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
		_, _, promptErr := prompt.Run()

		chosedBranchs := []string{}
		for _, chosed := range chosedIndex {
			chosedBranchs = append(chosedBranchs, branchs[chosed])
		}
		if len(chosedBranchs) == 0 {
			return nil
		}
		branch.DeleteBranches(chosedBranchs)
		return promptErr
	}
	app.Run(os.Args)
}
