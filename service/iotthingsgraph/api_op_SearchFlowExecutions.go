// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Searches for AWS IoT Things Graph workflow execution instances.
func (c *Client) SearchFlowExecutions(ctx context.Context, params *SearchFlowExecutionsInput, optFns ...func(*Options)) (*SearchFlowExecutionsOutput, error) {
	if params == nil {
		params = &SearchFlowExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchFlowExecutions", params, optFns, c.addOperationSearchFlowExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchFlowExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchFlowExecutionsInput struct {

	// The ID of the system instance that contains the flow.
	//
	// This member is required.
	SystemInstanceId *string

	// The date and time of the latest flow execution to return.
	EndTime *time.Time

	// The ID of a flow execution.
	FlowExecutionId *string

	// The maximum number of results to return in the response.
	MaxResults *int32

	// The string that specifies the next page of results. Use this when you're
	// paginating results.
	NextToken *string

	// The date and time of the earliest flow execution to return.
	StartTime *time.Time

	noSmithyDocumentSerde
}

type SearchFlowExecutionsOutput struct {

	// The string to specify as nextToken when you request the next page of results.
	NextToken *string

	// An array of objects that contain summary information about each workflow
	// execution in the result set.
	Summaries []types.FlowExecutionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchFlowExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchFlowExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchFlowExecutions{}, middleware.After)
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
	if err = addOpSearchFlowExecutionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchFlowExecutions(options.Region), middleware.Before); err != nil {
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

// SearchFlowExecutionsAPIClient is a client that implements the
// SearchFlowExecutions operation.
type SearchFlowExecutionsAPIClient interface {
	SearchFlowExecutions(context.Context, *SearchFlowExecutionsInput, ...func(*Options)) (*SearchFlowExecutionsOutput, error)
}

var _ SearchFlowExecutionsAPIClient = (*Client)(nil)

// SearchFlowExecutionsPaginatorOptions is the paginator options for
// SearchFlowExecutions
type SearchFlowExecutionsPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchFlowExecutionsPaginator is a paginator for SearchFlowExecutions
type SearchFlowExecutionsPaginator struct {
	options   SearchFlowExecutionsPaginatorOptions
	client    SearchFlowExecutionsAPIClient
	params    *SearchFlowExecutionsInput
	nextToken *string
	firstPage bool
}

// NewSearchFlowExecutionsPaginator returns a new SearchFlowExecutionsPaginator
func NewSearchFlowExecutionsPaginator(client SearchFlowExecutionsAPIClient, params *SearchFlowExecutionsInput, optFns ...func(*SearchFlowExecutionsPaginatorOptions)) *SearchFlowExecutionsPaginator {
	if params == nil {
		params = &SearchFlowExecutionsInput{}
	}

	options := SearchFlowExecutionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchFlowExecutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchFlowExecutionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchFlowExecutions page.
func (p *SearchFlowExecutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchFlowExecutionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.SearchFlowExecutions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opSearchFlowExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "SearchFlowExecutions",
	}
}