// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a branch for an Amplify app.
func (c *Client) UpdateBranch(ctx context.Context, params *UpdateBranchInput, optFns ...func(*Options)) (*UpdateBranchOutput, error) {
	if params == nil {
		params = &UpdateBranchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateBranch", params, optFns, c.addOperationUpdateBranchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateBranchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the update branch request.
type UpdateBranchInput struct {

	//  The unique ID for an Amplify app.
	//
	// This member is required.
	AppId *string

	// The name of the branch.
	//
	// This member is required.
	BranchName *string

	// The backend for a Branch of an Amplify app. Use for a backend created from an
	// CloudFormation stack.
	//
	// This field is available to Amplify Gen 2 apps only. When you deploy an
	// application with Amplify Gen 2, you provision the app's backend infrastructure
	// using Typescript code.
	Backend *types.Backend

	// The Amazon Resource Name (ARN) for a backend environment that is part of a Gen
	// 1 Amplify app.
	//
	// This field is available to Amplify Gen 1 apps only where the backend is created
	// using Amplify Studio or the Amplify command line interface (CLI).
	BackendEnvironmentArn *string

	//  The basic authorization credentials for the branch. You must base64-encode the
	// authorization credentials and provide them in the format user:password .
	BasicAuthCredentials *string

	//  The build specification (build spec) for the branch.
	BuildSpec *string

	//  The description for the branch.
	Description *string

	//  The display name for a branch. This is used as the default domain prefix.
	DisplayName *string

	//  Enables auto building for the branch.
	EnableAutoBuild *bool

	//  Enables basic authorization for the branch.
	EnableBasicAuth *bool

	//  Enables notifications for the branch.
	EnableNotification *bool

	// Enables performance mode for the branch.
	//
	// Performance mode optimizes for faster hosting performance by keeping content
	// cached at the edge for a longer interval. When performance mode is enabled,
	// hosting configuration or code changes can take up to 10 minutes to roll out.
	EnablePerformanceMode *bool

	//  Enables pull request previews for this branch.
	EnablePullRequestPreview *bool

	//  The environment variables for the branch.
	EnvironmentVariables map[string]string

	//  The framework for the branch.
	Framework *string

	//  The Amplify environment name for the pull request.
	PullRequestEnvironmentName *string

	//  Describes the current stage for the branch.
	Stage types.Stage

	//  The content Time to Live (TTL) for the website in seconds.
	Ttl *string

	noSmithyDocumentSerde
}

// The result structure for the update branch request.
type UpdateBranchOutput struct {

	//  The branch for an Amplify app, which maps to a third-party repository branch.
	//
	// This member is required.
	Branch *types.Branch

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateBranchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateBranch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateBranch{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateBranch"); err != nil {
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
	if err = addOpUpdateBranchValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBranch(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateBranch(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateBranch",
	}
}
