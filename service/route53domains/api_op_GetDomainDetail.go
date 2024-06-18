// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This operation returns detailed information about a specified domain that is
// associated with the current Amazon Web Services account. Contact information for
// the domain is also returned as part of the output.
func (c *Client) GetDomainDetail(ctx context.Context, params *GetDomainDetailInput, optFns ...func(*Options)) (*GetDomainDetailOutput, error) {
	if params == nil {
		params = &GetDomainDetailInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDomainDetail", params, optFns, c.addOperationGetDomainDetailMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDomainDetailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The GetDomainDetail request includes the following element.
type GetDomainDetailInput struct {

	// The name of the domain that you want to get detailed information about.
	//
	// This member is required.
	DomainName *string

	noSmithyDocumentSerde
}

// The GetDomainDetail response includes the following elements.
type GetDomainDetailOutput struct {

	// Email address to contact to report incorrect contact information for a domain,
	// to report that the domain is being used to send spam, to report that someone is
	// cybersquatting on a domain name, or report some other type of abuse.
	AbuseContactEmail *string

	// Phone number for reporting abuse.
	AbuseContactPhone *string

	// Provides details about the domain administrative contact.
	AdminContact *types.ContactDetail

	// Specifies whether contact information is concealed from WHOIS queries. If the
	// value is true , WHOIS ("who is") queries return contact information either for
	// Amazon Registrar or for our registrar associate, Gandi. If the value is false ,
	// WHOIS queries return the information that you entered for the admin contact.
	AdminPrivacy *bool

	// Specifies whether the domain registration is set to renew automatically.
	AutoRenew *bool

	// Provides details about the domain billing contact.
	BillingContact *types.ContactDetail

	// Specifies whether contact information is concealed from WHOIS queries. If the
	// value is true , WHOIS ("who is") queries return contact information either for
	// Amazon Registrar or for our registrar associate, Gandi. If the value is false ,
	// WHOIS queries return the information that you entered for the billing contact.
	BillingPrivacy *bool

	// The date when the domain was created as found in the response to a WHOIS query.
	// The date and time is in Unix time format and Coordinated Universal time (UTC).
	CreationDate *time.Time

	// Deprecated.
	DnsSec *string

	// A complex type that contains information about the DNSSEC configuration.
	DnssecKeys []types.DnssecKey

	// The name of a domain.
	DomainName *string

	// The date when the registration for the domain is set to expire. The date and
	// time is in Unix time format and Coordinated Universal time (UTC).
	ExpirationDate *time.Time

	// The name servers of the domain.
	Nameservers []types.Nameserver

	// Provides details about the domain registrant.
	RegistrantContact *types.ContactDetail

	// Specifies whether contact information is concealed from WHOIS queries. If the
	// value is true , WHOIS ("who is") queries return contact information either for
	// Amazon Registrar or for our registrar associate, Gandi. If the value is false ,
	// WHOIS queries return the information that you entered for the registrant contact
	// (domain owner).
	RegistrantPrivacy *bool

	// Name of the registrar of the domain as identified in the registry.
	RegistrarName *string

	// Web address of the registrar.
	RegistrarUrl *string

	// Reserved for future use.
	RegistryDomainId *string

	// Reseller of the domain. Domains registered or transferred using Route 53
	// domains will have "Amazon" as the reseller.
	Reseller *string

	// An array of domain name status codes, also known as Extensible Provisioning
	// Protocol (EPP) status codes.
	//
	// ICANN, the organization that maintains a central database of domain names, has
	// developed a set of domain name status codes that tell you the status of a
	// variety of operations on a domain name, for example, registering a domain name,
	// transferring a domain name to another registrar, renewing the registration for a
	// domain name, and so on. All registrars use this same set of status codes.
	//
	// For a current list of domain name status codes and an explanation of what each
	// code means, go to the [ICANN website]and search for epp status codes . (Search on the ICANN
	// website; web searches sometimes return an old version of the document.)
	//
	// [ICANN website]: https://www.icann.org/
	StatusList []string

	// Provides details about the domain technical contact.
	TechContact *types.ContactDetail

	// Specifies whether contact information is concealed from WHOIS queries. If the
	// value is true , WHOIS ("who is") queries return contact information either for
	// Amazon Registrar or for our registrar associate, Gandi. If the value is false ,
	// WHOIS queries return the information that you entered for the technical contact.
	TechPrivacy *bool

	// The last updated date of the domain as found in the response to a WHOIS query.
	// The date and time is in Unix time format and Coordinated Universal time (UTC).
	UpdatedDate *time.Time

	// The fully qualified name of the WHOIS server that can answer the WHOIS query
	// for the domain.
	WhoIsServer *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDomainDetailMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDomainDetail{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDomainDetail{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDomainDetail"); err != nil {
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
	if err = addOpGetDomainDetailValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDomainDetail(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDomainDetail(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDomainDetail",
	}
}
