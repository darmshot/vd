package git

import (
	"errors"
	"github.com/darmshot/vd/util"
)

func FeatureStart(featureName string) error {
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

func FeatureFinish() error {
	stdout, err := gitStatus()
	if err != nil {
		return errors.New("error git status")
	}

	if util.IsBranchFeature(stdout) == false {
		println("switch to the feature branch")
		return errors.New("error is not feature branch")
	}

	featureName, err := util.GetFeatureName(stdout)
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
