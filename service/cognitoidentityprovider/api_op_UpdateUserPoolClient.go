// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified user pool app client with the specified attributes. You
// can get a list of the current user pool app client settings using
// DescribeUserPoolClient
// (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_DescribeUserPoolClient.html).
// If you don't provide a value for an attribute, it will be set to the default
// value. You can also use this operation to enable token revocation for user pool
// clients. For more information about revoking tokens, see RevokeToken
// (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_RevokeToken.html).
func (c *Client) UpdateUserPoolClient(ctx context.Context, params *UpdateUserPoolClientInput, optFns ...func(*Options)) (*UpdateUserPoolClientOutput, error) {
	if params == nil {
		params = &UpdateUserPoolClientInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateUserPoolClient", params, optFns, c.addOperationUpdateUserPoolClientMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateUserPoolClientOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to update the user pool client.
type UpdateUserPoolClientInput struct {

	// The ID of the client associated with the user pool.
	//
	// This member is required.
	ClientId *string

	// The user pool ID for the user pool where you want to update the user pool
	// client.
	//
	// This member is required.
	UserPoolId *string

	// The time limit after which the access token is no longer valid and can't be
	// used.
	AccessTokenValidity *int32

	// The allowed OAuth flows. Set to code to initiate a code grant flow, which
	// provides an authorization code as the response. This code can be exchanged for
	// access tokens with the token endpoint. Set to implicit to specify that the
	// client should get the access token (and, optionally, ID token, based on scopes)
	// directly. Set to client_credentials to specify that the client should get the
	// access token (and, optionally, ID token, based on scopes) from the token
	// endpoint using a combination of client and client_secret.
	AllowedOAuthFlows []types.OAuthFlowType

	// Set to true if the client is allowed to follow the OAuth protocol when
	// interacting with Amazon Cognito user pools.
	AllowedOAuthFlowsUserPoolClient bool

	// The allowed OAuth scopes. Possible values provided by OAuth are: phone, email,
	// openid, and profile. Possible values provided by Amazon Web Services are:
	// aws.cognito.signin.user.admin. Custom scopes created in Resource Servers are
	// also supported.
	AllowedOAuthScopes []string

	// The Amazon Pinpoint analytics configuration for collecting metrics for this user
	// pool. In Amazon Web Services Regions where isn't available, User Pools only
	// supports sending events to Amazon Pinpoint projects in us-east-1. In Regions
	// where Pinpoint is available, User Pools will support sending events to Amazon
	// Pinpoint projects within that same Region.
	AnalyticsConfiguration *types.AnalyticsConfigurationType

	// A list of allowed redirect (callback) URLs for the identity providers. A
	// redirect URI must:
	//
	// * Be an absolute URI.
	//
	// * Be registered with the
	// authorization server.
	//
	// * Not include a fragment component.
	//
	// See OAuth 2.0 -
	// Redirection Endpoint (https://tools.ietf.org/html/rfc6749#section-3.1.2). Amazon
	// Cognito requires HTTPS over HTTP except for http://localhost for testing
	// purposes only. App callback URLs such as myapp://example are also supported.
	CallbackURLs []string

	// The client name from the update user pool client request.
	ClientName *string

	// The default redirect URI. Must be in the CallbackURLs list. A redirect URI
	// must:
	//
	// * Be an absolute URI.
	//
	// * Be registered with the authorization server.
	//
	// *
	// Not include a fragment component.
	//
	// See OAuth 2.0 - Redirection Endpoint
	// (https://tools.ietf.org/html/rfc6749#section-3.1.2). Amazon Cognito requires
	// HTTPS over HTTP except for http://localhost for testing purposes only. App
	// callback URLs such as myapp://example are also supported.
	DefaultRedirectURI *string

	// Activates or deactivates token revocation. For more information about revoking
	// tokens, see RevokeToken
	// (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_RevokeToken.html).
	EnableTokenRevocation *bool

	// The authentication flows that are supported by the user pool clients. Flow names
	// without the ALLOW_ prefix are no longer supported in favor of new names with the
	// ALLOW_ prefix. Note that values with ALLOW_ prefix must be used only along with
	// values with the ALLOW_ prefix. Valid values include:
	//
	// *
	// ALLOW_ADMIN_USER_PASSWORD_AUTH: Enable admin based user password authentication
	// flow ADMIN_USER_PASSWORD_AUTH. This setting replaces the ADMIN_NO_SRP_AUTH
	// setting. With this authentication flow, Amazon Cognito receives the password in
	// the request instead of using the Secure Remote Password (SRP) protocol to verify
	// passwords.
	//
	// * ALLOW_CUSTOM_AUTH: Enable Lambda trigger based authentication.
	//
	// *
	// ALLOW_USER_PASSWORD_AUTH: Enable user password-based authentication. In this
	// flow, Amazon Cognito receives the password in the request instead of using the
	// SRP protocol to verify passwords.
	//
	// * ALLOW_USER_SRP_AUTH: Enable SRP-based
	// authentication.
	//
	// * ALLOW_REFRESH_TOKEN_AUTH: Enable authflow to refresh tokens.
	ExplicitAuthFlows []types.ExplicitAuthFlowsType

	// The time limit after which the ID token is no longer valid and can't be used.
	IdTokenValidity *int32

	// A list of allowed logout URLs for the identity providers.
	LogoutURLs []string

	// Errors and responses that you want Amazon Cognito APIs to return during
	// authentication, account confirmation, and password recovery when the user
	// doesn't exist in the user pool. When set to ENABLED and the user doesn't exist,
	// authentication returns an error indicating either the username or password was
	// incorrect. Account confirmation and password recovery return a response
	// indicating a code was sent to a simulated destination. When set to LEGACY, those
	// APIs return a UserNotFoundException exception if the user doesn't exist in the
	// user pool. Valid values include:
	//
	// * ENABLED - This prevents user
	// existence-related errors.
	//
	// * LEGACY - This represents the early behavior of
	// Amazon Cognito where user existence related errors aren't prevented.
	PreventUserExistenceErrors types.PreventUserExistenceErrorTypes

	// The read-only attributes of the user pool.
	ReadAttributes []string

	// The time limit, in days, after which the refresh token is no longer valid and
	// can't be used.
	RefreshTokenValidity int32

	// A list of provider names for the identity providers that are supported on this
	// client.
	SupportedIdentityProviders []string

	// The units in which the validity times are represented. Default for RefreshToken
	// is days, and default for ID and access tokens is hours.
	TokenValidityUnits *types.TokenValidityUnitsType

	// The writeable attributes of the user pool.
	WriteAttributes []string

	noSmithyDocumentSerde
}

// Represents the response from the server to the request to update the user pool
// client.
type UpdateUserPoolClientOutput struct {

	// The user pool client value from the response from the server when you request to
	// update the user pool client.
	UserPoolClient *types.UserPoolClientType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateUserPoolClientMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateUserPoolClient{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateUserPoolClient{}, middleware.After)
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
	if err = addOpUpdateUserPoolClientValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUserPoolClient(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateUserPoolClient(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "UpdateUserPoolClient",
	}
}