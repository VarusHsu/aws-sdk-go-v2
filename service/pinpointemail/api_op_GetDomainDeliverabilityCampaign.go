// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieve all the deliverability data for a specific campaign. This data is
// available for a campaign only if the campaign sent email by using a domain that
// the Deliverability dashboard is enabled for (PutDeliverabilityDashboardOption
// operation).
func (c *Client) GetDomainDeliverabilityCampaign(ctx context.Context, params *GetDomainDeliverabilityCampaignInput, optFns ...func(*Options)) (*GetDomainDeliverabilityCampaignOutput, error) {
	stack := middleware.NewStack("GetDomainDeliverabilityCampaign", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetDomainDeliverabilityCampaignMiddlewares(stack)
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
	addOpGetDomainDeliverabilityCampaignValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDomainDeliverabilityCampaign(options.Region), middleware.Before)
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
			OperationName: "GetDomainDeliverabilityCampaign",
			Err:           err,
		}
	}
	out := result.(*GetDomainDeliverabilityCampaignOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Retrieve all the deliverability data for a specific campaign. This data is
// available for a campaign only if the campaign sent email by using a domain that
// the Deliverability dashboard is enabled for (PutDeliverabilityDashboardOption
// operation).
type GetDomainDeliverabilityCampaignInput struct {

	// The unique identifier for the campaign. Amazon Pinpoint automatically generates
	// and assigns this identifier to a campaign. This value is not the same as the
	// campaign identifier that Amazon Pinpoint assigns to campaigns that you create
	// and manage by using the Amazon Pinpoint API or the Amazon Pinpoint console.
	//
	// This member is required.
	CampaignId *string
}

// An object that contains all the deliverability data for a specific campaign.
// This data is available for a campaign only if the campaign sent email by using a
// domain that the Deliverability dashboard is enabled for
// (PutDeliverabilityDashboardOption operation).
type GetDomainDeliverabilityCampaignOutput struct {

	// An object that contains the deliverability data for the campaign.
	//
	// This member is required.
	DomainDeliverabilityCampaign *types.DomainDeliverabilityCampaign

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetDomainDeliverabilityCampaignMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetDomainDeliverabilityCampaign{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDomainDeliverabilityCampaign{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDomainDeliverabilityCampaign(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetDomainDeliverabilityCampaign",
	}
}
