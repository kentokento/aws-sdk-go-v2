// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new component that can be used to build, validate, test, and assess
// your image. The component is based on a YAML document that you specify using
// exactly one of the following methods:
//
// * Inline, using the data property in the
// request body.
//
// * A URL that points to a YAML document file stored in Amazon S3,
// using the uri property in the request body.
func (c *Client) CreateComponent(ctx context.Context, params *CreateComponentInput, optFns ...func(*Options)) (*CreateComponentOutput, error) {
	if params == nil {
		params = &CreateComponentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateComponent", params, optFns, c.addOperationCreateComponentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateComponentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateComponentInput struct {

	// The idempotency token of the component.
	//
	// This member is required.
	ClientToken *string

	// The name of the component.
	//
	// This member is required.
	Name *string

	// The platform of the component.
	//
	// This member is required.
	Platform types.Platform

	// The semantic version of the component. This version follows the semantic version
	// syntax. The semantic version has four nodes: ../. You can assign values for the
	// first three, and can filter on all of them. Assignment: For the first three
	// nodes you can assign any positive integer value, including zero, with an upper
	// limit of 2^30-1, or 1073741823 for each node. Image Builder automatically
	// assigns the build number to the fourth node. Patterns: You can use any numeric
	// pattern that adheres to the assignment requirements for the nodes that you can
	// assign. For example, you might choose a software version pattern, such as 1.0.0,
	// or a date, such as 2021.01.01.
	//
	// This member is required.
	SemanticVersion *string

	// The change description of the component. Describes what change has been made in
	// this version, or what makes this version different from other versions of this
	// component.
	ChangeDescription *string

	// Component data contains inline YAML document content for the component.
	// Alternatively, you can specify the uri of a YAML document file stored in Amazon
	// S3. However, you cannot specify both properties.
	Data *string

	// The description of the component. Describes the contents of the component.
	Description *string

	// The ID of the KMS key that should be used to encrypt this component.
	KmsKeyId *string

	// The operating system (OS) version supported by the component. If the OS
	// information is available, a prefix match is performed against the base image OS
	// version during image recipe creation.
	SupportedOsVersions []string

	// The tags of the component.
	Tags map[string]string

	// The uri of a YAML component document file. This must be an S3 URL
	// (s3://bucket/key), and the requester must have permission to access the S3
	// bucket it points to. If you use Amazon S3, you can specify component content up
	// to your service quota. Alternatively, you can specify the YAML document inline,
	// using the component data property. You cannot specify both properties.
	Uri *string

	noSmithyDocumentSerde
}

type CreateComponentOutput struct {

	// The idempotency token used to make this request idempotent.
	ClientToken *string

	// The Amazon Resource Name (ARN) of the component that was created by this
	// request.
	ComponentBuildVersionArn *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateComponentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateComponent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateComponent{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateComponentMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateComponentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateComponent(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpCreateComponent struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateComponent) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateComponent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateComponentInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateComponentInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateComponentMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateComponent{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateComponent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "CreateComponent",
	}
}