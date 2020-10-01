// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all approval rule templates that are associated with a specified
// repository.
func (c *Client) ListAssociatedApprovalRuleTemplatesForRepository(ctx context.Context, params *ListAssociatedApprovalRuleTemplatesForRepositoryInput, optFns ...func(*Options)) (*ListAssociatedApprovalRuleTemplatesForRepositoryOutput, error) {
	stack := middleware.NewStack("ListAssociatedApprovalRuleTemplatesForRepository", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListAssociatedApprovalRuleTemplatesForRepositoryMiddlewares(stack)
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
	addOpListAssociatedApprovalRuleTemplatesForRepositoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAssociatedApprovalRuleTemplatesForRepository(options.Region), middleware.Before)
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
			OperationName: "ListAssociatedApprovalRuleTemplatesForRepository",
			Err:           err,
		}
	}
	out := result.(*ListAssociatedApprovalRuleTemplatesForRepositoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssociatedApprovalRuleTemplatesForRepositoryInput struct {

	// A non-zero, non-negative integer used to limit the number of returned results.
	MaxResults *int32

	// An enumeration token that, when provided in a request, returns the next batch of
	// the results.
	NextToken *string

	// The name of the repository for which you want to list all associated approval
	// rule templates.
	//
	// This member is required.
	RepositoryName *string
}

type ListAssociatedApprovalRuleTemplatesForRepositoryOutput struct {

	// The names of all approval rule templates associated with the repository.
	ApprovalRuleTemplateNames []*string

	// An enumeration token that allows the operation to batch the next results of the
	// operation.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListAssociatedApprovalRuleTemplatesForRepositoryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListAssociatedApprovalRuleTemplatesForRepository{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAssociatedApprovalRuleTemplatesForRepository{}, middleware.After)
}

func newServiceMetadataMiddleware_opListAssociatedApprovalRuleTemplatesForRepository(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "ListAssociatedApprovalRuleTemplatesForRepository",
	}
}
