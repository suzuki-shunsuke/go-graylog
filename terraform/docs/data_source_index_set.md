# Data source graylog_index_set

https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/data_source_index_set.go

```hcl
data "graylog_index_set" "test-index-set" {
  index_prefix = "terraform-test"
}
```

## Required Argument

One of `index_set_id` or `title` or `index_prefix` must be set.

## Attributes

name | type | description
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
description | string |
replicas | int |
index_optimization_disabled | bool |
writable | bool |
default | bool |
creation_date | string |
id | string |
