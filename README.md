<p align="center">
  <img alt="GitHub Logo" src="docs/github_logo.png" height="140" />
  <h3 align="center">Go Git</h3>
  <p align="center">Automatically tag your repository!</p>
  <p align="center">
    <a href="https://github.com/oshalygin/go-tag/releases/latest"><img alt="Release" src="https://img.shields.io/github/release/oshalygin/go-tag.svg?style=flat-square"></a>
    <a href="https://travis-ci.org/oshalygin/go-tag"><img alt="Travis" src="https://travis-ci.org/oshalygin/go-tag.svg?branch=master"></a>
    <a href="/LICENSE.md"><img alt="Software License" src="https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square"></a>
    <a href="https://coveralls.io/github/oshalygin/go-tag?branch=master"><img alt="Coveralls" src="https://coveralls.io/repos/github/oshalygin/go-tag/badge.svg?branch=master"></a>
    <a href="https://codeclimate.com/repos/59bede4e2bfc96025600026b/feed"><img alt="Code Climate Issue Count" src="https://codeclimate.com/repos/59bede4e2bfc96025600026b/badges/d8e88772201d137ea8b7/issue_count.svg"></a>
    <a href="https://goreportcard.com/report/github.com/oshalygin/go-tag"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/oshalygin/go-tag"></a>
    <a href="https://godoc.org/github.com/oshalygin/go-tag"><img src="https://godoc.org/github.com/oshalygin/go-tag?status.svg" alt="GoDoc"></a>
  </p>
</p>

# Introduction

This is a simple and straightforward CLI utility that automatically tags your repository.

## Motivation

Commonly, you'll want to automate your tagging process.  As you release features, you'll add a version somewhere within your application, for Node.js projects this is `package.json`.  This whole process is time consuming and repetitive, why not let a utility handle this automatically?

* Ensuring your repository is properly tagged based on release versions

## Currently Supported Frameworks

This is the current state of the utility and what it supports.

| Framework   | File Read         |   Status     |   Examples     |
|-------------|-------------------|--------------|----------------|
| Generic     | **VERSION**       | In Progress  | [VERSION](test-files/VERSION) |
| Node.js     | **package.json**  | In Progress  | [package.json](test-files/package.json) |

## Requirements

Without bloating the utility with a ton of _required_ options, the `go-tag` utility depends on certain aspects being present in each of your git repositories.

* You are authorized to push against the repository.
* All of the credentials are properly configured in your shell.
* You have one of the files described in the [Currently Supported Frameworks](#currently-supported-frameworks).  This file lives in the root of where the utility is ran.

### Installation

```bash
go get -u github.com/oshalygin/go-tag
```

### Usage

```bash
# Call the utility and enjoy the automation!

go-tag

```

### Command Line Arguments

None, this utility is called within the current git repository and it will make the rest happen!

### Limitations

This utility depends on _you_ having the right credentials set.

Additionally, this utility only works with Node.js projects(for now), namely projects with a `package.json` file.  The reason for this is because the application looks into this configuration file to determine what tag version to set.

In the future, tagging will extend to other languages and frameworks.

# License

[MIT](LICENSE)
