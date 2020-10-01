// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Get information about a dedicated IP address, including the name of the
// dedicated IP pool that it's associated with, as well information about the
// automatic warm-up process for the address.
func (c *Client) GetDedicatedIp(ctx context.Context, params *GetDedicatedIpInput, optFns ...func(*Options)) (*GetDedicatedIpOutput, error) {
	stack := middleware.NewStack("GetDedicatedIp", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetDedicatedIpMiddlewares(stack)
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
	addOpGetDedicatedIpValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDedicatedIp(options.Region), middleware.Before)
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
			OperationName: "GetDedicatedIp",
			Err:           err,
		}
	}
	out := result.(*GetDedicatedIpOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to obtain more information about a dedicated IP address.
type GetDedicatedIpInput struct {

	// The IP address that you want to obtain more information about. The value you
	// specify has to be a dedicated IP address that's assocaited with your AWS
	// account.
	//
	// This member is required.
	Ip *string
}

// Information about a dedicated IP address.
type GetDedicatedIpOutput struct {

	// An object that contains information about a dedicated IP address.
	DedicatedIp *types.DedicatedIp

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetDedicatedIpMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetDedicatedIp{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDedicatedIp{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDedicatedIp(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetDedicatedIp",
	}
}
