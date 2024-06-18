// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns an Amazon S3 presigned URL for an update file associated with a
// specified JobId .
func (c *Client) GetSoftwareUpdates(ctx context.Context, params *GetSoftwareUpdatesInput, optFns ...func(*Options)) (*GetSoftwareUpdatesOutput, error) {
	if params == nil {
		params = &GetSoftwareUpdatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSoftwareUpdates", params, optFns, c.addOperationGetSoftwareUpdatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSoftwareUpdatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSoftwareUpdatesInput struct {

	// The ID for a job that you want to get the software update file for, for example
	// JID123e4567-e89b-12d3-a456-426655440000 .
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

type GetSoftwareUpdatesOutput struct {

	// The Amazon S3 presigned URL for the update file associated with the specified
	// JobId value. The software update will be available for 2 days after this request
	// is made. To access an update after the 2 days have passed, you'll have to make
	// another call to GetSoftwareUpdates .
	UpdatesURI *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSoftwareUpdatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSoftwareUpdates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSoftwareUpdates{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetSoftwareUpdates"); err != nil {
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
	if err = addOpGetSoftwareUpdatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSoftwareUpdates(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSoftwareUpdates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetSoftwareUpdates",
	}
}
