// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enable or disable collection of reputation metrics for emails that you send
// using a particular configuration set in a specific AWS Region.
func (c *Client) PutConfigurationSetReputationOptions(ctx context.Context, params *PutConfigurationSetReputationOptionsInput, optFns ...func(*Options)) (*PutConfigurationSetReputationOptionsOutput, error) {
	stack := middleware.NewStack("PutConfigurationSetReputationOptions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutConfigurationSetReputationOptionsMiddlewares(stack)
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
	addOpPutConfigurationSetReputationOptionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutConfigurationSetReputationOptions(options.Region), middleware.Before)
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
			OperationName: "PutConfigurationSetReputationOptions",
			Err:           err,
		}
	}
	out := result.(*PutConfigurationSetReputationOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to enable or disable tracking of reputation metrics for a
// configuration set.
type PutConfigurationSetReputationOptionsInput struct {

	// The name of the configuration set that you want to enable or disable reputation
	// metric tracking for.
	//
	// This member is required.
	ConfigurationSetName *string

	// If true, tracking of reputation metrics is enabled for the configuration set. If
	// false, tracking of reputation metrics is disabled for the configuration set.
	ReputationMetricsEnabled *bool
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type PutConfigurationSetReputationOptionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutConfigurationSetReputationOptionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutConfigurationSetReputationOptions{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutConfigurationSetReputationOptions{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutConfigurationSetReputationOptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutConfigurationSetReputationOptions",
	}
}
