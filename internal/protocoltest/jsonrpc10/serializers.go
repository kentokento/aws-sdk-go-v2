// Code generated by smithy-go-codegen DO NOT EDIT.

package jsonrpc10

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/jsonrpc10/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"math"
	"path"
	"strings"
)

type awsAwsjson10_serializeOpEmptyInputAndEmptyOutput struct {
}

func (*awsAwsjson10_serializeOpEmptyInputAndEmptyOutput) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpEmptyInputAndEmptyOutput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*EmptyInputAndEmptyOutputInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("JsonRpc10.EmptyInputAndEmptyOutput")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentEmptyInputAndEmptyOutputInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpEndpointOperation struct {
}

func (*awsAwsjson10_serializeOpEndpointOperation) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpEndpointOperation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*EndpointOperationInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("JsonRpc10.EndpointOperation")

	if request, err = request.SetStream(strings.NewReader(`{}`)); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpEndpointWithHostLabelOperation struct {
}

func (*awsAwsjson10_serializeOpEndpointWithHostLabelOperation) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpEndpointWithHostLabelOperation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*EndpointWithHostLabelOperationInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("JsonRpc10.EndpointWithHostLabelOperation")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentEndpointWithHostLabelOperationInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpGreetingWithErrors struct {
}

func (*awsAwsjson10_serializeOpGreetingWithErrors) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpGreetingWithErrors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GreetingWithErrorsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("JsonRpc10.GreetingWithErrors")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentGreetingWithErrorsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpHostWithPathOperation struct {
}

func (*awsAwsjson10_serializeOpHostWithPathOperation) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpHostWithPathOperation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*HostWithPathOperationInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("JsonRpc10.HostWithPathOperation")

	if request, err = request.SetStream(strings.NewReader(`{}`)); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpJsonUnions struct {
}

func (*awsAwsjson10_serializeOpJsonUnions) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpJsonUnions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*JsonUnionsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("JsonRpc10.JsonUnions")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentJsonUnionsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpNoInputAndNoOutput struct {
}

func (*awsAwsjson10_serializeOpNoInputAndNoOutput) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpNoInputAndNoOutput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*NoInputAndNoOutputInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("JsonRpc10.NoInputAndNoOutput")

	if request, err = request.SetStream(strings.NewReader(`{}`)); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpNoInputAndOutput struct {
}

func (*awsAwsjson10_serializeOpNoInputAndOutput) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpNoInputAndOutput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*NoInputAndOutputInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("JsonRpc10.NoInputAndOutput")

	if request, err = request.SetStream(strings.NewReader(`{}`)); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpSimpleScalarProperties struct {
}

func (*awsAwsjson10_serializeOpSimpleScalarProperties) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpSimpleScalarProperties) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*SimpleScalarPropertiesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("JsonRpc10.SimpleScalarProperties")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentSimpleScalarPropertiesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsAwsjson10_serializeDocumentMyUnion(v types.MyUnion, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	switch uv := v.(type) {
	case *types.MyUnionMemberBlobValue:
		av := object.Key("blobValue")
		av.Base64EncodeBytes(uv.Value)

	case *types.MyUnionMemberBooleanValue:
		av := object.Key("booleanValue")
		av.Boolean(uv.Value)

	case *types.MyUnionMemberEnumValue:
		av := object.Key("enumValue")
		av.String(string(uv.Value))

	case *types.MyUnionMemberListValue:
		av := object.Key("listValue")
		if err := awsAwsjson10_serializeDocumentStringList(uv.Value, av); err != nil {
			return err
		}

	case *types.MyUnionMemberMapValue:
		av := object.Key("mapValue")
		if err := awsAwsjson10_serializeDocumentStringMap(uv.Value, av); err != nil {
			return err
		}

	case *types.MyUnionMemberNumberValue:
		av := object.Key("numberValue")
		av.Integer(uv.Value)

	case *types.MyUnionMemberStringValue:
		av := object.Key("stringValue")
		av.String(uv.Value)

	case *types.MyUnionMemberStructureValue:
		av := object.Key("structureValue")
		if err := awsAwsjson10_serializeDocumentGreetingStruct(&uv.Value, av); err != nil {
			return err
		}

	case *types.MyUnionMemberTimestampValue:
		av := object.Key("timestampValue")
		av.Double(smithytime.FormatEpochSeconds(uv.Value))

	default:
		return fmt.Errorf("attempted to serialize unknown member type %T for union %T", uv, v)

	}
	return nil
}

func awsAwsjson10_serializeDocumentGreetingStruct(v *types.GreetingStruct, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Hi != nil {
		ok := object.Key("hi")
		ok.String(*v.Hi)
	}

	return nil
}

func awsAwsjson10_serializeDocumentStringList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson10_serializeDocumentStringMap(v map[string]string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		om.String(v[key])
	}
	return nil
}

func awsAwsjson10_serializeOpDocumentEmptyInputAndEmptyOutputInput(v *EmptyInputAndEmptyOutputInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	return nil
}

func awsAwsjson10_serializeOpDocumentEndpointWithHostLabelOperationInput(v *EndpointWithHostLabelOperationInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Label != nil {
		ok := object.Key("label")
		ok.String(*v.Label)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentGreetingWithErrorsInput(v *GreetingWithErrorsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Greeting != nil {
		ok := object.Key("greeting")
		ok.String(*v.Greeting)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentJsonUnionsInput(v *JsonUnionsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Contents != nil {
		ok := object.Key("contents")
		if err := awsAwsjson10_serializeDocumentMyUnion(v.Contents, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentSimpleScalarPropertiesInput(v *SimpleScalarPropertiesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DoubleValue != nil {
		ok := object.Key("doubleValue")
		switch {
		case math.IsNaN(*v.DoubleValue):
			ok.String("NaN")

		case math.IsInf(*v.DoubleValue, 1):
			ok.String("Infinity")

		case math.IsInf(*v.DoubleValue, -1):
			ok.String("-Infinity")

		default:
			ok.Double(*v.DoubleValue)

		}
	}

	if v.FloatValue != nil {
		ok := object.Key("floatValue")
		switch {
		case math.IsNaN(float64(*v.FloatValue)):
			ok.String("NaN")

		case math.IsInf(float64(*v.FloatValue), 1):
			ok.String("Infinity")

		case math.IsInf(float64(*v.FloatValue), -1):
			ok.String("-Infinity")

		default:
			ok.Float(*v.FloatValue)

		}
	}

	return nil
}