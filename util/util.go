package util

import (
	"errors"
	"github.com/darmshot/vd/api/jira"
	"github.com/darmshot/vd/api/youtrack"
	"github.com/darmshot/vd/config"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func IsBranchFeature(gitStatus string) bool {
	re := regexp.MustCompile(`On branch feature/.*`)

	return re.MatchString(gitStatus)
}

func IsBranchRelease(gitStatus string) bool {
	re := regexp.MustCompile(`On branch release/v(\d+)\.(\d+)\.(\d+)`)

	return re.MatchString(gitStatus)
}

func IsBranchHotfix(gitStatus string) bool {
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

func GetLastVersion(branches string) (major int, minor int, patch int, err error) {
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

		majorPrepare := major * 100000

		minorPrepare := minor * 10000

		number := majorPrepare + minorPrepare + patch

		versionItems = append(versionItems, versionItem{versions[i], number, version{major, minor, patch}})
	}

	sort.Slice(versionItems, func(i, j int) bool { return versionItems[i].Sort < versionItems[j].Sort })

	lastVersion := versionItems[count-1]

	return lastVersion.Version.Major, lastVersion.Version.Minor, lastVersion.Version.Patch, nil
}

func GetFeatureName(gitStatus string) (string, error) {
	re := regexp.MustCompile(`On branch feature/(.*)`)

	result := re.FindStringSubmatch(gitStatus)
	if len(result) < 2 {
		return "", errors.New("feature name not found")
	}

	return result[1], nil
}

func GetNameBranchFromVersion(major int, minor int, patch int) string {
	return "v" + strconv.Itoa(major) + "." + strconv.Itoa(minor) + "." + strconv.Itoa(patch)
}

func GetNumbersFromName(name string) ([]int, error) {
	var numbers []int

	re := regexp.MustCompile(`\d+`)
	numbersMatch := re.FindAllString(name, -1)
	count := len(numbersMatch)

	if count == 0 {
		return numbers, errors.New("not found numbers")
	}

	for i := 0; i < count; i++ {
		convertNumber, _ := strconv.Atoi(numbersMatch[i])
		numbers = append(numbers, convertNumber)
	}

	return numbers, nil
}

func GetLastPartFromUrl(url string) string {

	parts := strings.Split(url, "/")
	end := parts[len(parts)-1]

	return end
}

func GetCommitMessage(name string, numbers []int, message string) string {
	var summary string
	var prefix string
	var summaryList []string
	var issueList []string
	var messageBlockList []string

	count := len(numbers)

	for i := 0; i < count; i++ {
		if config.CommitMessagePrefix == "" {
			prefix = name
		} else {
			prefix = config.CommitMessagePrefix + strconv.Itoa(numbers[i])
		}

		if config.TaskDriver == "jira" {
			summary = " " + jira.GetIssueSummary(GetLastPartFromUrl(prefix))
		} else if config.TaskDriver == "youtrack" {
			summary = " " + youtrack.GetIssueSummary(GetLastPartFromUrl(prefix))
		}

		if len(summary) != 0 {
			summaryList = append(summaryList, summary)
		}

		if len(prefix) != 0 {
			issueList = append(issueList, prefix)
		}
	}

	if len(message) != 0 {
		messageBlockList = append(messageBlockList, message)
	}

	if len(summaryList) != 0 {
		messageBlockList = append(messageBlockList, strings.Join(summaryList, "\n"))
	}

	if len(issueList) != 0 {
		messageBlockList = append(messageBlockList, "Closes: "+strings.Join(issueList, ", "))
	}

	result := strings.Join(messageBlockList, "\n\n")

	return result
}
