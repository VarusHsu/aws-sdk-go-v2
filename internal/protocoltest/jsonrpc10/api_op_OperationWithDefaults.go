// Code generated by smithy-go-codegen DO NOT EDIT.

package jsonrpc10

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/jsonrpc10/document"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/jsonrpc10/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

func (c *Client) OperationWithDefaults(ctx context.Context, params *OperationWithDefaultsInput, optFns ...func(*Options)) (*OperationWithDefaultsOutput, error) {
	if params == nil {
		params = &OperationWithDefaultsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "OperationWithDefaults", params, optFns, c.addOperationOperationWithDefaultsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*OperationWithDefaultsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type OperationWithDefaultsInput struct {
	ClientOptionalDefaults *types.ClientOptionalDefaults

	Defaults *types.Defaults

	OtherTopLevelDefault int32

	TopLevelDefault *string

	noSmithyDocumentSerde
}

type OperationWithDefaultsOutput struct {
	DefaultBlob []byte

	DefaultBoolean *bool

	DefaultByte *int8

	DefaultDocumentBoolean document.Interface

	DefaultDocumentList document.Interface

	DefaultDocumentMap document.Interface

	DefaultDocumentString document.Interface

	DefaultDouble *float64

	DefaultEnum types.TestEnum

	DefaultFloat *float32

	DefaultIntEnum types.TestIntEnum

	DefaultInteger *int32

	DefaultList []string

	DefaultLong *int64

	DefaultMap map[string]string

	DefaultNullDocument document.Interface

	DefaultShort *int16

	DefaultString *string

	DefaultTimestamp *time.Time

	EmptyBlob []byte

	EmptyString *string

	FalseBoolean bool

	ZeroByte int8

	ZeroDouble float64

	ZeroFloat float32

	ZeroInteger int32

	ZeroLong int64

	ZeroShort int16

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationOperationWithDefaultsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpOperationWithDefaults{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpOperationWithDefaults{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "OperationWithDefaults"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opOperationWithDefaults(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opOperationWithDefaults(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "OperationWithDefaults",
	}
}
