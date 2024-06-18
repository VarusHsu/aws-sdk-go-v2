// Code generated by smithy-go-codegen DO NOT EDIT.

package pipes

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pipes/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Get the information about an existing pipe. For more information about pipes,
// see [Amazon EventBridge Pipes]in the Amazon EventBridge User Guide.
//
// [Amazon EventBridge Pipes]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes.html
func (c *Client) DescribePipe(ctx context.Context, params *DescribePipeInput, optFns ...func(*Options)) (*DescribePipeOutput, error) {
	if params == nil {
		params = &DescribePipeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePipe", params, optFns, c.addOperationDescribePipeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePipeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePipeInput struct {

	// The name of the pipe.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type DescribePipeOutput struct {

	// The ARN of the pipe.
	Arn *string

	// The time the pipe was created.
	CreationTime *time.Time

	// The state the pipe is in.
	CurrentState types.PipeState

	// A description of the pipe.
	Description *string

	// The state the pipe should be in.
	DesiredState types.RequestedPipeStateDescribeResponse

	// The ARN of the enrichment resource.
	Enrichment *string

	// The parameters required to set up enrichment on your pipe.
	EnrichmentParameters *types.PipeEnrichmentParameters

	// When the pipe was last updated, in [ISO-8601 format] (YYYY-MM-DDThh:mm:ss.sTZD).
	//
	// [ISO-8601 format]: https://www.w3.org/TR/NOTE-datetime
	LastModifiedTime *time.Time

	// The logging configuration settings for the pipe.
	LogConfiguration *types.PipeLogConfiguration

	// The name of the pipe.
	Name *string

	// The ARN of the role that allows the pipe to send data to the target.
	RoleArn *string

	// The ARN of the source resource.
	Source *string

	// The parameters required to set up a source for your pipe.
	SourceParameters *types.PipeSourceParameters

	// The reason the pipe is in its current state.
	StateReason *string

	// The list of key-value pairs to associate with the pipe.
	Tags map[string]string

	// The ARN of the target resource.
	Target *string

	// The parameters required to set up a target for your pipe.
	//
	// For more information about pipe target parameters, including how to use dynamic
	// path parameters, see [Target parameters]in the Amazon EventBridge User Guide.
	//
	// [Target parameters]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-event-target.html
	TargetParameters *types.PipeTargetParameters

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePipeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribePipe{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribePipe{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribePipe"); err != nil {
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
	if err = addOpDescribePipeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePipe(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribePipe(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribePipe",
	}
}
