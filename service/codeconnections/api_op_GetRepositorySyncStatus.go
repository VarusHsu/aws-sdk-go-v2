// Code generated by smithy-go-codegen DO NOT EDIT.

package codeconnections

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codeconnections/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns details about the sync status for a repository. A repository sync uses
// Git sync to push and pull changes from your remote repository.
func (c *Client) GetRepositorySyncStatus(ctx context.Context, params *GetRepositorySyncStatusInput, optFns ...func(*Options)) (*GetRepositorySyncStatusOutput, error) {
	if params == nil {
		params = &GetRepositorySyncStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRepositorySyncStatus", params, optFns, c.addOperationGetRepositorySyncStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRepositorySyncStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRepositorySyncStatusInput struct {

	// The branch of the repository link for the requested repository sync status.
	//
	// This member is required.
	Branch *string

	// The repository link ID for the requested repository sync status.
	//
	// This member is required.
	RepositoryLinkId *string

	// The sync type of the requested sync status.
	//
	// This member is required.
	SyncType types.SyncConfigurationType

	noSmithyDocumentSerde
}

type GetRepositorySyncStatusOutput struct {

	// The status of the latest sync returned for a specified repository and branch.
	//
	// This member is required.
	LatestSync *types.RepositorySyncAttempt

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRepositorySyncStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetRepositorySyncStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetRepositorySyncStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetRepositorySyncStatus"); err != nil {
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
	if err = addOpGetRepositorySyncStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRepositorySyncStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRepositorySyncStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetRepositorySyncStatus",
	}
}
