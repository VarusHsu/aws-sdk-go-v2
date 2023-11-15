// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation updates the specified domain contact's privacy setting. When
// privacy protection is enabled, your contact information is replaced with contact
// information for the registrar or with the phrase "REDACTED FOR PRIVACY", or "On
// behalf of owner." While some domains may allow different privacy settings per
// contact, we recommend specifying the same privacy setting for all contacts. This
// operation affects only the contact information for the specified contact type
// (administrative, registrant, or technical). If the request succeeds, Amazon
// Route 53 returns an operation ID that you can use with GetOperationDetail (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html)
// to track the progress and completion of the action. If the request doesn't
// complete successfully, the domain registrant will be notified by email. By
// disabling the privacy service via API, you consent to the publication of the
// contact information provided for this domain via the public WHOIS database. You
// certify that you are the registrant of this domain name and have the authority
// to make this decision. You may withdraw your consent at any time by enabling
// privacy protection using either UpdateDomainContactPrivacy or the Route 53
// console. Enabling privacy protection removes the contact information provided
// for this domain from the WHOIS database. For more information on our privacy
// practices, see https://aws.amazon.com/privacy/ (https://aws.amazon.com/privacy/)
// .
func (c *Client) UpdateDomainContactPrivacy(ctx context.Context, params *UpdateDomainContactPrivacyInput, optFns ...func(*Options)) (*UpdateDomainContactPrivacyOutput, error) {
	if params == nil {
		params = &UpdateDomainContactPrivacyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDomainContactPrivacy", params, optFns, c.addOperationUpdateDomainContactPrivacyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDomainContactPrivacyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The UpdateDomainContactPrivacy request includes the following elements.
type UpdateDomainContactPrivacyInput struct {

	// The name of the domain that you want to update the privacy setting for.
	//
	// This member is required.
	DomainName *string

	// Whether you want to conceal contact information from WHOIS queries. If you
	// specify true , WHOIS ("who is") queries return contact information either for
	// Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If you specify false , WHOIS queries
	// return the information that you entered for the admin contact. You must specify
	// the same privacy setting for the administrative, registrant, and technical
	// contacts.
	AdminPrivacy *bool

	// Whether you want to conceal contact information from WHOIS queries. If you
	// specify true , WHOIS ("who is") queries return contact information either for
	// Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If you specify false , WHOIS queries
	// return the information that you entered for the registrant contact (domain
	// owner). You must specify the same privacy setting for the administrative,
	// registrant, and technical contacts.
	RegistrantPrivacy *bool

	// Whether you want to conceal contact information from WHOIS queries. If you
	// specify true , WHOIS ("who is") queries return contact information either for
	// Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If you specify false , WHOIS queries
	// return the information that you entered for the technical contact. You must
	// specify the same privacy setting for the administrative, registrant, and
	// technical contacts.
	TechPrivacy *bool

	noSmithyDocumentSerde
}

// The UpdateDomainContactPrivacy response includes the following element.
type UpdateDomainContactPrivacyOutput struct {

	// Identifier for tracking the progress of the request. To use this ID to query
	// the operation status, use GetOperationDetail.
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDomainContactPrivacyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDomainContactPrivacy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDomainContactPrivacy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateDomainContactPrivacy"); err != nil {
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
	if err = addOpUpdateDomainContactPrivacyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDomainContactPrivacy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDomainContactPrivacy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateDomainContactPrivacy",
	}
}
