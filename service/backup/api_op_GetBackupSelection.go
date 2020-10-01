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

// Returns selection metadata and a document in JSON format that specifies a list
// of resources that are associated with a backup plan.
func (c *Client) GetBackupSelection(ctx context.Context, params *GetBackupSelectionInput, optFns ...func(*Options)) (*GetBackupSelectionOutput, error) {
	stack := middleware.NewStack("GetBackupSelection", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetBackupSelectionMiddlewares(stack)
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
	addOpGetBackupSelectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBackupSelection(options.Region), middleware.Before)
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
			OperationName: "GetBackupSelection",
			Err:           err,
		}
	}
	out := result.(*GetBackupSelectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBackupSelectionInput struct {

	// Uniquely identifies the body of a request to assign a set of resources to a
	// backup plan.
	//
	// This member is required.
	SelectionId *string

	// Uniquely identifies a backup plan.
	//
	// This member is required.
	BackupPlanId *string
}

type GetBackupSelectionOutput struct {

	// The date and time a backup selection is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time

	// Uniquely identifies the body of a request to assign a set of resources to a
	// backup plan.
	SelectionId *string

	// A unique string that identifies the request and allows failed requests to be
	// retried without the risk of executing the operation twice.
	CreatorRequestId *string

	// Uniquely identifies a backup plan.
	BackupPlanId *string

	// Specifies the body of a request to assign a set of resources to a backup plan.
	BackupSelection *types.BackupSelection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetBackupSelectionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetBackupSelection{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBackupSelection{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetBackupSelection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "GetBackupSelection",
	}
}
