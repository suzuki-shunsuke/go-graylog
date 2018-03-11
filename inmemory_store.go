package graylog

type InMemoryStore struct {
	users             map[string]User                  `json:"users"`
	roles             map[string]Role                  `json:"roles"`
	inputs            map[string]Input                 `json:"inputs"`
	indexSets         map[string]IndexSet              `json:"index_sets"`
	defaultIndexSetID string                           `json:"default_index_set_id"`
	indexSetStats     map[string]IndexSetStats         `json:"index_set_stats"`
	streams           map[string]Stream                `json:"streams"`
	streamRules       map[string]map[string]StreamRule `json:"stream_rules"`
	dataPath          string                           `json:"-"`
}
