// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This request will send changes to your specified restore testing plan.
// RestoreTestingPlanName cannot be updated after it is created.
//
// RecoveryPointSelection can contain:
//
//   - Algorithm
//
//   - ExcludeVaults
//
//   - IncludeVaults
//
//   - RecoveryPointTypes
//
//   - SelectionWindowDays
func (c *Client) UpdateRestoreTestingPlan(ctx context.Context, params *UpdateRestoreTestingPlanInput, optFns ...func(*Options)) (*UpdateRestoreTestingPlanOutput, error) {
	if params == nil {
		params = &UpdateRestoreTestingPlanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRestoreTestingPlan", params, optFns, c.addOperationUpdateRestoreTestingPlanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRestoreTestingPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRestoreTestingPlanInput struct {

	// Specifies the body of a restore testing plan.
	//
	// This member is required.
	RestoreTestingPlan *types.RestoreTestingPlanForUpdate

	// This is the restore testing plan name you wish to update.
	//
	// This member is required.
	RestoreTestingPlanName *string

	noSmithyDocumentSerde
}

type UpdateRestoreTestingPlanOutput struct {

	// This is the time the resource testing plan was created.
	//
	// This member is required.
	CreationTime *time.Time

	// Unique ARN (Amazon Resource Name) of the restore testing plan.
	//
	// This member is required.
	RestoreTestingPlanArn *string

	// The name cannot be changed after creation. The name consists of only
	// alphanumeric characters and underscores. Maximum length is 50.
	//
	// This member is required.
	RestoreTestingPlanName *string

	// This is the time the update completed for the restore testing plan.
	//
	// This member is required.
	UpdateTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRestoreTestingPlanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRestoreTestingPlan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRestoreTestingPlan{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateRestoreTestingPlan"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateRestoreTestingPlanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRestoreTestingPlan(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateRestoreTestingPlan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateRestoreTestingPlan",
	}
}
