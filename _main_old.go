package main

import (
	"fmt"
	"os"
)

func main() {
	var arg string
	var arg2 string

	args := os.Args
	count := len(args)

	if count <= 1 {
		printInfo()
		return
	}

	command := args[1]

	if count >= 3 {
		arg = args[2]
	} else {
		arg = ""
	}

	if count >= 4 {
		arg2 = args[3]
	} else {
		arg2 = ""
	}

	makeCommand(command, arg, arg2)
}

func makeCommand(command string, arg string, arg2 string) {
	var releaseType string

	if command == "hs" {
		err := hotfixStart()
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	if command == "hf" {
		err := hotfixFinish()
		if err != nil {
			fmt.Println(err)
			return
		}
		pushDevelopMasterTags()
		return
	}

	if command == "rs" {
		releaseType = "minor"

		if arg == "--major" {
			releaseType = "major"
		}

		if arg == "--first" {
			releaseType = "first"
		}

		err := releaseStart(releaseType)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if command == "rf" {
		releaseType = "minor"

		if arg == "--major" {
			releaseType = "major"
		}

		if arg == "--first" {
			releaseType = "first"
		}

		err := releaseFinish(releaseType)
		if err != nil {
			fmt.Println(err)
			return
		}

		pushDevelopMasterTags()
		return
	}

	if command == "fs" {
		if arg == "" {
			println("name feature require")
			return
		}

		err := featureStart(arg)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if command == "ff" {
		err := featureFinish()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if command == "c" {
		/*	if arg == "" {
			println("name feature require")
			return
		}*/

		err := commit(arg, arg2)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func printInfo() {
	println("VD 0.3.4")
	println("")
	println("Available commands:")
	println("hs - Hotfix Start")
	println("hf - Hotfix Finish")
	println("")
	println("rs - Release Start")
	println("rs --major - Release Start and upgrade major version")
	println("rs --first - Release Start first")
	println("rf - Release Finish")
	println("rf --major - Release Finish and upgrade major version")
	println("rf --first - Release Finish first")
	println("")
	println("fs - Feature Start")
	println("ff - Feature Finish")
	println("")
	println("c [?tasks] [?message] - Commit")
}

/*func featureStart(featureName string) error {
	_, err := git.gitPullDevelop()
	if err != nil {
		return err
	}

	_, err = git.gitCreateFeatureBranch(featureName)
	if err != nil {
		return errors.New("error create feature")
	}

	_, err = git.gitPushFeature(featureName)
	if err != nil {
		return errors.New("error push feature")
	}

	return nil
}

func featureFinish() error {
	stdout, err := git.gitStatus()
	if err != nil {
		return errors.New("error git status")
	}

	if util.isBranchFeature(stdout) == false {
		println("switch to the feature branch")
		return errors.New("error is not feature branch")
	}

	featureName, err := util.getFeatureName(stdout)
	if err != nil {
		return err
	}

	_, err = git.gitPullFeature(featureName)
	if err != nil {
		return err
	}

	_, err = git.gitCheckout("develop")
	if err != nil {
		return err
	}

	_, err = git.gitPullDevelop()
	if err != nil {
		return err
	}

	_, err = git.gitMergeFeature(featureName)
	if err != nil {
		return err
	}

	_, err = git.remoteGitDeleteFeature(featureName)
	if err != nil {
		return err
	}

	_, err = git.gitDeleteFeature(featureName)
	if err != nil {
		return err
	}

	_, err = git.gitPushDevelop()
	if err != nil {
		return err
	}

	return nil
}*/

/*func releaseStart(releaseType string) error {
	var releaseName string

	tags, err := git.remoteGitTags()
	if err != nil {
		return errors.New("error get version from remote tags")
	}

	remoteMajor, remoteMinor, _, err := util.getLastVersion(tags)
	if err != nil && releaseType != "first" {
		return errors.New("error get version from branch")
	}

	if releaseType == "major" {
		releaseName = util.getNameBranchFromVersion(remoteMajor+1, 0, 0)
	} else {
		releaseName = util.getNameBranchFromVersion(remoteMajor, remoteMinor+1, 0)
	}

	_, err = git.gitPullDevelop()
	if err != nil {
		return err
	}

	_, err = git.gitCreateReleaseBranch(releaseName)
	if err != nil {
		return errors.New("error create release branch: release/" + releaseName)
	}
	_, err = git.gitPushRelease(releaseName)
	if err != nil {
		return errors.New("error push release branch: release/" + releaseName)
	}

	return nil
}

func releaseFinish(releaseType string) error {
	var remoteReleaseName string
	var remoteMajor, remoteMinor = 0, 0

	stdout, err := git.gitStatus()
	if err != nil {
		return errors.New("error git status")
	}

	if util.isBranchRelease(stdout) == false {
		println("switch to the release branch")
		return errors.New("error is not release branch")
	}

	currentMajor, currentMinor, currentPatch, err := util.getLastVersion(stdout)
	if err != nil {
		return errors.New("error get version from branch")
	}

	releaseName := util.getNameBranchFromVersion(currentMajor, currentMinor, currentPatch)

	_, err = git.gitPullRelease(releaseName)
	if err != nil {
		return errors.New("error pull release: " + releaseName)
	}

	_, err = git.gitCheckout("master")
	if err != nil {
		return errors.New("error checkout to master")
	}

	_, err = git.gitPullMaster()
	if err != nil {
		return errors.New("error pull master")
	}

	_, err = git.gitMergeRelease(releaseName)
	if err != nil {
		return errors.New("error merge master")
	}

	if releaseType != "first" {
		tags, err := git.remoteGitTags()
		if err != nil {
			return errors.New("error get remote tags")
		}
		remoteMajor, remoteMinor, _, err = util.getLastVersion(tags)
	}

	if err != nil && (currentMajor >= 1 || currentMinor >= 2) {
		return errors.New("error get version from remote tags")
	}

	if releaseType == "major" {
		remoteReleaseName = util.getNameBranchFromVersion(remoteMajor+1, 0, 0)
	} else {
		remoteReleaseName = util.getNameBranchFromVersion(remoteMajor, remoteMinor+1, 0)
	}

	_, err = git.gitCreateTag(remoteReleaseName, "release from local branch: release/"+releaseName)
	if err != nil {
		return errors.New("error create tag")
	}

	_, err = git.gitCheckout("develop")
	if err != nil {
		return errors.New("error checkout to develop")
	}

	_, err = git.gitPullDevelop()
	if err != nil {
		return errors.New("error pull develop")
	}

	_, err = git.gitMergeRelease(releaseName)
	if err != nil {
		return errors.New("error merge develop")
	}

	_, err = git.remoteGitDeleteRelease(releaseName)
	if err != nil {
		return errors.New("error delete remote branch: release/" + releaseName)
	}

	_, err = git.gitDeleteRelease(releaseName)
	if err != nil {
		return errors.New("error delete local branch: release/" + releaseName)
	}

	return nil
}*/

/*func hotfixStart() error {
	tags, err := git.remoteGitTags()
	if err != nil {
		return err
	}

	major, minor, patch, err := util.getLastVersion(tags)
	if err != nil {
		return err
	}

	hotfixName := util.getNameBranchFromVersion(major, minor, patch+1)

	_, err = git.gitPullDevelop()
	if err != nil {
		return err
	}

	_, err = git.gitPullMaster()
	if err != nil {
		return err
	}

	_, err = git.gitCreateHotfixBranch(hotfixName)
	if err != nil {
		return err
	}

	return nil
}

func hotfixFinish() error {
	stdout, err := git.gitStatus()
	if err != nil {
		return errors.New("error git status")
	}

	if util.isBranchHotfix(stdout) == false {
		return errors.New("error is not hotfix branch")
	}

	currentMajor, currentMinor, currentPatch, err := util.getLastVersion(stdout)
	if err != nil {
		return errors.New("error get version from branch")
	}

	hotfixName := util.getNameBranchFromVersion(currentMajor, currentMinor, currentPatch)

	_, err = git.gitCheckout("master")
	if err != nil {
		return errors.New("error checkout to master")
	}

	_, err = git.gitPullMaster()
	if err != nil {
		return errors.New("error pull master")
	}

	_, err = git.gitMergeHotfix(hotfixName)
	if err != nil {
		return errors.New("error merge master")
	}

	tags, err := git.remoteGitTags()
	if err != nil {
		return errors.New("error get remote tags")
	}

	remoteMajor, remoteMinor, remotePatch, err := util.getLastVersion(tags)
	if err != nil {
		return errors.New("error get version from tags")
	}

	_, err = git.gitCreateTag(util.getNameBranchFromVersion(remoteMajor, remoteMinor, remotePatch+1), "hotfix from local branch: hotfix/"+hotfixName)
	if err != nil {
		return errors.New("error create tag")
	}

	_, err = git.gitCheckout("develop")
	if err != nil {
		return errors.New("error checkout to develop")
	}

	_, err = git.gitPullDevelop()
	if err != nil {
		return errors.New("error pull develop")
	}

	_, err = git.gitMergeHotfix(hotfixName)
	if err != nil {
		return errors.New("error merge develop")
	}

	_, err = git.gitDeleteHotfix(hotfixName)
	if err != nil {
		return errors.New("error delete local branch: hotfix/" + hotfixName)
	}
	return nil
}*/

/*func pushDevelopMasterTags() {
	_, err := git.gitPushMaster()
	if err != nil {
		return
	}

	_, err = git.gitPushDevelop()
	if err != nil {
		return
	}

	_, err = git.gitPushTags()
	if err != nil {
		return
	}
}*/

/*func commit(tasks string, message string) error {
	//https://ontid.atlassian.net/browse/CREOS-647
	var name string
	var fullMessage string
	var numbers []int

	stdout, err := git.gitStatus()
	if err != nil {
		return errors.New("error git status")
	}

	if tasks != "" {
		name = tasks
		numbers, err = util.getNumbersFromName(name)
	}

	if len(numbers) == 0 && util.isBranchFeature(stdout) {
		name, err = util.getFeatureName(stdout)
		if err != nil {
			return err
		}

		numbers, err = util.getNumbersFromName(name)
	}

	if len(numbers) == 0 {
		println("tasks is not valid")
		return errors.New("error is not feature branch")
	}

	count := len(numbers)

	for i := 0; i < count; i++ {
		fullMessage += "https://ontid.atlassian.net/browse/CREOS-" + strconv.Itoa(numbers[i]) + "\n"
	}

	fullMessage += message

	_, err = git.gitAdd()
	if err != nil {
		return err
	}

	_, err = git.gitCommit(fullMessage)
	if err != nil {
		return err
	}

	return nil
}*/
