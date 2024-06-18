// Code generated by smithy-go-codegen DO NOT EDIT.

package dlm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/dlm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified lifecycle policy.
//
// For more information about updating a policy, see [Modify lifecycle policies].
//
// [Modify lifecycle policies]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/view-modify-delete.html#modify
func (c *Client) UpdateLifecyclePolicy(ctx context.Context, params *UpdateLifecyclePolicyInput, optFns ...func(*Options)) (*UpdateLifecyclePolicyOutput, error) {
	if params == nil {
		params = &UpdateLifecyclePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLifecyclePolicy", params, optFns, c.addOperationUpdateLifecyclePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLifecyclePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLifecyclePolicyInput struct {

	// The identifier of the lifecycle policy.
	//
	// This member is required.
	PolicyId *string

	//  [Default policies only] Indicates whether the policy should copy tags from the
	// source resource to the snapshot or AMI.
	CopyTags *bool

	//  [Default policies only] Specifies how often the policy should run and create
	// snapshots or AMIs. The creation frequency can range from 1 to 7 days.
	CreateInterval *int32

	//  [Default policies only] Specifies destination Regions for snapshot or AMI
	// copies. You can specify up to 3 destination Regions. If you do not want to
	// create cross-Region copies, omit this parameter.
	CrossRegionCopyTargets []types.CrossRegionCopyTarget

	// A description of the lifecycle policy.
	Description *string

	//  [Default policies only] Specifies exclusion parameters for volumes or
	// instances for which you do not want to create snapshots or AMIs. The policy will
	// not create snapshots or AMIs for target resources that match any of the
	// specified exclusion parameters.
	Exclusions *types.Exclusions

	// The Amazon Resource Name (ARN) of the IAM role used to run the operations
	// specified by the lifecycle policy.
	ExecutionRoleArn *string

	//  [Default policies only] Defines the snapshot or AMI retention behavior for the
	// policy if the source volume or instance is deleted, or if the policy enters the
	// error, disabled, or deleted state.
	//
	// By default (ExtendDeletion=false):
	//
	//   - If a source resource is deleted, Amazon Data Lifecycle Manager will
	//   continue to delete previously created snapshots or AMIs, up to but not including
	//   the last one, based on the specified retention period. If you want Amazon Data
	//   Lifecycle Manager to delete all snapshots or AMIs, including the last one,
	//   specify true .
	//
	//   - If a policy enters the error, disabled, or deleted state, Amazon Data
	//   Lifecycle Manager stops deleting snapshots and AMIs. If you want Amazon Data
	//   Lifecycle Manager to continue deleting snapshots or AMIs, including the last
	//   one, if the policy enters one of these states, specify true .
	//
	// If you enable extended deletion (ExtendDeletion=true), you override both
	// default behaviors simultaneously.
	//
	// Default: false
	ExtendDeletion *bool

	// The configuration of the lifecycle policy. You cannot update the policy type or
	// the resource type.
	PolicyDetails *types.PolicyDetails

	//  [Default policies only] Specifies how long the policy should retain snapshots
	// or AMIs before deleting them. The retention period can range from 2 to 14 days,
	// but it must be greater than the creation frequency to ensure that the policy
	// retains at least 1 snapshot or AMI at any given time.
	RetainInterval *int32

	// The desired activation state of the lifecycle policy after creation.
	State types.SettablePolicyStateValues

	noSmithyDocumentSerde
}

type UpdateLifecyclePolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLifecyclePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateLifecyclePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateLifecyclePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateLifecyclePolicy"); err != nil {
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
	if err = addOpUpdateLifecyclePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLifecyclePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLifecyclePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateLifecyclePolicy",
	}
}
