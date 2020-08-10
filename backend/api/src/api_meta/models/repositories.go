package models

type (
	UpdateModel struct {
		Hash      string      `json:"hash" validation:"md5-hash"`
		FieldName string      `json:"field_name" validation:"model-field-name"`
		NewValue  interface{} `json:"new_value"`
	}
)
