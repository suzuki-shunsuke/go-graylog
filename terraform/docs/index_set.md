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
