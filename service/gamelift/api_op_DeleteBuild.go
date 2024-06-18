// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a build. This operation permanently deletes the build resource and any
// uploaded build files. Deleting a build does not affect the status of any active
// fleets using the build, but you can no longer create new fleets with the deleted
// build.
//
// To delete a build, specify the build ID.
//
// # Learn more
//
// [Upload a Custom Server Build]
//
// [All APIs by task]
//
// [Upload a Custom Server Build]: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-build-intro.html
// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
func (c *Client) DeleteBuild(ctx context.Context, params *DeleteBuildInput, optFns ...func(*Options)) (*DeleteBuildOutput, error) {
	if params == nil {
		params = &DeleteBuildInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBuild", params, optFns, c.addOperationDeleteBuildMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBuildOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBuildInput struct {

	// A unique identifier for the build to delete. You can use either the build ID or
	// ARN value.
	//
	// This member is required.
	BuildId *string

	noSmithyDocumentSerde
}

type DeleteBuildOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteBuildMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteBuild{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteBuild{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteBuild"); err != nil {
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
	if err = addOpDeleteBuildValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBuild(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteBuild(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteBuild",
	}
}
