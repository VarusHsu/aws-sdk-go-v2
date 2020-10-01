// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the user import job.
func (c *Client) DescribeUserImportJob(ctx context.Context, params *DescribeUserImportJobInput, optFns ...func(*Options)) (*DescribeUserImportJobOutput, error) {
	stack := middleware.NewStack("DescribeUserImportJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeUserImportJobMiddlewares(stack)
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
	addOpDescribeUserImportJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUserImportJob(options.Region), middleware.Before)
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
			OperationName: "DescribeUserImportJob",
			Err:           err,
		}
	}
	out := result.(*DescribeUserImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to describe the user import job.
type DescribeUserImportJobInput struct {

	// The job ID for the user import job.
	//
	// This member is required.
	JobId *string

	// The user pool ID for the user pool that the users are being imported into.
	//
	// This member is required.
	UserPoolId *string
}

// Represents the response from the server to the request to describe the user
// import job.
type DescribeUserImportJobOutput struct {

	// The job object that represents the user import job.
	UserImportJob *types.UserImportJobType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeUserImportJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeUserImportJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeUserImportJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeUserImportJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "DescribeUserImportJob",
	}
}
