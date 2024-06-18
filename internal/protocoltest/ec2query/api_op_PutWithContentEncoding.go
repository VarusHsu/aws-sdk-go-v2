// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2query

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyrequestcompression "github.com/aws/smithy-go/private/requestcompression"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

func (c *Client) PutWithContentEncoding(ctx context.Context, params *PutWithContentEncodingInput, optFns ...func(*Options)) (*PutWithContentEncodingOutput, error) {
	if params == nil {
		params = &PutWithContentEncodingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutWithContentEncoding", params, optFns, c.addOperationPutWithContentEncodingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutWithContentEncodingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutWithContentEncodingInput struct {
	Data *string

	Encoding *string

	noSmithyDocumentSerde
}

type PutWithContentEncodingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutWithContentEncodingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpPutWithContentEncoding{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpPutWithContentEncoding{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutWithContentEncoding"); err != nil {
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
	if err = addIsRequestCompressionUserAgent(stack, options); err != nil {
		return err
	}
	if err = addOperationPutWithContentEncodingRequestCompressionMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutWithContentEncoding(options.Region), middleware.Before); err != nil {
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

func addOperationPutWithContentEncodingRequestCompressionMiddleware(stack *middleware.Stack, options Options) error {
	return smithyrequestcompression.AddRequestCompression(stack, options.DisableRequestCompression, options.RequestMinCompressSizeBytes,
		[]string{
			"gzip",
		})
}

func newServiceMetadataMiddleware_opPutWithContentEncoding(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutWithContentEncoding",
	}
}
