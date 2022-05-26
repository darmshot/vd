package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
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

func featureStart(featureName string) error {
	_, err := gitPullDevelop()
	if err != nil {
		return err
	}

	_, err = gitCreateFeatureBranch(featureName)
	if err != nil {
		return errors.New("error create feature")
	}

	_, err = gitPushFeature(featureName)
	if err != nil {
		return errors.New("error push feature")
	}

	return nil
}

func featureFinish() error {
	stdout, err := gitStatus()
	if err != nil {
		return errors.New("error git status")
	}

	if isBranchFeature(stdout) == false {
		println("switch to the feature branch")
		return errors.New("error is not feature branch")
	}

	featureName, err := getFeatureName(stdout)
	if err != nil {
		return err
	}

	_, err = gitPullFeature(featureName)
	if err != nil {
		return err
	}

	_, err = gitCheckout("develop")
	if err != nil {
		return err
	}

	_, err = gitPullDevelop()
	if err != nil {
		return err
	}

	_, err = gitMergeFeature(featureName)
	if err != nil {
		return err
	}

	_, err = remoteGitDeleteFeature(featureName)
	if err != nil {
		return err
	}

	_, err = gitDeleteFeature(featureName)
	if err != nil {
		return err
	}

	_, err = gitPushDevelop()
	if err != nil {
		return err
	}

	return nil
}

func releaseStart(releaseType string) error {
	var releaseName string

	tags, err := remoteGitTags()
	if err != nil {
		return errors.New("error get version from remote tags")
	}

	remoteMajor, remoteMinor, _, err := getLastVersion(tags)
	if err != nil && releaseType != "first" {
		return errors.New("error get version from branch")
	}

	if releaseType == "major" {
		releaseName = getNameBranchFromVersion(remoteMajor+1, 0, 0)
	} else {
		releaseName = getNameBranchFromVersion(remoteMajor, remoteMinor+1, 0)
	}

	_, err = gitPullDevelop()
	if err != nil {
		return err
	}

	_, err = gitCreateReleaseBranch(releaseName)
	if err != nil {
		return errors.New("error create release branch: release/" + releaseName)
	}
	_, err = gitPushRelease(releaseName)
	if err != nil {
		return errors.New("error push release branch: release/" + releaseName)
	}

	return nil
}

func releaseFinish(releaseType string) error {
	var remoteReleaseName string
	var remoteMajor, remoteMinor = 0, 0

	stdout, err := gitStatus()
	if err != nil {
		return errors.New("error git status")
	}

	if isBranchRelease(stdout) == false {
		println("switch to the release branch")
		return errors.New("error is not release branch")
	}

	currentMajor, currentMinor, currentPatch, err := getLastVersion(stdout)
	if err != nil {
		return errors.New("error get version from branch")
	}

	releaseName := getNameBranchFromVersion(currentMajor, currentMinor, currentPatch)

	_, err = gitPullRelease(releaseName)
	if err != nil {
		return errors.New("error pull release: " + releaseName)
	}

	_, err = gitCheckout("master")
	if err != nil {
		return errors.New("error checkout to master")
	}

	_, err = gitPullMaster()
	if err != nil {
		return errors.New("error pull master")
	}

	_, err = gitMergeRelease(releaseName)
	if err != nil {
		return errors.New("error merge master")
	}

	if releaseType != "first" {
		tags, err := remoteGitTags()
		if err != nil {
			return errors.New("error get remote tags")
		}
		remoteMajor, remoteMinor, _, err = getLastVersion(tags)
	}

	if err != nil && (currentMajor >= 1 || currentMinor >= 2) {
		return errors.New("error get version from remote tags")
	}

	if releaseType == "major" {
		remoteReleaseName = getNameBranchFromVersion(remoteMajor+1, 0, 0)
	} else {
		remoteReleaseName = getNameBranchFromVersion(remoteMajor, remoteMinor+1, 0)
	}

	_, err = gitCreateTag(remoteReleaseName, "release from local branch: release/"+releaseName)
	if err != nil {
		return errors.New("error create tag")
	}

	_, err = gitCheckout("develop")
	if err != nil {
		return errors.New("error checkout to develop")
	}

	_, err = gitPullDevelop()
	if err != nil {
		return errors.New("error pull develop")
	}

	_, err = gitMergeRelease(releaseName)
	if err != nil {
		return errors.New("error merge develop")
	}

	_, err = remoteGitDeleteRelease(releaseName)
	if err != nil {
		return errors.New("error delete remote branch: release/" + releaseName)
	}

	_, err = gitDeleteRelease(releaseName)
	if err != nil {
		return errors.New("error delete local branch: release/" + releaseName)
	}

	return nil
}

func hotfixStart() error {
	tags, err := remoteGitTags()
	if err != nil {
		return err
	}

	major, minor, patch, err := getLastVersion(tags)
	if err != nil {
		return err
	}

	hotfixName := getNameBranchFromVersion(major, minor, patch+1)

	_, err = gitPullDevelop()
	if err != nil {
		return err
	}

	_, err = gitPullMaster()
	if err != nil {
		return err
	}

	_, err = gitCreateHotfixBranch(hotfixName)
	if err != nil {
		return err
	}

	return nil
}

func hotfixFinish() error {
	stdout, err := gitStatus()
	if err != nil {
		return errors.New("error git status")
	}

	if isBranchHotfix(stdout) == false {
		return errors.New("error is not hotfix branch")
	}

	currentMajor, currentMinor, currentPatch, err := getLastVersion(stdout)
	if err != nil {
		return errors.New("error get version from branch")
	}

	hotfixName := getNameBranchFromVersion(currentMajor, currentMinor, currentPatch)

	_, err = gitCheckout("master")
	if err != nil {
		return errors.New("error checkout to master")
	}

	_, err = gitPullMaster()
	if err != nil {
		return errors.New("error pull master")
	}

	_, err = gitMergeHotfix(hotfixName)
	if err != nil {
		return errors.New("error merge master")
	}

	tags, err := remoteGitTags()
	if err != nil {
		return errors.New("error get remote tags")
	}

	remoteMajor, remoteMinor, remotePatch, err := getLastVersion(tags)
	if err != nil {
		return errors.New("error get version from tags")
	}

	_, err = gitCreateTag(getNameBranchFromVersion(remoteMajor, remoteMinor, remotePatch+1), "hotfix from local branch: hotfix/"+hotfixName)
	if err != nil {
		return errors.New("error create tag")
	}

	_, err = gitCheckout("develop")
	if err != nil {
		return errors.New("error checkout to develop")
	}

	_, err = gitPullDevelop()
	if err != nil {
		return errors.New("error pull develop")
	}

	_, err = gitMergeHotfix(hotfixName)
	if err != nil {
		return errors.New("error merge develop")
	}

	_, err = gitDeleteHotfix(hotfixName)
	if err != nil {
		return errors.New("error delete local branch: hotfix/" + hotfixName)
	}
	return nil
}

func pushDevelopMasterTags() {
	_, err := gitPushMaster()
	if err != nil {
		return
	}

	_, err = gitPushDevelop()
	if err != nil {
		return
	}

	_, err = gitPushTags()
	if err != nil {
		return
	}
}

func commit(tasks string, message string) error {
	//https://ontid.atlassian.net/browse/CREOS-647
	var name string
	var fullMessage string
	stdout, err := gitStatus()
	if err != nil {
		return errors.New("error git status")
	}

	if tasks != "" {
		name = tasks
	} else if isBranchFeature(stdout) {
		name, err = getFeatureName(stdout)
		if err != nil {
			return err
		}
	} else {
		println("tasks is not valid")
		return errors.New("error is not feature branch")
	}

	fmt.Println(tasks)
	fmt.Println(message)

	numbers, err := getNumbersFromName(name)

	count := len(numbers)

	for i := 0; i < count; i++ {
		fullMessage += "https://ontid.atlassian.net/browse/CREOS-" + strconv.Itoa(numbers[i]) + "\n"
	}

	fullMessage += message

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
