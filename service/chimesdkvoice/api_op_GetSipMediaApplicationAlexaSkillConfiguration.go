// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkvoice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkvoice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the Alexa Skill configuration for the SIP media application.
//
// Due to changes made by the Amazon Alexa service, this API is no longer
// available for use. For more information, refer to the [Alexa Smart Properties]page.
//
// Deprecated: Due to changes made by the Amazon Alexa service, this API is no
// longer available for use. For more information, refer to the Alexa Smart
// Properties page(https://developer.amazon.com/en-US/alexa/alexasmartproperties).
//
// [Alexa Smart Properties]: https://developer.amazon.com/en-US/alexa/alexasmartproperties
func (c *Client) GetSipMediaApplicationAlexaSkillConfiguration(ctx context.Context, params *GetSipMediaApplicationAlexaSkillConfigurationInput, optFns ...func(*Options)) (*GetSipMediaApplicationAlexaSkillConfigurationOutput, error) {
	if params == nil {
		params = &GetSipMediaApplicationAlexaSkillConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSipMediaApplicationAlexaSkillConfiguration", params, optFns, c.addOperationGetSipMediaApplicationAlexaSkillConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSipMediaApplicationAlexaSkillConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSipMediaApplicationAlexaSkillConfigurationInput struct {

	// The SIP media application ID.
	//
	// This member is required.
	SipMediaApplicationId *string

	noSmithyDocumentSerde
}

type GetSipMediaApplicationAlexaSkillConfigurationOutput struct {

	// Returns the Alexa Skill configuration.
	SipMediaApplicationAlexaSkillConfiguration *types.SipMediaApplicationAlexaSkillConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSipMediaApplicationAlexaSkillConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSipMediaApplicationAlexaSkillConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSipMediaApplicationAlexaSkillConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetSipMediaApplicationAlexaSkillConfiguration"); err != nil {
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
	if err = addOpGetSipMediaApplicationAlexaSkillConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSipMediaApplicationAlexaSkillConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSipMediaApplicationAlexaSkillConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetSipMediaApplicationAlexaSkillConfiguration",
	}
}
