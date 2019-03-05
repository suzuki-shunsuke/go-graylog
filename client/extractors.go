package client

import (
	"context"

	"github.com/suzuki-shunsuke/go-graylog"
)

// GetInputExtractors returns all extractors of an input.
func (client *Client) GetInputExtractors(inputID string) ([]graylog.Extractor, *ErrorInfo, error) {
	return client.GetInputExtractorsContext(context.Background(), inputID)
}

// GetInputExtractorsContext returns all extractors of an input with a context.
func (client *Client) GetInputExtractorsContext(ctx context.Context, inputID string) ([]graylog.Extractor, *ErrorInfo, error) {
	extractors := &graylog.ExtractorsBody{}
	u, err := client.Endpoints().Extractors(inputID)
	if err != nil {
		return nil, nil, err
	}
	ei, err := client.callGet(ctx, u.String(), nil, extractors)
	return extractors.Extractors, ei, err
}

// GetInputExtractor returns an extractors of an input.
func (client *Client) GetInputExtractor(inputID string, extractorID string) (*graylog.Extractor, *ErrorInfo, error) {
	return client.GetInputExtractorContext(context.Background(), inputID, extractorID)
}

// GetInputExtractorContext returns an extractor of an input with a context.
func (client *Client) GetInputExtractorContext(ctx context.Context, inputID string, extractorID string) (*graylog.Extractor, *ErrorInfo, error) {
	extractor := &graylog.Extractor{}
	u, err := client.Endpoints().Extractor(inputID, extractorID)
	if err != nil {
		return nil, nil, err
	}
	ei, err := client.callGet(ctx, u.String(), nil, extractor)
	return extractor, ei, err
}

// CreateInputExtractor creates multiple extractors of an input.
func (client *Client) CreateInputExtractor(extractor *graylog.Extractor, inputID string) (*ErrorInfo, error) {
	return client.CreateInputExtractorContext(context.Background(), extractor, inputID)
}

// CreateInputExtractorContext creates multiple extractors of an input with a context.
func (client *Client) CreateInputExtractorContext(ctx context.Context, extractor *graylog.Extractor, inputID string) (*ErrorInfo, error) {
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

// UpdateInputExtractor updates an extractor of an input.
func (client *Client) UpdateInputExtractor(extractor *graylog.Extractor, inputID string) (*ErrorInfo, error) {
	return client.UpdateInputExtractorContext(context.Background(), extractor, inputID)
}

// UpdateInputExtractorContext updates an extractor of an input with context.
func (client *Client) UpdateInputExtractorContext(ctx context.Context, extractor *graylog.Extractor, inputID string) (*ErrorInfo, error) {
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

// DeleteInputExtractor deletes an extractor of an input.
func (client *Client) DeleteInputExtractor(inputID string, extractorID string) (*ErrorInfo, error) {
	return client.DeleteInputExtractorContext(context.Background(), inputID, extractorID)
}

// DeleteInputExtractorContext deletes an extractor of an input with context.
func (client *Client) DeleteInputExtractorContext(ctx context.Context, inputID string, extractorID string) (*ErrorInfo, error) {
	u, err := client.Endpoints().Extractor(inputID, extractorID)
	if err != nil {
		return nil, err
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}
