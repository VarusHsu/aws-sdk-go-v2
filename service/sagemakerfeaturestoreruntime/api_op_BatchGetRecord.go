// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakerfeaturestoreruntime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemakerfeaturestoreruntime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a batch of Records from a FeatureGroup .
func (c *Client) BatchGetRecord(ctx context.Context, params *BatchGetRecordInput, optFns ...func(*Options)) (*BatchGetRecordOutput, error) {
	if params == nil {
		params = &BatchGetRecordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetRecord", params, optFns, c.addOperationBatchGetRecordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetRecordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetRecordInput struct {

	// A list containing the name or Amazon Resource Name (ARN) of the FeatureGroup ,
	// the list of names of Feature s to be retrieved, and the corresponding
	// RecordIdentifier values as strings.
	//
	// This member is required.
	Identifiers []types.BatchGetRecordIdentifier

	// Parameter to request ExpiresAt in response. If Enabled , BatchGetRecord will
	// return the value of ExpiresAt , if it is not null. If Disabled and null,
	// BatchGetRecord will return null.
	ExpirationTimeResponse types.ExpirationTimeResponse

	noSmithyDocumentSerde
}

type BatchGetRecordOutput struct {

	// A list of errors that have occurred when retrieving a batch of Records.
	//
	// This member is required.
	Errors []types.BatchGetRecordError

	// A list of Records you requested to be retrieved in batch.
	//
	// This member is required.
	Records []types.BatchGetRecordResultDetail

	// A unprocessed list of FeatureGroup names, with their corresponding
	// RecordIdentifier value, and Feature name.
	//
	// This member is required.
	UnprocessedIdentifiers []types.BatchGetRecordIdentifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetRecordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchGetRecord{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchGetRecord{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchGetRecord"); err != nil {
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
	if err = addOpBatchGetRecordValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetRecord(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetRecord(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchGetRecord",
	}
}
