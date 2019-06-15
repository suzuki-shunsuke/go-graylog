# Contributing

## Requirements

* [npm](https://www.npmjs.com/): to validate a commit message and generate the Change Log
* [Golang](https://golang.org/)
* [golangci-lint](https://github.com/golangci/golangci-lint)

```console
$ npm i
```

## Test

```console
$ npm t
```

## Commit Message Format

The commit message format of this project conforms to the [AngularJS Commit Message Format](https://github.com/angular/angular.js/blob/master/CONTRIBUTING.md#commit-message-format).
By conforming its format, we can generate the [Change Log](CHANGELOG.md) and conform the [semantic versioning](http://semver.org/) automatically by [standard-version](https://www.npmjs.com/package/standard-version).
We validate the commit message with git's `commit-msg` hook using [commitlint](http://marionebl.github.io/commitlint/#/) and [husky](https://www.npmjs.com/package/husky).

## Coding Guide

* https://github.com/golang/go/wiki/CodeReviewComments

```console
$ npm run lint
```

## docker-compose.yml

To run graylog using docker for development, we prepare the template of `docker-compose.yml`.

```console
$ cp docker-compose.yml.tmpl docker-compose.yml
$ docker-compose up -d
```

## env.sh

To set environment variables for development, we prepare the template of `setenv.sh` .

```console
$ cp env.sh.tmpl env.sh
```

## Develop terraform provider

* https://www.terraform.io/docs/plugins/provider.html 
* https://www.terraform.io/guides/writing-custom-terraform-providers.html
* https://godoc.org/github.com/hashicorp/terraform/helper/schema
* https://godoc.org/github.com/hashicorp/terraform/helper/resource#TestCase
