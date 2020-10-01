// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a hosted zone.  <p>If the hosted zone was created by another service,
// such as AWS Cloud Map, see <a
// href="https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DeleteHostedZone.html#delete-public-hosted-zone-created-by-another-service">Deleting
// Public Hosted Zones That Were Created by Another Service</a> in the <i>Amazon
// Route 53 Developer Guide</i> for information about how to delete it. (The
// process is the same for public and private hosted zones that were created by
// another service.)</p> <p>If you want to keep your domain registration but you
// want to stop routing internet traffic to your website or web application, we
// recommend that you delete resource record sets in the hosted zone instead of
// deleting the hosted zone.</p> <important> <p>If you delete a hosted zone, you
// can't undelete it. You must create a new hosted zone and update the name servers
// for your domain registration, which can require up to 48 hours to take effect.
// (If you delegated responsibility for a subdomain to a hosted zone and you delete
// the child hosted zone, you must update the name servers in the parent hosted
// zone.) In addition, if you delete a hosted zone, someone could hijack the domain
// and route traffic to their own resources using your domain name.</p>
// </important> <p>If you want to avoid the monthly charge for the hosted zone, you
// can transfer DNS service for the domain to a free DNS service. When you transfer
// DNS service, you have to update the name servers for the domain registration. If
// the domain is registered with Route 53, see <a
// href="https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_UpdateDomainNameservers.html">UpdateDomainNameservers</a>
// for information about how to replace Route 53 name servers with name servers for
// the new DNS service. If the domain is registered with another registrar, use the
// method provided by the registrar to update name servers for the domain
// registration. For more information, perform an internet search on "free DNS
// service."</p> <p>You can delete a hosted zone only if it contains only the
// default SOA record and NS resource record sets. If the hosted zone contains
// other resource record sets, you must delete them before you can delete the
// hosted zone. If you try to delete a hosted zone that contains other resource
// record sets, the request fails, and Route 53 returns a
// <code>HostedZoneNotEmpty</code> error. For information about deleting records
// from your hosted zone, see <a
// href="https://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets.html">ChangeResourceRecordSets</a>.</p>
// <p>To verify that the hosted zone has been deleted, do one of the following:</p>
// <ul> <li> <p>Use the <code>GetHostedZone</code> action to request information
// about the hosted zone.</p> </li> <li> <p>Use the <code>ListHostedZones</code>
// action to get a list of the hosted zones associated with the current AWS
// account.</p> </li> </ul>
func (c *Client) DeleteHostedZone(ctx context.Context, params *DeleteHostedZoneInput, optFns ...func(*Options)) (*DeleteHostedZoneOutput, error) {
	stack := middleware.NewStack("DeleteHostedZone", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpDeleteHostedZoneMiddlewares(stack)
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
	addOpDeleteHostedZoneValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteHostedZone(options.Region), middleware.Before)
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
			OperationName: "DeleteHostedZone",
			Err:           err,
		}
	}
	out := result.(*DeleteHostedZoneOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to delete a hosted zone.
type DeleteHostedZoneInput struct {

	// The ID of the hosted zone you want to delete.
	//
	// This member is required.
	Id *string
}

// A complex type that contains the response to a DeleteHostedZone request.
type DeleteHostedZoneOutput struct {

	// A complex type that contains the ID, the status, and the date and time of a
	// request to delete a hosted zone.
	//
	// This member is required.
	ChangeInfo *types.ChangeInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpDeleteHostedZoneMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpDeleteHostedZone{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteHostedZone{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteHostedZone(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "DeleteHostedZone",
	}
}
