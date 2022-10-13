package git

import (
	"errors"
	"github.com/darmshot/vd/util"
)

func HotfixStart() error {
	tags, err := remoteGitTags()
	if err != nil {
		return err
	}

	major, minor, patch, err := util.GetLastVersion(tags)
	if err != nil {
		return err
	}

	hotfixName := util.GetNameBranchFromVersion(major, minor, patch+1)

	_, err = gitCheckout("develop")
	if err != nil {
		return errors.New("error checkout to develop")
	}

	_, err = gitPullDevelop()
	if err != nil {
		return err
	}

	_, err = gitCheckout("master")
	if err != nil {
		return errors.New("error checkout to master")
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

func HotfixFinish() error {
	stdout, err := gitStatus()
	if err != nil {
		return errors.New("error git status")
	}

	if util.IsBranchHotfix(stdout) == false {
		return errors.New("error is not hotfix branch")
	}

	currentMajor, currentMinor, currentPatch, err := util.GetLastVersion(stdout)
	if err != nil {
		return errors.New("error get version from branch")
	}

	hotfixName := util.GetNameBranchFromVersion(currentMajor, currentMinor, currentPatch)

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

	remoteMajor, remoteMinor, remotePatch, err := util.GetLastVersion(tags)
	if err != nil {
		return errors.New("error get version from tags")
	}

	_, err = gitCreateTag(util.GetNameBranchFromVersion(remoteMajor, remoteMinor, remotePatch+1), "hotfix from local branch: hotfix/"+hotfixName)
	if err != nil {
		return errors.New("error create tag")
	}

	_, err = gitPushTags()
	if err != nil {
		return errors.New("error push tags")
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

	_, err = gitPushDevelop()
	if err != nil {
		return errors.New("error push develop")
	}	

	_, err = gitDeleteHotfix(hotfixName)
	if err != nil {
		return errors.New("error delete local branch: hotfix/" + hotfixName)
	}
	return nil
}
