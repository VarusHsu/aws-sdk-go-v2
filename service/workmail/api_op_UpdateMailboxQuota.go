// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a user's current mailbox quota for a specified organization and user.
func (c *Client) UpdateMailboxQuota(ctx context.Context, params *UpdateMailboxQuotaInput, optFns ...func(*Options)) (*UpdateMailboxQuotaOutput, error) {
	stack := middleware.NewStack("UpdateMailboxQuota", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateMailboxQuotaMiddlewares(stack)
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
	addOpUpdateMailboxQuotaValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMailboxQuota(options.Region), middleware.Before)
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
			OperationName: "UpdateMailboxQuota",
			Err:           err,
		}
	}
	out := result.(*UpdateMailboxQuotaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMailboxQuotaInput struct {

	// The identifer for the user for whom to update the mailbox quota.
	//
	// This member is required.
	UserId *string

	// The identifier for the organization that contains the user for whom to update
	// the mailbox quota.
	//
	// This member is required.
	OrganizationId *string

	// The updated mailbox quota, in MB, for the specified user.
	//
	// This member is required.
	MailboxQuota *int32
}

type UpdateMailboxQuotaOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateMailboxQuotaMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateMailboxQuota{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateMailboxQuota{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateMailboxQuota(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "UpdateMailboxQuota",
	}
}
