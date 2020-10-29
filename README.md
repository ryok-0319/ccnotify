# ccnotify

ccnotify posts a comment to GitHub from CircleCI job.

## Motivation

There are many cases where developers want to notify some test result, test coverage and etc to GitHub.
Although some test tools provide such function, it's not always the case.
Using this command, you can post a comment from CircleCI job, to the pull-request or the commit the job is associated with.

## Usage

ccnotify is intended to be run on CircleCI job.

### Preerquisite
First of all, ccnotify requires a GitHub token that has a proper access right.

ccnotify automatically reads the environment variable `$GITHUB_ACCESS_TOKEN`.

Note this behavior can be overridden (described below).

Thus, please set it somehow, e.g.:
```sh
export GITHUB_ACCESS_TOKEN=XXXXX
```

### Basic
ccnotify reads from the standard input, so you do something like the following

```sh
echo foo | ccnotify
```

Then ccnotify posts a comment `foo` to GitHub.

### Configute GitHub token
As mentioned above, ccnotify automatically reads `$GITHUB_ACCESS_TOKEN`.

However, you can also give a token using a flag.

```sh
echo foo | ccnotify --token XXXXX

# This also works
echo foo | ccnotify -t XXXXX
```



