// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the account preferences settings for the Amazon Web Services account
// associated with the user making the request, in the current Amazon Web Services
// Region.
func (c *Client) DescribeAccountPreferences(ctx context.Context, params *DescribeAccountPreferencesInput, optFns ...func(*Options)) (*DescribeAccountPreferencesOutput, error) {
	if params == nil {
		params = &DescribeAccountPreferencesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAccountPreferences", params, optFns, c.addOperationDescribeAccountPreferencesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAccountPreferencesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccountPreferencesInput struct {

	// (Optional) When retrieving account preferences, you can optionally specify the
	// MaxItems parameter to limit the number of objects returned in a response. The
	// default value is 100.
	MaxResults *int32

	// (Optional) You can use NextToken in a subsequent request to fetch the next page
	// of Amazon Web Services account preferences if the response payload was
	// paginated.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeAccountPreferencesOutput struct {

	// Present if there are more records than returned in the response. You can use
	// the NextToken in the subsequent request to fetch the additional descriptions.
	NextToken *string

	// Describes the resource ID preference setting for the Amazon Web Services
	// account associated with the user making the request, in the current Amazon Web
	// Services Region.
	ResourceIdPreference *types.ResourceIdPreference

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAccountPreferencesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAccountPreferences{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAccountPreferences{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeAccountPreferences"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccountPreferences(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAccountPreferences(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeAccountPreferences",
	}
}
