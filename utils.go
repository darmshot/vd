package main

import (
	"errors"
	"regexp"
	"strconv"
)

func buildNameBranchFromVersion(major int, minor int, patch int) string {
	return "v" + strconv.Itoa(major) + "." + strconv.Itoa(minor) + "." + strconv.Itoa(patch)
}
func getFeatureName(gitStatus string) (string, error) {
	re := regexp.MustCompile(`On branch feature/(.*)`)

	result := re.FindStringSubmatch(gitStatus)
	if len(result) < 2 {
		return "", errors.New("feature name not found")
	}

	return result[1], nil
}

func isBranchFeature(gitStatus string) bool {
	re := regexp.MustCompile(`On branch feature/.*`)

	return re.MatchString(gitStatus)
}

func isBranchRelease(gitStatus string) bool {
	re := regexp.MustCompile(`On branch release/v(\d+)\.(\d+)\.(\d+)`)

	return re.MatchString(gitStatus)
}

func isBranchHotfix(gitStatus string) bool {
	re := regexp.MustCompile(`On branch hotfix/v(\d+)\.(\d+)\.(\d+)`)

	return re.MatchString(gitStatus)
}

func getLastVersion(branches string) (major int, minor int, patch int, err error) {
	re := regexp.MustCompile(`/v(\d+)\.(\d+)\.(\d+)`)
	versions := re.FindAllString(branches, -1)
	count := len(versions)
	if count == 0 {
		return 0, 0, 0, errors.New("version not found")
	}

	lastVersion := versions[count-1]

	re = regexp.MustCompile(`/v(\d+)\.(\d+)\.(\d+)`)

	numbers := re.FindStringSubmatch(lastVersion)

	if len(numbers) < 2 {
		return 0, 0, 0, errors.New("version not found")
	}

	major, _ = strconv.Atoi(numbers[1])
	minor, _ = strconv.Atoi(numbers[2])
	patch, _ = strconv.Atoi(numbers[3])

	return major, minor, patch, nil
}
