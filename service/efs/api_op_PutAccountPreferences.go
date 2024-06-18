// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this operation to set the account preference in the current Amazon Web
// Services Region to use long 17 character (63 bit) or short 8 character (32 bit)
// resource IDs for new EFS file system and mount target resources. All existing
// resource IDs are not affected by any changes you make. You can set the ID
// preference during the opt-in period as EFS transitions to long resource IDs. For
// more information, see [Managing Amazon EFS resource IDs].
//
// Starting in October, 2021, you will receive an error if you try to set the
// account preference to use the short 8 character format resource ID. Contact
// Amazon Web Services support if you receive an error and must use short IDs for
// file system and mount target resources.
//
// [Managing Amazon EFS resource IDs]: https://docs.aws.amazon.com/efs/latest/ug/manage-efs-resource-ids.html
func (c *Client) PutAccountPreferences(ctx context.Context, params *PutAccountPreferencesInput, optFns ...func(*Options)) (*PutAccountPreferencesOutput, error) {
	if params == nil {
		params = &PutAccountPreferencesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAccountPreferences", params, optFns, c.addOperationPutAccountPreferencesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAccountPreferencesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAccountPreferencesInput struct {

	// Specifies the EFS resource ID preference to set for the user's Amazon Web
	// Services account, in the current Amazon Web Services Region, either LONG_ID (17
	// characters), or SHORT_ID (8 characters).
	//
	// Starting in October, 2021, you will receive an error when setting the account
	// preference to SHORT_ID . Contact Amazon Web Services support if you receive an
	// error and must use short IDs for file system and mount target resources.
	//
	// This member is required.
	ResourceIdType types.ResourceIdType

	noSmithyDocumentSerde
}

type PutAccountPreferencesOutput struct {

	// Describes the resource type and its ID preference for the user's Amazon Web
	// Services account, in the current Amazon Web Services Region.
	ResourceIdPreference *types.ResourceIdPreference

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutAccountPreferencesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutAccountPreferences{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutAccountPreferences{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutAccountPreferences"); err != nil {
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
	if err = addOpPutAccountPreferencesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAccountPreferences(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutAccountPreferences(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutAccountPreferences",
	}
}
