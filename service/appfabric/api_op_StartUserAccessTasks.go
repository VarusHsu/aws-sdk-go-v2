// Code generated by smithy-go-codegen DO NOT EDIT.

package appfabric

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/appfabric/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts the tasks to search user access status for a specific email address.
//
// The tasks are stopped when the user access status data is found. The tasks are
// terminated when the API calls to the application time out.
func (c *Client) StartUserAccessTasks(ctx context.Context, params *StartUserAccessTasksInput, optFns ...func(*Options)) (*StartUserAccessTasksOutput, error) {
	if params == nil {
		params = &StartUserAccessTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartUserAccessTasks", params, optFns, c.addOperationStartUserAccessTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartUserAccessTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartUserAccessTasksInput struct {

	// The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app
	// bundle to use for the request.
	//
	// This member is required.
	AppBundleIdentifier *string

	// The email address of the target user.
	//
	// This member is required.
	Email *string

	noSmithyDocumentSerde
}

type StartUserAccessTasksOutput struct {

	// Contains a list of user access task information.
	UserAccessTasksList []types.UserAccessTaskItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartUserAccessTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartUserAccessTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartUserAccessTasks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartUserAccessTasks"); err != nil {
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
	if err = addOpStartUserAccessTasksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartUserAccessTasks(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartUserAccessTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartUserAccessTasks",
	}
}
