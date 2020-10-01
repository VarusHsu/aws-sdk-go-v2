// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Get the details for a program in a multiplex.
func (c *Client) DescribeMultiplexProgram(ctx context.Context, params *DescribeMultiplexProgramInput, optFns ...func(*Options)) (*DescribeMultiplexProgramOutput, error) {
	stack := middleware.NewStack("DescribeMultiplexProgram", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeMultiplexProgramMiddlewares(stack)
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
	addOpDescribeMultiplexProgramValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMultiplexProgram(options.Region), middleware.Before)
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
			OperationName: "DescribeMultiplexProgram",
			Err:           err,
		}
	}
	out := result.(*DescribeMultiplexProgramOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DescribeMultiplexProgramRequest
type DescribeMultiplexProgramInput struct {

	// The ID of the multiplex that the program belongs to.
	//
	// This member is required.
	MultiplexId *string

	// The name of the program.
	//
	// This member is required.
	ProgramName *string
}

// Placeholder documentation for DescribeMultiplexProgramResponse
type DescribeMultiplexProgramOutput struct {

	// The MediaLive channel associated with the program.
	ChannelId *string

	// The name of the multiplex program.
	ProgramName *string

	// The settings for this multiplex program.
	MultiplexProgramSettings *types.MultiplexProgramSettings

	// The packet identifier map for this multiplex program.
	PacketIdentifiersMap *types.MultiplexProgramPacketIdentifiersMap

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeMultiplexProgramMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeMultiplexProgram{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeMultiplexProgram{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeMultiplexProgram(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "DescribeMultiplexProgram",
	}
}
