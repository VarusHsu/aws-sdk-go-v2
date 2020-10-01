// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops a running synchronization job. You can't stop a scheduled synchronization
// job.
func (c *Client) StopDataSourceSyncJob(ctx context.Context, params *StopDataSourceSyncJobInput, optFns ...func(*Options)) (*StopDataSourceSyncJobOutput, error) {
	stack := middleware.NewStack("StopDataSourceSyncJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStopDataSourceSyncJobMiddlewares(stack)
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
	addOpStopDataSourceSyncJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopDataSourceSyncJob(options.Region), middleware.Before)
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
			OperationName: "StopDataSourceSyncJob",
			Err:           err,
		}
	}
	out := result.(*StopDataSourceSyncJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopDataSourceSyncJobInput struct {

	// The identifier of the data source for which to stop the synchronization jobs.
	//
	// This member is required.
	Id *string

	// The identifier of the index that contains the data source.
	//
	// This member is required.
	IndexId *string
}

type StopDataSourceSyncJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStopDataSourceSyncJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStopDataSourceSyncJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopDataSourceSyncJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opStopDataSourceSyncJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "StopDataSourceSyncJob",
	}
}
