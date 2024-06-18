// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the contents of a specified folder in a repository.
func (c *Client) GetFolder(ctx context.Context, params *GetFolderInput, optFns ...func(*Options)) (*GetFolderOutput, error) {
	if params == nil {
		params = &GetFolderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFolder", params, optFns, c.addOperationGetFolderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFolderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFolderInput struct {

	// The fully qualified path to the folder whose contents are returned, including
	// the folder name. For example, /examples is a fully-qualified path to a folder
	// named examples that was created off of the root directory (/) of a repository.
	//
	// This member is required.
	FolderPath *string

	// The name of the repository.
	//
	// This member is required.
	RepositoryName *string

	// A fully qualified reference used to identify a commit that contains the version
	// of the folder's content to return. A fully qualified reference can be a commit
	// ID, branch name, tag, or reference such as HEAD. If no specifier is provided,
	// the folder content is returned as it exists in the HEAD commit.
	CommitSpecifier *string

	noSmithyDocumentSerde
}

type GetFolderOutput struct {

	// The full commit ID used as a reference for the returned version of the folder
	// content.
	//
	// This member is required.
	CommitId *string

	// The fully qualified path of the folder whose contents are returned.
	//
	// This member is required.
	FolderPath *string

	// The list of files in the specified folder, if any.
	Files []types.File

	// The list of folders that exist under the specified folder, if any.
	SubFolders []types.Folder

	// The list of submodules in the specified folder, if any.
	SubModules []types.SubModule

	// The list of symbolic links to other files and folders in the specified folder,
	// if any.
	SymbolicLinks []types.SymbolicLink

	// The full SHA-1 pointer of the tree information for the commit that contains the
	// folder.
	TreeId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFolderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetFolder{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetFolder{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetFolder"); err != nil {
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
	if err = addOpGetFolderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFolder(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetFolder(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetFolder",
	}
}
