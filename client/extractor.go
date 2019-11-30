package client

import (
	"context"
	"errors"

	"github.com/suzuki-shunsuke/go-graylog/v8"
)

// GetExtractors returns all extractors.
func (client *Client) GetExtractors(ctx context.Context, inputID string) (
	[]graylog.Extractor, int, *ErrorInfo, error,
) {
	if inputID == "" {
		return nil, 0, nil, errors.New("input id is required")
	}
	body := &graylog.ExtractorsBody{}
	ei, err := client.callGet(ctx, client.Endpoints().Extractors(inputID), nil, body)
	return body.Extractors, body.Total, ei, err
}

// GetExtractor returns an extractor.
func (client *Client) GetExtractor(
	ctx context.Context, inputID, extractorID string,
) (
	*graylog.Extractor, *ErrorInfo, error,
) {
	if inputID == "" {
		return nil, nil, errors.New("input id is required")
	}
	if extractorID == "" {
		return nil, nil, errors.New("extractor id is required")
	}
	ext := &graylog.Extractor{}
	ei, err := client.callGet(ctx, client.Endpoints().Extractor(inputID, extractorID), nil, ext)
	return ext, ei, err
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

// CreateExtractor adds an extractor to an input.
func (client *Client) CreateExtractor(
	ctx context.Context, inputID string, extractor *graylog.Extractor,
) (
	*ErrorInfo, error,
) {
	if inputID == "" {
		return nil, errors.New("input id is required")
	}
	resp := map[string]string{}
	ei, err := client.callPost(
		ctx, client.Endpoints().Extractors(inputID), convertExtractorForPostAndPut(extractor), &resp)
	if err != nil {
		return ei, err
	}
	id, ok := resp["extractor_id"]
	if !ok {
		return ei, errors.New(`response doesn't have the field "extractor_id""`)
	}
	extractor.ID = id
	return ei, nil
}

// UpdateExtractor updates an extractor.
func (client *Client) UpdateExtractor(
	ctx context.Context, inputID string, extractor *graylog.Extractor,
) (
	*ErrorInfo, error,
) {
	if inputID == "" {
		return nil, errors.New("input id is required")
	}
	return client.callPut(
		ctx, client.Endpoints().Extractor(inputID, extractor.ID),
		convertExtractorForPostAndPut(extractor), extractor)
}

// DeleteExtractor updates an extractor.
func (client *Client) DeleteExtractor(
	ctx context.Context, inputID, extractorID string,
) (
	*ErrorInfo, error,
) {
	if inputID == "" {
		return nil, errors.New("input id is required")
	}
	if extractorID == "" {
		return nil, errors.New("extractor id is required")
	}
	return client.callDelete(ctx, client.Endpoints().Extractor(inputID, extractorID), nil, nil)
}
