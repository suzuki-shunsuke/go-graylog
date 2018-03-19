package graylog

// IndexSet represents a Graylog's Index Set.
type IndexSet struct {
	// required
	Title string `json:"title,omitempty" v-create:"required" v-update:"required"`
	// ^[a-z0-9][a-z0-9_+-]*$
	IndexPrefix string `json:"index_prefix,omitempty" v-create:"required,indexprefixregexp" v-update:"required,indexprefixregexp"`
	// ex. "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy"
	RotationStrategyClass string            `json:"rotation_strategy_class,omitempty" v-create:"required" v-update:"required"`
	RotationStrategy      *RotationStrategy `json:"rotation_strategy,omitempty" v-create:"required" v-update:"required"`
	// ex. "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy"
	RetentionStrategyClass string             `json:"retention_strategy_class,omitempty" v-create:"required" v-update:"required"`
	RetentionStrategy      *RetentionStrategy `json:"retention_strategy,omitempty" v-create:"required" v-update:"required"`
	// ex. "2018-02-20T11:37:19.305Z"
	CreationDate                    string `json:"creation_date,omitempty" v-create:"required" v-update:"required"`
	IndexAnalyzer                   string `json:"index_analyzer,omitempty" v-create:"required" v-update:"required"`
	Shards                          int    `json:"shards,omitempty" v-create:"required" v-update:"required"`
	IndexOptimizationMaxNumSegments int    `json:"index_optimization_max_num_segments,omitempty" v-create:"required" v-update:"required"`

	ID string `json:"id,omitempty" v-create:"isdefault" v-update:"required"`

	Description               string `json:"description,omitempty"`
	Replicas                  int    `json:"replicas,omitempty"`
	IndexOptimizationDisabled bool   `json:"index_optimization_disabled,omitempty"`
	Writable                  bool   `json:"writable,omitempty"`
	Default                   bool   `json:"default,omitempty"`
}

// RotationStrategy represents a Graylog's Index Set Rotation Strategy.
type RotationStrategy struct {
	// ex. "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig"
	Type string `json:"type,omitempty"`
	// ex. 20000000
	MaxDocsPerIndex int `json:"max_docs_per_index,omitempty"`
}

// RetentionStrategy represents a Graylog's Index Set Retention Strategy.
type RetentionStrategy struct {
	// ex. "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig"
	Type string `json:"type,omitempty"`
	// ex. 20
	MaxNumberOfIndices int `json:"max_number_of_indices,omitempty"`
}

type IndexSetsBody struct {
	IndexSets []IndexSet     `json:"index_sets"`
	Stats     *IndexSetStats `json:"stats"`
	Total     int            `json:"total"`
}
