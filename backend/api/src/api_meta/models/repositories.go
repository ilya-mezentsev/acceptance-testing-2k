package models

type (
	UpdateModel struct {
		Hash      string      `json:"hash"`
		FieldName string      `json:"field_name"`
		NewValue  interface{} `json:"new_value"`
	}
)
