# Change Log

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

<a name="4.0.0"></a>
# [4.0.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v3.3.0...v4.0.0) (2019-07-03)


### Features

* add graylog_dashboard_widget resource ([ebf7443](https://github.com/suzuki-shunsuke/go-graylog/commit/ebf7443))
* support terraform resource graylog_dashboard_widget ([6188000](https://github.com/suzuki-shunsuke/go-graylog/commit/6188000)), closes [#20](https://github.com/suzuki-shunsuke/go-graylog/issues/20)
* support update dashboard cache time and description API ([6be8303](https://github.com/suzuki-shunsuke/go-graylog/commit/6be8303))
* support update dashboard widget API ([11ce42f](https://github.com/suzuki-shunsuke/go-graylog/commit/11ce42f))


### BREAKING CHANGES

* the type of DashboardWidget and Widget.Config change



<a name="3.3.0"></a>
# [3.3.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v3.2.0...v3.3.0) (2019-06-27)


### Features

* support StaticFields API ([b96fe9a](https://github.com/suzuki-shunsuke/go-graylog/commit/b96fe9a)), closes [#119](https://github.com/suzuki-shunsuke/go-graylog/issues/119)



<a name="3.2.0"></a>
# [3.2.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v3.1.0...v3.2.0) (2019-06-26)


### Features

* add a "static_fields" to Input ([79123e5](https://github.com/suzuki-shunsuke/go-graylog/commit/79123e5)), closes [#118](https://github.com/suzuki-shunsuke/go-graylog/issues/118)



<a name="3.1.0"></a>
# [3.1.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v3.0.0...v3.1.0) (2019-06-21)


### Features

* add the data source of graylog_stream ([465e4a1](https://github.com/suzuki-shunsuke/go-graylog/commit/465e4a1)), closes [#113](https://github.com/suzuki-shunsuke/go-graylog/issues/113)



<a name="3.0.0"></a>
# [3.0.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v2.6.0...v3.0.0) (2019-06-21)


### Code Refactoring

* replace the package mockserver to graylog-mock-server ([cbb3545](https://github.com/suzuki-shunsuke/go-graylog/commit/cbb3545)), closes [#114](https://github.com/suzuki-shunsuke/go-graylog/issues/114)


### BREAKING CHANGES

* remove the package mockserver

https://github.com/suzuki-shunsuke/graylog-mock-server



<a name="2.6.0"></a>
# [2.6.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v2.5.0...v2.6.0) (2019-06-20)


### Features

* support data source of graylog_index_set ([d4ad234](https://github.com/suzuki-shunsuke/go-graylog/commit/d4ad234)), closes [#110](https://github.com/suzuki-shunsuke/go-graylog/issues/110)



<a name="2.5.0"></a>
# [2.5.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v2.4.0...v2.5.0) (2019-05-25)


### Features

* support Terraform 0.12 ([7dba535](https://github.com/suzuki-shunsuke/go-graylog/commit/7dba535)), closes [#103](https://github.com/suzuki-shunsuke/go-graylog/issues/103)



<a name="2.4.0"></a>
# [2.4.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v2.3.3...v2.4.0) (2019-05-23)


### Bug Fixes

* fix Extractor API ([b865d5e](https://github.com/suzuki-shunsuke/go-graylog/commit/b865d5e))
* fix Extractor API and add tests ([dc5ecfd](https://github.com/suzuki-shunsuke/go-graylog/commit/dc5ecfd))


### Features

* [#69](https://github.com/suzuki-shunsuke/go-graylog/issues/69) support Extractor API ([ebfab91](https://github.com/suzuki-shunsuke/go-graylog/commit/ebfab91))



<a name="2.3.3"></a>
## [2.3.3](https://github.com/suzuki-shunsuke/go-graylog/compare/v2.3.2...v2.3.3) (2019-05-20)


### Bug Fixes

* make some terraform resource fields's ForceNew true ([2ad4551](https://github.com/suzuki-shunsuke/go-graylog/commit/2ad4551)), closes [#101](https://github.com/suzuki-shunsuke/go-graylog/issues/101)



<a name="2.3.2"></a>
## [2.3.2](https://github.com/suzuki-shunsuke/go-graylog/compare/v2.3.1...v2.3.2) (2019-05-19)


### Bug Fixes

* add the input type InputTypeRawKafka ([83132bd](https://github.com/suzuki-shunsuke/go-graylog/commit/83132bd)), closes [#73](https://github.com/suzuki-shunsuke/go-graylog/issues/73)



<a name="2.3.1"></a>
## [2.3.1](https://github.com/suzuki-shunsuke/go-graylog/compare/v2.3.0...v2.3.1) (2019-05-19)


### Bug Fixes

* fix import of input ([de718e1](https://github.com/suzuki-shunsuke/go-graylog/commit/de718e1)), closes [#97](https://github.com/suzuki-shunsuke/go-graylog/issues/97)



<a name="2.3.0"></a>
# [2.3.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v2.2.0...v2.3.0) (2019-05-19)


### Features

* add the field "field_type_refresh_interval" to IndexSet ([ab45d2d](https://github.com/suzuki-shunsuke/go-graylog/commit/ab45d2d)), closes [#92](https://github.com/suzuki-shunsuke/go-graylog/issues/92)



<a name="2.2.0"></a>
# [2.2.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v2.1.0...v2.2.0) (2019-05-19)


### Bug Fixes

* add graylog_pipeline_connection resource ([c29afd5](https://github.com/suzuki-shunsuke/go-graylog/commit/c29afd5))
* remove sync_pipelines ([6486d68](https://github.com/suzuki-shunsuke/go-graylog/commit/6486d68))


### Features

* add the parameter api_version to support v3 ([ab5fa4f](https://github.com/suzuki-shunsuke/go-graylog/commit/ab5fa4f))
* support Plugins/Pipelines/Connections API ([5ac0342](https://github.com/suzuki-shunsuke/go-graylog/commit/5ac0342)), closes [#91](https://github.com/suzuki-shunsuke/go-graylog/issues/91)



<a name="2.1.0"></a>
# [2.1.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v2.0.1...v2.1.0) (2019-05-18)


### Features

* support Plugins/Pipelines/Pipelines API ([7be7523](https://github.com/suzuki-shunsuke/go-graylog/commit/7be7523)), closes [#67](https://github.com/suzuki-shunsuke/go-graylog/issues/67)



<a name="2.0.1"></a>
## [2.0.1](https://github.com/suzuki-shunsuke/go-graylog/compare/v2.0.0...v2.0.1) (2019-05-18)


### Bug Fixes

* fix CreatePipelineRule ([70e1185](https://github.com/suzuki-shunsuke/go-graylog/commit/70e1185))



<a name="2.0.0"></a>
# [2.0.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v2.0.0-5...v2.0.0) (2019-05-09)



<a name="2.0.0-5"></a>
# [2.0.0-5](https://github.com/suzuki-shunsuke/go-graylog/compare/v2.0.0-4...v2.0.0-5) (2019-04-29)


### Bug Fixes

* fix type conversion error ([d4b1777](https://github.com/suzuki-shunsuke/go-graylog/commit/d4b1777)), closes [#84](https://github.com/suzuki-shunsuke/go-graylog/issues/84)



<a name="2.0.0-4"></a>
# [2.0.0-4](https://github.com/suzuki-shunsuke/go-graylog/compare/v2.0.0-3...v2.0.0-4) (2019-04-11)


### Bug Fixes

* add repeat_notifications to terraform's alert condition resource ([d7fbb04](https://github.com/suzuki-shunsuke/go-graylog/commit/d7fbb04)), closes [#82](https://github.com/suzuki-shunsuke/go-graylog/issues/82)



<a name="2.0.0-3"></a>
# [2.0.0-3](https://github.com/suzuki-shunsuke/go-graylog/compare/v2.0.0-2...v2.0.0-3) (2019-04-11)


### Features

* support terraform import of Alert Condition ([45ac871](https://github.com/suzuki-shunsuke/go-graylog/commit/45ac871)), closes [#55](https://github.com/suzuki-shunsuke/go-graylog/issues/55)



<a name="2.0.0-2"></a>
# [2.0.0-2](https://github.com/suzuki-shunsuke/go-graylog/compare/v2.0.0-1...v2.0.0-2) (2019-04-10)



<a name="2.0.0-1"></a>
# [2.0.0-1](https://github.com/suzuki-shunsuke/go-graylog/compare/v2.0.0-0...v2.0.0-1) (2019-04-10)



<a name="2.0.0-0"></a>
# [2.0.0-0](https://github.com/suzuki-shunsuke/go-graylog/compare/v1.4.0...v2.0.0-0) (2019-04-10)


### Features

* migrate terraform AlertCondition resource ([23dc1eb](https://github.com/suzuki-shunsuke/go-graylog/commit/23dc1eb)), closes [#76](https://github.com/suzuki-shunsuke/go-graylog/issues/76)



<a name="1.4.0"></a>
# [1.4.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v1.3.2...v1.4.0) (2019-02-27)


### Features

* [#68](https://github.com/suzuki-shunsuke/go-graylog/issues/68) support create a pipeline rule API ([5134501](https://github.com/suzuki-shunsuke/go-graylog/commit/5134501))
* [#68](https://github.com/suzuki-shunsuke/go-graylog/issues/68) support Get all pipeline rules API ([be29881](https://github.com/suzuki-shunsuke/go-graylog/commit/be29881))
* [#68](https://github.com/suzuki-shunsuke/go-graylog/issues/68) support get pipeline rule API ([3b0439a](https://github.com/suzuki-shunsuke/go-graylog/commit/3b0439a))
* [#68](https://github.com/suzuki-shunsuke/go-graylog/issues/68) support terraform pipeline_rule resource ([cea5ca7](https://github.com/suzuki-shunsuke/go-graylog/commit/cea5ca7))
* [#68](https://github.com/suzuki-shunsuke/go-graylog/issues/68) support update pipeline API ([49e0221](https://github.com/suzuki-shunsuke/go-graylog/commit/49e0221))
* support delete pipeline rule API ([5768980](https://github.com/suzuki-shunsuke/go-graylog/commit/5768980))



<a name="1.3.2"></a>
## [1.3.2](https://github.com/suzuki-shunsuke/go-graylog/compare/v1.3.1...v1.3.2) (2019-01-27)


### Bug Fixes

* fix error handling of terraform provider ([985dc90](https://github.com/suzuki-shunsuke/go-graylog/commit/985dc90))



<a name="1.3.1"></a>
## [1.3.1](https://github.com/suzuki-shunsuke/go-graylog/compare/v1.3.0...v1.3.1) (2019-01-26)


### Bug Fixes

* handle error of terraform resource read ([4719ba3](https://github.com/suzuki-shunsuke/go-graylog/commit/4719ba3)), closes [#63](https://github.com/suzuki-shunsuke/go-graylog/issues/63)



<a name="1.3.0"></a>
# [1.3.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v1.2.0...v1.3.0) (2019-01-25)


### Features

* support terraform import of stream rule ([c68778a](https://github.com/suzuki-shunsuke/go-graylog/commit/c68778a)), closes [#57](https://github.com/suzuki-shunsuke/go-graylog/issues/57)



<a name="1.2.0"></a>
# [1.2.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v1.1.1...v1.2.0) (2019-01-25)


### Features

* support terraform import of alarm callback ([baee116](https://github.com/suzuki-shunsuke/go-graylog/commit/baee116)), closes [#56](https://github.com/suzuki-shunsuke/go-graylog/issues/56)



<a name="1.1.1"></a>
## [1.1.1](https://github.com/suzuki-shunsuke/go-graylog/compare/v1.1.0...v1.1.1) (2019-01-08)


### Bug Fixes

* fix terraform resource_role ([dda5782](https://github.com/suzuki-shunsuke/go-graylog/commit/dda5782))



<a name="1.1.0"></a>
# [1.1.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v1.0.0...v1.1.0) (2018-12-31)


### Features

* support alram callbacks ([dfd4f8e](https://github.com/suzuki-shunsuke/go-graylog/commit/dfd4f8e))



<a name="1.0.0"></a>
# [1.0.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v1.0.0-0...v1.0.0) (2018-12-30)



<a name="1.0.0-0"></a>
# [1.0.0-0](https://github.com/suzuki-shunsuke/go-graylog/compare/v0.12.1...v1.0.0-0) (2018-12-30)


### Bug Fixes

* change AlertConditionParameter to interface ([69daf46](https://github.com/suzuki-shunsuke/go-graylog/commit/69daf46))
* fix resource_alert_condition ([fa70c3a](https://github.com/suzuki-shunsuke/go-graylog/commit/fa70c3a))
* fix resource_alert_condition's arguments ([6c05135](https://github.com/suzuki-shunsuke/go-graylog/commit/6c05135))
* fix resource_alert_condition's parameters ([b029d5b](https://github.com/suzuki-shunsuke/go-graylog/commit/b029d5b))


### Features

* add GeneralAlertConditionParameters ([3419450](https://github.com/suzuki-shunsuke/go-graylog/commit/3419450))
* support Alert Condition API ([91238a3](https://github.com/suzuki-shunsuke/go-graylog/commit/91238a3))



<a name="0.12.1"></a>
## [0.12.1](https://github.com/suzuki-shunsuke/go-graylog/compare/v0.12.0...v0.12.1) (2018-12-23)


### Bug Fixes

* make index set prefix force new ([07c24b9](https://github.com/suzuki-shunsuke/go-graylog/commit/07c24b9))



<a name="0.12.0"></a>
# [0.12.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v0.11.0...v0.12.0) (2018-12-03)


### Features

* support terraform resource stream_rule ([65b921b](https://github.com/suzuki-shunsuke/go-graylog/commit/65b921b))



<a name="0.11.0"></a>
# [0.11.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v0.10.0...v0.11.0) (2018-12-02)


### Bug Fixes

* support the custom header "X-Requested-By" ([ff66bf1](https://github.com/suzuki-shunsuke/go-graylog/commit/ff66bf1)), closes [#42](https://github.com/suzuki-shunsuke/go-graylog/issues/42)



<a name="0.10.0"></a>
# [0.10.0](https://github.com/suzuki-shunsuke/go-graylog/compare/v0.9.0-0...v0.10.0) (2018-11-06)


### Bug Fixes

* fix UpdateLDAPSetting and terraform ldap_setting provider ([7883a76](https://github.com/suzuki-shunsuke/go-graylog/commit/7883a76))


### Features

* support LDAP Groups API ([efb9a96](https://github.com/suzuki-shunsuke/go-graylog/commit/efb9a96))


### BREAKING CHANGES

* * remove LDAPSettingUpdateParams
* make some ldap_setting terraform resource's parameter required



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
