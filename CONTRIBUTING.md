# Contributing

## Summary

* [AngularJS Commit Message Format](https://github.com/angular/angular.js/blob/master/CONTRIBUTING.md#commit-message-format)
* Format with `go fmt`
* Limit all lines to a maximum of 79 characters as much as possible
* Write tests and comments
* Pay attention to the test covarage

## Steps

1. Fork (https://github.com/suzuki-shunsuke/go-graylog/fork)
2. Clone the forked repository
3. Install npm dependencies
4. Create a feature branch
5. Make your change
6. Test
7. Commit your change
8. Rebase your local changes against the master branch
9. Create a new Pull Request

## Requirements

* [npm](https://www.npmjs.com/): to validate a commit message and generate the Change Log
* [Golang](https://golang.org/)
* [dep](https://golang.github.io/dep/)

We use some node modules to validate commit messages and generate Change Log.

Install node modules by

```
$ npm
```

## Test

```
$ make test
```

Confirm coverage at the web browser.

```
$ make cover
```

## Commit Message Format

The commit message format of this project conforms to the [AngularJS Commit Message Format](https://github.com/angular/angular.js/blob/master/CONTRIBUTING.md#commit-message-format).
By conforming its format, we can generate the [Change Log](CHANGELOG.md) and conform the [semantic versioning](http://semver.org/) automatically by [standard-version](https://www.npmjs.com/package/standard-version).
We validate the commit message with git's `commit-msg` hook using [validate-commit-msg](https://www.npmjs.com/package/validate-commit-msg) and [husky](https://www.npmjs.com/package/husky).

## Release

We generate the [Change Log](CHANGELOG.md) by [standard-version](https://www.npmjs.com/package/standard-version).
After merge your PR at the master branch,
The author ([suzuki-shunsuke](https://github.com/suzuki-shunsuke)) will generate the release tag.

```
$ npm run release
$ git push origin master --follow-tags
```

## Coding Guide

* Format with `go fmt`
* Limit all lines to a maximum of 79 characters as much as possible

### Variable name Guide

base | good | bad
--- | --- | ---
id | Id | ID
user_name | UserName | -
username | Username | UserName
