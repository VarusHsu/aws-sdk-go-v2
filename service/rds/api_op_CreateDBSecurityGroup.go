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

// Creates a new DB security group. DB security groups control access to a DB
// instance. A DB security group controls access to EC2-Classic DB instances that
// are not in a VPC.
func (c *Client) CreateDBSecurityGroup(ctx context.Context, params *CreateDBSecurityGroupInput, optFns ...func(*Options)) (*CreateDBSecurityGroupOutput, error) {
	stack := middleware.NewStack("CreateDBSecurityGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateDBSecurityGroupMiddlewares(stack)
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
	addOpCreateDBSecurityGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBSecurityGroup(options.Region), middleware.Before)
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
			OperationName: "CreateDBSecurityGroup",
			Err:           err,
		}
	}
	out := result.(*CreateDBSecurityGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateDBSecurityGroupInput struct {

	// Tags to assign to the DB security group.
	Tags []*types.Tag

	// The description for the DB security group.
	//
	// This member is required.
	DBSecurityGroupDescription *string

	// The name for the DB security group. This value is stored as a lowercase string.
	// Constraints:
	//
	//     * Must be 1 to 255 letters, numbers, or hyphens.
	//
	//     * First
	// character must be a letter
	//
	//     * Can't end with a hyphen or contain two
	// consecutive hyphens
	//
	//     * Must not be "Default"
	//
	// Example: mysecuritygroup
	//
	// This member is required.
	DBSecurityGroupName *string
}

type CreateDBSecurityGroupOutput struct {

	// Contains the details for an Amazon RDS DB security group. This data type is used
	// as a response element in the DescribeDBSecurityGroups action.
	DBSecurityGroup *types.DBSecurityGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateDBSecurityGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBSecurityGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBSecurityGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDBSecurityGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateDBSecurityGroup",
	}
}
