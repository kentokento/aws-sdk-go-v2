// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists iSCSI initiators that are connected to a volume. You can use this
// operation to determine whether a volume is being used or not. This operation is
// only supported in the cached volume and stored volume gateway types.
func (c *Client) ListVolumeInitiators(ctx context.Context, params *ListVolumeInitiatorsInput, optFns ...func(*Options)) (*ListVolumeInitiatorsOutput, error) {
	if params == nil {
		params = &ListVolumeInitiatorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVolumeInitiators", params, optFns, c.addOperationListVolumeInitiatorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVolumeInitiatorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// ListVolumeInitiatorsInput
type ListVolumeInitiatorsInput struct {

	// The Amazon Resource Name (ARN) of the volume. Use the ListVolumes operation to
	// return a list of gateway volumes for the gateway.
	//
	// This member is required.
	VolumeARN *string

	noSmithyDocumentSerde
}

// ListVolumeInitiatorsOutput
type ListVolumeInitiatorsOutput struct {

	// The host names and port numbers of all iSCSI initiators that are connected to
	// the gateway.
	Initiators []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVolumeInitiatorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListVolumeInitiators{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListVolumeInitiators{}, middleware.After)
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
	if err = addOpListVolumeInitiatorsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVolumeInitiators(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListVolumeInitiators(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "ListVolumeInitiators",
	}
}