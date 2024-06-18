// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// The example tests how requests serialize different timestamp formats in the URI
// path.
func (c *Client) HttpRequestWithLabelsAndTimestampFormat(ctx context.Context, params *HttpRequestWithLabelsAndTimestampFormatInput, optFns ...func(*Options)) (*HttpRequestWithLabelsAndTimestampFormatOutput, error) {
	if params == nil {
		params = &HttpRequestWithLabelsAndTimestampFormatInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "HttpRequestWithLabelsAndTimestampFormat", params, optFns, c.addOperationHttpRequestWithLabelsAndTimestampFormatMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*HttpRequestWithLabelsAndTimestampFormatOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpRequestWithLabelsAndTimestampFormatInput struct {

	// This member is required.
	DefaultFormat *time.Time

	// This member is required.
	MemberDateTime *time.Time

	// This member is required.
	MemberEpochSeconds *time.Time

	// This member is required.
	MemberHttpDate *time.Time

	// This member is required.
	TargetDateTime *time.Time

	// This member is required.
	TargetEpochSeconds *time.Time

	// This member is required.
	TargetHttpDate *time.Time

	noSmithyDocumentSerde
}

type HttpRequestWithLabelsAndTimestampFormatOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationHttpRequestWithLabelsAndTimestampFormatMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpHttpRequestWithLabelsAndTimestampFormat{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpHttpRequestWithLabelsAndTimestampFormat{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "HttpRequestWithLabelsAndTimestampFormat"); err != nil {
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
	if err = addOpHttpRequestWithLabelsAndTimestampFormatValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opHttpRequestWithLabelsAndTimestampFormat(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opHttpRequestWithLabelsAndTimestampFormat(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "HttpRequestWithLabelsAndTimestampFormat",
	}
}
