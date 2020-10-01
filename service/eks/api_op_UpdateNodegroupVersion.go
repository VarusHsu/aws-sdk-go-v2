// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the Kubernetes version or AMI version of an Amazon EKS managed node
// group. You can update to the latest available AMI version of a node group's
// current Kubernetes version by not specifying a Kubernetes version in the
// request. You can update to the latest AMI version of your cluster's current
// Kubernetes version by specifying your cluster's Kubernetes version in the
// request. For more information, see Amazon EKS-Optimized Linux AMI Versions
// (https://docs.aws.amazon.com/eks/latest/userguide/eks-linux-ami-versions.html)
// in the Amazon EKS User Guide. You cannot roll back a node group to an earlier
// Kubernetes version or AMI version. When a node in a managed node group is
// terminated due to a scaling action or update, the pods in that node are drained
// first. Amazon EKS attempts to drain the nodes gracefully and will fail if it is
// unable to do so. You can force the update if Amazon EKS is unable to drain the
// nodes as a result of a pod disruption budget issue.
func (c *Client) UpdateNodegroupVersion(ctx context.Context, params *UpdateNodegroupVersionInput, optFns ...func(*Options)) (*UpdateNodegroupVersionOutput, error) {
	stack := middleware.NewStack("UpdateNodegroupVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateNodegroupVersionMiddlewares(stack)
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
	addIdempotencyToken_opUpdateNodegroupVersionMiddleware(stack, options)
	addOpUpdateNodegroupVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNodegroupVersion(options.Region), middleware.Before)
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
			OperationName: "UpdateNodegroupVersion",
			Err:           err,
		}
	}
	out := result.(*UpdateNodegroupVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateNodegroupVersionInput struct {

	// The AMI version of the Amazon EKS-optimized AMI to use for the update. By
	// default, the latest available AMI version for the node group's Kubernetes
	// version is used. For more information, see Amazon EKS-Optimized Linux AMI
	// Versions
	// (https://docs.aws.amazon.com/eks/latest/userguide/eks-linux-ami-versions.html)
	// in the Amazon EKS User Guide.
	ReleaseVersion *string

	// Force the update if the existing node group's pods are unable to be drained due
	// to a pod disruption budget issue. If an update fails because pods could not be
	// drained, you can force the update after it fails to terminate the old node
	// whether or not any pods are running on the node.
	Force *bool

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string

	// The Kubernetes version to update to. If no version is specified, then the
	// Kubernetes version of the node group does not change. You can specify the
	// Kubernetes version of the cluster to update the node group to the latest AMI
	// version of the cluster's Kubernetes version.
	Version *string

	// The name of the Amazon EKS cluster that is associated with the managed node
	// group to update.
	//
	// This member is required.
	ClusterName *string

	// The name of the managed node group to update.
	//
	// This member is required.
	NodegroupName *string
}

type UpdateNodegroupVersionOutput struct {

	// An object representing an asynchronous update.
	Update *types.Update

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateNodegroupVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateNodegroupVersion{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateNodegroupVersion{}, middleware.After)
}

type idempotencyToken_initializeOpUpdateNodegroupVersion struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateNodegroupVersion) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateNodegroupVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateNodegroupVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateNodegroupVersionInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateNodegroupVersionMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpUpdateNodegroupVersion{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateNodegroupVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "eks",
		OperationName: "UpdateNodegroupVersion",
	}
}
