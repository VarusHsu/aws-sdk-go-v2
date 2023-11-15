// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates a time series (data stream) from an asset property.
func (c *Client) DisassociateTimeSeriesFromAssetProperty(ctx context.Context, params *DisassociateTimeSeriesFromAssetPropertyInput, optFns ...func(*Options)) (*DisassociateTimeSeriesFromAssetPropertyOutput, error) {
	if params == nil {
		params = &DisassociateTimeSeriesFromAssetPropertyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateTimeSeriesFromAssetProperty", params, optFns, c.addOperationDisassociateTimeSeriesFromAssetPropertyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateTimeSeriesFromAssetPropertyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateTimeSeriesFromAssetPropertyInput struct {

	// The alias that identifies the time series.
	//
	// This member is required.
	Alias *string

	// The ID of the asset in which the asset property was created.
	//
	// This member is required.
	AssetId *string

	// The ID of the asset property.
	//
	// This member is required.
	PropertyId *string

	// A unique case-sensitive identifier that you can provide to ensure the
	// idempotency of the request. Don't reuse this client token if a new idempotent
	// request is required.
	ClientToken *string

	noSmithyDocumentSerde
}

type DisassociateTimeSeriesFromAssetPropertyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateTimeSeriesFromAssetPropertyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateTimeSeriesFromAssetProperty{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateTimeSeriesFromAssetProperty{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DisassociateTimeSeriesFromAssetProperty"); err != nil {
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
	if err = addEndpointPrefix_opDisassociateTimeSeriesFromAssetPropertyMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opDisassociateTimeSeriesFromAssetPropertyMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDisassociateTimeSeriesFromAssetPropertyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateTimeSeriesFromAssetProperty(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDisassociateTimeSeriesFromAssetPropertyMiddleware struct {
}

func (*endpointPrefix_opDisassociateTimeSeriesFromAssetPropertyMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDisassociateTimeSeriesFromAssetPropertyMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opDisassociateTimeSeriesFromAssetPropertyMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opDisassociateTimeSeriesFromAssetPropertyMiddleware{}, "ResolveEndpointV2", middleware.After)
}

type idempotencyToken_initializeOpDisassociateTimeSeriesFromAssetProperty struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDisassociateTimeSeriesFromAssetProperty) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDisassociateTimeSeriesFromAssetProperty) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DisassociateTimeSeriesFromAssetPropertyInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DisassociateTimeSeriesFromAssetPropertyInput ")
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
func addIdempotencyToken_opDisassociateTimeSeriesFromAssetPropertyMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpDisassociateTimeSeriesFromAssetProperty{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDisassociateTimeSeriesFromAssetProperty(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DisassociateTimeSeriesFromAssetProperty",
	}
}
