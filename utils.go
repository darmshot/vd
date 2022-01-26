package main

import (
	"errors"
	"regexp"
	"sort"
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

type version struct {
	Major int
	Minor int
	Patch int
}

type versionItem struct {
	Text    string
	Sort    int
	Version version
}

func getLastVersion(branches string) (major int, minor int, patch int, err error) {
	var versionItems []versionItem

	re := regexp.MustCompile(`/v(\d+)\.(\d+)\.(\d+)`)
	versions := re.FindAllString(branches, -1)
	count := len(versions)
	if count == 0 {
		return 0, 0, 0, errors.New("version not found")
	}

	for i := 0; i < count; i++ {
		re = regexp.MustCompile(`/v(\d+)\.(\d+)\.(\d+)`)

		numbers := re.FindStringSubmatch(versions[i])

		if len(numbers) < 2 {
			return 0, 0, 0, errors.New("version not found")
		}

		major, _ := strconv.Atoi(numbers[1])
		minor, _ := strconv.Atoi(numbers[2])
		patch, _ := strconv.Atoi(numbers[3])

		number, _ := strconv.Atoi(numbers[1] + numbers[2] + numbers[3])

		versionItems = append(versionItems, versionItem{versions[i], number, version{major, minor, patch}})
	}

	sort.Slice(versionItems, func(i, j int) bool { return versionItems[i].Sort < versionItems[j].Sort })

	lastVersion := versionItems[count-1]

	return lastVersion.Version.Major, lastVersion.Version.Minor, lastVersion.Version.Patch, nil
}
