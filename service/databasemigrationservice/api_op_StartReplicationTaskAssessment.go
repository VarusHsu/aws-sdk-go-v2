// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts the replication task assessment for unsupported data types in the source
// database.
func (c *Client) StartReplicationTaskAssessment(ctx context.Context, params *StartReplicationTaskAssessmentInput, optFns ...func(*Options)) (*StartReplicationTaskAssessmentOutput, error) {
	stack := middleware.NewStack("StartReplicationTaskAssessment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartReplicationTaskAssessmentMiddlewares(stack)
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
	addOpStartReplicationTaskAssessmentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartReplicationTaskAssessment(options.Region), middleware.Before)
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
			OperationName: "StartReplicationTaskAssessment",
			Err:           err,
		}
	}
	out := result.(*StartReplicationTaskAssessmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type StartReplicationTaskAssessmentInput struct {

	// The Amazon Resource Name (ARN) of the replication task.
	//
	// This member is required.
	ReplicationTaskArn *string
}

//
type StartReplicationTaskAssessmentOutput struct {

	// The assessed replication task.
	ReplicationTask *types.ReplicationTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartReplicationTaskAssessmentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartReplicationTaskAssessment{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartReplicationTaskAssessment{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartReplicationTaskAssessment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "StartReplicationTaskAssessment",
	}
}
