package proto

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestParseOption(t *testing.T) {
	// New a new parser
	ctx := NewContext()

	var (
		validationFieldValidationTypeEnum = &Enum{
			enumName: "FieldValidationType",
			enumValues: []*EnumValue{
				{identifier: "ALPHA", number: "0"},
				{identifier: "ALPHA_NUM", number: "1"},
				{identifier: "ALPHA_NUM_UNICODE", number: "2"},
				{identifier: "ALPHA_UNICODE", number: "3"},
				{identifier: "ASCII", number: "4"},
				{identifier: "MULTIBYTE", number: "6"},
				{identifier: "PRINT_ASCII", number: "7"},
				{identifier: "LOWERCASE", number: "5"},
				{identifier: "UPPERCASE", number: "8"},
				{identifier: "BASE64", number: "9"},
				{identifier: "DATETIME", number: "10"},
				{identifier: "HEXADECIMAL", number: "11"},
				{identifier: "HEX_COLOR", number: "12"},
				{identifier: "LATITUDE", number: "13"},
				{identifier: "LONGITUDE", number: "14"},
				{identifier: "UUID", number: "15"},
			},
		}
		validationComparableValidationTypeEnum = &Enum{
			enumName: "ComparableValidationType",
			enumValues: []*EnumValue{
				{identifier: "EQ", number: "0"},
				{identifier: "GT", number: "1"},
				{identifier: "GTE", number: "2"},
				{identifier: "LT", number: "3"},
				{identifier: "LTE", number: "4"},
				{identifier: "NE", number: "5"},
				{identifier: "MAX", number: "6"},
				{identifier: "MIN", number: "7"},
			},
		}
		validationFormatComparisonMessage = &Message{
			msgName: "Comparison",
			fields: []*Field{
				{
					fieldName:   "comparator",
					fieldNumber: "4",
					t:           enumRef(ctx, "validation.", validationComparableValidationTypeEnum),
				},
				{
					fieldName:   "value",
					fieldNumber: "5",
					t:           NativeStringField,
				},
			},
			Options: nil,
		}
		validationFormatMessage = &Message{
			msgName:  "FormatValidation",
			messages: []*Message{validationFormatComparisonMessage},
			fields: []*Field{
				{
					fieldName:   "required",
					fieldNumber: "1",
					t:           NativeBoolField,
				},
				{
					fieldName:   "type",
					fieldNumber: "2",
					t:           enumRef(ctx, "validation.", validationFieldValidationTypeEnum),
				},
				{
					fieldName:   "compare",
					fieldNumber: "3",
					isRepeated:  true,
					t:           msgRef(ctx, "validation.", validationFormatComparisonMessage),
				},
			},
		}
	)

	optionText := "{required:true,type:ALPHA_NUM_UNICODE}"
	parsedOption := ParseOption(optionText, msgRef(ctx, "validation.", validationFormatMessage))
	//assert.NilError(t, err)
	assert.DeepEqual(t, parsedOption, map[string]interface{}{
		"type":     "ALPHA_NUM_UNICODE",
		"required": true,
	})
}

func TestParseOptionArray(t *testing.T) {
	// New a new parser
	ctx := NewContext()

	var (
		validationFieldValidationTypeEnum = &Enum{
			enumName: "FieldValidationType",
			enumValues: []*EnumValue{
				{identifier: "ALPHA", number: "0"},
				{identifier: "ALPHA_NUM", number: "1"},
				{identifier: "ALPHA_NUM_UNICODE", number: "2"},
				{identifier: "ALPHA_UNICODE", number: "3"},
				{identifier: "ASCII", number: "4"},
				{identifier: "MULTIBYTE", number: "6"},
				{identifier: "PRINT_ASCII", number: "7"},
				{identifier: "LOWERCASE", number: "5"},
				{identifier: "UPPERCASE", number: "8"},
				{identifier: "BASE64", number: "9"},
				{identifier: "DATETIME", number: "10"},
				{identifier: "HEXADECIMAL", number: "11"},
				{identifier: "HEX_COLOR", number: "12"},
				{identifier: "LATITUDE", number: "13"},
				{identifier: "LONGITUDE", number: "14"},
				{identifier: "UUID", number: "15"},
			},
		}
		validationComparableValidationTypeEnum = &Enum{
			enumName: "ComparableValidationType",
			enumValues: []*EnumValue{
				{identifier: "EQ", number: "0"},
				{identifier: "GT", number: "1"},
				{identifier: "GTE", number: "2"},
				{identifier: "LT", number: "3"},
				{identifier: "LTE", number: "4"},
				{identifier: "NE", number: "5"},
				{identifier: "MAX", number: "6"},
				{identifier: "MIN", number: "7"},
			},
		}
		validationFormatComparisonMessage = &Message{
			msgName: "Comparison",
			fields: []*Field{
				{
					fieldName:   "comparator",
					fieldNumber: "4",
					t:           enumRef(ctx, "validation.", validationComparableValidationTypeEnum),
				},
				{
					fieldName:   "value",
					fieldNumber: "5",
					t:           NativeStringField,
				},
			},
			Options: nil,
		}
		validationFormatMessage = &Message{
			msgName:  "FormatValidation",
			messages: []*Message{validationFormatComparisonMessage},
			fields: []*Field{
				{
					fieldName:   "required",
					fieldNumber: "1",
					t:           NativeBoolField,
				},
				{
					fieldName:   "type",
					fieldNumber: "2",
					t:           enumRef(ctx, "validation.", validationFieldValidationTypeEnum),
				},
				{
					fieldName:   "compare",
					fieldNumber: "3",
					isRepeated:  true,
					t:           msgRef(ctx, "validation.", validationFormatComparisonMessage),
				},
			},
		}
	)

	optionArrayText := `{required:false;type:ALPHA_NUM_UNICODE;compare:[{comparator:MIN;value:"0";},{comparator:MAX;value:"100";}]}`
	parsedOption := ParseOption(optionArrayText, msgRef(ctx, "validation.", validationFormatMessage))
	//assert.NilError(t, err)
	assert.DeepEqual(t, parsedOption, map[string]interface{}{
		"type":     "ALPHA_NUM_UNICODE",
		"required": false,
		"compare": []interface{}{
			map[string]interface{}{
				"comparator": "MIN",
				"value":      "0",
			},
			map[string]interface{}{
				"comparator": "MAX",
				"value":      "100",
			},
		},
	})
}
