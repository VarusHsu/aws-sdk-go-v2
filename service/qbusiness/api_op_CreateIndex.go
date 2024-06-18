// Code generated by smithy-go-codegen DO NOT EDIT.

package qbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Q Business index.
//
// To determine if index creation has completed, check the Status field returned
// from a call to DescribeIndex . The Status field is set to ACTIVE when the index
// is ready to use.
//
// Once the index is active, you can index your documents using the [BatchPutDocument]
// BatchPutDocument API or the [CreateDataSource]CreateDataSource API.
//
// [BatchPutDocument]: https://docs.aws.amazon.com/amazonq/latest/api-reference/API_BatchPutDocument.html
// [CreateDataSource]: https://docs.aws.amazon.com/amazonq/latest/api-reference/API_CreateDataSource.html
func (c *Client) CreateIndex(ctx context.Context, params *CreateIndexInput, optFns ...func(*Options)) (*CreateIndexOutput, error) {
	if params == nil {
		params = &CreateIndexInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateIndex", params, optFns, c.addOperationCreateIndexMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateIndexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIndexInput struct {

	// The identifier of the Amazon Q Business application using the index.
	//
	// This member is required.
	ApplicationId *string

	// A name for the Amazon Q Business index.
	//
	// This member is required.
	DisplayName *string

	// The capacity units you want to provision for your index. You can add and remove
	// capacity to fit your usage needs.
	CapacityConfiguration *types.IndexCapacityConfiguration

	// A token that you provide to identify the request to create an index. Multiple
	// calls to the CreateIndex API with the same client token will create only one
	// index.
	ClientToken *string

	// A description for the Amazon Q Business index.
	Description *string

	// A list of key-value pairs that identify or categorize the index. You can also
	// use tags to help control access to the index. Tag keys and values can consist of
	// Unicode letters, digits, white space, and any of the following symbols: _ . : /
	// = + - @.
	Tags []types.Tag

	// The index type that's suitable for your needs. For more information on what's
	// included in each type of index or index tier, see [Amazon Q Business tiers].
	//
	// [Amazon Q Business tiers]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/what-is.html#tiers
	Type types.IndexType

	noSmithyDocumentSerde
}

type CreateIndexOutput struct {

	//  The Amazon Resource Name (ARN) of an Amazon Q Business index.
	IndexArn *string

	// The identifier for the Amazon Q Business index.
	IndexId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateIndexMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateIndex{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateIndex{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateIndex"); err != nil {
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
	if err = addIdempotencyToken_opCreateIndexMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateIndexValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIndex(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateIndex struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateIndex) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateIndex) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateIndexInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateIndexInput ")
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
func addIdempotencyToken_opCreateIndexMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateIndex{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateIndex(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateIndex",
	}
}
