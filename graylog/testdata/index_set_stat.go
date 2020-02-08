package testdata

import (
	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/graylog"
)

func IndexSetStats() graylog.IndexSetStats {
	return graylog.IndexSetStats{
		Indices:   1,
		Documents: 0,
		Size:      1044,
	}
}
