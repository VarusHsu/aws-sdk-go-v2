// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the resource policy that is set on a repository. After a resource policy
// is deleted, the permissions allowed and denied by the deleted policy are
// removed. The effect of deleting a resource policy might not be immediate. Use
// DeleteRepositoryPermissionsPolicy with caution. After a policy is deleted, AWS
// users, roles, and accounts lose permissions to perform the repository actions
// granted by the deleted policy.
func (c *Client) DeleteRepositoryPermissionsPolicy(ctx context.Context, params *DeleteRepositoryPermissionsPolicyInput, optFns ...func(*Options)) (*DeleteRepositoryPermissionsPolicyOutput, error) {
	stack := middleware.NewStack("DeleteRepositoryPermissionsPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteRepositoryPermissionsPolicyMiddlewares(stack)
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
	addOpDeleteRepositoryPermissionsPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRepositoryPermissionsPolicy(options.Region), middleware.Before)
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
			OperationName: "DeleteRepositoryPermissionsPolicy",
			Err:           err,
		}
	}
	out := result.(*DeleteRepositoryPermissionsPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRepositoryPermissionsPolicyInput struct {

	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string

	// The name of the repository that is associated with the resource policy to be
	// deleted
	//
	// This member is required.
	Repository *string

	// The revision of the repository's resource policy to be deleted. This revision is
	// used for optimistic locking, which prevents others from accidentally overwriting
	// your changes to the repository's resource policy.
	PolicyRevision *string

	// The name of the domain that contains the repository associated with the resource
	// policy to be deleted.
	//
	// This member is required.
	Domain *string
}

type DeleteRepositoryPermissionsPolicyOutput struct {

	// Information about the deleted policy after processing the request.
	Policy *types.ResourcePolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteRepositoryPermissionsPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteRepositoryPermissionsPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteRepositoryPermissionsPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteRepositoryPermissionsPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "DeleteRepositoryPermissionsPolicy",
	}
}
