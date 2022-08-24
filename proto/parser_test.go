package proto

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gotest.tools/v3/assert"
	"os"
	"testing"
)

// Shortcut vars
var (
	NativeStringField = nativeTypes["string"]
	NativeInt32Field  = nativeTypes["int32"]
	NativeInt64Field  = nativeTypes["int64"]
	NativeBoolField   = nativeTypes["bool"]
)

// options for the comparer
var options = []cmp.Option{
	cmp.AllowUnexported(Type{}, Service{}, RPC{}, Message{}, Field{}, Options{}, Option{}, Enum{}, EnumValue{}),
	cmpopts.IgnoreFields(Message{}, "types", "msgType"),
	cmpopts.IgnoreFields(Field{}, "types"),
	cmpopts.IgnoreFields(Service{}, "types"),
	cmpopts.IgnoreFields(RPC{}, "types"),
	cmpopts.IgnoreFields(Enum{}, "types"),
	cmpopts.IgnoreFields(Options{}, "types", "parent"),
}

func TestSimpleProtobuf(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	// New a new parser
	parser := NewParser([]string{"./testdata/simple.proto"}, []string{})

	// Parse
	err := parser.Parse()
	assert.NilError(t, err)

	ctx := parser.GetContext()

	var (
		productionMessage = &Message{
			msgName: "Production",
			fields: []*Field{
				{
					fieldName:   "id",
					fieldNumber: "1",
					t:           NativeStringField,
					Options:     &Options{},
				},
				{
					fieldName:   "title",
					fieldNumber: "2",
					t:           NativeStringField,
					Options:     &Options{},
				},
			},
			Options: &Options{},
		}
		getProductionRequestMessage = &Message{
			msgName: "GetProductionRequest",
			fields: []*Field{
				{
					fieldName:   "production_id",
					fieldNumber: "1",
					t:           NativeStringField,
					Options:     &Options{},
				},
			},
			Options: &Options{},
		}
		getProductionResponseMessage = &Message{
			msgName: "GetProductionResponse",
			fields: []*Field{
				{
					fieldName:   "production",
					fieldNumber: "1",
					t:           msgRef(ctx, ".", productionMessage),
					Options:     &Options{},
				},
			},
			Options: &Options{},
		}
		productionService = &Service{
			serviceName: "ProductionService",
			rpcs: []*RPC{
				{
					RPCName:             "GetProduction",
					RequestMessageType:  msgRef(ctx, ".", getProductionRequestMessage),
					ResponseMessageType: msgRef(ctx, ".", getProductionResponseMessage),
					Options:             &Options{},
				},
			},
			Options: &Options{},
		}
	)

	assert.Assert(t, len(ctx.Packages()) == 0)
	assert.Assert(t, len(ctx.GlobalScope().Messages()) == 3)

	// Production message
	assert.DeepEqual(t, productionMessage, ctx.GlobalScope().Message("Production"), options...)
	assert.DeepEqual(t, getProductionRequestMessage, ctx.GlobalScope().Message("GetProductionRequest"), options...)
	assert.DeepEqual(t, getProductionResponseMessage, ctx.GlobalScope().Message("GetProductionResponse"), options...)

	// Production service
	assert.DeepEqual(t, productionService, ctx.GlobalScope().Service("ProductionService"), options...)
}

func TestComplexProtobuf(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	// New a new parser
	parser := NewParser([]string{
		"./testdata/complex.proto",
	}, []string{
		"./testdata/",
		os.Getenv("PROTOBUF_INCLUDE"),
	})

	// Parse
	err := parser.Parse()
	assert.NilError(t, err)

	ctx := parser.GetContext()
	assert.Assert(t, len(ctx.Packages()) == 4)
	assert.Assert(t, len(ctx.GlobalScope().Messages()) == 0)

	var (
		httpCustomPatternMessage = &Message{
			msgName: "CustomHttpPattern",
			fields: []*Field{
				{
					fieldName:   "kind",
					fieldNumber: "1",
					t:           NativeStringField,
					Options:     &Options{},
				},
				{
					fieldName:   "path",
					fieldNumber: "2",
					t:           NativeStringField,
					Options:     &Options{},
				},
			},
			Options: &Options{
				Options: []*Option{
					{
						OptionName:  "(http.skip_generation)",
						OptionValue: "true",
					},
				},
			},
		}
		httpRulePatternMessage = &Message{
			msgName: "pattern",
			oneOf:   true,
			fields: []*Field{
				{
					fieldName:   "get",
					fieldNumber: "2",
					t:           NativeStringField,
					Options:     &Options{},
				},
				{
					fieldName:   "put",
					fieldNumber: "3",
					t:           NativeStringField,
					Options:     &Options{},
				},
				{
					fieldName:   "post",
					fieldNumber: "4",
					t:           NativeStringField,
					Options:     &Options{},
				},
				{
					fieldName:   "delete",
					fieldNumber: "5",
					t:           NativeStringField,
					Options:     &Options{},
				},
				{
					fieldName:   "patch",
					fieldNumber: "6",
					t:           NativeStringField,
					Options:     &Options{},
				},
				{
					fieldName:   "custom",
					fieldNumber: "8",
					t:           msgRef(ctx, "http.", httpCustomPatternMessage),
					Options:     &Options{},
				},
			},
			Options: &Options{},
		}
		httpRuleMessage = &Message{
			msgName: "HttpRule",
			fields: []*Field{
				{
					fieldName:   "selector",
					fieldNumber: "1",
					t:           NativeStringField,
					Options:     &Options{},
				},
				{
					fieldName: "pattern",
					t:         msgRef(ctx, "http.HttpRule.", httpRulePatternMessage),
					Options:   &Options{},
				},
				{
					fieldName:   "body",
					fieldNumber: "7",
					t:           NativeStringField,
					Options:     &Options{},
				},
				{
					fieldName:   "response_body",
					fieldNumber: "12",
					t:           NativeStringField,
					Options:     &Options{},
				},
				{
					fieldName:   "additional_bindings",
					fieldNumber: "11",
					isRepeated:  true,
					Options:     &Options{},
					//t: is set later
				},
				{
					fieldName:   "requires_auth",
					fieldNumber: "25",
					t:           NativeBoolField,
					Options:     &Options{},
				},
			},
			Options: &Options{
				Options: []*Option{
					{
						OptionName:  "(http.skip_generation)",
						OptionValue: "true",
					},
				},
			},
		}
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
			Options: &Options{},
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
			Options: &Options{},
		}
		validationFormatComparisonMessage = &Message{
			msgName: "Comparison",
			fields: []*Field{
				{
					fieldName:   "comparator",
					fieldNumber: "4",
					t:           enumRef(ctx, "validation.", validationComparableValidationTypeEnum),
					Options:     &Options{},
				},
				{
					fieldName:   "value",
					fieldNumber: "5",
					t:           NativeStringField,
					Options:     &Options{},
				},
			},
			Options: &Options{},
		}
		validationFormatMessage = &Message{
			msgName:  "FormatValidation",
			messages: []*Message{validationFormatComparisonMessage},
			fields: []*Field{
				{
					fieldName:   "required",
					fieldNumber: "1",
					t:           NativeBoolField,
					Options:     &Options{},
				},
				{
					fieldName:   "type",
					fieldNumber: "2",
					t:           enumRef(ctx, "validation.", validationFieldValidationTypeEnum),
					Options:     &Options{},
				},
				{
					fieldName:   "compare",
					fieldNumber: "3",
					isRepeated:  true,
					t:           msgRef(ctx, "validation.FormatValidation.", validationFormatComparisonMessage),
					Options:     &Options{},
				},
			},
			Options: &Options{},
		}
		apiDataStatusEnum = &Enum{
			enumName: "DataStatus",
			enumValues: []*EnumValue{
				{identifier: "DATA_STATUS_UNKNOWN", number: "0"},
				{identifier: "DATA_STATUS_SCHEDULED", number: "1"},
				{identifier: "DATA_STATUS_RUNNING", number: "2"},
				{identifier: "DATA_STATUS_FAILED", number: "3"},
				{identifier: "DATA_STATUS_SUSPENDED", number: "4"},
			},
			Options: &Options{
				Options: []*Option{
					{
						OptionName:  "(validation.test_bool)",
						OptionValue: "true",
					},
				},
			},
		}
		apiDataFooType = &Enum{
			enumName: "DataFooType",
			enumValues: []*EnumValue{
				{identifier: "DATA_FOO_TYPE_UNKNOWN", number: "0"},
				{identifier: "DATA_FOO_TYPE_TIMED", number: "1"},
				{identifier: "DATA_FOO_TYPE_DEPENDENCY", number: "2"},
			},
			Options: &Options{},
		}
		apiNestedMessageOneOneOfMsg = &Message{
			msgName: "NestedOneOfMessage",
			oneOf:   false,
			fields: []*Field{
				{
					fieldName:   "value",
					fieldNumber: "1",
					isRepeated:  true,
					t:           nativeTypes["string"],
					Options:     &Options{},
				},
			},
			Options:  &Options{},
			messages: nil,
		}
		apiNestedMessageOne = &Message{
			msgName: "NestedMessageOne",
			oneOf:   false,
			fields: []*Field{
				{
					fieldName:   "type",
					fieldNumber: "1",
					t:           enumRef(ctx, "api.", apiDataFooType),
					Options:     &Options{},
				},
				{
					fieldName: "one_type",
					t: msgRef(ctx, "api.NestedMessageOne.", &Message{
						msgName: "one_type",
						oneOf:   true,
						fields: []*Field{
							{
								fieldName:   "string_property",
								fieldNumber: "2",
								t:           NativeStringField,
								Options:     &Options{},
							},
							{
								fieldName:   "nested_message",
								fieldNumber: "3",
								t:           msgRef(ctx, "api.NestedMessageOne.", apiNestedMessageOneOneOfMsg),
								Options:     &Options{},
							},
						},
						Options: &Options{},
					}),
					Options: &Options{},
				},
			},
			Options:  &Options{},
			messages: []*Message{apiNestedMessageOneOneOfMsg},
		}
		apiNestedMessageTwo = &Message{
			Options: &Options{},
			msgName: "NestedMessageTwo",
			fields: []*Field{
				{
					fieldName:   "prop1",
					fieldNumber: "1",
					t:           NativeStringField,
					Options:     &Options{},
				},
				{
					fieldName:   "prop2",
					fieldNumber: "2",
					t:           NativeStringField,
					Options:     &Options{},
				},
				{
					fieldName:   "prop3",
					fieldNumber: "3",
					t:           NativeStringField,
					Options:     &Options{},
				},
				{
					fieldName:   "prop4",
					fieldNumber: "4",
					t:           NativeStringField,
					Options:     &Options{},
				},
			},
		}
		apiCreateDataRequest = &Message{
			msgName: "CreateDataRequest",
			oneOf:   false,
			fields: []*Field{
				{
					fieldName:   "name",
					fieldNumber: "1",
					isRepeated:  false,
					t:           nativeTypes["string"],
					Options: &Options{
						Options: []*Option{
							{
								OptionName:  "(validation.format).type",
								OptionValue: "ALPHA",
							},
						},
					},
				},
				{
					fieldName:   "repeated_string_property",
					fieldNumber: "2",
					isRepeated:  true,
					t:           NativeStringField,
					Options: &Options{
						Options: []*Option{
							{
								OptionName:  "(validation.format)",
								OptionValue: "{required:false,type:ALPHA_NUM_UNICODE}",
							},
						},
					},
				},
				{
					fieldName:   "msg_one",
					fieldNumber: "3",
					isRepeated:  false,
					t:           msgRef(ctx, "api.", apiNestedMessageOne),
					Options:     &Options{},
				},
				{
					fieldName:   "msg_two",
					fieldNumber: "4",
					t:           msgRef(ctx, "api.", apiNestedMessageTwo),
					Options:     &Options{},
				},
				{
					fieldName:   "field",
					fieldNumber: "5",
					isRepeated:  false,
					t:           NativeStringField,
					Options:     &Options{},
				},
			},
			Options:  &Options{},
			messages: nil,
		}
		apiDataListResponseData = &Message{
			msgName: "Data",
			fields: []*Field{
				{
					fieldName:   "id",
					fieldNumber: "1",
					t:           NativeStringField,
					Options:     &Options{},
				},
				{
					fieldName:   "name",
					fieldNumber: "2",
					t:           NativeStringField,
					Options:     &Options{},
				},
				{
					fieldName:   "status",
					fieldNumber: "3",
					t:           enumRef(ctx, "api.", apiDataStatusEnum),
					Options:     &Options{},
				},
				{
					fieldName:   "string_property",
					fieldNumber: "4",
					t:           NativeStringField,
					Options:     &Options{},
				},
				{
					fieldName:   "int_property",
					fieldNumber: "5",
					t:           NativeInt64Field,
					Options:     &Options{},
				},
			},
			Options:  &Options{},
			messages: nil,
		}
		apiDataListResponse = &Message{
			msgName: "DataListResponse",
			fields: []*Field{
				{
					fieldName:   "count",
					fieldNumber: "1",
					t:           NativeInt32Field,
					Options:     &Options{},
				},
				{
					fieldName:   "data",
					fieldNumber: "2",
					isRepeated:  true,
					t:           msgRef(ctx, "api.DataListResponse.", apiDataListResponseData),
					Options:     &Options{},
				},
			},
			Options: &Options{},
			messages: []*Message{
				apiDataListResponseData,
			},
		}
		apiRpcOne = &RPC{
			RPCName: "rpcOne",
			RequestMessageType: msgRef(ctx, "google.protobuf.", &Message{
				msgName: "Empty",
				Options: &Options{},
			}),
			ResponseMessageType: msgRef(ctx, "api.", apiDataListResponse),
			Options: &Options{
				Options: []*Option{
					{
						OptionName:  "(http.rule)",
						OptionValue: "{get:\"/v1/data\";body:\"*\";requires_auth:true;}",
					},
				},
			},
		}
	)

	// Test HTTP package
	{
		// Update self references for messages
		httpRuleMessage.Field("additional_bindings").t = msgRef(ctx, "http.", httpRuleMessage)

		// Verify the HTTP package
		httpPackage := ctx.Package("http")
		assert.Assert(t, len(httpPackage.Messages()) == 2)
		assert.DeepEqual(t, httpCustomPatternMessage, httpPackage.Message("CustomHttpPattern"), options...)
		assert.Equal(t, true, httpPackage.Message("CustomHttpPattern").Has("http.skip_generation"))
		assert.Equal(t, true, httpPackage.Message("CustomHttpPattern").Get("http.skip_generation"))
		assert.DeepEqual(t, httpRuleMessage, httpPackage.Message("HttpRule"), options...)
		assert.Equal(t, true, httpPackage.Message("HttpRule").Field("pattern").t.GetMessage().IsOneOf())
		assert.DeepEqual(t, httpRuleMessage.fields, httpPackage.Message("HttpRule").Fields(), options...)
	}

	{
		// Verify the validation package
		validationPackage := ctx.Package("validation")
		assert.Assert(t, len(validationPackage.Enums()) == 2)
		assert.DeepEqual(t, validationFieldValidationTypeEnum, validationPackage.Enum("FieldValidationType"), options...)
		assert.DeepEqual(t, validationComparableValidationTypeEnum, validationPackage.Enum("ComparableValidationType"), options...)

		assert.Assert(t, len(validationPackage.Messages()) == 1)
		assert.DeepEqual(t, validationFormatMessage, validationPackage.Message("FormatValidation"), options...)
		assert.DeepEqual(t, validationFormatMessage.fields, validationPackage.Message("FormatValidation").Fields(), options...)
		assert.DeepEqual(t, validationFormatComparisonMessage, validationPackage.Message("FormatValidation").Message("Comparison"), options...)
		assert.DeepEqual(t, validationFormatComparisonMessage.fields, validationPackage.Message("FormatValidation").Message("Comparison").Fields(), options...)
	}

	// Verify the API package
	apiPackage := ctx.Package("api")
	assert.Assert(t, len(apiPackage.Enums()) == 2)
	assert.DeepEqual(t, apiDataStatusEnum, apiPackage.Enum("DataStatus"), options...)
	assert.DeepEqual(t, apiDataFooType, apiPackage.Enum("DataFooType"), options...)

	assert.Assert(t, len(apiPackage.Messages()) == 6)

	// Verify the overall create data request message
	assert.DeepEqual(t, apiNestedMessageOne, apiPackage.Message("NestedMessageOne"), options...)
	assert.DeepEqual(t, apiNestedMessageOne.fields, apiPackage.Message("NestedMessageOne").Fields(), options...)
	assert.DeepEqual(t, apiNestedMessageTwo, apiPackage.Message("NestedMessageTwo"), options...)
	assert.DeepEqual(t, apiNestedMessageTwo.fields, apiPackage.Message("NestedMessageTwo").Fields(), options...)

	createDataRequestMessage := apiPackage.Message("CreateDataRequest")
	assert.DeepEqual(t, apiCreateDataRequest, createDataRequestMessage, options...)
	assert.DeepEqual(t, apiCreateDataRequest.fields, createDataRequestMessage.Fields(), options...)

	createDataRequestNameField := createDataRequestMessage.Field("name")
	assert.Equal(t, true, createDataRequestNameField.Has("validation.format"))
	assert.DeepEqual(t, map[string]interface{}{
		"type": "ALPHA",
	}, createDataRequestNameField.Get("validation.format"))

	repeatedStringPropertyField := createDataRequestMessage.Field("repeated_string_property")
	assert.Equal(t, true, repeatedStringPropertyField.Has("validation.format"))
	assert.DeepEqual(t, map[string]interface{}{
		"type":     "ALPHA_NUM_UNICODE",
		"required": false,
	}, repeatedStringPropertyField.Get("validation.format"))

	// Check services
	assert.Assert(t, len(apiPackage.Services()) == 1)
	jobsService := apiPackage.Service("Jobs")

	assert.Assert(t, len(jobsService.RPCs()) == 2)
	rpcOneRpc := jobsService.RPC("rpcOne")
	assert.DeepEqual(t, rpcOneRpc, apiRpcOne, options...)
}

func msgRef(c *Context, scope string, m *Message) *Type {
	return c.typeDictionary.newMsgTypeForTesting(scope, m)
}

func enumRef(c *Context, scope string, e *Enum) *Type {
	return c.typeDictionary.newEnumTypeForTesting(scope, e)
}
