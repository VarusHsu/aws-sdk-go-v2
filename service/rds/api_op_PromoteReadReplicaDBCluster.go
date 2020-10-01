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

// Promotes a read replica DB cluster to a standalone DB cluster. This action only
// applies to Aurora DB clusters.
func (c *Client) PromoteReadReplicaDBCluster(ctx context.Context, params *PromoteReadReplicaDBClusterInput, optFns ...func(*Options)) (*PromoteReadReplicaDBClusterOutput, error) {
	stack := middleware.NewStack("PromoteReadReplicaDBCluster", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpPromoteReadReplicaDBClusterMiddlewares(stack)
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
	addOpPromoteReadReplicaDBClusterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPromoteReadReplicaDBCluster(options.Region), middleware.Before)
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
			OperationName: "PromoteReadReplicaDBCluster",
			Err:           err,
		}
	}
	out := result.(*PromoteReadReplicaDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type PromoteReadReplicaDBClusterInput struct {

	// The identifier of the DB cluster read replica to promote. This parameter isn't
	// case-sensitive. Constraints:
	//
	//     * Must match the identifier of an existing DB
	// cluster read replica.
	//
	// Example: my-cluster-replica1
	//
	// This member is required.
	DBClusterIdentifier *string
}

type PromoteReadReplicaDBClusterOutput struct {

	// Contains the details of an Amazon Aurora DB cluster. This data type is used as a
	// response element in the DescribeDBClusters, StopDBCluster, and StartDBCluster
	// actions.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpPromoteReadReplicaDBClusterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpPromoteReadReplicaDBCluster{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpPromoteReadReplicaDBCluster{}, middleware.After)
}

func newServiceMetadataMiddleware_opPromoteReadReplicaDBCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "PromoteReadReplicaDBCluster",
	}
}
