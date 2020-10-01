// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Tests a custom data identifier.
func (c *Client) TestCustomDataIdentifier(ctx context.Context, params *TestCustomDataIdentifierInput, optFns ...func(*Options)) (*TestCustomDataIdentifierOutput, error) {
	stack := middleware.NewStack("TestCustomDataIdentifier", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpTestCustomDataIdentifierMiddlewares(stack)
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
	addOpTestCustomDataIdentifierValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTestCustomDataIdentifier(options.Region), middleware.Before)
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
			OperationName: "TestCustomDataIdentifier",
			Err:           err,
		}
	}
	out := result.(*TestCustomDataIdentifierOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TestCustomDataIdentifierInput struct {

	// The regular expression (regex) that defines the pattern to match. The expression
	// can contain as many as 512 characters.
	//
	// This member is required.
	Regex *string

	// An array that lists specific character sequences (ignore words) to exclude from
	// the results. If the text matched by the regular expression is the same as any
	// string in this array, Amazon Macie ignores it. The array can contain as many as
	// 10 ignore words. Each ignore word can contain 4 - 90 characters.
	IgnoreWords []*string

	// An array that lists specific character sequences (keywords), one of which must
	// be within proximity (maximumMatchDistance) of the regular expression to match.
	// The array can contain as many as 50 keywords. Each keyword can contain 4 - 90
	// characters.
	Keywords []*string

	// The maximum number of characters that can exist between text that matches the
	// regex pattern and the character sequences specified by the keywords array. Macie
	// includes or excludes a result based on the proximity of a keyword to text that
	// matches the regex pattern. The distance can be 1 - 300 characters. The default
	// value is 50.
	MaximumMatchDistance *int32

	// The sample text to inspect by using the custom data identifier. The text can
	// contain as many as 1,000 characters.
	//
	// This member is required.
	SampleText *string
}

type TestCustomDataIdentifierOutput struct {

	// The number of instances of sample text that matched the detection criteria
	// specified in the custom data identifier.
	MatchCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpTestCustomDataIdentifierMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpTestCustomDataIdentifier{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpTestCustomDataIdentifier{}, middleware.After)
}

func newServiceMetadataMiddleware_opTestCustomDataIdentifier(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "TestCustomDataIdentifier",
	}
}
