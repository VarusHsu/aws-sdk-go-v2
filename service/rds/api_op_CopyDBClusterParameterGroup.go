// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Copies the specified DB cluster parameter group. This action only applies to
// Aurora DB clusters.
func (c *Client) CopyDBClusterParameterGroup(ctx context.Context, params *CopyDBClusterParameterGroupInput, optFns ...func(*Options)) (*CopyDBClusterParameterGroupOutput, error) {
	stack := middleware.NewStack("CopyDBClusterParameterGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCopyDBClusterParameterGroupMiddlewares(stack)
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
	addOpCopyDBClusterParameterGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCopyDBClusterParameterGroup(options.Region), middleware.Before)
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
			OperationName: "CopyDBClusterParameterGroup",
			Err:           err,
		}
	}
	out := result.(*CopyDBClusterParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopyDBClusterParameterGroupInput struct {

	// A description for the copied DB cluster parameter group.
	//
	// This member is required.
	TargetDBClusterParameterGroupDescription *string

	// The identifier for the copied DB cluster parameter group. Constraints:
	//
	//     *
	// Can't be null, empty, or blank
	//
	//     * Must contain from 1 to 255 letters,
	// numbers, or hyphens
	//
	//     * First character must be a letter
	//
	//     * Can't end
	// with a hyphen or contain two consecutive hyphens
	//
	// Example:
	// my-cluster-param-group1
	//
	// This member is required.
	TargetDBClusterParameterGroupIdentifier *string

	// The identifier or Amazon Resource Name (ARN) for the source DB cluster parameter
	// group. For information about creating an ARN, see  Constructing an ARN for
	// Amazon RDS
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Tagging.ARN.html#USER_Tagging.ARN.Constructing)
	// in the Amazon Aurora User Guide. Constraints:
	//
	//     * Must specify a valid DB
	// cluster parameter group.
	//
	//     * If the source DB cluster parameter group is in
	// the same AWS Region as the copy, specify a valid DB parameter group identifier,
	// for example my-db-cluster-param-group, or a valid ARN.
	//
	//     * If the source DB
	// parameter group is in a different AWS Region than the copy, specify a valid DB
	// cluster parameter group ARN, for example
	// arn:aws:rds:us-east-1:123456789012:cluster-pg:custom-cluster-group1.
	//
	// This member is required.
	SourceDBClusterParameterGroupIdentifier *string

	// A list of tags. For more information, see Tagging Amazon RDS Resources
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in
	// the Amazon RDS User Guide.
	Tags []*types.Tag
}

type CopyDBClusterParameterGroupOutput struct {

	// Contains the details of an Amazon RDS DB cluster parameter group. This data type
	// is used as a response element in the DescribeDBClusterParameterGroups action.
	DBClusterParameterGroup *types.DBClusterParameterGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCopyDBClusterParameterGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCopyDBClusterParameterGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCopyDBClusterParameterGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCopyDBClusterParameterGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CopyDBClusterParameterGroup",
	}
}
