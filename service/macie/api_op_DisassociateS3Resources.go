// Code generated by smithy-go-codegen DO NOT EDIT.

package macie

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes specified S3 resources from being monitored by Amazon Macie Classic. If
// memberAccountId isn't specified, the action removes specified S3 resources from
// Macie Classic for the current master account. If memberAccountId is specified,
// the action removes specified S3 resources from Macie Classic for the specified
// member account.
func (c *Client) DisassociateS3Resources(ctx context.Context, params *DisassociateS3ResourcesInput, optFns ...func(*Options)) (*DisassociateS3ResourcesOutput, error) {
	stack := middleware.NewStack("DisassociateS3Resources", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDisassociateS3ResourcesMiddlewares(stack)
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
	addOpDisassociateS3ResourcesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateS3Resources(options.Region), middleware.Before)
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
			OperationName: "DisassociateS3Resources",
			Err:           err,
		}
	}
	out := result.(*DisassociateS3ResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateS3ResourcesInput struct {

	// The ID of the Amazon Macie Classic member account whose resources you want to
	// remove from being monitored by Amazon Macie Classic.
	MemberAccountId *string

	// The S3 resources (buckets or prefixes) that you want to remove from being
	// monitored and classified by Amazon Macie Classic.
	//
	// This member is required.
	AssociatedS3Resources []*types.S3Resource
}

type DisassociateS3ResourcesOutput struct {

	// S3 resources that couldn't be removed from being monitored and classified by
	// Amazon Macie Classic. An error code and an error message are provided for each
	// failed item.
	FailedS3Resources []*types.FailedS3Resource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDisassociateS3ResourcesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateS3Resources{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateS3Resources{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisassociateS3Resources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie",
		OperationName: "DisassociateS3Resources",
	}
}
