// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkvoice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the emergency calling details from the specified Amazon Chime SDK Voice
// Connector.
func (c *Client) DeleteVoiceConnectorEmergencyCallingConfiguration(ctx context.Context, params *DeleteVoiceConnectorEmergencyCallingConfigurationInput, optFns ...func(*Options)) (*DeleteVoiceConnectorEmergencyCallingConfigurationOutput, error) {
	if params == nil {
		params = &DeleteVoiceConnectorEmergencyCallingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteVoiceConnectorEmergencyCallingConfiguration", params, optFns, c.addOperationDeleteVoiceConnectorEmergencyCallingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteVoiceConnectorEmergencyCallingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteVoiceConnectorEmergencyCallingConfigurationInput struct {

	// The Voice Connector ID.
	//
	// This member is required.
	VoiceConnectorId *string

	noSmithyDocumentSerde
}

type DeleteVoiceConnectorEmergencyCallingConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteVoiceConnectorEmergencyCallingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteVoiceConnectorEmergencyCallingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteVoiceConnectorEmergencyCallingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteVoiceConnectorEmergencyCallingConfiguration"); err != nil {
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
	if err = addOpDeleteVoiceConnectorEmergencyCallingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVoiceConnectorEmergencyCallingConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteVoiceConnectorEmergencyCallingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteVoiceConnectorEmergencyCallingConfiguration",
	}
}
