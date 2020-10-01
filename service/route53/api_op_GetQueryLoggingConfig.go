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

// Gets information about a specified configuration for DNS query logging.  <p>For
// more information about DNS query logs, see <a
// href="https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateQueryLoggingConfig.html">CreateQueryLoggingConfig</a>
// and <a
// href="https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/query-logs.html">Logging
// DNS Queries</a>.</p>
func (c *Client) GetQueryLoggingConfig(ctx context.Context, params *GetQueryLoggingConfigInput, optFns ...func(*Options)) (*GetQueryLoggingConfigOutput, error) {
	stack := middleware.NewStack("GetQueryLoggingConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetQueryLoggingConfigMiddlewares(stack)
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
	addOpGetQueryLoggingConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetQueryLoggingConfig(options.Region), middleware.Before)
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
			OperationName: "GetQueryLoggingConfig",
			Err:           err,
		}
	}
	out := result.(*GetQueryLoggingConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetQueryLoggingConfigInput struct {

	// The ID of the configuration for DNS query logging that you want to get
	// information about.
	//
	// This member is required.
	Id *string
}

type GetQueryLoggingConfigOutput struct {

	// A complex type that contains information about the query logging configuration
	// that you specified in a GetQueryLoggingConfig
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetQueryLoggingConfig.html)
	// request.
	//
	// This member is required.
	QueryLoggingConfig *types.QueryLoggingConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetQueryLoggingConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetQueryLoggingConfig{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetQueryLoggingConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetQueryLoggingConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "GetQueryLoggingConfig",
	}
}
