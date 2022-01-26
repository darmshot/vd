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

	tagsThree := "fefc973ae33bbc5733bbade6cea3af985ba23a0c        refs/tags/v0.1.0\nec81481fb740329e70065874cd2c3783de08bf35        refs/tags/v0.1.0^{}\n78e78932613f36238499256d01b59bb552f4abfe        refs/tags/v0.1.1\n668d185f47de0fbcc6ee85646329b646d51931f5        refs/tags/v0.1.1^{}\nfb4487415898e86db252323f784030f520831139        refs/tags/v0.1.2\na060986cf5055f31acb8c3d2c447874c90b77c18        refs/tags/v0.1.2^{}\nffb3fc8679bd3e83497eed9e8eddd3398b3597d1        refs/tags/v0.1.3\nd0f89dca907ae47c07af8ca7b82cb2a15e04e61b        refs/tags/v0.1.3^{}\nd42f680746c7283eafef287b21c659b4946d2aef        refs/tags/v0.1.4\n682a021c55bbe6dc27cf68ccb48f5bf29c6df314        refs/tags/v0.1.4^{}\na8be5af2a9ae9462d20d3f25e42685bc71d2d330        refs/tags/v0.1.5\nf9984f3ee1f21083736166cd3ab0206145d71a18        refs/tags/v0.1.5^{}\nedd7753e8ccf67fea66f87d82af4bea5535aad15        refs/tags/v0.1.6\nf9984f3ee1f21083736166cd3ab0206145d71a18        refs/tags/v0.1.6^{}\n447e6ae9d888d6cb6200c8d130033ffa47a85fcc        refs/tags/v0.2.0\nd90e2ee41528c71ef8f75fd29d0fa0241e6a1fc8        refs/tags/v0.2.0^{}\n6d22fc8c3f417c1606fec725149a2d87634d22c7        refs/tags/v0.2.1\nd90e2ee41528c71ef8f75fd29d0fa0241e6a1fc8        refs/tags/v0.2.1^{}\n61888c1ce0283a743664412d0b9dca6608d426f4        refs/tags/v0.2.10\nb2ae5c2ba9a17b27df5920a5f180a9e6aa944d19        refs/tags/v0.2.10^{}\nf378a7b3e20f5fc7f87d30929350af7958c03b7b        refs/tags/v0.2.2\n3b869d75d0c7a30ebcbcb584ac82df8e636abb9a        refs/tags/v0.2.2^{}\n6fb6b9268d362050defd1c2f25613b70f03c9033        refs/tags/v0.2.3\n74aa2cc2d6a20ea2303b2a82c6809b17c7b00475        refs/tags/v0.2.3^{}\n5f5505602a889a7d662024f1a724bde0f9a8a51e        refs/tags/v0.2.4\n74aa2cc2d6a20ea2303b2a82c6809b17c7b00475        refs/tags/v0.2.4^{}\n813b6e7ef67d2d80c437205b16a80deac25ee884        refs/tags/v0.2.5\nd680a9b6d9d2682bedea62ade4cd906f4764e1ba        refs/tags/v0.2.5^{}\ndf542a005905cd4f2285d3c124c11f44a49679ae        refs/tags/v0.2.6\n327d0fa6236515ae163e66840e9e5611ceb17bc8        refs/tags/v0.2.6^{}\n8ef7961d4760210a424fbf19800338173a1bbcf5        refs/tags/v0.2.7\n7a07c5748014585c13b8fd10e590d65c0cd45ba1        refs/tags/v0.2.7^{}\nb29be4c596cfb96daa3550c2f1c57aede64277d0        refs/tags/v0.2.8\n3a23491d7f1732bfdd8ade27b71fd76cfbbb0471        refs/tags/v0.2.8^{}\nd4f072cd9a6716f3fb4635c6d1ccbac7308b9272        refs/tags/v0.2.9\n3b5f5c7b84a53e35041f00b5c5770bb309c7d46a        refs/tags/v0.2.9^{}"

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

	// test tagsThree
	major, minor, patch, _ = getLastVersion(tagsThree)

	got = fmt.Sprintf("v%d.%d.%d", major, minor, patch)

	if got != "v0.2.10" {
		t.Errorf("buildNameBranchFromVersion(string) = %s; want v0.2.10", got)
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
