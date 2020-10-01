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

// Modifies the specified attribute of the specified AMI. You can specify only one
// attribute at a time. You can use the Attribute parameter to specify the
// attribute or one of the following parameters: Description, LaunchPermission, or
// ProductCode. AWS Marketplace product codes cannot be modified. Images with an
// AWS Marketplace product code cannot be made public. To enable the
// SriovNetSupport enhanced networking attribute of an image, enable
// SriovNetSupport on an instance and create an AMI from the instance.
func (c *Client) ModifyImageAttribute(ctx context.Context, params *ModifyImageAttributeInput, optFns ...func(*Options)) (*ModifyImageAttributeOutput, error) {
	stack := middleware.NewStack("ModifyImageAttribute", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpModifyImageAttributeMiddlewares(stack)
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
	addOpModifyImageAttributeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyImageAttribute(options.Region), middleware.Before)
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
			OperationName: "ModifyImageAttribute",
			Err:           err,
		}
	}
	out := result.(*ModifyImageAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for ModifyImageAttribute.
type ModifyImageAttributeInput struct {

	// The value of the attribute being modified. This parameter can be used only when
	// the Attribute parameter is description or productCodes.
	Value *string

	// The operation type. This parameter can be used only when the Attribute parameter
	// is launchPermission.
	OperationType types.OperationType

	// The ID of the AMI.
	//
	// This member is required.
	ImageId *string

	// A new description for the AMI.
	Description *types.AttributeValue

	// The DevPay product codes. After you add a product code to an AMI, it can't be
	// removed.
	ProductCodes []*string

	// The name of the attribute to modify. The valid values are description,
	// launchPermission, and productCodes.
	Attribute *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The user groups. This parameter can be used only when the Attribute parameter is
	// launchPermission.
	UserGroups []*string

	// A new launch permission for the AMI.
	LaunchPermission *types.LaunchPermissionModifications

	// The AWS account IDs. This parameter can be used only when the Attribute
	// parameter is launchPermission.
	UserIds []*string
}

type ModifyImageAttributeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpModifyImageAttributeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpModifyImageAttribute{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpModifyImageAttribute{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyImageAttribute(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyImageAttribute",
	}
}
