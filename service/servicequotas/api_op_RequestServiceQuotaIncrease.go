// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the details of a service quota increase request. The response to this
// command provides the details in the RequestedServiceQuotaChange () object.
func (c *Client) RequestServiceQuotaIncrease(ctx context.Context, params *RequestServiceQuotaIncreaseInput, optFns ...func(*Options)) (*RequestServiceQuotaIncreaseOutput, error) {
	stack := middleware.NewStack("RequestServiceQuotaIncrease", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRequestServiceQuotaIncreaseMiddlewares(stack)
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
	addOpRequestServiceQuotaIncreaseValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRequestServiceQuotaIncrease(options.Region), middleware.Before)
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
			OperationName: "RequestServiceQuotaIncrease",
			Err:           err,
		}
	}
	out := result.(*RequestServiceQuotaIncreaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RequestServiceQuotaIncreaseInput struct {

	// Specifies the value submitted in the service quota increase request.
	//
	// This member is required.
	DesiredValue *float64

	// Specifies the service that you want to use.
	//
	// This member is required.
	ServiceCode *string

	// Specifies the service quota that you want to use.
	//
	// This member is required.
	QuotaCode *string
}

type RequestServiceQuotaIncreaseOutput struct {

	// Returns a list of service quota requests.
	RequestedQuota *types.RequestedServiceQuotaChange

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRequestServiceQuotaIncreaseMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRequestServiceQuotaIncrease{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRequestServiceQuotaIncrease{}, middleware.After)
}

func newServiceMetadataMiddleware_opRequestServiceQuotaIncrease(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "RequestServiceQuotaIncrease",
	}
}
