// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns high-level aggregated patch compliance state for a patch group.
func (c *Client) DescribePatchGroupState(ctx context.Context, params *DescribePatchGroupStateInput, optFns ...func(*Options)) (*DescribePatchGroupStateOutput, error) {
	stack := middleware.NewStack("DescribePatchGroupState", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribePatchGroupStateMiddlewares(stack)
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
	addOpDescribePatchGroupStateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePatchGroupState(options.Region), middleware.Before)
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
			OperationName: "DescribePatchGroupState",
			Err:           err,
		}
	}
	out := result.(*DescribePatchGroupStateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePatchGroupStateInput struct {

	// The name of the patch group whose patch snapshot should be retrieved.
	//
	// This member is required.
	PatchGroup *string
}

type DescribePatchGroupStateOutput struct {

	// The number of instances with NotApplicable patches beyond the supported limit,
	// which are not reported by name to Systems Manager Inventory.
	InstancesWithUnreportedNotApplicablePatches *int32

	// The number of instances with patches from the patch baseline that failed to
	// install.
	InstancesWithFailedPatches *int32

	// The number of instances with installed patches.
	InstancesWithInstalledPatches *int32

	// The number of instances with patches installed by Patch Manager that have not
	// been rebooted after the patch installation. The status of these instances is
	// NON_COMPLIANT.
	InstancesWithInstalledPendingRebootPatches *int32

	// The number of instances with patches installed that are specified in a
	// RejectedPatches list. Patches with a status of INSTALLED_REJECTED were typically
	// installed before they were added to a RejectedPatches list. If
	// ALLOW_AS_DEPENDENCY is the specified option for RejectedPatchesAction, the value
	// of InstancesWithInstalledRejectedPatches will always be 0 (zero).
	InstancesWithInstalledRejectedPatches *int32

	// The number of instances with patches that aren't applicable.
	InstancesWithNotApplicablePatches *int32

	// The number of instances with patches installed that aren't defined in the patch
	// baseline.
	InstancesWithInstalledOtherPatches *int32

	// The number of instances with missing patches from the patch baseline.
	InstancesWithMissingPatches *int32

	// The number of instances in the patch group.
	Instances *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribePatchGroupStateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePatchGroupState{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePatchGroupState{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribePatchGroupState(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribePatchGroupState",
	}
}
