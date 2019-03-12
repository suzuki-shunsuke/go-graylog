package client

import (
	"context"

	"github.com/suzuki-shunsuke/go-graylog"
)

// GetExtractors returns all extractors of an input.
func (client *Client) GetExtractors(inputID string) ([]graylog.Extractor, int, *ErrorInfo, error) {
	return client.GetExtractorsContext(context.Background(), inputID)
}

// GetExtractorsContext returns all extractors of an input with a context.
func (client *Client) GetExtractorsContext(ctx context.Context, inputID string) ([]graylog.Extractor, int, *ErrorInfo, error) {
	extractors := &graylog.ExtractorsBody{}
	u, err := client.Endpoints().Extractors(inputID)
	if err != nil {
		return nil, 0, nil, err
	}
	ei, err := client.callGet(ctx, u.String(), nil, extractors)
	return extractors.Extractors, extractors.Total, ei, err
}

// GetExtractor returns an extractors of an input.
func (client *Client) GetExtractor(inputID, extractorID string) (*graylog.Extractor, *ErrorInfo, error) {
	return client.GetExtractorContext(context.Background(), inputID, extractorID)
}

// GetExtractorContext returns an extractor of an input with a context.
func (client *Client) GetExtractorContext(ctx context.Context, inputID, extractorID string) (*graylog.Extractor, *ErrorInfo, error) {
	extractor := &graylog.Extractor{}
	u, err := client.Endpoints().Extractor(inputID, extractorID)
	if err != nil {
		return nil, nil, err
	}
	ei, err := client.callGet(ctx, u.String(), nil, extractor)
	return extractor, ei, err
}

// CreateExtractor creates multiple extractors of an input.
func (client *Client) CreateExtractor(extractor *graylog.Extractor, inputID string) (*ErrorInfo, error) {
	return client.CreateExtractorContext(context.Background(), extractor, inputID)
}

// CreateExtractorContext creates multiple extractors of an input with a context.
func (client *Client) CreateExtractorContext(ctx context.Context, extractor *graylog.Extractor, inputID string) (*ErrorInfo, error) {
	u, err := client.Endpoints().Extractors(inputID)
	if err != nil {
		return nil, err
	}
	d := map[string]interface{}{
		"title":            extractor.Title,
		"cut_or_copy":      extractor.CursorStrategy,
		"source_field":     extractor.SourceField,
		"target_field":     extractor.TargetField,
		"extractor_type":   extractor.Type,
		"extractor_config": extractor.ExtractorConfig,
		"converters":       map[string]interface{}{},
		"condition_type":   extractor.ConditionType,
		"condition_value":  extractor.ConditionValue,
		"order":            extractor.Order,
	}
	return client.callPost(ctx, u.String(), d, nil)
}

// UpdateExtractor updates an extractor of an input.
func (client *Client) UpdateExtractor(extractor *graylog.Extractor, inputID string) (*ErrorInfo, error) {
	return client.UpdateExtractorContext(context.Background(), extractor, inputID)
}

// UpdateExtractorContext updates an extractor of an input with context.
func (client *Client) UpdateExtractorContext(ctx context.Context, extractor *graylog.Extractor, inputID string) (*ErrorInfo, error) {
	u, err := client.Endpoints().Extractor(inputID, extractor.ID)
	if err != nil {
		return nil, err
	}
	d := map[string]interface{}{
		"title":            extractor.Title,
		"cut_or_copy":      extractor.CursorStrategy,
		"source_field":     extractor.SourceField,
		"target_field":     extractor.TargetField,
		"extractor_type":   extractor.Type,
		"extractor_config": extractor.ExtractorConfig,
		"converters":       map[string]interface{}{},
		"condition_type":   extractor.ConditionType,
		"condition_value":  extractor.ConditionValue,
		"order":            extractor.Order,
	}
	ei, err := client.callPut(ctx, u.String(), d, nil)
	return ei, err
}

// DeleteExtractor deletes an extractor of an input.
func (client *Client) DeleteExtractor(inputID, extractorID string) (*ErrorInfo, error) {
	return client.DeleteExtractorContext(context.Background(), inputID, extractorID)
}

// DeleteExtractorContext deletes an extractor of an input with context.
func (client *Client) DeleteExtractorContext(ctx context.Context, inputID, extractorID string) (*ErrorInfo, error) {
	u, err := client.Endpoints().Extractor(inputID, extractorID)
	if err != nil {
		return nil, err
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}
