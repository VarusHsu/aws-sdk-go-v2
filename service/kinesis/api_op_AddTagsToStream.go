// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds or updates tags for the specified Kinesis data stream. Each time you invoke
// this operation, you can specify up to 10 tags. If you want to add more than 10
// tags to your stream, you can invoke this operation multiple times. In total,
// each stream can have up to 50 tags. If tags have already been assigned to the
// stream, AddTagsToStream overwrites any existing tags that correspond to the
// specified tag keys. AddTagsToStream () has a limit of five transactions per
// second per account.
func (c *Client) AddTagsToStream(ctx context.Context, params *AddTagsToStreamInput, optFns ...func(*Options)) (*AddTagsToStreamOutput, error) {
	stack := middleware.NewStack("AddTagsToStream", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAddTagsToStreamMiddlewares(stack)
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
	addOpAddTagsToStreamValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddTagsToStream(options.Region), middleware.Before)
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
			OperationName: "AddTagsToStream",
			Err:           err,
		}
	}
	out := result.(*AddTagsToStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for AddTagsToStream.
type AddTagsToStreamInput struct {

	// The name of the stream.
	//
	// This member is required.
	StreamName *string

	// A set of up to 10 key-value pairs to use to create the tags.
	//
	// This member is required.
	Tags map[string]*string
}

type AddTagsToStreamOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAddTagsToStreamMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAddTagsToStream{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddTagsToStream{}, middleware.After)
}

func newServiceMetadataMiddleware_opAddTagsToStream(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "AddTagsToStream",
	}
}
