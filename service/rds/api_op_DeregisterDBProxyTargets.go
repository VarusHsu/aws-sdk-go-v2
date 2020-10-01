// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Remove the association between one or more DBProxyTarget data structures and a
// DBProxyTargetGroup.
func (c *Client) DeregisterDBProxyTargets(ctx context.Context, params *DeregisterDBProxyTargetsInput, optFns ...func(*Options)) (*DeregisterDBProxyTargetsOutput, error) {
	stack := middleware.NewStack("DeregisterDBProxyTargets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeregisterDBProxyTargetsMiddlewares(stack)
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
	addOpDeregisterDBProxyTargetsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterDBProxyTargets(options.Region), middleware.Before)
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
			OperationName: "DeregisterDBProxyTargets",
			Err:           err,
		}
	}
	out := result.(*DeregisterDBProxyTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterDBProxyTargetsInput struct {

	// The identifier of the DBProxyTargetGroup.
	TargetGroupName *string

	// One or more DB cluster identifiers.
	DBClusterIdentifiers []*string

	// The identifier of the DBProxy that is associated with the DBProxyTargetGroup.
	//
	// This member is required.
	DBProxyName *string

	// One or more DB instance identifiers.
	DBInstanceIdentifiers []*string
}

type DeregisterDBProxyTargetsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeregisterDBProxyTargetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeregisterDBProxyTargets{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeregisterDBProxyTargets{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeregisterDBProxyTargets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DeregisterDBProxyTargets",
	}
}
