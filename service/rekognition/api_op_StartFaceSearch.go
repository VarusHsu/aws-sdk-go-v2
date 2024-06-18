// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts the asynchronous search for faces in a collection that match the faces
// of persons detected in a stored video.
//
// The video must be stored in an Amazon S3 bucket. Use Video to specify the bucket
// name and the filename of the video. StartFaceSearch returns a job identifier (
// JobId ) which you use to get the search results once the search has completed.
// When searching is finished, Amazon Rekognition Video publishes a completion
// status to the Amazon Simple Notification Service topic that you specify in
// NotificationChannel . To get the search results, first check that the status
// value published to the Amazon SNS topic is SUCCEEDED . If so, call GetFaceSearch and pass
// the job identifier ( JobId ) from the initial call to StartFaceSearch . For more
// information, see [Searching stored videos for faces].
//
// [Searching stored videos for faces]: https://docs.aws.amazon.com/rekognition/latest/dg/procedure-person-search-videos.html
func (c *Client) StartFaceSearch(ctx context.Context, params *StartFaceSearchInput, optFns ...func(*Options)) (*StartFaceSearchOutput, error) {
	if params == nil {
		params = &StartFaceSearchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartFaceSearch", params, optFns, c.addOperationStartFaceSearchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartFaceSearchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartFaceSearchInput struct {

	// ID of the collection that contains the faces you want to search for.
	//
	// This member is required.
	CollectionId *string

	// The video you want to search. The video must be stored in an Amazon S3 bucket.
	//
	// This member is required.
	Video *types.Video

	// Idempotent token used to identify the start request. If you use the same token
	// with multiple StartFaceSearch requests, the same JobId is returned. Use
	// ClientRequestToken to prevent the same job from being accidently started more
	// than once.
	ClientRequestToken *string

	// The minimum confidence in the person match to return. For example, don't return
	// any matches where confidence in matches is less than 70%. The default value is
	// 80%.
	FaceMatchThreshold *float32

	// An identifier you specify that's returned in the completion notification that's
	// published to your Amazon Simple Notification Service topic. For example, you can
	// use JobTag to group related jobs and identify them in the completion
	// notification.
	JobTag *string

	// The ARN of the Amazon SNS topic to which you want Amazon Rekognition Video to
	// publish the completion status of the search. The Amazon SNS topic must have a
	// topic name that begins with AmazonRekognition if you are using the
	// AmazonRekognitionServiceRole permissions policy to access the topic.
	NotificationChannel *types.NotificationChannel

	noSmithyDocumentSerde
}

type StartFaceSearchOutput struct {

	// The identifier for the search job. Use JobId to identify the job in a
	// subsequent call to GetFaceSearch .
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartFaceSearchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartFaceSearch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartFaceSearch{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartFaceSearch"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpStartFaceSearchValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartFaceSearch(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opStartFaceSearch(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartFaceSearch",
	}
}
