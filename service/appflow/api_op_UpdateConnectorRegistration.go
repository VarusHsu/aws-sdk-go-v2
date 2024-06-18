// Code generated by smithy-go-codegen DO NOT EDIT.

package appflow

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/appflow/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a custom connector that you've previously registered. This operation
// updates the connector with one of the following:
//
//   - The latest version of the AWS Lambda function that's assigned to the
//     connector
//
//   - A new AWS Lambda function that you specify
func (c *Client) UpdateConnectorRegistration(ctx context.Context, params *UpdateConnectorRegistrationInput, optFns ...func(*Options)) (*UpdateConnectorRegistrationOutput, error) {
	if params == nil {
		params = &UpdateConnectorRegistrationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateConnectorRegistration", params, optFns, c.addOperationUpdateConnectorRegistrationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateConnectorRegistrationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateConnectorRegistrationInput struct {

	// The name of the connector. The name is unique for each connector registration
	// in your AWS account.
	//
	// This member is required.
	ConnectorLabel *string

	// The clientToken parameter is an idempotency token. It ensures that your
	// UpdateConnectorRegistration request completes only once. You choose the value to
	// pass. For example, if you don't receive a response from your request, you can
	// safely retry the request with the same clientToken parameter value.
	//
	// If you omit a clientToken value, the Amazon Web Services SDK that you are using
	// inserts a value for you. This way, the SDK can safely retry requests multiple
	// times after a network error. You must provide your own value for other use
	// cases.
	//
	// If you specify input parameters that differ from your first request, an error
	// occurs. If you use a different value for clientToken , Amazon AppFlow considers
	// it a new call to UpdateConnectorRegistration . The token is active for 8 hours.
	ClientToken *string

	// Contains information about the configuration of the connector being registered.
	ConnectorProvisioningConfig *types.ConnectorProvisioningConfig

	// A description about the update that you're applying to the connector.
	Description *string

	noSmithyDocumentSerde
}

type UpdateConnectorRegistrationOutput struct {

	// The ARN of the connector being updated.
	ConnectorArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateConnectorRegistrationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateConnectorRegistration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateConnectorRegistration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateConnectorRegistration"); err != nil {
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
	if err = addIdempotencyToken_opUpdateConnectorRegistrationMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateConnectorRegistrationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateConnectorRegistration(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateConnectorRegistration struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateConnectorRegistration) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateConnectorRegistration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateConnectorRegistrationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateConnectorRegistrationInput ")
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
func addIdempotencyToken_opUpdateConnectorRegistrationMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateConnectorRegistration{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateConnectorRegistration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateConnectorRegistration",
	}
}
