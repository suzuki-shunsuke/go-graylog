# graylog_index_set

https://github.com/suzuki-shunsuke/terraform-provider-graylog/blob/master/resource_index_set.go

```
resource "graylog_index_set" "test-index-set" {
  title = "terraform test index set"
  index_prefix = "terraform-test"
  rotation_strategy_class = "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy"
  rotation_strategy = {
    type = "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig"
  }
  retention_strategy_class = "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy"
  retention_strategy = {
    type = "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig"
  }
  creation_date = "2018-02-20T11:37:29.305Z"
  index_analyzer = "standard"
  shards = 4
  index_optimization_max_num_segments = 1
}
```

## Argument Reference

### Required Argument

name | type | description
--- | --- | ---
title | string |
index_prefix | string |
rotation_strategy_class | string |
rotation_strategy | map[string] |
rotation_strategy.type | string |
retention_strategy_class | string |
retention_strategy | map[string] |
retention_strategy.type | string |
creation_date | string |
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

## Attributes Reference

name | type | etc
--- | --- | ---
id | string |
