// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers a new migration task which represents a server, database, etc., being
// migrated to AWS by a migration tool. This API is a prerequisite to calling the
// NotifyMigrationTaskState API as the migration tool must first register the
// migration task with Migration Hub.
func (c *Client) ImportMigrationTask(ctx context.Context, params *ImportMigrationTaskInput, optFns ...func(*Options)) (*ImportMigrationTaskOutput, error) {
	stack := middleware.NewStack("ImportMigrationTask", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpImportMigrationTaskMiddlewares(stack)
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
	addOpImportMigrationTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opImportMigrationTask(options.Region), middleware.Before)
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
			OperationName: "ImportMigrationTask",
			Err:           err,
		}
	}
	out := result.(*ImportMigrationTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportMigrationTaskInput struct {

	// Unique identifier that references the migration task. Do not store personal data
	// in this field.
	//
	// This member is required.
	MigrationTaskName *string

	// The name of the ProgressUpdateStream. >
	//
	// This member is required.
	ProgressUpdateStream *string

	// Optional boolean flag to indicate whether any effect should take place. Used to
	// test if the caller has permission to make the call.
	DryRun *bool
}

type ImportMigrationTaskOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpImportMigrationTaskMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpImportMigrationTask{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpImportMigrationTask{}, middleware.After)
}

func newServiceMetadataMiddleware_opImportMigrationTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgh",
		OperationName: "ImportMigrationTask",
	}
}
