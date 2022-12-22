// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanagerlinuxsubscriptions

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/licensemanagerlinuxsubscriptions/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the Linux subscriptions service settings.
func (c *Client) GetServiceSettings(ctx context.Context, params *GetServiceSettingsInput, optFns ...func(*Options)) (*GetServiceSettingsOutput, error) {
	if params == nil {
		params = &GetServiceSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetServiceSettings", params, optFns, c.addOperationGetServiceSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetServiceSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServiceSettingsInput struct {
	noSmithyDocumentSerde
}

type GetServiceSettingsOutput struct {

	// The Region in which License Manager displays the aggregated data for Linux
	// subscriptions.
	HomeRegions []string

	// Lists if discovery has been enabled for Linux subscriptions.
	LinuxSubscriptionsDiscovery types.LinuxSubscriptionsDiscovery

	// Lists the settings defined for Linux subscriptions discovery. The settings
	// include if Organizations integration has been enabled, and which Regions data
	// will be aggregated from.
	LinuxSubscriptionsDiscoverySettings *types.LinuxSubscriptionsDiscoverySettings

	// Indicates the status of Linux subscriptions settings being applied.
	Status types.Status

	// A message which details the Linux subscriptions service settings current status.
	StatusMessage map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetServiceSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetServiceSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetServiceSettings{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetServiceSettings(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opGetServiceSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager-linux-subscriptions",
		OperationName: "GetServiceSettings",
	}
}