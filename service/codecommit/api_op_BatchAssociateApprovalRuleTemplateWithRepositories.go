// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an association between an approval rule template and one or more
// specified repositories.
func (c *Client) BatchAssociateApprovalRuleTemplateWithRepositories(ctx context.Context, params *BatchAssociateApprovalRuleTemplateWithRepositoriesInput, optFns ...func(*Options)) (*BatchAssociateApprovalRuleTemplateWithRepositoriesOutput, error) {
	stack := middleware.NewStack("BatchAssociateApprovalRuleTemplateWithRepositories", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpBatchAssociateApprovalRuleTemplateWithRepositoriesMiddlewares(stack)
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
	addOpBatchAssociateApprovalRuleTemplateWithRepositoriesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchAssociateApprovalRuleTemplateWithRepositories(options.Region), middleware.Before)
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
			OperationName: "BatchAssociateApprovalRuleTemplateWithRepositories",
			Err:           err,
		}
	}
	out := result.(*BatchAssociateApprovalRuleTemplateWithRepositoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchAssociateApprovalRuleTemplateWithRepositoriesInput struct {

	// The name of the template you want to associate with one or more repositories.
	//
	// This member is required.
	ApprovalRuleTemplateName *string

	// The names of the repositories you want to associate with the template. The
	// length constraint limit is for each string in the array. The array itself can be
	// empty.
	//
	// This member is required.
	RepositoryNames []*string
}

type BatchAssociateApprovalRuleTemplateWithRepositoriesOutput struct {

	// A list of any errors that might have occurred while attempting to create the
	// association between the template and the repositories.
	//
	// This member is required.
	Errors []*types.BatchAssociateApprovalRuleTemplateWithRepositoriesError

	// A list of names of the repositories that have been associated with the template.
	//
	// This member is required.
	AssociatedRepositoryNames []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpBatchAssociateApprovalRuleTemplateWithRepositoriesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpBatchAssociateApprovalRuleTemplateWithRepositories{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchAssociateApprovalRuleTemplateWithRepositories{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchAssociateApprovalRuleTemplateWithRepositories(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "BatchAssociateApprovalRuleTemplateWithRepositories",
	}
}
