package client

import (
	"context"
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
)

// GetExtractors returns all extractors of an input
func (client *Client) GetExtractors(inputID string) (
	[]graylog.Extractor, int, *ErrorInfo, error,
) {
	return client.GetExtractorsContext(context.Background(), inputID)
}

// GetExtractorsContext returns all extractors with a context.
func (client *Client) GetExtractorsContext(ctx context.Context, inputID string) (
	[]graylog.Extractor, int, *ErrorInfo, error,
) {
	if inputID == "" {
		return nil, 0, nil, fmt.Errorf("input id is required")
	}
	body := &graylog.ExtractorsBody{}
	u, err := client.Endpoints().Extractors(inputID)
	if err != nil {
		return nil, 0, nil, err
	}
	ei, err := client.callGet(ctx, u.String(), nil, body)
	return body.Extractors, body.Total, ei, err
}

// GetExtractor returns an extractor.
func (client *Client) GetExtractor(inputID, extractorID string) (
	*graylog.Extractor, *ErrorInfo, error,
) {
	return client.GetExtractorContext(context.Background(), inputID, extractorID)
}

// GetExtractorContext returns an extractor with a context.
func (client *Client) GetExtractorContext(
	ctx context.Context, inputID, extractorID string,
) (
	*graylog.Extractor, *ErrorInfo, error,
) {
	if inputID == "" {
		return nil, nil, fmt.Errorf("input id is required")
	}
	if extractorID == "" {
		return nil, nil, fmt.Errorf("extractor id is required")
	}
	ext := &graylog.Extractor{}
	u, err := client.Endpoints().Extractor(inputID, extractorID)
	if err != nil {
		return nil, nil, err
	}
	ei, err := client.callGet(ctx, u.String(), nil, ext)
	return ext, ei, err
}

// CreateExtractor adds an extractor to an input.
func (client *Client) CreateExtractor(
	inputID string, extractor *graylog.Extractor,
) (
	*ErrorInfo, error,
) {
	return client.CreateExtractorContext(context.Background(), inputID, extractor)
}

func convertExtractorForPostAndPut(extractor *graylog.Extractor) interface{} {
	converters := make(map[string]interface{}, len(extractor.Converters))
	for _, converter := range extractor.Converters {
		converters[converter.Type] = converter.Config
	}
	return map[string]interface{}{
		"title":            extractor.Title,
		"cut_or_copy":      extractor.CursorStrategy,
		"source_field":     extractor.SourceField,
		"target_field":     extractor.TargetField,
		"extractor_type":   extractor.Type,
		"extractor_config": extractor.ExtractorConfig,
		"converters":       converters,
		"condition_type":   extractor.ConditionType,
		"condition_value":  extractor.ConditionValue,
		"order":            extractor.Order,
	}
}

// CreateExtractorContext adds an extractor to an input with a context.
func (client *Client) CreateExtractorContext(
	ctx context.Context, inputID string, extractor *graylog.Extractor,
) (
	*ErrorInfo, error,
) {
	if inputID == "" {
		return nil, fmt.Errorf("input id is required")
	}
	u, err := client.Endpoints().Extractors(inputID)
	if err != nil {
		return nil, err
	}
	resp := map[string]string{}
	ei, err := client.callPost(
		ctx, u.String(), convertExtractorForPostAndPut(extractor), &resp)
	if err != nil {
		return ei, err
	}
	id, ok := resp["extractor_id"]
	if !ok {
		return ei, fmt.Errorf(`response doesn't have the field "extractor_id""`)
	}
	extractor.ID = id
	return ei, nil
}

// UpdateExtractor updates an extractor.
func (client *Client) UpdateExtractor(
	inputID string, extractor *graylog.Extractor,
) (
	*ErrorInfo, error,
) {
	return client.UpdateExtractorContext(context.Background(), inputID, extractor)
}

// UpdateExtractorContext updates an extractor with a context.
func (client *Client) UpdateExtractorContext(
	ctx context.Context, inputID string, extractor *graylog.Extractor,
) (
	*ErrorInfo, error,
) {
	if inputID == "" {
		return nil, fmt.Errorf("input id is required")
	}
	u, err := client.Endpoints().Extractor(inputID, extractor.ID)
	if err != nil {
		return nil, err
	}
	return client.callPut(
		ctx, u.String(), convertExtractorForPostAndPut(extractor), extractor)
}

// DeleteExtractor updates an extractor.
func (client *Client) DeleteExtractor(
	inputID, extractorID string,
) (
	*ErrorInfo, error,
) {
	return client.DeleteExtractorContext(context.Background(), inputID, extractorID)
}

// DeleteExtractorContext updates an extractor with a context.
func (client *Client) DeleteExtractorContext(
	ctx context.Context, inputID, extractorID string,
) (
	*ErrorInfo, error,
) {
	if inputID == "" {
		return nil, fmt.Errorf("input id is required")
	}
	if extractorID == "" {
		return nil, fmt.Errorf("extractor id is required")
	}
	u, err := client.Endpoints().Extractor(inputID, extractorID)
	if err != nil {
		return nil, err
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}
