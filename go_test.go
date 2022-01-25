package main

import (
	"fmt"
	"testing"
)

func TestBuildNameBranchFromVersion(t *testing.T) {
	got := buildNameBranchFromVersion(0, 1, 0)
	if got != "v0.1.0" {
		t.Errorf("buildNameBranchFromVersion(0,1,0) = %s; want v0.1.0", got)
	}
}

func TestGetLastVersion(t *testing.T) {
	//status, _ := gitStatus()
	//println(status)

	tags := "2c482b1eb49403b7cd75d4445eb4baffec2d4ecd        refs/tags/1.2.0\n190db65c604bc993036c7d054c29fbd5aac23043        refs/tags/1.2.0^{}\n9501436276e390b4a83f1356ea4f917ec3c10164        refs/tags/v0.1\n62f63696a533a49a5c3361c7c4ac1eea7f9ba28a        refs/tags/v0.1^{}\n08eda95b0f1a3c93542a8d1a5924c8e8b57f0bb2        refs/tags/v1.2.1\n5e35b8da63f26bb4865713d0cd25d678cce9e1c0        refs/tags/v1.2.1^{}\n29832a86d144f6f37b0da40633accf20e5af43d2        refs/tags/v1.2.2\n91479c2ad667a52268ce64033a4ab72863deff78        refs/tags/v1.2.2^{}\n\n"

	tagsTwo := "2c482b1eb49403b7cd75d4445eb4baffec2d4ecd        refs/tags/1.2.0\n62f63696a533a49a5c3361c7c4ac1eea7f9ba28a        refs/tags/v0.1^{}\n29832a86d144f6f37b0da40633accf20e5af43d2        refs/tags/v1.2.2\n\n91479c2ad667a52268ce64033a4ab72863deff78        refs/tags/v1.2.2^{}\n08eda95b0f1a3c93542a8d1a5924c8e8b57f0bb2        refs/tags/v1.2.1\n5e35b8da63f26bb4865713d0cd25d678cce9e1c0        refs/tags/v1.2.1^{}\n190db65c604bc993036c7d054c29fbd5aac23043        refs/tags/1.2.0^{}\n9501436276e390b4a83f1356ea4f917ec3c10164        refs/tags/v0.1"

	statusWrong := "On branch develop\nYour branch is up to date with 'origin/develop'.\n\nUntracked files:\n  (use \"git add <file>...\" to include in what will be committed)\n\tcommands.go\n\tgo.mod\n\tgo_test.go\n\tutils.go\n\nnothing added to commit but untracked files present (use \"git add\" to track)\n"

	statusTrue := "On branch hotfix/v1.2.3\nUntracked files:\n  (use \"git add <file>...\" to include in what will be committed)\n        commands.go\n        go.mod\n        go_test.go\n        main.go\n        readme.md\n        utils.go\n        vd\n\nnothing added to commit but untracked files present (use \"git add\" to track)\n\n"

	// test tags
	major, minor, patch, _ := getLastVersion(tags)

	got := fmt.Sprintf("v%d.%d.%d", major, minor, patch)

	if got != "v1.2.2" {
		t.Errorf("buildNameBranchFromVersion(string) = %s; want v1.2.2", got)
	}

	// test tagsTwo
	major, minor, patch, _ = getLastVersion(tagsTwo)

	got = fmt.Sprintf("v%d.%d.%d", major, minor, patch)

	if got != "v1.2.2" {
		t.Errorf("buildNameBranchFromVersion(string) = %s; want v1.2.2", got)
	}

	// test statusTrue
	major, minor, patch, _ = getLastVersion(statusTrue)

	got = fmt.Sprintf("v%d.%d.%d", major, minor, patch)

	if got != "v1.2.3" {
		t.Errorf("buildNameBranchFromVersion(string) = %s; want v1.2.3", got)
	}

	// test statusWrong
	major, minor, patch, _ = getLastVersion(statusWrong)

	got = fmt.Sprintf("v%d.%d.%d", major, minor, patch)

	if got != "v0.0.0" {
		t.Errorf("buildNameBranchFromVersion(sting) = %s; want v0.0.0", got)
	}
}
