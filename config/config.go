package config

import (
	"github.com/darmshot/vd/functions"
)

var CommitMessagePrefix = functions.Env("COMMIT_MESSAGE_PREFIX")

var TaskDriver = functions.Env("TASK_DRIVER")

var JiraKey = functions.Env("JIRA_KEY")

var YoutrackBaseUrl = functions.Env("YOUTRACK_BASE_URL")
var YoutrackKey = functions.Env("YOUTRACK_KEY")
