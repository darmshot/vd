package git

import (
	"errors"
	"github.com/darmshot/vd/util"
)

func ReleaseStart(releaseType string) error {
	var releaseName string

	tags, err := remoteGitTags()
	if err != nil {
		return errors.New("error get version from remote tags")
	}

	remoteMajor, remoteMinor, _, err := util.GetLastVersion(tags)
	if err != nil && releaseType != "first" {
		return errors.New("error get version from branch")
	}

	if releaseType == "major" {
		releaseName = util.GetNameBranchFromVersion(remoteMajor+1, 0, 0)
	} else {
		releaseName = util.GetNameBranchFromVersion(remoteMajor, remoteMinor+1, 0)
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

func ReleaseFinish(releaseType string) error {
	var remoteReleaseName string
	var remoteMajor, remoteMinor = 0, 0

	stdout, err := gitStatus()
	if err != nil {
		return errors.New("error git status")
	}

	if util.IsBranchRelease(stdout) == false {
		println("switch to the release branch")
		return errors.New("error is not release branch")
	}

	currentMajor, currentMinor, currentPatch, err := util.GetLastVersion(stdout)
	if err != nil {
		return errors.New("error get version from branch")
	}

	releaseName := util.GetNameBranchFromVersion(currentMajor, currentMinor, currentPatch)

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
		remoteMajor, remoteMinor, _, err = util.GetLastVersion(tags)
	}

	if err != nil && (currentMajor >= 1 || currentMinor >= 2) {
		return errors.New("error get version from remote tags")
	}

	if releaseType == "major" {
		remoteReleaseName = util.GetNameBranchFromVersion(remoteMajor+1, 0, 0)
	} else {
		remoteReleaseName = util.GetNameBranchFromVersion(remoteMajor, remoteMinor+1, 0)
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
