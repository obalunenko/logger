// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Schedules the deletion of a customer master key (CMK). By default, AWS KMS
// applies a waiting period of 30 days, but you can specify a waiting period of
// 7-30 days. When this operation is successful, the key state of the CMK changes
// to PendingDeletion and the key can't be used in any cryptographic operations. It
// remains in this state for the duration of the waiting period. Before the waiting
// period ends, you can use CancelKeyDeletion to cancel the deletion of the CMK.
// After the waiting period ends, AWS KMS deletes the CMK, its key material, and
// all AWS KMS data associated with it, including all aliases that refer to it.
// Deleting a CMK is a destructive and potentially dangerous operation. When a CMK
// is deleted, all data that was encrypted under the CMK is unrecoverable. (The
// only exception is a multi-Region replica key.) To prevent the use of a CMK
// without deleting it, use DisableKey. If you schedule deletion of a CMK from a
// custom key store
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html),
// when the waiting period expires, ScheduleKeyDeletion deletes the CMK from AWS
// KMS. Then AWS KMS makes a best effort to delete the key material from the
// associated AWS CloudHSM cluster. However, you might need to manually delete the
// orphaned key material
// (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html#fix-keystore-orphaned-key)
// from the cluster and its backups. You can schedule the deletion of a
// multi-Region primary key and its replica keys at any time. However, AWS KMS will
// not delete a multi-Region primary key with existing replica keys. If you
// schedule the deletion of a primary key with replicas, its key state changes to
// PendingReplicaDeletion and it cannot be replicated or used in cryptographic
// operations. This status can continue indefinitely. When the last of its replicas
// keys is deleted (not just scheduled), the key state of the primary key changes
// to PendingDeletion and its waiting period (PendingWindowInDays) begins. For
// details, see Deleting multi-Region keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-delete.html)
// in the AWS Key Management Service Developer Guide. For more information about
// scheduling a CMK for deletion, see Deleting Customer Master Keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html) in
// the AWS Key Management Service Developer Guide. The CMK that you use for this
// operation must be in a compatible key state. For details, see Key state: Effect
// on your CMK
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// AWS Key Management Service Developer Guide. Cross-account use: No. You cannot
// perform this operation on a CMK in a different AWS account. Required
// permissions: kms:ScheduleKeyDeletion (key policy) Related operations
//
// *
// CancelKeyDeletion
//
// * DisableKey
func (c *Client) ScheduleKeyDeletion(ctx context.Context, params *ScheduleKeyDeletionInput, optFns ...func(*Options)) (*ScheduleKeyDeletionOutput, error) {
	if params == nil {
		params = &ScheduleKeyDeletionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ScheduleKeyDeletion", params, optFns, c.addOperationScheduleKeyDeletionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ScheduleKeyDeletionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ScheduleKeyDeletionInput struct {

	// The unique identifier of the customer master key (CMK) to delete. Specify the
	// key ID or key ARN of the CMK. For example:
	//
	// * Key ID:
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To
	// get the key ID and key ARN for a CMK, use ListKeys or DescribeKey.
	//
	// This member is required.
	KeyId *string

	// The waiting period, specified in number of days. After the waiting period ends,
	// AWS KMS deletes the customer master key (CMK). If the CMK is a multi-Region
	// primary key with replicas, the waiting period begins when the last of its
	// replica keys is deleted. Otherwise, the waiting period begins immediately. This
	// value is optional. If you include a value, it must be between 7 and 30,
	// inclusive. If you do not include a value, it defaults to 30.
	PendingWindowInDays *int32

	noSmithyDocumentSerde
}

type ScheduleKeyDeletionOutput struct {

	// The date and time after which AWS KMS deletes the customer master key (CMK). If
	// the CMK is a multi-Region primary key with replica keys, this field does not
	// appear. The deletion date for the primary key isn't known until its last replica
	// key is deleted.
	DeletionDate *time.Time

	// The Amazon Resource Name (key ARN
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN))
	// of the CMK whose deletion is scheduled.
	KeyId *string

	// The current status of the CMK. For more information about how key state affects
	// the use of a CMK, see Key state: Effect on your CMK
	// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
	// AWS Key Management Service Developer Guide.
	KeyState types.KeyState

	// The waiting period before the CMK is deleted. If the CMK is a multi-Region
	// primary key with replicas, the waiting period begins when the last of its
	// replica keys is deleted. Otherwise, the waiting period begins immediately.
	PendingWindowInDays *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationScheduleKeyDeletionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpScheduleKeyDeletion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpScheduleKeyDeletion{}, middleware.After)
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
	if err = addOpScheduleKeyDeletionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opScheduleKeyDeletion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opScheduleKeyDeletion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "ScheduleKeyDeletion",
	}
}
