// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation initiates a job of the specified type, which can be a select, an
// archival retrieval, or a vault retrieval. For more information about using this
// operation, see the documentation for the underlying REST API Initiate a Job
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-initiate-job-post.html).
func (c *Client) InitiateJob(ctx context.Context, params *InitiateJobInput, optFns ...func(*Options)) (*InitiateJobOutput, error) {
	stack := middleware.NewStack("InitiateJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpInitiateJobMiddlewares(stack)
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
	addOpInitiateJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opInitiateJob(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	glaciercust.AddTreeHashMiddleware(stack)
	glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion)
	glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID)

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
			OperationName: "InitiateJob",
			Err:           err,
		}
	}
	out := result.(*InitiateJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options for initiating an Amazon S3 Glacier job.
type InitiateJobInput struct {

	// The name of the vault.
	//
	// This member is required.
	VaultName *string

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen), in
	// which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// Provides options for specifying job information.
	JobParameters *types.JobParameters
}

// Contains the Amazon S3 Glacier response to your request.
type InitiateJobOutput struct {

	// The relative URI path of the job.
	Location *string

	// The ID of the job.
	JobId *string

	// The path to the location of where the select results are stored.
	JobOutputPath *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpInitiateJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpInitiateJob{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpInitiateJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opInitiateJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "InitiateJob",
	}
}
