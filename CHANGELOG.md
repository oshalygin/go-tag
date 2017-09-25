# CHANGELOG

## 1.0.0 (September 24, 2017)

- First release:
  - Node.js and Generic `VERSION` files are supported out of the gate.
  - Users of the application should reference the documentation in the README to get started.
- Updates to the Git Service to Push and Tag in a single call, to be used by consumers.

## 0.0.12 (September 24, 2017)

- Removal of `main` tests.
  - This was causing a slew of tags getting pushed through tests.  The true integration test will be tested in CI.

## 0.0.11 (September 24, 2017)

- Version Service:  Updated tests to support the default case

## 0.0.10 (September 24, 2017)

- Additional build out of the version service to support Node.js and properly return the version.

## 0.0.9 (September 24, 2017)

- Addition of Version Service
  - Base implementation against the generic `VERSION` file

## 0.0.8 (September 24, 2017)

- Cleaned up Limitations section in README.md to clearly explain the supported frameworks.

## 0.0.7 (September 24, 2017)

- Documentation updates to reflect examples of supported frameworks.

## 0.0.6 (September 24, 2017)

- Addition of a struct representation of a typical `package.json` file.

## 0.0.5 (September 24, 2017)

- Addition of a Framework resolver which determines the framework based on the files present in the directory.
  - Addition of a framework constants export, which provides the currently supported versions.

## 0.0.4 (September 23, 2017)

- Addition of a git service to handle tagging and pushing the tags

## 0.0.3 (September 23, 2017)

- Addition version flag to determine the current version of `go-tag`

## 0.0.2 (September 23, 2017)

- Addition of test files and a root VERSION file.

## 0.0.1 (September 23, 2017)

- Expansion of README.md to cover additional planned support
