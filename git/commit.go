package git

import (
	"errors"
	"github.com/darmshot/vd/util"
)

func Commit(tasks string, message string) error {
	//https://ontid.atlassian.net/browse/CREOS-647
	var name string
	var fullMessage string
	var numbers []int // - numbers of task

	stdout, err := gitStatus()
	if err != nil {
		return errors.New("error git status")
	}

	if tasks != "" {
		name = tasks
		numbers, err = util.GetNumbersFromName(name)
	}

	if len(numbers) == 0 && util.IsBranchFeature(stdout) {
		name, err = util.GetFeatureName(stdout)
		if err != nil {
			return err
		}

		numbers, err = util.GetNumbersFromName(name)
	}

	if len(numbers) == 0 {
		println("feature is not valid")
		return errors.New("error is not feature branch")
	}

	fullMessage = util.GetCommitMessage(name, numbers, message)

	_, err = gitAdd()
	if err != nil {
		return err
	}

	_, err = gitCommit(fullMessage)
	if err != nil {
		return err
	}

	return nil
}
