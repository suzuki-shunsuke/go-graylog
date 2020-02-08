package endpoint

// InputStaticFields returns the StaticFields API's endpoint url.
func (ep *Endpoints) InputStaticFields(inputID string) string {
	// /system/inputs/{inputId}/staticfields
	return ep.inputs + "/" + inputID + "/staticfields"
}

// InputStaticField returns the StaticFields API's endpoint url.
func (ep *Endpoints) InputStaticField(inputID, key string) string {
	// /system/inputs/{inputId}/staticfields/{key}
	return ep.inputs + "/" + inputID + "/staticfields/" + key
}
