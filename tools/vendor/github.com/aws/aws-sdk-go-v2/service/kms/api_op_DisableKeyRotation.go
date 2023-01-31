// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disables automatic rotation of the key material
// (https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html) of the
// specified symmetric encryption KMS key. Automatic key rotation is supported only
// on symmetric encryption KMS keys. You cannot enable automatic rotation of
// asymmetric KMS keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html),
// HMAC KMS keys (https://docs.aws.amazon.com/kms/latest/developerguide/hmac.html),
// KMS keys with imported key material
// (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html), or
// KMS keys in a custom key store
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html).
// To enable or disable automatic rotation of a set of related multi-Region keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-manage.html#multi-region-rotate),
// set the property on the primary key. You can enable (EnableKeyRotation) and
// disable automatic rotation of the key material in customer managed KMS keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk).
// Key material rotation of Amazon Web Services managed KMS keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk)
// is not configurable. KMS always rotates the key material for every year.
// Rotation of Amazon Web Services owned KMS keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-owned-cmk)
// varies. In May 2022, KMS changed the rotation schedule for Amazon Web Services
// managed keys from every three years to every year. For details, see
// EnableKeyRotation. The KMS key that you use for this operation must be in a
// compatible key state. For details, see Key states of KMS keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// Key Management Service Developer Guide. Cross-account use: No. You cannot
// perform this operation on a KMS key in a different Amazon Web Services account.
// Required permissions: kms:DisableKeyRotation
// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// (key policy) Related operations:
//
// * EnableKeyRotation
//
// * GetKeyRotationStatus
func (c *Client) DisableKeyRotation(ctx context.Context, params *DisableKeyRotationInput, optFns ...func(*Options)) (*DisableKeyRotationOutput, error) {
	if params == nil {
		params = &DisableKeyRotationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableKeyRotation", params, optFns, c.addOperationDisableKeyRotationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableKeyRotationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableKeyRotationInput struct {

	// Identifies a symmetric encryption KMS key. You cannot enable or disable
	// automatic rotation of asymmetric KMS keys
	// (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html#asymmetric-cmks),
	// HMAC KMS keys (https://docs.aws.amazon.com/kms/latest/developerguide/hmac.html),
	// KMS keys with imported key material
	// (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html), or
	// KMS keys in a custom key store
	// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html).
	// Specify the key ID or key ARN of the KMS key. For example:
	//
	// * Key ID:
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To
	// get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey.
	//
	// This member is required.
	KeyId *string

	noSmithyDocumentSerde
}

type DisableKeyRotationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisableKeyRotationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisableKeyRotation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisableKeyRotation{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDisableKeyRotationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisableKeyRotation(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opDisableKeyRotation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "DisableKeyRotation",
	}
}
