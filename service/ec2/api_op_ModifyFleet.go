// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the specified EC2 Fleet. You can only modify an EC2 Fleet request of
// type maintain. While the EC2 Fleet is being modified, it is in the modifying
// state. To scale up your EC2 Fleet, increase its target capacity. The EC2 Fleet
// launches the additional Spot Instances according to the allocation strategy for
// the EC2 Fleet request. If the allocation strategy is lowest-price, the EC2 Fleet
// launches instances using the Spot Instance pool with the lowest price. If the
// allocation strategy is diversified, the EC2 Fleet distributes the instances
// across the Spot Instance pools. If the allocation strategy is
// capacity-optimized, EC2 Fleet launches instances from Spot Instance pools with
// optimal capacity for the number of instances that are launching. To scale down
// your EC2 Fleet, decrease its target capacity. First, the EC2 Fleet cancels any
// open requests that exceed the new target capacity. You can request that the EC2
// Fleet terminate Spot Instances until the size of the fleet no longer exceeds the
// new target capacity. If the allocation strategy is lowest-price, the EC2 Fleet
// terminates the instances with the highest price per unit. If the allocation
// strategy is capacity-optimized, the EC2 Fleet terminates the instances in the
// Spot Instance pools that have the least available Spot Instance capacity. If the
// allocation strategy is diversified, the EC2 Fleet terminates instances across
// the Spot Instance pools. Alternatively, you can request that the EC2 Fleet keep
// the fleet at its current size, but not replace any Spot Instances that are
// interrupted or that you terminate manually. If you are finished with your EC2
// Fleet for now, but will use it again later, you can set the target capacity to
// 0.
func (c *Client) ModifyFleet(ctx context.Context, params *ModifyFleetInput, optFns ...func(*Options)) (*ModifyFleetOutput, error) {
	stack := middleware.NewStack("ModifyFleet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpModifyFleetMiddlewares(stack)
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
	addOpModifyFleetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyFleet(options.Region), middleware.Before)
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
			OperationName: "ModifyFleet",
			Err:           err,
		}
	}
	out := result.(*ModifyFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyFleetInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The ID of the EC2 Fleet.
	//
	// This member is required.
	FleetId *string

	// Indicates whether running instances should be terminated if the total target
	// capacity of the EC2 Fleet is decreased below the current size of the EC2 Fleet.
	ExcessCapacityTerminationPolicy types.FleetExcessCapacityTerminationPolicy

	// The size of the EC2 Fleet.
	//
	// This member is required.
	TargetCapacitySpecification *types.TargetCapacitySpecificationRequest
}

type ModifyFleetOutput struct {

	// Is true if the request succeeds, and an error otherwise.
	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpModifyFleetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpModifyFleet{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpModifyFleet{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyFleet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyFleet",
	}
}
