// Code generated by smithy-go-codegen DO NOT EDIT.

package grafana

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/grafana/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a token that can be used to authenticate and authorize Grafana HTTP API
// operations for the given [workspace service account]. The service account acts as a user for the API
// operations, and defines the permissions that are used by the API.
//
// When you create the service account token, you will receive a key that is used
// when calling Grafana APIs. Do not lose this key, as it will not be retrievable
// again.
//
// If you do lose the key, you can delete the token and recreate it to receive a
// new key. This will disable the initial key.
//
// Service accounts are only available for workspaces that are compatible with
// Grafana version 9 and above.
//
// [workspace service account]: https://docs.aws.amazon.com/grafana/latest/userguide/service-accounts.html
func (c *Client) CreateWorkspaceServiceAccountToken(ctx context.Context, params *CreateWorkspaceServiceAccountTokenInput, optFns ...func(*Options)) (*CreateWorkspaceServiceAccountTokenOutput, error) {
	if params == nil {
		params = &CreateWorkspaceServiceAccountTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWorkspaceServiceAccountToken", params, optFns, c.addOperationCreateWorkspaceServiceAccountTokenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWorkspaceServiceAccountTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWorkspaceServiceAccountTokenInput struct {

	// A name for the token to create.
	//
	// This member is required.
	Name *string

	// Sets how long the token will be valid, in seconds. You can set the time up to
	// 30 days in the future.
	//
	// This member is required.
	SecondsToLive *int32

	// The ID of the service account for which to create a token.
	//
	// This member is required.
	ServiceAccountId *string

	// The ID of the workspace the service account resides within.
	//
	// This member is required.
	WorkspaceId *string

	noSmithyDocumentSerde
}

type CreateWorkspaceServiceAccountTokenOutput struct {

	// The ID of the service account where the token was created.
	//
	// This member is required.
	ServiceAccountId *string

	// Information about the created token, including the key. Be sure to store the
	// key securely.
	//
	// This member is required.
	ServiceAccountToken *types.ServiceAccountTokenSummaryWithKey

	// The ID of the workspace where the token was created.
	//
	// This member is required.
	WorkspaceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWorkspaceServiceAccountTokenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateWorkspaceServiceAccountToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateWorkspaceServiceAccountToken{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateWorkspaceServiceAccountToken"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpCreateWorkspaceServiceAccountTokenValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWorkspaceServiceAccountToken(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateWorkspaceServiceAccountToken(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateWorkspaceServiceAccountToken",
	}
}
