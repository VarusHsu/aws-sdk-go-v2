// Code generated by smithy-go-codegen DO NOT EDIT.

package pcaconnectorscep

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pcaconnectorscep/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// For general-purpose connectors. Creates a challenge password for the specified
// connector. The SCEP protocol uses a challenge password to authenticate a request
// before issuing a certificate from a certificate authority (CA). Your SCEP
// clients include the challenge password as part of their certificate request to
// Connector for SCEP. To retrieve the connector Amazon Resource Names (ARNs) for
// the connectors in your account, call [ListConnectors].
//
// To create additional challenge passwords for the connector, call CreateChallenge
// again. We recommend frequently rotating your challenge passwords.
//
// [ListConnectors]: https://docs.aws.amazon.com/C4SCEP_API/pca-connector-scep/latest/APIReference/API_ListConnectors.html
func (c *Client) CreateChallenge(ctx context.Context, params *CreateChallengeInput, optFns ...func(*Options)) (*CreateChallengeOutput, error) {
	if params == nil {
		params = &CreateChallengeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateChallenge", params, optFns, c.addOperationCreateChallengeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateChallengeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateChallengeInput struct {

	// The Amazon Resource Name (ARN) of the connector that you want to create a
	// challenge for.
	//
	// This member is required.
	ConnectorArn *string

	// Custom string that can be used to distinguish between calls to the [CreateChallenge] action.
	// Client tokens for CreateChallenge time out after five minutes. Therefore, if
	// you call CreateChallenge multiple times with the same client token within five
	// minutes, Connector for SCEP recognizes that you are requesting only one
	// challenge and will only respond with one. If you change the client token for
	// each call, Connector for SCEP recognizes that you are requesting multiple
	// challenge passwords.
	//
	// [CreateChallenge]: https://docs.aws.amazon.com/C4SCEP_API/pca-connector-scep/latest/APIReference/API_CreateChallenge.html
	ClientToken *string

	// The key-value pairs to associate with the resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateChallengeOutput struct {

	// Returns the challenge details for the specified connector.
	Challenge *types.Challenge

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateChallengeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateChallenge{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateChallenge{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateChallenge"); err != nil {
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
	if err = addIdempotencyToken_opCreateChallengeMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateChallengeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateChallenge(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateChallenge struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateChallenge) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateChallenge) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateChallengeInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateChallengeInput ")
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
func addIdempotencyToken_opCreateChallengeMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateChallenge{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateChallenge(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateChallenge",
	}
}
