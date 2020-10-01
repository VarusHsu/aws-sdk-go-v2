// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Configures the access rules that control access to the domain's document and
// search endpoints. For more information, see  Configuring Access for an Amazon
// CloudSearch Domain
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-access.html).
func (c *Client) UpdateServiceAccessPolicies(ctx context.Context, params *UpdateServiceAccessPoliciesInput, optFns ...func(*Options)) (*UpdateServiceAccessPoliciesOutput, error) {
	stack := middleware.NewStack("UpdateServiceAccessPolicies", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpUpdateServiceAccessPoliciesMiddlewares(stack)
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
	addOpUpdateServiceAccessPoliciesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateServiceAccessPolicies(options.Region), middleware.Before)
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
			OperationName: "UpdateServiceAccessPolicies",
			Err:           err,
		}
	}
	out := result.(*UpdateServiceAccessPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the UpdateServiceAccessPolicies () operation.
// Specifies the name of the domain you want to update and the access rules you
// want to configure.
type UpdateServiceAccessPoliciesInput struct {

	// A string that represents the name of a domain. Domain names are unique across
	// the domains owned by an account within an AWS region. Domain names start with a
	// letter or number and can contain the following characters: a-z (lowercase), 0-9,
	// and - (hyphen).
	//
	// This member is required.
	DomainName *string

	// The access rules you want to configure. These rules replace any existing rules.
	//
	// This member is required.
	AccessPolicies *string
}

// The result of an UpdateServiceAccessPolicies request. Contains the new access
// policies.
type UpdateServiceAccessPoliciesOutput struct {

	// The access rules configured for the domain.
	//
	// This member is required.
	AccessPolicies *types.AccessPoliciesStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpUpdateServiceAccessPoliciesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpUpdateServiceAccessPolicies{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateServiceAccessPolicies{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateServiceAccessPolicies(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "UpdateServiceAccessPolicies",
	}
}
