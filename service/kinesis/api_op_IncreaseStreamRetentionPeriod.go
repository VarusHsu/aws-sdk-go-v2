// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Increases the Kinesis data stream's retention period, which is the length of
// time data records are accessible after they are added to the stream. The maximum
// value of a stream's retention period is 8760 hours (365 days). When invoking
// this API, it is recommended you use the StreamARN input parameter rather than
// the StreamName input parameter. If you choose a longer stream retention period,
// this operation increases the time period during which records that have not yet
// expired are accessible. However, it does not make previous, expired data (older
// than the stream's previous retention period) accessible after the operation has
// been called. For example, if a stream's retention period is set to 24 hours and
// is increased to 168 hours, any data that is older than 24 hours remains
// inaccessible to consumer applications.
func (c *Client) IncreaseStreamRetentionPeriod(ctx context.Context, params *IncreaseStreamRetentionPeriodInput, optFns ...func(*Options)) (*IncreaseStreamRetentionPeriodOutput, error) {
	if params == nil {
		params = &IncreaseStreamRetentionPeriodInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "IncreaseStreamRetentionPeriod", params, optFns, c.addOperationIncreaseStreamRetentionPeriodMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*IncreaseStreamRetentionPeriodOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for IncreaseStreamRetentionPeriod .
type IncreaseStreamRetentionPeriodInput struct {

	// The new retention period of the stream, in hours. Must be more than the current
	// retention period.
	//
	// This member is required.
	RetentionPeriodHours *int32

	// The ARN of the stream.
	StreamARN *string

	// The name of the stream to modify.
	StreamName *string

	noSmithyDocumentSerde
}

func (in *IncreaseStreamRetentionPeriodInput) bindEndpointParams(p *EndpointParameters) {
	p.StreamARN = in.StreamARN
	p.OperationType = ptr.String("control")
}

type IncreaseStreamRetentionPeriodOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationIncreaseStreamRetentionPeriodMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpIncreaseStreamRetentionPeriod{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpIncreaseStreamRetentionPeriod{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "IncreaseStreamRetentionPeriod"); err != nil {
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
	if err = addOpIncreaseStreamRetentionPeriodValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opIncreaseStreamRetentionPeriod(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opIncreaseStreamRetentionPeriod(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "IncreaseStreamRetentionPeriod",
	}
}
