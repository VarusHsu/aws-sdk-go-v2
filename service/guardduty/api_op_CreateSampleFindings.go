// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Generates example findings of types specified by the list of finding types. If
// 'NULL' is specified for findingTypes, the API generates example findings of all
// supported finding types.
func (c *Client) CreateSampleFindings(ctx context.Context, params *CreateSampleFindingsInput, optFns ...func(*Options)) (*CreateSampleFindingsOutput, error) {
	stack := middleware.NewStack("CreateSampleFindings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateSampleFindingsMiddlewares(stack)
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
	addOpCreateSampleFindingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSampleFindings(options.Region), middleware.Before)
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
			OperationName: "CreateSampleFindings",
			Err:           err,
		}
	}
	out := result.(*CreateSampleFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSampleFindingsInput struct {

	// The ID of the detector to create sample findings for.
	//
	// This member is required.
	DetectorId *string

	// The types of sample findings to generate.
	FindingTypes []*string
}

type CreateSampleFindingsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateSampleFindingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateSampleFindings{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSampleFindings{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateSampleFindings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "CreateSampleFindings",
	}
}
