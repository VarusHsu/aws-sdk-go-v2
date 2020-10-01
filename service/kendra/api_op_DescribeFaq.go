// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Gets information about an FAQ list.
func (c *Client) DescribeFaq(ctx context.Context, params *DescribeFaqInput, optFns ...func(*Options)) (*DescribeFaqOutput, error) {
	stack := middleware.NewStack("DescribeFaq", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeFaqMiddlewares(stack)
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
	addOpDescribeFaqValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFaq(options.Region), middleware.Before)
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
			OperationName: "DescribeFaq",
			Err:           err,
		}
	}
	out := result.(*DescribeFaqOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFaqInput struct {

	// The unique identifier of the FAQ.
	//
	// This member is required.
	Id *string

	// The identifier of the index that contains the FAQ.
	//
	// This member is required.
	IndexId *string
}

type DescribeFaqOutput struct {

	// The identifier of the FAQ.
	Id *string

	// The description of the FAQ that you provided when it was created.
	Description *string

	// Information required to find a specific file in an Amazon S3 bucket.
	S3Path *types.S3Path

	// The date and time that the FAQ was created.
	CreatedAt *time.Time

	// The identifier of the index that contains the FAQ.
	IndexId *string

	// The status of the FAQ. It is ready to use when the status is ACTIVE.
	Status types.FaqStatus

	// The name that you gave the FAQ when it was created.
	Name *string

	// The Amazon Resource Name (ARN) of the role that provides access to the S3 bucket
	// containing the input files for the FAQ.
	RoleArn *string

	// If the Status field is FAILED, the ErrorMessage field contains the reason why
	// the FAQ failed.
	ErrorMessage *string

	// The date and time that the FAQ was last updated.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeFaqMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFaq{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFaq{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeFaq(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "DescribeFaq",
	}
}
