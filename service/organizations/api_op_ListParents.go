// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the root or organizational units (OUs) that serve as the immediate parent
// of the specified child OU or account. This operation, along with ListChildren
// enables you to traverse the tree structure that makes up this root. Always check
// the NextToken response parameter for a null value when calling a List*
// operation. These operations can occasionally return an empty set of results even
// when there are more results available. The NextToken response parameter value is
// null only when there are no more results to display. This operation can be
// called only from the organization's management account or by a member account
// that is a delegated administrator for an AWS service. In the current release, a
// child can have only a single parent.
func (c *Client) ListParents(ctx context.Context, params *ListParentsInput, optFns ...func(*Options)) (*ListParentsOutput, error) {
	if params == nil {
		params = &ListParentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListParents", params, optFns, c.addOperationListParentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListParentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListParentsInput struct {

	// The unique identifier (ID) of the OU or account whose parent containers you want
	// to list. Don't specify a root. The regex pattern
	// (http://wikipedia.org/wiki/regex) for a child ID string requires one of the
	// following:
	//
	// * Account - A string that consists of exactly 12 digits.
	//
	// *
	// Organizational unit (OU) - A string that begins with "ou-" followed by from 4 to
	// 32 lowercase letters or digits (the ID of the root that contains the OU). This
	// string is followed by a second "-" dash and from 8 to 32 additional lowercase
	// letters or digits.
	//
	// This member is required.
	ChildId *string

	// The total number of results that you want included on each page of the response.
	// If you do not include this parameter, it defaults to a value that is specific to
	// the operation. If additional items exist beyond the maximum you specify, the
	// NextToken response element is present and has a value (is not null). Include
	// that value as the NextToken request parameter in the next call to the operation
	// to get the next part of the results. Note that Organizations might return fewer
	// results than the maximum even when there are more results available. You should
	// check NextToken after every operation to ensure that you receive all of the
	// results.
	MaxResults *int32

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value of the previous call's NextToken
	// response to indicate where the output should continue from.
	NextToken *string

	noSmithyDocumentSerde
}

type ListParentsOutput struct {

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null.
	NextToken *string

	// A list of parents for the specified child account or OU.
	Parents []types.Parent

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListParentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListParents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListParents{}, middleware.After)
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
	if err = addOpListParentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListParents(options.Region), middleware.Before); err != nil {
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

// ListParentsAPIClient is a client that implements the ListParents operation.
type ListParentsAPIClient interface {
	ListParents(context.Context, *ListParentsInput, ...func(*Options)) (*ListParentsOutput, error)
}

var _ ListParentsAPIClient = (*Client)(nil)

// ListParentsPaginatorOptions is the paginator options for ListParents
type ListParentsPaginatorOptions struct {
	// The total number of results that you want included on each page of the response.
	// If you do not include this parameter, it defaults to a value that is specific to
	// the operation. If additional items exist beyond the maximum you specify, the
	// NextToken response element is present and has a value (is not null). Include
	// that value as the NextToken request parameter in the next call to the operation
	// to get the next part of the results. Note that Organizations might return fewer
	// results than the maximum even when there are more results available. You should
	// check NextToken after every operation to ensure that you receive all of the
	// results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListParentsPaginator is a paginator for ListParents
type ListParentsPaginator struct {
	options   ListParentsPaginatorOptions
	client    ListParentsAPIClient
	params    *ListParentsInput
	nextToken *string
	firstPage bool
}

// NewListParentsPaginator returns a new ListParentsPaginator
func NewListParentsPaginator(client ListParentsAPIClient, params *ListParentsInput, optFns ...func(*ListParentsPaginatorOptions)) *ListParentsPaginator {
	if params == nil {
		params = &ListParentsInput{}
	}

	options := ListParentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListParentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListParentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListParents page.
func (p *ListParentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListParentsOutput, error) {
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

	result, err := p.client.ListParents(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListParents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "ListParents",
	}
}