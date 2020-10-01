// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the list of AWS Identity and Access Management (IAM) roles that can be
// used by the cluster to access other AWS services. A cluster can have up to 10
// IAM roles associated at any time.
func (c *Client) ModifyClusterIamRoles(ctx context.Context, params *ModifyClusterIamRolesInput, optFns ...func(*Options)) (*ModifyClusterIamRolesOutput, error) {
	stack := middleware.NewStack("ModifyClusterIamRoles", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpModifyClusterIamRolesMiddlewares(stack)
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
	addOpModifyClusterIamRolesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyClusterIamRoles(options.Region), middleware.Before)
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
			OperationName: "ModifyClusterIamRoles",
			Err:           err,
		}
	}
	out := result.(*ModifyClusterIamRolesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ModifyClusterIamRolesInput struct {

	// The unique identifier of the cluster for which you want to associate or
	// disassociate IAM roles.
	//
	// This member is required.
	ClusterIdentifier *string

	// Zero or more IAM roles to associate with the cluster. The roles must be in their
	// Amazon Resource Name (ARN) format. You can associate up to 10 IAM roles with a
	// single cluster in a single request.
	AddIamRoles []*string

	// Zero or more IAM roles in ARN format to disassociate from the cluster. You can
	// disassociate up to 10 IAM roles from a single cluster in a single request.
	RemoveIamRoles []*string
}

type ModifyClusterIamRolesOutput struct {

	// Describes a cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpModifyClusterIamRolesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpModifyClusterIamRoles{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyClusterIamRoles{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyClusterIamRoles(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "ModifyClusterIamRoles",
	}
}
