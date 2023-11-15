// Code generated by smithy-go-codegen DO NOT EDIT.

package grafana

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Grafana API key for the workspace. This key can be used to
// authenticate requests sent to the workspace's HTTP API. See
// https://docs.aws.amazon.com/grafana/latest/userguide/Using-Grafana-APIs.html (https://docs.aws.amazon.com/grafana/latest/userguide/Using-Grafana-APIs.html)
// for available APIs and example requests.
func (c *Client) CreateWorkspaceApiKey(ctx context.Context, params *CreateWorkspaceApiKeyInput, optFns ...func(*Options)) (*CreateWorkspaceApiKeyOutput, error) {
	if params == nil {
		params = &CreateWorkspaceApiKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWorkspaceApiKey", params, optFns, c.addOperationCreateWorkspaceApiKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWorkspaceApiKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWorkspaceApiKeyInput struct {

	// Specifies the name of the key. Keynames must be unique to the workspace.
	//
	// This member is required.
	KeyName *string

	// Specifies the permission level of the key. Valid values: VIEWER | EDITOR | ADMIN
	//
	// This member is required.
	KeyRole *string

	// Specifies the time in seconds until the key expires. Keys can be valid for up
	// to 30 days.
	//
	// This member is required.
	SecondsToLive *int32

	// The ID of the workspace to create an API key.
	//
	// This member is required.
	WorkspaceId *string

	noSmithyDocumentSerde
}

type CreateWorkspaceApiKeyOutput struct {

	// The key token. Use this value as a bearer token to authenticate HTTP requests
	// to the workspace.
	//
	// This member is required.
	Key *string

	// The name of the key that was created.
	//
	// This member is required.
	KeyName *string

	// The ID of the workspace that the key is valid for.
	//
	// This member is required.
	WorkspaceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWorkspaceApiKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateWorkspaceApiKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateWorkspaceApiKey{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateWorkspaceApiKey"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpCreateWorkspaceApiKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWorkspaceApiKey(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opCreateWorkspaceApiKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateWorkspaceApiKey",
	}
}
