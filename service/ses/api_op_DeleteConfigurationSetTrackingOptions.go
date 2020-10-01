// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an association between a configuration set and a custom domain for open
// and click event tracking. By default, images and links used for tracking open
// and click events are hosted on domains operated by Amazon SES. You can configure
// a subdomain of your own to handle these events. For information about using
// custom domains, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/configure-custom-open-click-domains.html).
// Deleting this kind of association will result in emails sent using the specified
// configuration set to capture open and click events using the standard, Amazon
// SES-operated domains.
func (c *Client) DeleteConfigurationSetTrackingOptions(ctx context.Context, params *DeleteConfigurationSetTrackingOptionsInput, optFns ...func(*Options)) (*DeleteConfigurationSetTrackingOptionsOutput, error) {
	stack := middleware.NewStack("DeleteConfigurationSetTrackingOptions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteConfigurationSetTrackingOptionsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeleteConfigurationSetTrackingOptionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteConfigurationSetTrackingOptions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "DeleteConfigurationSetTrackingOptions",
			Err:           err,
		}
	}
	out := result.(*DeleteConfigurationSetTrackingOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to delete open and click tracking options in a
// configuration set.
type DeleteConfigurationSetTrackingOptionsInput struct {

	// The name of the configuration set from which you want to delete the tracking
	// options.
	//
	// This member is required.
	ConfigurationSetName *string
}

// An empty element returned on a successful request.
type DeleteConfigurationSetTrackingOptionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteConfigurationSetTrackingOptionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteConfigurationSetTrackingOptions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteConfigurationSetTrackingOptions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteConfigurationSetTrackingOptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "DeleteConfigurationSetTrackingOptions",
	}
}
