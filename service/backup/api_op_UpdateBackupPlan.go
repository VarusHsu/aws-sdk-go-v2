// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Replaces the body of a saved backup plan identified by its backupPlanId with the
// input document in JSON format. The new version is uniquely identified by a
// VersionId.
func (c *Client) UpdateBackupPlan(ctx context.Context, params *UpdateBackupPlanInput, optFns ...func(*Options)) (*UpdateBackupPlanOutput, error) {
	stack := middleware.NewStack("UpdateBackupPlan", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateBackupPlanMiddlewares(stack)
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
	addOpUpdateBackupPlanValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBackupPlan(options.Region), middleware.Before)
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
			OperationName: "UpdateBackupPlan",
			Err:           err,
		}
	}
	out := result.(*UpdateBackupPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateBackupPlanInput struct {

	// Uniquely identifies a backup plan.
	//
	// This member is required.
	BackupPlanId *string

	// Specifies the body of a backup plan. Includes a BackupPlanName and one or more
	// sets of Rules.
	//
	// This member is required.
	BackupPlan *types.BackupPlanInput
}

type UpdateBackupPlanOutput struct {

	// Uniquely identifies a backup plan.
	BackupPlanId *string

	// An Amazon Resource Name (ARN) that uniquely identifies a backup plan; for
	// example,
	// arn:aws:backup:us-east-1:123456789012:plan:8F81F553-3A74-4A3F-B93D-B3360DC80C50.
	BackupPlanArn *string

	// Unique, randomly generated, Unicode, UTF-8 encoded strings that are at most
	// 1,024 bytes long. Version Ids cannot be edited.
	VersionId *string

	// The date and time a backup plan is updated, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateBackupPlanMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateBackupPlan{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateBackupPlan{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateBackupPlan(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "UpdateBackupPlan",
	}
}
