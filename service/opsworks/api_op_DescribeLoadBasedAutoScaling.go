// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes load-based auto scaling configurations for specified layers. You must
// specify at least one of the parameters. Required Permissions: To use this
// action, an IAM user must have a Show, Deploy, or Manage permissions level for
// the stack, or an attached policy that explicitly grants permissions. For more
// information about user permissions, see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) DescribeLoadBasedAutoScaling(ctx context.Context, params *DescribeLoadBasedAutoScalingInput, optFns ...func(*Options)) (*DescribeLoadBasedAutoScalingOutput, error) {
	if params == nil {
		params = &DescribeLoadBasedAutoScalingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLoadBasedAutoScaling", params, optFns, c.addOperationDescribeLoadBasedAutoScalingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLoadBasedAutoScalingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLoadBasedAutoScalingInput struct {

	// An array of layer IDs.
	//
	// This member is required.
	LayerIds []string

	noSmithyDocumentSerde
}

// Contains the response to a DescribeLoadBasedAutoScaling request.
type DescribeLoadBasedAutoScalingOutput struct {

	// An array of LoadBasedAutoScalingConfiguration objects that describe each layer's
	// configuration.
	LoadBasedAutoScalingConfigurations []types.LoadBasedAutoScalingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLoadBasedAutoScalingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeLoadBasedAutoScaling{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeLoadBasedAutoScaling{}, middleware.After)
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
	if err = addOpDescribeLoadBasedAutoScalingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLoadBasedAutoScaling(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeLoadBasedAutoScaling(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "DescribeLoadBasedAutoScaling",
	}
}