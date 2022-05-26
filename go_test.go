package main

import (
	"fmt"
	"testing"
)

func TestBuildNameBranchFromVersion(t *testing.T) {
	got := getNameBranchFromVersion(0, 1, 0)
	if got != "v0.1.0" {
		t.Errorf("getNameBranchFromVersion(0,1,0) = %s; want v0.1.0", got)
	}
}

func TestGetLastVersion(t *testing.T) {
	//status, _ := gitStatus()
	//println(status)

	tags := "2c482b1eb49403b7cd75d4445eb4baffec2d4ecd        refs/tags/1.2.0\n190db65c604bc993036c7d054c29fbd5aac23043        refs/tags/1.2.0^{}\n9501436276e390b4a83f1356ea4f917ec3c10164        refs/tags/v0.1\n62f63696a533a49a5c3361c7c4ac1eea7f9ba28a        refs/tags/v0.1^{}\n08eda95b0f1a3c93542a8d1a5924c8e8b57f0bb2        refs/tags/v1.2.1\n5e35b8da63f26bb4865713d0cd25d678cce9e1c0        refs/tags/v1.2.1^{}\n29832a86d144f6f37b0da40633accf20e5af43d2        refs/tags/v1.2.2\n91479c2ad667a52268ce64033a4ab72863deff78        refs/tags/v1.2.2^{}\n\n"

	tagsTwo := "2c482b1eb49403b7cd75d4445eb4baffec2d4ecd        refs/tags/1.2.0\n62f63696a533a49a5c3361c7c4ac1eea7f9ba28a        refs/tags/v0.1^{}\n29832a86d144f6f37b0da40633accf20e5af43d2        refs/tags/v1.2.2\n\n91479c2ad667a52268ce64033a4ab72863deff78        refs/tags/v1.2.2^{}\n08eda95b0f1a3c93542a8d1a5924c8e8b57f0bb2        refs/tags/v1.2.1\n5e35b8da63f26bb4865713d0cd25d678cce9e1c0        refs/tags/v1.2.1^{}\n190db65c604bc993036c7d054c29fbd5aac23043        refs/tags/1.2.0^{}\n9501436276e390b4a83f1356ea4f917ec3c10164        refs/tags/v0.1"

	tagsThree := "fefc973ae33bbc5733bbade6cea3af985ba23a0c        refs/tags/v0.1.0\nec81481fb740329e70065874cd2c3783de08bf35        refs/tags/v0.1.0^{}\n78e78932613f36238499256d01b59bb552f4abfe        refs/tags/v0.1.1\n668d185f47de0fbcc6ee85646329b646d51931f5        refs/tags/v0.1.1^{}\nfb4487415898e86db252323f784030f520831139        refs/tags/v0.1.2\na060986cf5055f31acb8c3d2c447874c90b77c18        refs/tags/v0.1.2^{}\nffb3fc8679bd3e83497eed9e8eddd3398b3597d1        refs/tags/v0.1.3\nd0f89dca907ae47c07af8ca7b82cb2a15e04e61b        refs/tags/v0.1.3^{}\nd42f680746c7283eafef287b21c659b4946d2aef        refs/tags/v0.1.4\n682a021c55bbe6dc27cf68ccb48f5bf29c6df314        refs/tags/v0.1.4^{}\na8be5af2a9ae9462d20d3f25e42685bc71d2d330        refs/tags/v0.1.5\nf9984f3ee1f21083736166cd3ab0206145d71a18        refs/tags/v0.1.5^{}\nedd7753e8ccf67fea66f87d82af4bea5535aad15        refs/tags/v0.1.6\nf9984f3ee1f21083736166cd3ab0206145d71a18        refs/tags/v0.1.6^{}\n447e6ae9d888d6cb6200c8d130033ffa47a85fcc        refs/tags/v0.2.0\nd90e2ee41528c71ef8f75fd29d0fa0241e6a1fc8        refs/tags/v0.2.0^{}\n6d22fc8c3f417c1606fec725149a2d87634d22c7        refs/tags/v0.2.1\nd90e2ee41528c71ef8f75fd29d0fa0241e6a1fc8        refs/tags/v0.2.1^{}\n61888c1ce0283a743664412d0b9dca6608d426f4        refs/tags/v0.2.10\nb2ae5c2ba9a17b27df5920a5f180a9e6aa944d19        refs/tags/v0.2.10^{}\nf378a7b3e20f5fc7f87d30929350af7958c03b7b        refs/tags/v0.2.2\n3b869d75d0c7a30ebcbcb584ac82df8e636abb9a        refs/tags/v0.2.2^{}\n6fb6b9268d362050defd1c2f25613b70f03c9033        refs/tags/v0.2.3\n74aa2cc2d6a20ea2303b2a82c6809b17c7b00475        refs/tags/v0.2.3^{}\n5f5505602a889a7d662024f1a724bde0f9a8a51e        refs/tags/v0.2.4\n74aa2cc2d6a20ea2303b2a82c6809b17c7b00475        refs/tags/v0.2.4^{}\n813b6e7ef67d2d80c437205b16a80deac25ee884        refs/tags/v0.2.5\nd680a9b6d9d2682bedea62ade4cd906f4764e1ba        refs/tags/v0.2.5^{}\ndf542a005905cd4f2285d3c124c11f44a49679ae        refs/tags/v0.2.6\n327d0fa6236515ae163e66840e9e5611ceb17bc8        refs/tags/v0.2.6^{}\n8ef7961d4760210a424fbf19800338173a1bbcf5        refs/tags/v0.2.7\n7a07c5748014585c13b8fd10e590d65c0cd45ba1        refs/tags/v0.2.7^{}\nb29be4c596cfb96daa3550c2f1c57aede64277d0        refs/tags/v0.2.8\n3a23491d7f1732bfdd8ade27b71fd76cfbbb0471        refs/tags/v0.2.8^{}\nd4f072cd9a6716f3fb4635c6d1ccbac7308b9272        refs/tags/v0.2.9\n3b5f5c7b84a53e35041f00b5c5770bb309c7d46a        refs/tags/v0.2.9^{}"

	tagsFour := "9dde9e0a3964729b4b541e5d47a0cce33ccd71fd        refs/tags/v0.1.0\n6deaa4f28d8abd64419c8e928e3d722a61c790f8        refs/tags/v0.1.0^{}\n5e0ae9048de39b1c2e981aa201524b84edbe38ea        refs/tags/v0.1.1\nfdb27506bf980f9e6f38c8c3ae488fa3e1f60988        refs/tags/v0.1.1^{}\n33c56ff4d82535cc21c9de8c89af80ecad1b1f91        refs/tags/v0.1.2\nbf9c7b461fd9ac5faa798013c88d80a099a75140        refs/tags/v0.1.2^{}\n33757fe33e550dcd310e9115ef78c808222997a6        refs/tags/v0.1.3\nfd6a296563d7a69fe175a0503f2bb064b74a15fb        refs/tags/v0.1.3^{}\n1c77c06e130f6e972189f6ff97be1baeeeecee6b        refs/tags/v0.1.4\nd8cc0dccf54394e3e4aa5eb47a830627fbf8b9c2        refs/tags/v0.1.4^{}\n445372bb1a8f983fe54f29d9e64b76d907aa3af2        refs/tags/v0.1.5\n4f60a1c7d0331dc41f06a6efd649c4cfba53f132        refs/tags/v0.1.5^{}\n53a155cabdcdbd888490d7518541d94d80803828        refs/tags/v0.1.6\n4f60a1c7d0331dc41f06a6efd649c4cfba53f132        refs/tags/v0.1.6^{}\n127dbf0c8b61c1d75f5deed5b5a1f3976a238d46        refs/tags/v0.2.0\n7caabe2a7d2318eb2e6e4baf20e8eee80e66c4c1        refs/tags/v0.2.0^{}\n86d6ec462181ce0643dee5dd82dba64f6b2c1521        refs/tags/v0.2.1\n4c539c975250ace620553fa0bf3ca724e6d1af97        refs/tags/v0.2.1^{}\n30c41375a70168cd6942d521847840da110417bf        refs/tags/v0.2.10\ndf229d262134c70f4372db889d07914727513ea1        refs/tags/v0.2.10^{}\ne932e6a52526c9569c887034d9955c844fde8feb        refs/tags/v0.2.11\n99e241446a20e5ffaff98f3dd7e564c249de0a0e        refs/tags/v0.2.11^{}\n6fa66f30e5b231690df365d03f2d2011c7e5efe5        refs/tags/v0.2.12\n6e6291c49c9cb1f1eadfcf3ec23d2e917569a53b        refs/tags/v0.2.12^{}\n1ec612585fd35887ba9da35bd72dcf7053c2c6bd        refs/tags/v0.2.13\n1ab4e3a537a99cb8f66ae3ec40533cf1d2f06676        refs/tags/v0.2.13^{}\nd01a561320e88f61e09f9a8feedfcce8ab73b944        refs/tags/v0.2.14\nbba47f6e7c1a4a3f3adee2de93aee62adac67287        refs/tags/v0.2.14^{}\n33df2095d87a0703600a3cfd2df8212eeaf95157        refs/tags/v0.2.15\naf1113c0e4aac7fcf6ab26729d407dd7bf87aba4        refs/tags/v0.2.15^{}\nc83601d3b4216faa5b34e71f96e04c5a36303e68        refs/tags/v0.2.16\n77ce241df57cef48d902045bb3dba95c5c0a43fc        refs/tags/v0.2.16^{}\n07277c02c7594f43b39641f20824ce36a7ceaf41        refs/tags/v0.2.17\n000fcf6d8081a08789e9a1b4f3ca4414fe1992f1        refs/tags/v0.2.17^{}\nc39e8896e3e9071dc7017f70bb94deae1bb713a5        refs/tags/v0.2.18\n9cf73cddfaadbbaacd9f5d0a7b9edeb8c3855021        refs/tags/v0.2.18^{}\nf9c8a8b42977915e6a6d281623e3e30806a3ae29        refs/tags/v0.2.19\ndd10e0dca99dee9b5101dd2d4a3a89a65e959607        refs/tags/v0.2.19^{}\n0588072743e2e690d13cc4bf2078a65ba035b5bb        refs/tags/v0.2.2\nfc9c65cfcc82555a9de31e81ac9dfa00d7204739        refs/tags/v0.2.2^{}\nb409906c6fdda1cf8125642e7cd773ab8c6a4538        refs/tags/v0.2.20\nb47c90f0bb9be4ae71aaf209d280750d317d99f3        refs/tags/v0.2.20^{}\n7bd49147c12d50c4e9c1549bb951a40ffb66fe82        refs/tags/v0.2.21\nfb923498ddc09498432df229ccc9b5f075fbd1c1        refs/tags/v0.2.21^{}\nb1b1d0379e69f97767bb43cf56e5adc503359cf7        refs/tags/v0.2.22\n078220789bb57158f2768ad1f9e2125f3192a6c8        refs/tags/v0.2.22^{}\n5b0913f698ff4a2fb15d80973b925fb158f91dcc        refs/tags/v0.2.23\nc54749d2cef6e8cabfc8600d98913f46c5d06b0b        refs/tags/v0.2.23^{}\n2bc2ceffd650eccc8f1c2ef6f270ef260e419e6c        refs/tags/v0.2.24\n2971fd6790bde8ea59e62ca8bc25e83e40a4b4f9        refs/tags/v0.2.24^{}\naa706b465f7fe39687bb2802aed77f99e0499741        refs/tags/v0.2.25\nfbd786b9d710b4a829e774f3927de5b13fbd08d7        refs/tags/v0.2.25^{}\nc4f110f071049a96cc0cece0febce60689be6c92        refs/tags/v0.2.26\nfbd786b9d710b4a829e774f3927de5b13fbd08d7        refs/tags/v0.2.26^{}\na62bd4770aeda3487cdc73da5c97de80a5d8843e        refs/tags/v0.2.3\nb5c293aeca925caf3617096daf817215d855dda0        refs/tags/v0.2.3^{}\n1197586c42851330d5f6c77b0fb03f07c018cf73        refs/tags/v0.2.4\nba449f542e652ffda3262d4c060f32ac6b9177cf        refs/tags/v0.2.4^{}\n8a06ae22ef36fa4d3ca82790265aa8fde45284c5        refs/tags/v0.2.5\n62d45af324ebab2b390f7fc6bc4931fdd045686b        refs/tags/v0.2.5^{}\nf3b71dd49e5168a538c9de764d7482ee0e16d2f9        refs/tags/v0.2.6\n62d45af324ebab2b390f7fc6bc4931fdd045686b        refs/tags/v0.2.6^{}\nd547e9e7f59abd4710f1b45cee82869514f03797        refs/tags/v0.2.7\n433d09187098dc17b744ef481ee816a476a0d003        refs/tags/v0.2.7^{}\nb4ee31dae326b9db3eac2a1abe48cc1e81da1dca        refs/tags/v0.2.8\nc0e6344ea512bda8749a546409f6b692afd27bf4        refs/tags/v0.2.8^{}\n7c096b591ae825f4d9f8a4c8505e3641bdc22195        refs/tags/v0.2.9\n96b382fb69fd56fde0c6b2745947940cdd4bdc4a        refs/tags/v0.2.9^{}\n105bde8504b6585346cc7b56776689c3de0ec192        refs/tags/v0.3.0\nfbd786b9d710b4a829e774f3927de5b13fbd08d7        refs/tags/v0.3.0^{}\n"

	statusWrong := "On branch develop\nYour branch is up to date with 'origin/develop'.\n\nUntracked files:\n  (use \"git add <file>...\" to include in what will be committed)\n\tcommands.go\n\tgo.mod\n\tgo_test.go\n\tutils.go\n\nnothing added to commit but untracked files present (use \"git add\" to track)\n"

	statusTrue := "On branch hotfix/v1.2.3\nUntracked files:\n  (use \"git add <file>...\" to include in what will be committed)\n        commands.go\n        go.mod\n        go_test.go\n        main.go\n        readme.md\n        utils.go\n        vd\n\nnothing added to commit but untracked files present (use \"git add\" to track)\n\n"

	// test tags
	major, minor, patch, _ := getLastVersion(tags)

	got := fmt.Sprintf("v%d.%d.%d", major, minor, patch)

	if got != "v1.2.2" {
		t.Errorf("getNameBranchFromVersion(string) = %s; want v1.2.2", got)
	}

	// test tagsTwo
	major, minor, patch, _ = getLastVersion(tagsTwo)

	got = fmt.Sprintf("v%d.%d.%d", major, minor, patch)

	if got != "v1.2.2" {
		t.Errorf("getNameBranchFromVersion(string) = %s; want v1.2.2", got)
	}

	// test tagsThree
	major, minor, patch, _ = getLastVersion(tagsThree)

	got = fmt.Sprintf("v%d.%d.%d", major, minor, patch)

	if got != "v0.2.10" {
		t.Errorf("getNameBranchFromVersion(string) = %s; want v0.2.10", got)
	}

	// test tagsFour
	major, minor, patch, _ = getLastVersion(tagsFour)

	got = fmt.Sprintf("v%d.%d.%d", major, minor, patch)

	if got != "v0.3.0" {
		t.Errorf("getNameBranchFromVersion(string) = %s; want v0.3.0", got)
	}

	// test statusTrue
	major, minor, patch, _ = getLastVersion(statusTrue)

	got = fmt.Sprintf("v%d.%d.%d", major, minor, patch)

	if got != "v1.2.3" {
		t.Errorf("getNameBranchFromVersion(string) = %s; want v1.2.3", got)
	}

	// test statusWrong
	major, minor, patch, _ = getLastVersion(statusWrong)

	got = fmt.Sprintf("v%d.%d.%d", major, minor, patch)

	if got != "v0.0.0" {
		t.Errorf("getNameBranchFromVersion(sting) = %s; want v0.0.0", got)
	}
}

func TestGetNumbersFromFeatureName(t *testing.T) {
	name := "CREOS-333_4353_355"
	numbers, _ := getNumbersFromName(name)
	fmt.Println(numbers)
	if len(numbers) == 0 {
		t.Errorf("getNumbersFromName(string) is empty; want fill number")
	}

}
