# Contributing

## Summary

* https://github.com/golang/go/wiki/CodeReviewComments
* [AngularJS Commit Message Format](https://github.com/angular/angular.js/blob/master/CONTRIBUTING.md#commit-message-format)
* Format with `npm run fmt`
* Write tests and comments

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
$ npm t
```

## Commit Message Format

The commit message format of this project conforms to the [AngularJS Commit Message Format](https://github.com/angular/angular.js/blob/master/CONTRIBUTING.md#commit-message-format).
By conforming its format, we can generate the [Change Log](CHANGELOG.md) and conform the [semantic versioning](http://semver.org/) automatically by [standard-version](https://www.npmjs.com/package/standard-version).
We validate the commit message with git's `commit-msg` hook using [commitlint](http://marionebl.github.io/commitlint/#/) and [husky](https://www.npmjs.com/package/husky).

## Release

We generate the [Change Log](CHANGELOG.md) by [standard-version](https://www.npmjs.com/package/standard-version).
After merge your PR at the master branch,
The author ([suzuki-shunsuke](https://github.com/suzuki-shunsuke)) will generate the release tag.

```
$ npm run release
$ git push origin master --follow-tags
```

## Coding Guide

* https://github.com/golang/go/wiki/CodeReviewComments
* Format with `npm run fmt`

## docker-compose.yml

```
$ cp docker-compose.yml.tmpl docker-compose.yml
$ docker-compose up -d
```

## env.sh

```
$ cp env.sh.tmpl env.sh
```

## Develop terraform provider

* https://www.terraform.io/docs/plugins/provider.html 
* https://www.terraform.io/guides/writing-custom-terraform-providers.html
* https://godoc.org/github.com/hashicorp/terraform/helper/schema
* https://godoc.org/github.com/hashicorp/terraform/helper/resource#TestCase

## Tools

These tools are not used at CI due to false positive, but we recommend use them at local before PR.

* [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck)
