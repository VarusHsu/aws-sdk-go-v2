// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all the roots, organizational units (OUs), and accounts that the specified
// policy is attached to. Always check the NextToken response parameter for a null
// value when calling a List* operation. These operations can occasionally return
// an empty set of results even when there are more results available. The
// NextToken response parameter value is null only when there are no more results
// to display. This operation can be called only from the organization's master
// account or by a member account that is a delegated administrator for an AWS
// service.
func (c *Client) ListTargetsForPolicy(ctx context.Context, params *ListTargetsForPolicyInput, optFns ...func(*Options)) (*ListTargetsForPolicyOutput, error) {
	stack := middleware.NewStack("ListTargetsForPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListTargetsForPolicyMiddlewares(stack)
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
	addOpListTargetsForPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTargetsForPolicy(options.Region), middleware.Before)
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
			OperationName: "ListTargetsForPolicy",
			Err:           err,
		}
	}
	out := result.(*ListTargetsForPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTargetsForPolicyInput struct {

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value of the previous call's NextToken
	// response to indicate where the output should continue from.
	NextToken *string

	// The total number of results that you want included on each page of the response.
	// If you do not include this parameter, it defaults to a value that is specific to
	// the operation. If additional items exist beyond the maximum you specify, the
	// NextToken response element is present and has a value (is not null). Include
	// that value as the NextToken request parameter in the next call to the operation
	// to get the next part of the results. Note that Organizations might return fewer
	// results than the maximum even when there are more results available. You should
	// check NextToken after every operation to ensure that you receive all of the
	// results.
	MaxResults *int32

	// The unique identifier (ID) of the policy whose attachments you want to know. The
	// regex pattern (http://wikipedia.org/wiki/regex) for a policy ID string requires
	// "p-" followed by from 8 to 128 lowercase or uppercase letters, digits, or the
	// underscore character (_).
	//
	// This member is required.
	PolicyId *string
}

type ListTargetsForPolicyOutput struct {

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null.
	NextToken *string

	// A list of structures, each of which contains details about one of the entities
	// to which the specified policy is attached.
	Targets []*types.PolicyTargetSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListTargetsForPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListTargetsForPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTargetsForPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opListTargetsForPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "ListTargetsForPolicy",
	}
}
