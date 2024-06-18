// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Deletes an existing opted out destination phone number from the specified
// opt-out list.
//
// Each destination phone number can only be deleted once every 30 days.
//
// If the specified destination phone number doesn't exist or if the opt-out list
// doesn't exist, an error is returned.
func (c *Client) DeleteOptedOutNumber(ctx context.Context, params *DeleteOptedOutNumberInput, optFns ...func(*Options)) (*DeleteOptedOutNumberOutput, error) {
	if params == nil {
		params = &DeleteOptedOutNumberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteOptedOutNumber", params, optFns, c.addOperationDeleteOptedOutNumberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteOptedOutNumberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteOptedOutNumberInput struct {

	// The OptOutListName or OptOutListArn to remove the phone number from.
	//
	// This member is required.
	OptOutListName *string

	// The phone number, in E.164 format, to remove from the OptOutList.
	//
	// This member is required.
	OptedOutNumber *string

	noSmithyDocumentSerde
}

type DeleteOptedOutNumberOutput struct {

	// This is true if it was the end user who requested their phone number be
	// removed.
	EndUserOptedOut bool

	// The OptOutListArn that the phone number was removed from.
	OptOutListArn *string

	// The OptOutListName that the phone number was removed from.
	OptOutListName *string

	// The phone number that was removed from the OptOutList.
	OptedOutNumber *string

	// The time that the number was removed at, in [UNIX epoch time] format.
	//
	// [UNIX epoch time]: https://www.epochconverter.com/
	OptedOutTimestamp *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteOptedOutNumberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteOptedOutNumber{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteOptedOutNumber{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteOptedOutNumber"); err != nil {
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
	if err = addOpDeleteOptedOutNumberValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteOptedOutNumber(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteOptedOutNumber(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteOptedOutNumber",
	}
}
