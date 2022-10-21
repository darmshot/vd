package git

import (
	"os"
	"os/exec"
)

func remoteGitTags() (string, error) {
	cmd := exec.Command("git", "ls-remote", "--tags", "origin")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), err
}

func gitStatus() (string, error) {
	cmd := exec.Command("git", "status")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), nil
}

func gitCreateFeatureBranch(featureName string) (string, error) {
	cmd := exec.Command("git", "checkout", "-b", "feature/"+featureName, "develop")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), nil
}

func gitCreateHotfixBranch(hotfixName string) (string, error) {
	cmd := exec.Command("git", "checkout", "-b", "hotfix/"+hotfixName, "master")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), nil
}

func gitCreateReleaseBranch(releaseName string) (string, error) {
	cmd := exec.Command("git", "checkout", "-b", "release/"+releaseName, "develop")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), nil
}

func gitCheckout(branch string) (string, error) {
	cmd := exec.Command("git", "checkout", branch)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), nil
}

func gitMergeFeature(featureName string) (string, error) {
	cmd := exec.Command("git", "merge", "--no-ff", "feature/"+featureName)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), nil
}

func gitMergeHotfix(hotfixName string) (string, error) {
	cmd := exec.Command("git", "merge", "--no-ff", "hotfix/"+hotfixName)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), nil
}

func gitMergeRelease(releaseName string) (string, error) {
	cmd := exec.Command("git", "merge", "--no-ff", "release/"+releaseName)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), nil
}

func gitCreateTag(name string, message string) (string, error) {
	cmd := exec.Command("git", "tag", name, "-m", message)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), nil
}

func remoteGitDeleteFeature(featureName string) (string, error) {
	cmd := exec.Command("git", "push", "origin", "--delete", "feature/"+featureName)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), nil
}

func remoteGitDeleteRelease(releaseName string) (string, error) {
	cmd := exec.Command("git", "push", "origin", "--delete", "release/"+releaseName)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), nil
}

func gitDeleteFeature(featureName string) (string, error) {
	cmd := exec.Command("git", "branch", "--delete", "feature/"+featureName)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), nil
}

func gitDeleteRelease(releaseName string) (string, error) {
	cmd := exec.Command("git", "branch", "--delete", "release/"+releaseName)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), nil
}

func gitDeleteHotfix(hotfixName string) (string, error) {
	cmd := exec.Command("git", "branch", "--delete", "hotfix/"+hotfixName)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), nil
}

func gitPushFeature(featureName string) (string, error) {
	cmd := exec.Command("git", "push", "origin", "feature/"+featureName)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), nil
}

func gitPushRelease(releaseName string) (string, error) {
	cmd := exec.Command("git", "push", "origin", "release/"+releaseName)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), nil
}

func gitPushMaster() (string, error) {
	cmd := exec.Command("git", "push", "origin", "master")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), nil
}

func gitPushDevelop() (string, error) {
	cmd := exec.Command("git", "push", "origin", "develop")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), nil
}

func gitPushTags() (string, error) {
	cmd := exec.Command("git", "push", "--follow-tags")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), nil
}

func gitPullMaster() (string, error) {
	cmd := exec.Command("git", "pull", "origin", "master")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), nil
}

func gitPullFeature(featureName string) (string, error) {
	cmd := exec.Command("git", "pull", "origin", "feature/"+featureName)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), nil
}

func gitPullDevelop() (string, error) {
	cmd := exec.Command("git", "pull", "origin", "develop")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), nil
}

func gitPullRelease(releaseName string) (string, error) {
	cmd := exec.Command("git", "pull", "origin", "release/"+releaseName)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), nil
}

func gitCommit(message string) (string, error) {
	cmd := exec.Command("git", "commit", "-m", message)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "error git commit", err
	}

	return string(stdout), nil
}

func gitAdd() (string, error) {
	cmd := exec.Command("git", "add", "-u")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(stdout), nil
}
