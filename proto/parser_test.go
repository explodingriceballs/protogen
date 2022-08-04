package proto

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSimpleProtobuf(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	var (
		productionMessage = func(c *Context) *Message {
			return &Message{
				types:       c.typeDictionary,
				MessageName: "Production",
				Fields: []*Field{
					{
						types:       c.typeDictionary,
						FieldName:   "id",
						FieldNumber: "1",
						IsRepeated:  false,
						Type:        nativeTypes["string"],
					},
					{
						types:       c.typeDictionary,
						FieldName:   "title",
						FieldNumber: "2",
						IsRepeated:  false,
						Type:        nativeTypes["string"],
					},
				},
			}
		}
		getProductionRequestMessage = func(c *Context) *Message {
			return &Message{
				types:       c.typeDictionary,
				MessageName: "GetProductionRequest",
				Fields: []*Field{
					{
						types:       c.typeDictionary,
						FieldName:   "production_id",
						FieldNumber: "1",
						IsRepeated:  false,
						Type:        nativeTypes["string"],
					},
				},
			}
		}
		getProductionResponseMessage = func(c *Context) *Message {
			return &Message{
				types:       c.typeDictionary,
				MessageName: "GetProductionResponse",
				Fields: []*Field{
					{
						types:       c.typeDictionary,
						FieldName:   "production",
						FieldNumber: "1",
						IsRepeated:  false,
						Type:        c.typeDictionary.newMsgTypeForTesting(productionMessage(c)),
					},
				},
			}
		}
		productionService = func(c *Context) *Service {
			return &Service{
				types:       c.typeDictionary,
				ServiceName: "ProductionService",
				RPCs: []*RPC{
					{
						types:               c.typeDictionary,
						RPCName:             "GetProduction",
						RequestMessageType:  c.typeDictionary.newMsgTypeForTesting(getProductionRequestMessage(c)),
						ResponseMessageType: c.typeDictionary.newMsgTypeForTesting(getProductionResponseMessage(c)),
					},
				},
			}
		}
	)

	// Create a new parser
	parser := NewParser([]string{"./testdata/simple.proto"}, []string{})

	// Parse
	err := parser.Parse()
	assert.NoError(t, err)

	ctx := parser.GetContext()
	assert.Len(t, ctx.Packages(), 0)
	assert.Len(t, ctx.GlobalScope().Messages(), 3)

	// Production message
	assert.Equal(t, productionMessage(ctx), ctx.GlobalScope().Messages()[0])
	assert.Equal(t, getProductionRequestMessage(ctx), ctx.GlobalScope().Messages()[1])
	assert.Equal(t, getProductionResponseMessage(ctx), ctx.GlobalScope().Messages()[2])

	// Production service
	assert.Equal(t, productionService(ctx), ctx.GlobalScope().Services()[0])
}

func TestComplexProtobuf(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	var (
		httpCustomPatternMessage = func(c *Context) *Message {
			return &Message{
				types:       c.typeDictionary,
				MessageName: "CustomHttpPattern",
				Fields: []*Field{
					{
						types:       c.typeDictionary,
						FieldName:   "kind",
						FieldNumber: "1",
						IsRepeated:  false,
						Type:        nativeTypes["string"],
					},
					{
						types:       c.typeDictionary,
						FieldName:   "path",
						FieldNumber: "2",
						IsRepeated:  false,
						Type:        nativeTypes["string"],
					},
				},
				Options: []*Option{
					{
						OptionName:  "(skip_generation)",
						OptionValue: "true",
					},
				},
			}
		}
		httpRuleMessage = func(c *Context) *Message {
			httpRulePattern := &Message{
				types:       c.typeDictionary,
				MessageName: "pattern",
				OneOf:       true,
				Fields: []*Field{
					{
						types:       c.typeDictionary,
						FieldName:   "get",
						FieldNumber: "2",
						IsRepeated:  false,
						Type:        nativeTypes["string"],
					},
					{
						types:       c.typeDictionary,
						FieldName:   "put",
						FieldNumber: "3",
						IsRepeated:  false,
						Type:        nativeTypes["string"],
					},
					{
						types:       c.typeDictionary,
						FieldName:   "post",
						FieldNumber: "4",
						IsRepeated:  false,
						Type:        nativeTypes["string"],
					},
					{
						types:       c.typeDictionary,
						FieldName:   "delete",
						FieldNumber: "5",
						IsRepeated:  false,
						Type:        nativeTypes["string"],
					},
					{
						types:       c.typeDictionary,
						FieldName:   "patch",
						FieldNumber: "6",
						IsRepeated:  false,
						Type:        nativeTypes["string"],
					},
					{
						types:       c.typeDictionary,
						FieldName:   "custom",
						FieldNumber: "8",
						IsRepeated:  false,
						Type:        c.typeDictionary.newMsgTypeForTesting(httpCustomPatternMessage(c)),
					},
				},
			}

			return &Message{
				types:       c.typeDictionary,
				MessageName: "HttpRule",
				Fields: []*Field{
					{
						types:       c.typeDictionary,
						FieldName:   "selector",
						FieldNumber: "1",
						IsRepeated:  false,
						Type:        nativeTypes["string"],
					},
					{
						types:     c.typeDictionary,
						FieldName: "pattern",
						Type:      c.typeDictionary.newMsgTypeForTesting(httpRulePattern),
					},
					{
						types:       c.typeDictionary,
						FieldName:   "body",
						FieldNumber: "7",
						IsRepeated:  false,
						Type:        nativeTypes["string"],
					},
					{
						types:       c.typeDictionary,
						FieldName:   "response_body",
						FieldNumber: "12",
						IsRepeated:  false,
						Type:        nativeTypes["string"],
					},
					{
						types:       c.typeDictionary,
						FieldName:   "additional_bindings",
						FieldNumber: "11",
						IsRepeated:  true,
					},
					{
						types:       c.typeDictionary,
						FieldName:   "requires_auth",
						FieldNumber: "25",
						IsRepeated:  false,
						Type:        nativeTypes["bool"],
					},
				},
				Options: []*Option{
					{
						OptionName:  "(skip_generation)",
						OptionValue: "true",
					},
				},
			}
		}
		validationFieldValidationTypeEnum = func(c *Context) *Enum {
			return &Enum{
				types:    c.typeDictionary,
				EnumName: "FieldValidationType",
				EnumValues: []*EnumValue{
					{Identifier: "ALPHA", Number: "0"},
					{Identifier: "ALPHA_NUM", Number: "1"},
					{Identifier: "ALPHA_NUM_UNICODE", Number: "2"},
					{Identifier: "ALPHA_UNICODE", Number: "3"},
					{Identifier: "ASCII", Number: "4"},
					{Identifier: "MULTIBYTE", Number: "6"},
					{Identifier: "PRINT_ASCII", Number: "7"},
					{Identifier: "LOWERCASE", Number: "5"},
					{Identifier: "UPPERCASE", Number: "8"},
					{Identifier: "BASE64", Number: "9"},
					{Identifier: "DATETIME", Number: "10"},
					{Identifier: "HEXADECIMAL", Number: "11"},
					{Identifier: "HEX_COLOR", Number: "12"},
					{Identifier: "LATITUDE", Number: "13"},
					{Identifier: "LONGITUDE", Number: "14"},
					{Identifier: "UUID", Number: "15"},
				},
			}
		}
		validationComparableValidationTypeEnum = func(c *Context) *Enum {
			return &Enum{
				types:    c.typeDictionary,
				EnumName: "ComparableValidationType",
				EnumValues: []*EnumValue{
					{Identifier: "EQ", Number: "0"},
					{Identifier: "GT", Number: "1"},
					{Identifier: "GTE", Number: "2"},
					{Identifier: "LT", Number: "3"},
					{Identifier: "LTE", Number: "4"},
					{Identifier: "NE", Number: "5"},
					{Identifier: "MAX", Number: "6"},
					{Identifier: "MIN", Number: "7"},
				},
			}
		}
		validationFormatComparisonMessage = func(c *Context) *Message {
			return &Message{
				types:       c.typeDictionary,
				MessageName: "Comparison",
				OneOf:       false,
				Fields: []*Field{
					{
						types:       c.typeDictionary,
						FieldName:   "comparator",
						FieldNumber: "4",
						IsRepeated:  false,
						Type:        c.typeDictionary.newEnumTypeForTesting(validationComparableValidationTypeEnum(c)),
					},
					{
						types:       c.typeDictionary,
						FieldName:   "value",
						FieldNumber: "5",
						IsRepeated:  false,
						Type:        nativeTypes["string"],
					},
				},
				Options: nil,
			}
		}
		validationFormatMessage = func(c *Context) *Message {
			return &Message{
				types:       c.typeDictionary,
				MessageName: "FormatValidation",
				Messages:    []*Message{validationFormatComparisonMessage(c)},
				Fields: []*Field{
					{
						types:       c.typeDictionary,
						FieldName:   "required",
						FieldNumber: "1",
						IsRepeated:  false,
						Type:        nativeTypes["bool"],
					},
					{
						types:       c.typeDictionary,
						FieldName:   "type",
						FieldNumber: "2",
						IsRepeated:  false,
						Type:        c.typeDictionary.newEnumTypeForTesting(validationFieldValidationTypeEnum(c)),
					},
					{
						types:       c.typeDictionary,
						FieldName:   "compare",
						FieldNumber: "3",
						IsRepeated:  true,
						Type:        c.typeDictionary.newMsgTypeForTesting(validationFormatComparisonMessage(c)),
					},
					// TODO: enum & nested message
				},
			}
		}
	)

	// Create a new parser
	parser := NewParser([]string{
		"./testdata/complex.proto",
	}, []string{
		"./testdata/",
		os.Getenv("PROTOBUF_INCLUDE"),
	})

	// Parse
	err := parser.Parse()
	assert.NoError(t, err)

	ctx := parser.GetContext()
	assert.Len(t, ctx.Packages(), 3)
	assert.Len(t, ctx.GlobalScope().Messages(), 0)

	// Update self references for messages
	httpRuleMessageRef := httpRuleMessage(ctx)
	httpRuleMessageRef.Field("additional_bindings").Type = ctx.typeDictionary.newMsgTypeForTesting(httpRuleMessageRef)

	// Verify the HTTP package
	httpPackage := ctx.GetPackage("http")
	assert.Len(t, httpPackage.Messages(), 2)
	assert.Equal(t, httpCustomPatternMessage(ctx), httpPackage.Message("CustomHttpPattern"))
	assert.Equal(t, httpRuleMessageRef, httpPackage.Message("HttpRule"))

	// Verify the validation package
	validationPackage := ctx.GetPackage("validation")
	assert.Len(t, validationPackage.Enums(), 2)
	assert.Equal(t, validationFieldValidationTypeEnum(ctx), validationPackage.Enum("FieldValidationType"))
	assert.Equal(t, validationComparableValidationTypeEnum(ctx), validationPackage.Enum("ComparableValidationType"))

	assert.Len(t, validationPackage.Messages(), 1)
	assert.Equal(t, validationFormatMessage(ctx), validationPackage.Message("FormatValidation"))
	assert.Equal(t, validationFormatComparisonMessage(ctx), validationPackage.Message("FormatValidation").Message("Comparison"))
}

//	// Create a new parser
//	parser := NewParser([]string{"./testdata/complex.proto"}, []string{"./testdata", os.Getenv("PROTOBUF_INCLUDE")})
//
//	// Parse
//	err := parser.Parse()
//	assert.NoError(t, err)
//
//	ctx := parser.GetContext()
//	assert.Len(t, ctx.GetPackages(), 4)
//
//	apiPackage := ctx.GetPackages()[3]
//	assert.Equal(t, "google.protobuf", ctx.GetPackages()[0].Name)
//	assert.Equal(t, "http", ctx.GetPackages()[1].Name)
//	assert.Equal(t, "validation", ctx.GetPackages()[2].Name)
//	assert.Equal(t, "api", apiPackage.Name)
//
//	assert.Equal(t, []*Enum{
//		enumDataStatus,
//		enumDataFooType,
//	}, apiPackage.Enums())
//
//	nestedMessageOne := NewMessageWithContext(ctx, "NestedMessageOne")
//	nestedMessageOne.SetFields(
//		&Field{
//			ctx:              ctx,
//			FieldName:        "type",
//			FieldNumber:      "1",
//			IsRepeated:       false,
//			FieldType:        EnumFieldType,
//			MessageFieldType: nil,
//			EnumFieldType:    enumDataFooType,
//		},
//		&Field{
//			ctx:              ctx,
//			FieldName:        "one_type",
//			FieldNumber:      "",
//			IsRepeated:       false,
//			FieldType:        OneOfType,
//			NativeFieldType:  "",
//			MessageFieldType: nil,
//			EnumFieldType:    nil,
//			OneOfType:        oneOfOneType(ctx),
//		})
//	assert.Equal(t, nestedMessageOne, apiPackage.Messages()[1])
//
//	nestedMessageTwo := NewMessageWithContext(ctx, "NestedMessageTwo")
//	nestedMessageTwo.SetFields(
//		&Field{
//			ctx:              ctx,
//			FieldName:        "prop1",
//			FieldNumber:      "1",
//			IsRepeated:       false,
//			FieldType:        NativeFieldType,
//			NativeFieldType:  "string",
//			MessageFieldType: nil,
//		},
//		&Field{
//			ctx:              ctx,
//			FieldName:        "prop2",
//			FieldNumber:      "2",
//			IsRepeated:       false,
//			FieldType:        NativeFieldType,
//			NativeFieldType:  "string",
//			MessageFieldType: nil,
//		},
//		&Field{
//			ctx:              ctx,
//			FieldName:        "prop3",
//			FieldNumber:      "3",
//			IsRepeated:       false,
//			FieldType:        NativeFieldType,
//			NativeFieldType:  "string",
//			MessageFieldType: nil,
//		},
//		&Field{
//			ctx:              ctx,
//			FieldName:        "prop4",
//			FieldNumber:      "4",
//			IsRepeated:       false,
//			FieldType:        NativeFieldType,
//			NativeFieldType:  "string",
//			MessageFieldType: nil,
//		})
//	assert.Contains(t, apiPackage.Messages(), nestedMessageTwo)
//
//	//assert.Equal(t, []*Message{
//	//	{},
//	//}, apiPackage.Messages())
//}
//
//var (
//	enumDataFooType = &Enum{
//		EnumName: "DataFooType",
//		Values: []*EnumValue{
//			{
//				Identifier: "DATA_FOO_TYPE_UNKNOWN",
//				Number:     0,
//			},
//			{
//				Identifier: "DATA_FOO_TYPE_TIMED",
//				Number:     1,
//			},
//			{
//				Identifier: "DATA_FOO_TYPE_DEPENDENCY",
//				Number:     2,
//			},
//		},
//	}
//	enumDataStatus = &Enum{
//		EnumName: "DataStatus",
//		Options: Options{
//			options: []Option{
//				{
//					OptionName:  "(validation.test_bool)",
//					OptionValue: "true",
//				},
//			},
//		},
//		Values: []*EnumValue{
//			{
//				Identifier: "DATA_STATUS_UNKNOWN",
//				Number:     0,
//				Options: Options{
//					options: []Option{
//						{
//							OptionName:  "deprecated",
//							OptionValue: "true",
//						},
//					},
//				},
//			},
//			{
//				Identifier: "DATA_STATUS_SCHEDULED",
//				Number:     1,
//			},
//			{
//				Identifier: "DATA_STATUS_RUNNING",
//				Number:     2,
//			},
//			{
//				Identifier: "DATA_STATUS_FAILED",
//				Number:     3,
//			},
//			{
//				Identifier: "DATA_STATUS_SUSPENDED",
//				Number:     4,
//			},
//		},
//	}
//	oneOfOneType = func(ctx *Context) *OneOf {
//		nestedOneOfMessage := NewMessageWithContext(ctx, "NestedOneOfMessage")
//		nestedOneOfMessage.SetFields(&Field{
//			ctx:              ctx,
//			FieldName:        "value",
//			FieldNumber:      "1",
//			IsRepeated:       true,
//			FieldType:        NativeFieldType,
//			NativeFieldType:  "string",
//			MessageFieldType: nil,
//			EnumFieldType:    nil,
//			OneOfType:        nil,
//		})
//
//		return &OneOf{
//			OneOfName: "one_type",
//			Fields: []*Field{
//				{
//					ctx:              ctx,
//					FieldName:        "string_property",
//					FieldNumber:      "2",
//					IsRepeated:       false,
//					FieldType:        NativeFieldType,
//					NativeFieldType:  "string",
//					MessageFieldType: nil,
//					EnumFieldType:    nil,
//					OneOfType:        nil,
//				},
//				{
//					ctx:              ctx,
//					FieldName:        "nested_message",
//					FieldNumber:      "3",
//					IsRepeated:       false,
//					FieldType:        MessageFieldType,
//					NativeFieldType:  "",
//					MessageFieldType: nestedOneOfMessage,
//					EnumFieldType:    nil,
//					OneOfType:        nil,
//				},
//			},
//		}
//	}
//)
