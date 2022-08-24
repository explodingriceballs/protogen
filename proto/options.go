package proto

import (
	"github.com/rs/zerolog/log"
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"strconv"
	"strings"
)

type Options struct {
	noopVisitor
	types  *TypeDictionary
	parent Element

	Options []*Option
}

func (o *Options) getOption(name string) *Option {
	var option *Option = nil
	for _, opt := range o.Options {
		if strings.Contains(opt.OptionName, name) {
			option = opt
			break
		}
	}
	return option
}

func (o *Options) getOptionType(name string) *Type {
	var (
		field *Field = nil
	)

	extType := getExtension(o.parent.GetElementType())
	if extType == UnknownExtension {
		log.Warn().Str("option", name).Msg("Unknown extension type")
		return nil
	}
	if _, ok := o.types.extensions[extType]; !ok {
		log.Warn().Str("option", name).Int("extType", int(extType)).Msg("No extensions registered")
		return nil
	}

	for extName, extensions := range o.types.extensions[extType] {
		if !strings.Contains(name, extName) {
			continue
		}
		for _, extension := range extensions {
			for _, extField := range extension.fields {
				if strings.Contains(name, extField.fieldName) {
					field = extField
					break
				}
			}
		}
	}

	if field == nil {
		log.Warn().Str("name", name).Msg("no extension found")
		return nil
	}

	return field.t
}

func (o *Options) Has(name string) bool {
	if o.getOption(name) == nil {
		return false
	}
	return true
}

func (o *Options) Get(name string) interface{} {
	var (
		optionType = o.getOptionType(name)
		option     = o.getOption(name)
	)
	if optionType == nil || option == nil {
		return nil
	}
	if optionType.isNative {
		return toNativeValue(option.OptionValue, optionType)
	}
	if option.OptionValue[0] != '{' {
		// Only a single field in the referenced message is being set
		strippedName := strings.Replace(option.OptionName, name, "", 1)
		// Strip the empty brackets, if there are any
		strippedName = strings.Replace(strippedName, "()", "", 1)
		// Split by dots
		parts := strings.Split(strippedName, ".")

		returnValue := map[string]interface{}{}

		// Loop through all parts
		for _, part := range parts {
			// If the part is empty - skip
			if part == "" {
				continue
			}

			// Loop through all the fields in the current message
			if field := optionType.message.Field(part); field != nil {
				if field.t.isNative {
					returnValue[part] = toNativeValue(option.OptionValue, field.t)
				}
				if field.t.isEnumType {
					returnValue[part] = toEnumValue(option.OptionValue, field.t)
				}
			}
		}
		return returnValue
	}

	return ParseOption(option.OptionValue, optionType)
}

func (o *Options) VisitOption(option *parser.Option) (next bool) {
	o.AddOption(option.OptionName, option.Constant)

	// Ignore comments
	return false
}

func (o *Options) AddOption(name string, value string) {
	o.Options = append(o.Options, &Option{
		OptionName:  name,
		OptionValue: value,
	})
}

func NewOptions(parent Element, types *TypeDictionary) *Options {
	return &Options{
		parent: parent,
		types:  types,
	}
}

type Option struct {
	noopVisitor
	OptionName  string
	OptionValue string
}

func getExtension(elementType ElementType) ExtensionType {
	switch elementType {
	case MessageElementType:
		return MessageExtension
	case FieldElementType:
		return FieldExtension
	case ServiceElementType:
		return ServiceExtension
	case RPCElementType:
		return MethodExtension
	}
	return UnknownExtension
}

func toValue(name string, option *Option, optionType *Type) interface{} {
	optionValue := option.OptionValue
	if optionType.isNative {
		return toNativeValue(optionValue, optionType)
	}
	if optionType.isMessageType {
		returnValue := map[string]interface{}{}

		// If the value begins with {, then we can assume multiple fields are being set
		if optionValue[0] == '{' {
			var (
				inIdentifier      = true
				currentIdentifier = ""
				inValue           = false
				currentValue      = ""
				readerIndex       = 1 // Skip the initial {
				currentType       = optionType
			)

			for {
				if readerIndex == len(optionValue) {
					break
				}
				currentChar := optionValue[readerIndex]

				// We are reading an identifier
				if inIdentifier {
					// Check for the first colon
					colonIndex := strings.Index(optionValue[readerIndex:], ":")
					currentIdentifier = optionValue[readerIndex : readerIndex+colonIndex]
					// Update the index & set the flags
					readerIndex = readerIndex + colonIndex + 1
					inIdentifier = false
					inValue = true
				}

				// We are reading a value
				if inValue {
					if currentChar == '"' {
						for charIdx := readerIndex + 1; charIdx <= len(currentValue); charIdx++ {
							if optionValue[charIdx] == '"' && optionValue[charIdx-1] != '\\' {
								currentValue = optionValue[readerIndex:charIdx]

								// Update the index & set the flags
								readerIndex = charIdx
								inIdentifier = true
								inValue = false
								break
							}
						}

					} else {
						semiColonIndex := strings.Index(optionValue[readerIndex:], ";")
						currentValue = optionValue[readerIndex : readerIndex+semiColonIndex]

						// Update the index & set the flags
						readerIndex = readerIndex + semiColonIndex + 1
						inIdentifier = true
						inValue = false
					}
					if currentType.isNative {
						returnValue[currentIdentifier] = toNativeValue(currentValue, currentType)
					}
				}
			}
			return returnValue
		} else {
			// Only a single field in the referenced message is being set
			strippedName := strings.Replace(option.OptionName, name, "", 1)
			// Strip the empty brackets, if there are any
			strippedName = strings.Replace(strippedName, "()", "", 1)
			// Split by dots
			parts := strings.Split(strippedName, ".")

			returnValue := map[string]interface{}{}

			// Loop through all parts
			for _, part := range parts {
				// If the part is empty - skip
				if part == "" {
					continue
				}

				// Loop through all the fields in the current message
				if field := optionType.message.Field(part); field != nil {
					if field.t.isNative {
						returnValue[part] = toNativeValue(optionValue, field.t)
					}
					if field.t.isEnumType {
						returnValue[part] = toEnumValue(optionValue, field.t)
					}
				}
			}
			return returnValue
		}
	}
	return nil
}

func toEnumValue(value string, t *Type) interface{} {
	for _, enum := range t.enum.enumValues {
		if enum.identifier == value {
			return value
		}
	}
	return nil
}

func toNativeValue(value string, optionType *Type) interface{} {
	switch optionType.nativeType {
	case "double":
		uintValue, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			log.Warn().Str("option", value).Err(err).Msg("failed to parse")
			return nil
		}
		return uintValue
	case "float":
	case "int32":
	case "int64":
	case "uint32":
	case "uint64":
	case "sint32":
	case "sint64":
	case "fixed32":
	case "fixed64":
	case "sfixed32":
	case "sfixed64":
	case "bool":
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			log.Warn().Str("option", value).Err(err).Msg("failed to parse")
			return nil
		}
		return boolValue
	case "string":
	case "bytes":
	}
	return nil
}
