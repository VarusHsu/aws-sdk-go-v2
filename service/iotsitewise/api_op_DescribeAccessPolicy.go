// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes an access policy, which specifies an AWS SSO user or group's access to
// an AWS IoT SiteWise Monitor portal or project.
func (c *Client) DescribeAccessPolicy(ctx context.Context, params *DescribeAccessPolicyInput, optFns ...func(*Options)) (*DescribeAccessPolicyOutput, error) {
	stack := middleware.NewStack("DescribeAccessPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeAccessPolicyMiddlewares(stack)
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
	addOpDescribeAccessPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccessPolicy(options.Region), middleware.Before)
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
			OperationName: "DescribeAccessPolicy",
			Err:           err,
		}
	}
	out := result.(*DescribeAccessPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccessPolicyInput struct {

	// The ID of the access policy.
	//
	// This member is required.
	AccessPolicyId *string
}

type DescribeAccessPolicyOutput struct {

	// The date the access policy was created, in Unix epoch time.
	//
	// This member is required.
	AccessPolicyCreationDate *time.Time

	// The date the access policy was last updated, in Unix epoch time.
	//
	// This member is required.
	AccessPolicyLastUpdateDate *time.Time

	// The access policy permission. Note that a project ADMINISTRATOR is also known as
	// a project owner.
	//
	// This member is required.
	AccessPolicyPermission types.Permission

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the access policy, which has the following format.
	// arn:${Partition}:iotsitewise:${Region}:${Account}:access-policy/${AccessPolicyId}
	//
	// This member is required.
	AccessPolicyArn *string

	// The AWS IoT SiteWise Monitor resource (portal or project) to which this access
	// policy provides access.
	//
	// This member is required.
	AccessPolicyResource *types.Resource

	// The ID of the access policy.
	//
	// This member is required.
	AccessPolicyId *string

	// The AWS SSO identity (user or group) to which this access policy applies.
	//
	// This member is required.
	AccessPolicyIdentity *types.Identity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeAccessPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAccessPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAccessPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAccessPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "DescribeAccessPolicy",
	}
}
