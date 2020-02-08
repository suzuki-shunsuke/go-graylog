data "graylog_index_set" "default" {
  index_prefix = "graylog"
}

resource "graylog_index_set" "test" {
  title                               = "test"
  index_prefix                        = "lai3o-test"
  rotation_strategy_class             = "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy"
  retention_strategy_class            = "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy"
  description                         = "test 2"
  index_analyzer                      = "standard"
  index_optimization_disabled         = true
  writable                            = true
  shards                              = 4
  replicas                            = 0
  index_optimization_max_num_segments = 1
  field_type_refresh_interval         = 5000

  retention_strategy = {
    max_number_of_indices = 20
    type                  = "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig"
  }

  rotation_strategy = {
    max_docs_per_index = 20000000
    max_size           = 0
    type               = "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig"
  }
}
