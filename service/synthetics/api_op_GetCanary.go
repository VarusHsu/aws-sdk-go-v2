// Code generated by smithy-go-codegen DO NOT EDIT.

package synthetics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/synthetics/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves complete information about one canary. You must specify the name of
// the canary that you want. To get a list of canaries and their names, use
// DescribeCanaries
// (https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_DescribeCanaries.html).
func (c *Client) GetCanary(ctx context.Context, params *GetCanaryInput, optFns ...func(*Options)) (*GetCanaryOutput, error) {
	stack := middleware.NewStack("GetCanary", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetCanaryMiddlewares(stack)
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
	addOpGetCanaryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCanary(options.Region), middleware.Before)
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
			OperationName: "GetCanary",
			Err:           err,
		}
	}
	out := result.(*GetCanaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCanaryInput struct {

	// The name of the canary that you want details for.
	//
	// This member is required.
	Name *string
}

type GetCanaryOutput struct {

	// A strucure that contains the full information about the canary.
	Canary *types.Canary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetCanaryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetCanary{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCanary{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCanary(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "synthetics",
		OperationName: "GetCanary",
	}
}
