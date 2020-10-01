// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Uploads segment documents to AWS X-Ray. The X-Ray SDK
// (https://docs.aws.amazon.com/xray/index.html) generates segment documents and
// sends them to the X-Ray daemon, which uploads them in batches. A segment
// document can be a completed segment, an in-progress segment, or an array of
// subsegments. Segments must include the following fields. For the full segment
// document schema, see AWS X-Ray Segment Documents
// (https://docs.aws.amazon.com/xray/latest/devguide/xray-api-segmentdocuments.html)
// in the AWS X-Ray Developer Guide. Required Segment Document Fields
//
//     * name -
// The name of the service that handled the request.
//
//     * id - A 64-bit
// identifier for the segment, unique among segments in the same trace, in 16
// hexadecimal digits.
//
//     * trace_id - A unique identifier that connects all
// segments and subsegments originating from a single client request.
//
//     *
// start_time - Time the segment or subsegment was created, in floating point
// seconds in epoch time, accurate to milliseconds. For example, 1480615200.010 or
// 1.480615200010E9.
//
//     * end_time - Time the segment or subsegment was closed.
// For example, 1480615200.090 or 1.480615200090E9. Specify either an end_time or
// in_progress.
//
//     * in_progress - Set to true instead of specifying an end_time
// to record that a segment has been started, but is not complete. Send an in
// progress segment when your application receives a request that will take a long
// time to serve, to trace the fact that the request was received. When the
// response is sent, send the complete segment to overwrite the in-progress
// segment.
//
// A trace_id consists of three numbers separated by hyphens. For
// example, 1-58406520-a006649127e371903a2de979. This includes: Trace ID Format
//
//
// * The version number, i.e. 1.
//
//     * The time of the original request, in Unix
// epoch time, in 8 hexadecimal digits. For example, 10:00AM December 2nd, 2016 PST
// in epoch time is 1480615200 seconds, or 58406520 in hexadecimal.
//
//     * A 96-bit
// identifier for the trace, globally unique, in 24 hexadecimal digits.
func (c *Client) PutTraceSegments(ctx context.Context, params *PutTraceSegmentsInput, optFns ...func(*Options)) (*PutTraceSegmentsOutput, error) {
	stack := middleware.NewStack("PutTraceSegments", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutTraceSegmentsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpPutTraceSegmentsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutTraceSegments(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "PutTraceSegments",
			Err:           err,
		}
	}
	out := result.(*PutTraceSegmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutTraceSegmentsInput struct {

	// A string containing a JSON document defining one or more segments or
	// subsegments.
	//
	// This member is required.
	TraceSegmentDocuments []*string
}

type PutTraceSegmentsOutput struct {

	// Segments that failed processing.
	UnprocessedTraceSegments []*types.UnprocessedTraceSegment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutTraceSegmentsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutTraceSegments{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutTraceSegments{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutTraceSegments(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "xray",
		OperationName: "PutTraceSegments",
	}
}
