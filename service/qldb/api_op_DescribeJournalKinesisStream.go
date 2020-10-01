// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns detailed information about a given Amazon QLDB journal stream. The
// output includes the Amazon Resource Name (ARN), stream name, current status,
// creation time, and the parameters of your original stream creation request.
func (c *Client) DescribeJournalKinesisStream(ctx context.Context, params *DescribeJournalKinesisStreamInput, optFns ...func(*Options)) (*DescribeJournalKinesisStreamOutput, error) {
	stack := middleware.NewStack("DescribeJournalKinesisStream", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeJournalKinesisStreamMiddlewares(stack)
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
	addOpDescribeJournalKinesisStreamValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeJournalKinesisStream(options.Region), middleware.Before)
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
			OperationName: "DescribeJournalKinesisStream",
			Err:           err,
		}
	}
	out := result.(*DescribeJournalKinesisStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeJournalKinesisStreamInput struct {

	// The unique ID that QLDB assigns to each QLDB journal stream.
	//
	// This member is required.
	StreamId *string

	// The name of the ledger.
	//
	// This member is required.
	LedgerName *string
}

type DescribeJournalKinesisStreamOutput struct {

	// Information about the QLDB journal stream returned by a DescribeJournalS3Export
	// request.
	Stream *types.JournalKinesisStreamDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeJournalKinesisStreamMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeJournalKinesisStream{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeJournalKinesisStream{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeJournalKinesisStream(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "DescribeJournalKinesisStream",
	}
}
