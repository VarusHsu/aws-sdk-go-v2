// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets detailed information about the specified license configuration.
func (c *Client) GetLicenseConfiguration(ctx context.Context, params *GetLicenseConfigurationInput, optFns ...func(*Options)) (*GetLicenseConfigurationOutput, error) {
	if params == nil {
		params = &GetLicenseConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLicenseConfiguration", params, optFns, c.addOperationGetLicenseConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLicenseConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLicenseConfigurationInput struct {

	// Amazon Resource Name (ARN) of the license configuration.
	//
	// This member is required.
	LicenseConfigurationArn *string

	noSmithyDocumentSerde
}

type GetLicenseConfigurationOutput struct {

	// Automated discovery information.
	AutomatedDiscoveryInformation *types.AutomatedDiscoveryInformation

	// Summaries of the licenses consumed by resources.
	ConsumedLicenseSummaryList []types.ConsumedLicenseSummary

	// Number of licenses assigned to resources.
	ConsumedLicenses *int64

	// Description of the license configuration.
	Description *string

	// When true, disassociates a resource when software is uninstalled.
	DisassociateWhenNotFound *bool

	// Amazon Resource Name (ARN) of the license configuration.
	LicenseConfigurationArn *string

	// Unique ID for the license configuration.
	LicenseConfigurationId *string

	// Number of available licenses.
	LicenseCount *int64

	// Sets the number of available licenses as a hard limit.
	LicenseCountHardLimit *bool

	// Dimension for which the licenses are counted.
	LicenseCountingType types.LicenseCountingType

	// License rules.
	LicenseRules []string

	// Summaries of the managed resources.
	ManagedResourceSummaryList []types.ManagedResourceSummary

	// Name of the license configuration.
	Name *string

	// Account ID of the owner of the license configuration.
	OwnerAccountId *string

	// Product information.
	ProductInformationList []types.ProductInformation

	// License configuration status.
	Status *string

	// Tags for the license configuration.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLicenseConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetLicenseConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetLicenseConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetLicenseConfiguration"); err != nil {
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
	if err = addOpGetLicenseConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLicenseConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLicenseConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetLicenseConfiguration",
	}
}
