// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Synchronizes a job from the source control repository. This operation takes the
// job artifacts that are located in the remote repository and updates the Glue
// internal stores with these artifacts.
//
// This API supports optional parameters which take in the repository information.
func (c *Client) UpdateJobFromSourceControl(ctx context.Context, params *UpdateJobFromSourceControlInput, optFns ...func(*Options)) (*UpdateJobFromSourceControlOutput, error) {
	if params == nil {
		params = &UpdateJobFromSourceControlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateJobFromSourceControl", params, optFns, c.addOperationUpdateJobFromSourceControlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateJobFromSourceControlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateJobFromSourceControlInput struct {

	// The type of authentication, which can be an authentication token stored in
	// Amazon Web Services Secrets Manager, or a personal access token.
	AuthStrategy types.SourceControlAuthStrategy

	// The value of the authorization token.
	AuthToken *string

	// An optional branch in the remote repository.
	BranchName *string

	// A commit ID for a commit in the remote repository.
	CommitId *string

	// An optional folder in the remote repository.
	Folder *string

	// The name of the Glue job to be synchronized to or from the remote repository.
	JobName *string

	//  The provider for the remote repository. Possible values: GITHUB,
	// AWS_CODE_COMMIT, GITLAB, BITBUCKET.
	Provider types.SourceControlProvider

	// The name of the remote repository that contains the job artifacts. For
	// BitBucket providers, RepositoryName should include WorkspaceName . Use the
	// format / .
	RepositoryName *string

	// The owner of the remote repository that contains the job artifacts.
	RepositoryOwner *string

	noSmithyDocumentSerde
}

type UpdateJobFromSourceControlOutput struct {

	// The name of the Glue job.
	JobName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateJobFromSourceControlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateJobFromSourceControl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateJobFromSourceControl{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateJobFromSourceControl"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateJobFromSourceControl(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateJobFromSourceControl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateJobFromSourceControl",
	}
}
