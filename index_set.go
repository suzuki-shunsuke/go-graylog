package graylog

import (
	"github.com/suzuki-shunsuke/go-ptr"
	"time"
)

const (
	// rotation_strategy_class
	MESSAGE_COUNT_ROTATION_STRATEGY string = "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy"
	SIZE_BASED_ROTATION_STRATEGY    string = "org.graylog2.indexer.rotation.strategies.SizeBasedRotationStrategy"
	TIME_BASED_ROTATION_STRATEGY    string = "org.graylog2.indexer.rotation.strategies.TimeBasedRotationStrategy"
	// rotation_strategy.type
	MESSAGE_COUNT_ROTATION_STRATEGY_CONFIG string = "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig"
	SIZE_BASED_ROTATION_STRATEGY_CONFIG    string = "org.graylog2.indexer.rotation.strategies.SizeBasedRotationStrategyConfig"
	TIME_BASED_ROTATION_STRATEGY_CONFIG    string = "org.graylog2.indexer.rotation.strategies.TimeBasedRotationStrategyConfig"

	// retention_strategy_class
	DELETION_RETENTION_STRATEGY_CLASS string = "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy"
	CLOSING_RETENTION_STRATEGY_CLASS  string = "org.graylog2.indexer.retention.strategies.ClosingRetentionStrategy"
	NOOP_RETENTION_STRATEGY_CLASS     string = "org.graylog2.indexer.retention.strategies.NoopRetentionStrategy"
	// retention_strategy_class.type
	DELETION_RETENTION_STRATEGY_CONFIG string = "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig"
	CLOSING_RETENTION_STRATEGY_CONFIG  string = "org.graylog2.indexer.retention.strategies.ClosingRetentionStrategyConfig"
	NOOP_RETENTION_STRATEGY_CONFIG     string = "org.graylog2.indexer.retention.strategies.NoopRetentionStrategyConfig"

	CREATION_DATE_FORMAT string = "2006-01-02T15:04:05.000Z"
)

// IndexSet represents a Graylog's Index Set.
// http://docs.graylog.org/en/2.4/pages/configuration/index_model.html#index-set-configuration
type IndexSet struct {
	// required
	Title string `json:"title,omitempty" v-create:"required"`
	// ^[a-z0-9][a-z0-9_+-]*$
	IndexPrefix string `json:"index_prefix,omitempty" v-create:"required,indexprefixregexp"`
	// ex. "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy"
	RotationStrategyClass string            `json:"rotation_strategy_class,omitempty" v-create:"required"`
	RotationStrategy      *RotationStrategy `json:"rotation_strategy,omitempty" v-create:"required"`
	// ex. "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy"
	RetentionStrategyClass string             `json:"retention_strategy_class,omitempty" v-create:"required"`
	RetentionStrategy      *RetentionStrategy `json:"retention_strategy,omitempty" v-create:"required"`
	// ex. "2018-02-20T11:37:19.305Z"
	CreationDate                    string `json:"creation_date,omitempty"`
	IndexAnalyzer                   string `json:"index_analyzer,omitempty" v-create:"required"`
	Shards                          int    `json:"shards,omitempty" v-create:"required"`
	IndexOptimizationMaxNumSegments int    `json:"index_optimization_max_num_segments,omitempty" v-create:"required"`

	ID string `json:"id,omitempty" v-create:"isdefault"`

	Description               string         `json:"description,omitempty"`
	Replicas                  int            `json:"replicas,omitempty"`
	IndexOptimizationDisabled bool           `json:"index_optimization_disabled,omitempty"`
	Writable                  bool           `json:"writable,omitempty"`
	Default                   bool           `json:"default,omitempty"`
	Stats                     *IndexSetStats `json:"-"`
}

func (is *IndexSet) NewUpdateParams() *IndexSetUpdateParams {
	return &IndexSetUpdateParams{
		Title:                  is.Title,
		IndexPrefix:            is.IndexPrefix,
		RotationStrategyClass:  is.RotationStrategyClass,
		RotationStrategy:       is.RotationStrategy,
		RetentionStrategyClass: is.RetentionStrategyClass,
		RetentionStrategy:      is.RetentionStrategy,
		IndexAnalyzer:          is.IndexAnalyzer,
		Shards:                 is.Shards,
		IndexOptimizationMaxNumSegments: is.IndexOptimizationMaxNumSegments,
		ID: is.ID,

		Description:               ptr.PStr(is.Description),
		Replicas:                  ptr.PInt(is.Replicas),
		IndexOptimizationDisabled: ptr.PBool(is.IndexOptimizationDisabled),
		Writable:                  ptr.PBool(is.Writable),
	}
}

// IndexSetUpdateParams represents a Graylog's Index Set Update API's parameter.
// http://docs.graylog.org/en/2.4/pages/configuration/index_model.html#index-set-configuration
type IndexSetUpdateParams struct {
	Title                           string             `json:"title" v-update:"required"`
	IndexPrefix                     string             `json:"index_prefix" v-update:"required,indexprefixregexp"`
	RotationStrategyClass           string             `json:"rotation_strategy_class" v-update:"required"`
	RotationStrategy                *RotationStrategy  `json:"rotation_strategy" v-update:"required"`
	RetentionStrategyClass          string             `json:"retention_strategy_class" v-update:"required"`
	RetentionStrategy               *RetentionStrategy `json:"retention_strategy" v-update:"required"`
	IndexAnalyzer                   string             `json:"index_analyzer" v-update:"required"`
	Shards                          int                `json:"shards" v-update:"required"`
	IndexOptimizationMaxNumSegments int                `json:"index_optimization_max_num_segments" v-update:"required"`
	ID                              string             `json:"id" v-update:"required,objectid"`

	Description               *string `json:"description,omitempty"`
	Replicas                  *int    `json:"replicas,omitempty"`
	IndexOptimizationDisabled *bool   `json:"index_optimization_disabled,omitempty"`
	Writable                  *bool   `json:"writable,omitempty"`
}

func (is *IndexSet) SetCreateDefaultValues() {
	if is.CreationDate == "" {
		is.SetCreationTime(time.Now())
	}
	if is.Shards == 0 {
		is.Shards = 4
	}
	if is.IndexAnalyzer == "" {
		is.IndexAnalyzer = "standard"
	}
}

// CreationTime returns a creation date converted to time.Time.
func (is *IndexSet) CreationTime() (time.Time, error) {
	return time.Parse(CREATION_DATE_FORMAT, is.CreationDate)
}

// SetCreationTime sets a creation date with time.Time.
func (is *IndexSet) SetCreationTime(t time.Time) {
	is.CreationDate = t.Format(CREATION_DATE_FORMAT)
}

// RotationStrategy represents a Graylog's Index Set Rotation Strategy.
type RotationStrategy struct {
	// ex. "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig"
	Type string `json:"type,omitempty"`
	// ex. 20000000
	// Maximum number of documents in an index before it gets rotated
	MaxDocsPerIndex int `json:"max_docs_per_index,omitempty"`
	// time based
	// How long an index gets written to before it is rotated. (i.e. "P1D" for 1 day, "PT6H" for 6 hours)
	RotationPeriod string `json:"rotation_period,omitempty"`
	// size based
	// Maximum size of an index before it gets rotated
	MaxSize int `json:"max_size,omitempty"`
}

// RetentionStrategy represents a Graylog's Index Set Retention Strategy.
type RetentionStrategy struct {
	// ex. "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig"
	Type string `json:"type,omitempty"`
	// ex. 20
	MaxNumberOfIndices int `json:"max_number_of_indices,omitempty"`
}

type IndexSetsBody struct {
	IndexSets []IndexSet               `json:"index_sets"`
	Stats     map[string]IndexSetStats `json:"stats"`
	Total     int                      `json:"total"`
}

// NewMessageCountRotationStrategy returns a new message count based RotationStrategy.
func NewMessageCountRotationStrategy(count int) *RotationStrategy {
	if count <= 0 {
		count = 20000000
	}
	return &RotationStrategy{
		Type:            MESSAGE_COUNT_ROTATION_STRATEGY_CONFIG,
		MaxDocsPerIndex: count,
	}
}

// NewSizeBasedRotationStrategy returns a new size based RotationStrategy.
func NewSizeBasedRotationStrategy(size int) *RotationStrategy {
	if size <= 0 {
		size = 1073741824
	}
	return &RotationStrategy{
		Type:    SIZE_BASED_ROTATION_STRATEGY_CONFIG,
		MaxSize: size,
	}
}

// NewTimeBasedRotationStrategy returns a new time based RotationStrategy.
func NewTimeBasedRotationStrategy(period string) *RotationStrategy {
	if period == "" {
		period = "P1D"
	}
	return &RotationStrategy{
		Type:           TIME_BASED_ROTATION_STRATEGY_CONFIG,
		RotationPeriod: period,
	}
}

// NewDeletionRetentionStrategy returns a new deletion RetentionStrategy.
func NewDeletionRetentionStrategy(num int) *RetentionStrategy {
	if num <= 0 {
		num = 20
	}
	return &RetentionStrategy{
		Type:               DELETION_RETENTION_STRATEGY_CONFIG,
		MaxNumberOfIndices: num,
	}
}

// NewClosingRetentionStrategy returns a new closing RetentionStrategy.
func NewClosingRetentionStrategy(num int) *RetentionStrategy {
	if num <= 0 {
		num = 20
	}
	return &RetentionStrategy{
		Type:               CLOSING_RETENTION_STRATEGY_CONFIG,
		MaxNumberOfIndices: num,
	}
}

// NewNoopRetentionStrategy returns a new noop RetentionStrategy.
func NewNoopRetentionStrategy(num int) *RetentionStrategy {
	if num <= 0 {
		num = 20
	}
	return &RetentionStrategy{
		Type:               NOOP_RETENTION_STRATEGY_CONFIG,
		MaxNumberOfIndices: num,
	}
}
