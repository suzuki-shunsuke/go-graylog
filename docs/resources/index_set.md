# graylog_index_set

* [Example](../../examples/v0.12/index_set.tf)
* [Source code](../../graylog/terraform/resource_index_set.go)

## Argument Reference

### Required Argument

name | type | etc
--- | --- | ---
title | string |
index_prefix | string | `force new`
rotation_strategy_class | string |
rotation_strategy | |
rotation_strategy.type | string |
rotation_strategy.max_docs_per_index | int |
rotation_strategy.max_size | int |
rotation_strategy.rotation_period | string |
retention_strategy_class | string |
retention_strategy | |
retention_strategy.type | string |
retention_strategy.max_number_of_indices | int |
index_analyzer | string |
shards | int |
index_optimization_max_num_segments | int |

### Optional Argument

name | default | type | description
--- | --- | --- | ---
description | "" | string |
replicas | 0 | int |
index_optimization_disabled | | bool |
writable | | bool |
default | | bool |
creation_date | computed | string |

## Attrs Reference

name | type | etc
--- | --- | ---
id | string |
