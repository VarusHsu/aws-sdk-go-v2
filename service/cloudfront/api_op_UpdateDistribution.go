// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the configuration for a web distribution. When you update a
// distribution, there are more required fields than when you create a
// distribution. When you update your distribution by using this API action, follow
// the steps here to get the current configuration and then make your updates, to
// make sure that you include all of the required fields. To view a summary, see
// Required Fields for Create Distribution and Update Distribution
// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-overview-required-fields.html)
// in the Amazon CloudFront Developer Guide. The update process includes getting
// the current distribution configuration, updating the XML document that is
// returned to make your changes, and then submitting an UpdateDistribution request
// to make the updates. For information about updating a distribution using the
// CloudFront console instead, see Creating a Distribution
// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-creating-console.html)
// in the Amazon CloudFront Developer Guide.  <p> <b>To update a web distribution
// using the CloudFront API</b> </p> <ol> <li> <p>Submit a <a
// href="https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_GetDistributionConfig.html">GetDistributionConfig</a>
// request to get the current configuration and an <code>Etag</code> header for the
// distribution.</p> <note> <p>If you update the distribution again, you must get a
// new <code>Etag</code> header.</p> </note> </li> <li> <p>Update the XML document
// that was returned in the response to your <code>GetDistributionConfig</code>
// request to include your changes. </p> <important> <p>When you edit the XML file,
// be aware of the following:</p> <ul> <li> <p>You must strip out the ETag
// parameter that is returned.</p> </li> <li> <p>Additional fields are required
// when you update a distribution. There may be fields included in the XML file for
// features that you haven't configured for your distribution. This is expected and
// required to successfully update the distribution.</p> </li> <li> <p>You can't
// change the value of <code>CallerReference</code>. If you try to change this
// value, CloudFront returns an <code>IllegalUpdate</code> error. </p> </li> <li>
// <p>The new configuration replaces the existing configuration; the values that
// you specify in an <code>UpdateDistribution</code> request are not merged into
// your existing configuration. When you add, delete, or replace values in an
// element that allows multiple values (for example, <code>CNAME</code>), you must
// specify all of the values that you want to appear in the updated distribution.
// In addition, you must update the corresponding <code>Quantity</code>
// element.</p> </li> </ul> </important> </li> <li> <p>Submit an
// <code>UpdateDistribution</code> request to update the configuration for your
// distribution:</p> <ul> <li> <p>In the request body, include the XML document
// that you updated in Step 2. The request body must include an XML document with a
// <code>DistributionConfig</code> element.</p> </li> <li> <p>Set the value of the
// HTTP <code>If-Match</code> header to the value of the <code>ETag</code> header
// that CloudFront returned when you submitted the
// <code>GetDistributionConfig</code> request in Step 1.</p> </li> </ul> </li> <li>
// <p>Review the response to the <code>UpdateDistribution</code> request to confirm
// that the configuration was successfully updated.</p> </li> <li> <p>Optional:
// Submit a <a
// href="https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_GetDistribution.html">GetDistribution</a>
// request to confirm that your changes have propagated. When propagation is
// complete, the value of <code>Status</code> is <code>Deployed</code>.</p> </li>
// </ol>
func (c *Client) UpdateDistribution(ctx context.Context, params *UpdateDistributionInput, optFns ...func(*Options)) (*UpdateDistributionOutput, error) {
	stack := middleware.NewStack("UpdateDistribution", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpUpdateDistributionMiddlewares(stack)
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
	addOpUpdateDistributionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDistribution(options.Region), middleware.Before)
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
			OperationName: "UpdateDistribution",
			Err:           err,
		}
	}
	out := result.(*UpdateDistributionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to update a distribution.
type UpdateDistributionInput struct {

	// The distribution's configuration information.
	//
	// This member is required.
	DistributionConfig *types.DistributionConfig

	// The distribution's id.
	//
	// This member is required.
	Id *string

	// The value of the ETag header that you received when retrieving the
	// distribution's configuration. For example: E2QWRUHAPOMQZL.
	IfMatch *string
}

// The returned result of the corresponding request.
type UpdateDistributionOutput struct {

	// The distribution's information.
	Distribution *types.Distribution

	// The current version of the configuration. For example: E2QWRUHAPOMQZL.
	ETag *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpUpdateDistributionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpUpdateDistribution{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpUpdateDistribution{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDistribution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "UpdateDistribution",
	}
}
