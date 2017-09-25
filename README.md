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

#### Golang Installed

There are numerous ways to consume this library.  If you have Golang installed on your machine, you can use the instructions in [Installation](#installation) followed by the following command:

```bash
# Call the utility and enjoy the automation!

go-tag
```

The following is an example of how you would consume this from a CI environment such as Travis.

#### From CI with Bash

```bash

readonly utility_version="1.0.0" # pick the release version you want

# Select the architecture and the OS
# Travis CI is generally configured for the following
# Operating System: linux
# Architecture: amd64
# For more information check out the #select-your-binary section
readonly utility_binary="linux_amd64_go-tag"

# If using wget, provide an output path of where the binary will live
readonly utility="${PWD}/dist/go-tag"

# Retrieve the latest binary
wget https://github.com/oshalygin/go-tag/releases/download/$utility_version/$utility_binary \
      -O $utility

# Update access permissions of the utility
chmod +x $utility

# Usage
sudo $utility
```

### Command Line Arguments

None, this utility is called within the current git repository and it will make the rest happen!

### Limitations

This utility depends on _you_ having the right credentials set.

Additionally, this utility only works with Node.js projects(for now), namely projects with a `package.json` file _AND_ generic version files. Please reference the [Currently Supported Frameworks](#currently-supported-frameworks) section for supported frameworks.

More frameworks are coming, so if you see something that isn't supported, post an Issue!

#### Select The Right Binary

If you're pulling down the binary manually, you'll need to match it to your OS and the instruction set.  Thankfully, if you pulled it via Go, the proper `$GOOS`(Operating System) and `$GOARCH` are set for you.

For exciting bathroom reading checkout the Go documentation for compilation architecture descriptions and overriding defaults.
[Go Documentation](https://golang.org/doc/install/source#environment)

# License

[MIT](LICENSE)
