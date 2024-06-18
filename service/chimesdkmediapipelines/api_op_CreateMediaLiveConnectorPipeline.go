// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmediapipelines

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a media live connector pipeline in an Amazon Chime SDK meeting.
func (c *Client) CreateMediaLiveConnectorPipeline(ctx context.Context, params *CreateMediaLiveConnectorPipelineInput, optFns ...func(*Options)) (*CreateMediaLiveConnectorPipelineOutput, error) {
	if params == nil {
		params = &CreateMediaLiveConnectorPipelineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMediaLiveConnectorPipeline", params, optFns, c.addOperationCreateMediaLiveConnectorPipelineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMediaLiveConnectorPipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMediaLiveConnectorPipelineInput struct {

	// The media live connector pipeline's data sinks.
	//
	// This member is required.
	Sinks []types.LiveConnectorSinkConfiguration

	// The media live connector pipeline's data sources.
	//
	// This member is required.
	Sources []types.LiveConnectorSourceConfiguration

	// The token assigned to the client making the request.
	ClientRequestToken *string

	// The tags associated with the media live connector pipeline.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateMediaLiveConnectorPipelineOutput struct {

	// The new media live connector pipeline.
	MediaLiveConnectorPipeline *types.MediaLiveConnectorPipeline

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMediaLiveConnectorPipelineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMediaLiveConnectorPipeline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMediaLiveConnectorPipeline{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateMediaLiveConnectorPipeline"); err != nil {
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
	if err = addIdempotencyToken_opCreateMediaLiveConnectorPipelineMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateMediaLiveConnectorPipelineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMediaLiveConnectorPipeline(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateMediaLiveConnectorPipeline struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateMediaLiveConnectorPipeline) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateMediaLiveConnectorPipeline) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateMediaLiveConnectorPipelineInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateMediaLiveConnectorPipelineInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateMediaLiveConnectorPipelineMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateMediaLiveConnectorPipeline{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateMediaLiveConnectorPipeline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateMediaLiveConnectorPipeline",
	}
}
