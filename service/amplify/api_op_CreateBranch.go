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

// Creates a new branch for an Amplify app.
func (c *Client) CreateBranch(ctx context.Context, params *CreateBranchInput, optFns ...func(*Options)) (*CreateBranchOutput, error) {
	if params == nil {
		params = &CreateBranchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBranch", params, optFns, c.addOperationCreateBranchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBranchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the create branch request.
type CreateBranchInput struct {

	//  The unique ID for an Amplify app.
	//
	// This member is required.
	AppId *string

	// The name for the branch.
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

	// The description for the branch.
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

	// Describes the current stage for the branch.
	Stage types.Stage

	//  The tag for the branch.
	Tags map[string]string

	//  The content Time To Live (TTL) for the website in seconds.
	Ttl *string

	noSmithyDocumentSerde
}

// The result structure for create branch request.
type CreateBranchOutput struct {

	//  Describes the branch for an Amplify app, which maps to a third-party
	// repository branch.
	//
	// This member is required.
	Branch *types.Branch

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBranchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateBranch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateBranch{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateBranch"); err != nil {
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
	if err = addOpCreateBranchValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBranch(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateBranch(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateBranch",
	}
}
