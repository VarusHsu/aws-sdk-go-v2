// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the things associated with the specified principal. A principal can be
// X.509 certificates, IAM users, groups, and roles, Amazon Cognito identities or
// federated identities.
func (c *Client) ListPrincipalThings(ctx context.Context, params *ListPrincipalThingsInput, optFns ...func(*Options)) (*ListPrincipalThingsOutput, error) {
	stack := middleware.NewStack("ListPrincipalThings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListPrincipalThingsMiddlewares(stack)
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
	addOpListPrincipalThingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPrincipalThings(options.Region), middleware.Before)
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
			OperationName: "ListPrincipalThings",
			Err:           err,
		}
	}
	out := result.(*ListPrincipalThingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ListPrincipalThings operation.
type ListPrincipalThingsInput struct {

	// The principal.
	//
	// This member is required.
	Principal *string

	// The maximum number of results to return in this operation.
	MaxResults *int32

	// The token to retrieve the next set of results.
	NextToken *string
}

// The output from the ListPrincipalThings operation.
type ListPrincipalThingsOutput struct {

	// The token used to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// The things.
	Things []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListPrincipalThingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListPrincipalThings{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListPrincipalThings{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPrincipalThings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListPrincipalThings",
	}
}
