# Change Log

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

<a name="0.9.0-0"></a>
# [0.9.0-0](https://github.com/suzuki-shunsuke/go-graylog/compare/v0.8.0...v0.9.0-0) (2018-11-05)


### Bug Fixes

* fix dashboard widget API ([42314db](https://github.com/suzuki-shunsuke/go-graylog/commit/42314db))


### BREAKING CHANGES

* change Widget.CacheTime's type



<a name="0.8.0"></a>
# [0.8.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v0.7.0...v0.8.0) (2018-11-03)


### Features

* support create and delete and get dashboard widget API ([38670a6](https://github.com/suzuki-shunsuke/go-graylog/commit/38670a6))



<a name="0.7.0"></a>
# [0.7.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v0.6.1...v0.7.0) (2018-09-09)


### Bug Fixes

* fix collector configuration client and mockserver ([e70a6ae](https://github.com/suzuki-shunsuke/go-graylog/commit/e70a6ae))
* give up support of some terraform resources ([b84d574](https://github.com/suzuki-shunsuke/go-graylog/commit/b84d574))
* make collector configuration output properties interface ([cd7e69a](https://github.com/suzuki-shunsuke/go-graylog/commit/cd7e69a))


### Features

* **client:** support Collector Configuration APIs ([b236665](https://github.com/suzuki-shunsuke/go-graylog/commit/b236665))
* **mockserver:** support collector configuration APIs ([4182322](https://github.com/suzuki-shunsuke/go-graylog/commit/4182322))
* **mockserver:** support collector configuration input APIs ([cfded83](https://github.com/suzuki-shunsuke/go-graylog/commit/cfded83))
* **mockserver:** support collector configuration output APIs ([be0d622](https://github.com/suzuki-shunsuke/go-graylog/commit/be0d622))
* **mockserver:** support collector configuration snippet APIs ([9419321](https://github.com/suzuki-shunsuke/go-graylog/commit/9419321))
* **terraform:** support collector configuration ([d05bd78](https://github.com/suzuki-shunsuke/go-graylog/commit/d05bd78))



<a name="0.6.1"></a>
## [0.6.1](https://github.com/suzuki-shunsuke/go-graylog/compare/v0.6.0...v0.6.1) (2018-09-04)


### Bug Fixes

* make LDAP Setting's user password sensitive ([7ca7177](https://github.com/suzuki-shunsuke/go-graylog/commit/7ca7177))
* **client:** fix wrong usage of net/http.Request#WithContext ([224a6ca](https://github.com/suzuki-shunsuke/go-graylog/commit/224a6ca))



<a name="0.6.0"></a>
# [0.6.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v0.5.1...v0.6.0) (2018-09-03)


### Bug Fixes

* make ldap setting's default group computed and add docs ([b4ba47b](https://github.com/suzuki-shunsuke/go-graylog/commit/b4ba47b))


### Features

* implement LDAP Setting API's client ([8cfd1e3](https://github.com/suzuki-shunsuke/go-graylog/commit/8cfd1e3))
* implement mockserver of LDAP Setting API ([3e0b18d](https://github.com/suzuki-shunsuke/go-graylog/commit/3e0b18d))
* implement terraform provider resource of LDAP Setting ([a12a493](https://github.com/suzuki-shunsuke/go-graylog/commit/a12a493))



<a name="0.5.1"></a>
## [0.5.1](https://github.com/suzuki-shunsuke/go-graylog/compare/v0.5.0...v0.5.1) (2018-08-18)


### Bug Fixes

* fix terraform provider's user resource ([9740b33](https://github.com/suzuki-shunsuke/go-graylog/commit/9740b33))


### Features

* improve client's error message ([95f0f8f](https://github.com/suzuki-shunsuke/go-graylog/commit/95f0f8f))



<a name="0.5.0"></a>
# [0.5.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v0.4.2...v0.5.0) (2018-08-05)


### Features

* support GET Alarm Callbacks API ([7d69c91](https://github.com/suzuki-shunsuke/go-graylog/commit/7d69c91))



<a name="0.4.2"></a>
## [0.4.2](https://github.com/suzuki-shunsuke/go-graylog/compare/v0.4.1...v0.4.2) (2018-08-05)


### Bug Fixes

* **mockserver:** fix store of stream rule ([fe1b391](https://github.com/suzuki-shunsuke/go-graylog/commit/fe1b391))



<a name="0.4.1"></a>
## [0.4.1](https://github.com/suzuki-shunsuke/go-graylog/compare/v0.4.0...v0.4.1) (2018-08-05)


### Bug Fixes

* change schema type from List to Set ([24cacf1](https://github.com/suzuki-shunsuke/go-graylog/commit/24cacf1))



<a name="0.4.0"></a>
# [0.4.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v0.3.0...v0.4.0) (2018-08-05)


### Bug Fixes

* fix schema.Type from List to Set ([7b89520](https://github.com/suzuki-shunsuke/go-graylog/commit/7b89520))


### Features

* add dashboard resource to terraform provider ([8392828](https://github.com/suzuki-shunsuke/go-graylog/commit/8392828))
* support Dashboard API ([cd1ec3c](https://github.com/suzuki-shunsuke/go-graylog/commit/cd1ec3c))



<a name="0.3.0"></a>
# [0.3.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v0.2.0...v0.3.0) (2018-08-04)


### Features

* support GET Alert and Alerts API ([88f8c83](https://github.com/suzuki-shunsuke/go-graylog/commit/88f8c83))



<a name="0.2.0"></a>
# [0.2.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v0.1.4...v0.2.0) (2018-07-16)


### Features

* support GET Alert Conditions API ([a35b7f5](https://github.com/suzuki-shunsuke/go-graylog/commit/a35b7f5))



<a name="0.1.4"></a>
## [0.1.4](https://github.com/suzuki-shunsuke/go-graylog/compare/v0.1.3...v0.1.4) (2018-07-02)


### Bug Fixes

* failed to read index set's data at terraform import ([de6bb73](https://github.com/suzuki-shunsuke/go-graylog/commit/de6bb73))



<a name="0.1.3"></a>
## [0.1.3](https://github.com/suzuki-shunsuke/go-graylog/compare/v0.1.2...v0.1.3) (2018-07-01)


### Bug Fixes

* failed to import input's attributes by Read ([d096650](https://github.com/suzuki-shunsuke/go-graylog/commit/d096650))



<a name="0.1.2"></a>
## [0.1.2](https://github.com/suzuki-shunsuke/go-graylog/compare/v0.1.1...v0.1.2) (2018-07-01)



<a name="0.1.1"></a>
## [0.1.1](https://github.com/suzuki-shunsuke/go-graylog/compare/v0.1.0...v0.1.1) (2018-07-01)



<a name="0.1.0"></a>
# 0.1.0 (2018-04-29)
