// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroups

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the service configuration associated with the specified resource group.
// For details about the service configuration syntax, see Service configurations
// for resource groups
// (https://docs.aws.amazon.com/ARG/latest/APIReference/about-slg.html). Minimum
// permissions To run this command, you must have the following permissions:
//
// *
// resource-groups:GetGroupConfiguration
func (c *Client) GetGroupConfiguration(ctx context.Context, params *GetGroupConfigurationInput, optFns ...func(*Options)) (*GetGroupConfigurationOutput, error) {
	if params == nil {
		params = &GetGroupConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetGroupConfiguration", params, optFns, c.addOperationGetGroupConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetGroupConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetGroupConfigurationInput struct {

	// The name or the ARN of the resource group.
	Group *string

	noSmithyDocumentSerde
}

type GetGroupConfigurationOutput struct {

	// The service configuration associated with the specified group. For details about
	// the service configuration syntax, see Service configurations for resource groups
	// (https://docs.aws.amazon.com/ARG/latest/APIReference/about-slg.html).
	GroupConfiguration *types.GroupConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetGroupConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetGroupConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetGroupConfiguration{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetGroupConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetGroupConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resource-groups",
		OperationName: "GetGroupConfiguration",
	}
}