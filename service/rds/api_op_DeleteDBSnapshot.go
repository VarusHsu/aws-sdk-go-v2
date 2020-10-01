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

// Deletes a DB snapshot. If the snapshot is being copied, the copy operation is
// terminated. The DB snapshot must be in the available state to be deleted.
func (c *Client) DeleteDBSnapshot(ctx context.Context, params *DeleteDBSnapshotInput, optFns ...func(*Options)) (*DeleteDBSnapshotOutput, error) {
	stack := middleware.NewStack("DeleteDBSnapshot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteDBSnapshotMiddlewares(stack)
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
	addOpDeleteDBSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDBSnapshot(options.Region), middleware.Before)
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
			OperationName: "DeleteDBSnapshot",
			Err:           err,
		}
	}
	out := result.(*DeleteDBSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DeleteDBSnapshotInput struct {

	// The DB snapshot identifier. Constraints: Must be the name of an existing DB
	// snapshot in the available state.
	//
	// This member is required.
	DBSnapshotIdentifier *string
}

type DeleteDBSnapshotOutput struct {

	// Contains the details of an Amazon RDS DB snapshot. This data type is used as a
	// response element in the DescribeDBSnapshots action.
	DBSnapshot *types.DBSnapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteDBSnapshotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteDBSnapshot{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteDBSnapshot{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteDBSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DeleteDBSnapshot",
	}
}
