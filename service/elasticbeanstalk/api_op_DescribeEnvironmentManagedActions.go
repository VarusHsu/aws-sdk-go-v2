// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists an environment's upcoming and in-progress managed actions.
func (c *Client) DescribeEnvironmentManagedActions(ctx context.Context, params *DescribeEnvironmentManagedActionsInput, optFns ...func(*Options)) (*DescribeEnvironmentManagedActionsOutput, error) {
	stack := middleware.NewStack("DescribeEnvironmentManagedActions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeEnvironmentManagedActionsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEnvironmentManagedActions(options.Region), middleware.Before)
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
			OperationName: "DescribeEnvironmentManagedActions",
			Err:           err,
		}
	}
	out := result.(*DescribeEnvironmentManagedActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to list an environment's upcoming and in-progress managed actions.
type DescribeEnvironmentManagedActionsInput struct {

	// The name of the target environment.
	EnvironmentName *string

	// The environment ID of the target environment.
	EnvironmentId *string

	// To show only actions with a particular status, specify a status.
	Status types.ActionStatus
}

// The result message containing a list of managed actions.
type DescribeEnvironmentManagedActionsOutput struct {

	// A list of upcoming and in-progress managed actions.
	ManagedActions []*types.ManagedAction

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeEnvironmentManagedActionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEnvironmentManagedActions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEnvironmentManagedActions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeEnvironmentManagedActions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "DescribeEnvironmentManagedActions",
	}
}
