// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds or updates an inline policy document that is embedded in the specified IAM
// role. When you embed an inline policy in a role, the inline policy is used as
// part of the role's access (permissions) policy. The role's trust policy is
// created at the same time as the role, using CreateRole (). You can update a
// role's trust policy using UpdateAssumeRolePolicy (). For more information about
// IAM roles, go to Using Roles to Delegate Permissions and Federate Identities
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/roles-toplevel.html). A role
// can also have a managed policy attached to it. To attach a managed policy to a
// role, use AttachRolePolicy (). To create a new managed policy, use CreatePolicy
// (). For information about policies, see Managed Policies and Inline Policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide. For information about limits on the number of inline
// policies that you can embed with a role, see Limitations on IAM Entities
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/LimitationsOnEntities.html) in
// the IAM User Guide. Because policy documents can be large, you should use POST
// rather than GET when calling PutRolePolicy. For general information about using
// the Query API with IAM, go to Making Query Requests
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/IAM_UsingQueryAPI.html) in the
// IAM User Guide.
func (c *Client) PutRolePolicy(ctx context.Context, params *PutRolePolicyInput, optFns ...func(*Options)) (*PutRolePolicyOutput, error) {
	stack := middleware.NewStack("PutRolePolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpPutRolePolicyMiddlewares(stack)
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
	addOpPutRolePolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutRolePolicy(options.Region), middleware.Before)
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
			OperationName: "PutRolePolicy",
			Err:           err,
		}
	}
	out := result.(*PutRolePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRolePolicyInput struct {

	// The name of the policy document. This parameter allows (through its regex
	// pattern (http://wikipedia.org/wiki/regex)) a string of characters consisting of
	// upper and lowercase alphanumeric characters with no spaces. You can also include
	// any of the following characters: _+=,.@-
	//
	// This member is required.
	PolicyName *string

	// The name of the role to associate the policy with. This parameter allows
	// (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	RoleName *string

	// The policy document. You must provide policies in JSON format in IAM. However,
	// for AWS CloudFormation templates formatted in YAML, you can provide the policy
	// in JSON or YAML format. AWS CloudFormation always converts a YAML policy to JSON
	// format before submitting it to IAM. The regex pattern
	// (http://wikipedia.org/wiki/regex) used to validate this parameter is a string of
	// characters consisting of the following:
	//
	//     * Any printable ASCII character
	// ranging from the space character (\u0020) through the end of the ASCII character
	// range
	//
	//     * The printable characters in the Basic Latin and Latin-1 Supplement
	// character set (through \u00FF)
	//
	//     * The special characters tab (\u0009), line
	// feed (\u000A), and carriage return (\u000D)
	//
	// This member is required.
	PolicyDocument *string
}

type PutRolePolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpPutRolePolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpPutRolePolicy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpPutRolePolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutRolePolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "PutRolePolicy",
	}
}
