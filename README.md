# vd


The simple Git Flow util for help you use git flow as this https://danielkummer.github.io/git-flow-cheatsheet/index.html.
Integration with Youtrack, Jira (set url and title to commit from task manager).
Automate add version tag to master branch.



## Install

### Linux
```bash
git clone https://github.com/darmshot/vd && cd vd
```

#### Go
*(required go 1.18)*

```bash
go build -ldflags "-X main.Version=$(git describe --tags)"
```

#### Docker

```bash
chmod 755 build.sh
```

```bash
./build.sh
```

#### Requires
* Require .env file in project folder.
* Master branch should be named "master".
* Require develop branch.
* Project should be placed in remote repository

## Available commands:

### Feature Start
 
Create branch from develop branch. Make feature branch `feature/task-1`.

```bash
vd fs --name <task-name>
```

**Flags:**
```
  -n, --name string   Feature name
```

### Feature Finish
Merge feature branch into develop and remove current feature branch.

```bash
vd ff
```

### Commit

This make speed commit with url of task and title of task.

Just run simple command `vd c` in feature branch or `vd c -t <number_task>` for any other branch.

Allowed make commit with several number of task like `vd c -t 001_002`.

Require setup TASK_DRIVER COMMIT_MESSAGE_PREFIX in .env.


```bash
vd c
```
_(only feature branch)_

**Flags:**
```
  -m, --message string   commit message
  -t, --task string      number task of numbers like 100 or 100_101
```

For example, if current git branch is feature/001. This automates create commit.

**_Result:_**
```
http://track-manager.url/path/TASK-001 Some title of task
```

**_Or for any other branch:_**

```bash
vd c -t 001
```

**_Result:_**
```
http://track-manager.url/path/TASK-001 Some title of task
```

Also, you can put message into description with flag: `--message`.

```bash
vd c -t 001 --message "Some additional message"
```

_Result:_
```
http://track-manager.url/path/001 Some title of task
Some additional message
```


### Release start

Create release branch `release/v0.1` and push to remote git repository.
First release should start with flag: `--first`.
By default, upper minor version. If you need up major use flag: `--major`.
Number tag will be got from last bigger version tag from you remote git repository. 

```bash
vd rs
```

### Release Finish
Finish you release and merge with master branch. Marks master branch tag.
First release should finish with flag: `--first`.

```bash
vd rf
```

### Hotfix Start
Create branch `hotfix/vX.X.X` from master (only in local repository).

```bash
vd hs
```


### Hotfix Finish
Merge hotfix branch into master and develop. Marks master branch tag.

```bash
vd hf
```
