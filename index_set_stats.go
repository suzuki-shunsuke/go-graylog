package graylog

// IndexSetStats represents a Graylog's Index Set Stats.
type IndexSetStats struct {
	Indices   int `json:"indices"`
	Documents int `json:"documents"`
	Size      int `json:"size"`
}
