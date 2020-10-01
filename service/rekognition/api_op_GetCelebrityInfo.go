// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the name and additional information about a celebrity based on his or her
// Amazon Rekognition ID. The additional information is returned as an array of
// URLs. If there is no additional information about the celebrity, this list is
// empty.  <p>For more information, see Recognizing Celebrities in an Image in the
// Amazon Rekognition Developer Guide.</p> <p>This operation requires permissions
// to perform the <code>rekognition:GetCelebrityInfo</code> action. </p>
func (c *Client) GetCelebrityInfo(ctx context.Context, params *GetCelebrityInfoInput, optFns ...func(*Options)) (*GetCelebrityInfoOutput, error) {
	stack := middleware.NewStack("GetCelebrityInfo", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetCelebrityInfoMiddlewares(stack)
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
	addOpGetCelebrityInfoValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCelebrityInfo(options.Region), middleware.Before)
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
			OperationName: "GetCelebrityInfo",
			Err:           err,
		}
	}
	out := result.(*GetCelebrityInfoOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCelebrityInfoInput struct {

	// The ID for the celebrity. You get the celebrity ID from a call to the
	// RecognizeCelebrities () operation, which recognizes celebrities in an image.
	//
	// This member is required.
	Id *string
}

type GetCelebrityInfoOutput struct {

	// The name of the celebrity.
	Name *string

	// An array of URLs pointing to additional celebrity information.
	Urls []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetCelebrityInfoMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetCelebrityInfo{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCelebrityInfo{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCelebrityInfo(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "GetCelebrityInfo",
	}
}
