// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates an IAM instance profile with a running or stopped instance. You
// cannot associate more than one IAM instance profile with an instance.
func (c *Client) AssociateIamInstanceProfile(ctx context.Context, params *AssociateIamInstanceProfileInput, optFns ...func(*Options)) (*AssociateIamInstanceProfileOutput, error) {
	stack := middleware.NewStack("AssociateIamInstanceProfile", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpAssociateIamInstanceProfileMiddlewares(stack)
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
	addOpAssociateIamInstanceProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateIamInstanceProfile(options.Region), middleware.Before)
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
			OperationName: "AssociateIamInstanceProfile",
			Err:           err,
		}
	}
	out := result.(*AssociateIamInstanceProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateIamInstanceProfileInput struct {

	// The ID of the instance.
	//
	// This member is required.
	InstanceId *string

	// The IAM instance profile.
	//
	// This member is required.
	IamInstanceProfile *types.IamInstanceProfileSpecification
}

type AssociateIamInstanceProfileOutput struct {

	// Information about the IAM instance profile association.
	IamInstanceProfileAssociation *types.IamInstanceProfileAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpAssociateIamInstanceProfileMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpAssociateIamInstanceProfile{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpAssociateIamInstanceProfile{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateIamInstanceProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AssociateIamInstanceProfile",
	}
}
