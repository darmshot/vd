package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var arg string

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

	makeCommand(command, arg)
}

func makeCommand(command string, arg string) {
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
}

func printInfo() {
	println("VD 0.2.2")
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

}

func featureStart(featureName string) error {
	_, err := gitCreateFeatureBranch(featureName)
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
		releaseName = buildNameBranchFromVersion(remoteMajor+1, 0, 0)
	} else {
		releaseName = buildNameBranchFromVersion(remoteMajor, remoteMinor+1, 0)
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

	releaseName := buildNameBranchFromVersion(currentMajor, currentMinor, currentPatch)

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
		remoteReleaseName = buildNameBranchFromVersion(remoteMajor+1, 0, 0)
	} else {
		remoteReleaseName = buildNameBranchFromVersion(remoteMajor, remoteMinor+1, 0)
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

	hotfixName := buildNameBranchFromVersion(major, minor, patch+1)

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

	hotfixName := buildNameBranchFromVersion(currentMajor, currentMinor, currentPatch)

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

	_, err = gitCreateTag(buildNameBranchFromVersion(remoteMajor, remoteMinor, remotePatch+1), "hotfix from local branch: hotfix/"+hotfixName)
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
