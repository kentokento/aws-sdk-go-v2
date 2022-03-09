// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Authorizes the Shield Response Team (SRT) to access the specified Amazon S3
// bucket containing log data such as Application Load Balancer access logs,
// CloudFront logs, or logs from third party sources. You can associate up to 10
// Amazon S3 buckets with your subscription. To use the services of the SRT and
// make an AssociateDRTLogBucket request, you must be subscribed to the Business
// Support plan (https://docs.aws.amazon.com/premiumsupport/business-support/) or
// the Enterprise Support plan
// (https://docs.aws.amazon.com/premiumsupport/enterprise-support/).
func (c *Client) AssociateDRTLogBucket(ctx context.Context, params *AssociateDRTLogBucketInput, optFns ...func(*Options)) (*AssociateDRTLogBucketOutput, error) {
	if params == nil {
		params = &AssociateDRTLogBucketInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateDRTLogBucket", params, optFns, c.addOperationAssociateDRTLogBucketMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateDRTLogBucketOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateDRTLogBucketInput struct {

	// The Amazon S3 bucket that contains the logs that you want to share.
	//
	// This member is required.
	LogBucket *string

	noSmithyDocumentSerde
}

type AssociateDRTLogBucketOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateDRTLogBucketMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateDRTLogBucket{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateDRTLogBucket{}, middleware.After)
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
	if err = addOpAssociateDRTLogBucketValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateDRTLogBucket(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateDRTLogBucket(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "AssociateDRTLogBucket",
	}
}