package graylog

type (
	// EventDefinition represents an event definition.
	EventDefinition struct {
		ID          string                              `json:"id"`
		Title       string                              `json:"title"`
		Description string                              `json:"description"`
		Priority    int                                 `json:"priority"`
		Alert       bool                                `json:"alert"`
		FieldSpec   map[string]EventDefinitionFieldSpec `json:"field_spec"`
		// KeySpec     []interface{}                       `json:"key_spec"`
		NotificationSettings EventDefinitionNotificationSettings `json:"notification_settings"`
		Notifications        []EventDefinitionNotification       `json:"notifications"`
		Storage              []interface{}                       `json:"storage"`
		Config               interface{}                         `json:"config"`
	}

	//	EventDefinitionStorage struct {
	//		Type    string   `json:"type"`
	//		Streams []string `json:"streams"`
	//	}

	EventDefinitionNotification struct {
		NotificationID string `json:"notification_id"`
		// NotificationParameters interface{} `json:"notification_parameters"`
	}

	EventDefinitionNotificationSettings struct {
		GracePeriodMS int `json:"grace_period_ms"`
		BacklogSize   int `json:"backlog_size"`
	}

	EventDefinitionFieldSpec struct {
		DataType  string        `json:"data_type"`
		Providers []interface{} `json:"providers"`
	}

	//	EventDefinitionFieldSpecProvider struct {
	//		Type          string `json:"type"`
	//		Template      string `json:"template"`
	//		RequireValues bool   `json:"require_values"`
	//	}

	//	EventDefinitionConfig struct {
	//		Type    string      `json:"type"`
	//		Query   string      `json:"query"`
	//		Streams *set.StrSet `json:"streams"`
	//		// GroupBy []string    `json:"group_by"`
	//		Series     *EventDefinitionConfigSeries    `json:"series"`
	//		Conditions EventDefinitionConfigConditions `json:"conditions"`
	//	}

	EventDefinitionConfigSeries struct {
		ID       string `json:"id"`
		Function string `json:"function"`
		// Field string `json:"field"`
	}

	EventDefinitionConfigConditions struct {
		Expression EventDefinitionConfigConditionsExpression `json:"expression"`
	}

	EventDefinitionConfigConditionsExpression struct {
		Expr  string                                          `json:"expr"`
		Left  *EventDefinitionConfigConditionsExpressionLeft  `json:"left"`
		Right *EventDefinitionConfigConditionsExpressionRight `json:"right"`
	}

	EventDefinitionConfigConditionsExpressionLeft struct {
		Expr string `json:"expr"`
		Ref  string `json:"ref"`
	}

	EventDefinitionConfigConditionsExpressionRight struct {
		Expr  string `json:"expr"`
		Value string `json:"value"`
	}

	// EventDefinitionsBody represents Get EventDefinitions API's response body.
	// Basically users don't use this struct, but this struct is public because some sub packages use this struct.
	EventDefinitionsBody struct {
		EventDefinitions []EventDefinition `json:"event_definitions"`
		Total            int               `json:"total"`
		Page             int               `json:"page"`
		PerPage          int               `json:"per_page"`
		Count            int               `json:"count"`
		GrandTotal       int               `json:"grand_total"`
		Query            string            `json:"query"`
	}
)
