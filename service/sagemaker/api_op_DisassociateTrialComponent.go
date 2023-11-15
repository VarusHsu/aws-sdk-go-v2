// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates a trial component from a trial. This doesn't effect other trials
// the component is associated with. Before you can delete a component, you must
// disassociate the component from all trials it is associated with. To associate a
// trial component with a trial, call the AssociateTrialComponent (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_AssociateTrialComponent.html)
// API. To get a list of the trials a component is associated with, use the Search (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_Search.html)
// API. Specify ExperimentTrialComponent for the Resource parameter. The list
// appears in the response under Results.TrialComponent.Parents .
func (c *Client) DisassociateTrialComponent(ctx context.Context, params *DisassociateTrialComponentInput, optFns ...func(*Options)) (*DisassociateTrialComponentOutput, error) {
	if params == nil {
		params = &DisassociateTrialComponentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateTrialComponent", params, optFns, c.addOperationDisassociateTrialComponentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateTrialComponentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateTrialComponentInput struct {

	// The name of the component to disassociate from the trial.
	//
	// This member is required.
	TrialComponentName *string

	// The name of the trial to disassociate from.
	//
	// This member is required.
	TrialName *string

	noSmithyDocumentSerde
}

type DisassociateTrialComponentOutput struct {

	// The Amazon Resource Name (ARN) of the trial.
	TrialArn *string

	// The Amazon Resource Name (ARN) of the trial component.
	TrialComponentArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateTrialComponentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateTrialComponent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateTrialComponent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DisassociateTrialComponent"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpDisassociateTrialComponentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateTrialComponent(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateTrialComponent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DisassociateTrialComponent",
	}
}
