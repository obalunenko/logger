// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a unique customer managed KMS key
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#kms-keys)
// in your Amazon Web Services account and Region. KMS is replacing the term
// customer master key (CMK) with KMS key and KMS key. The concept has not changed.
// To prevent breaking changes, KMS is keeping some variations of this term. You
// can use the CreateKey operation to create symmetric or asymmetric KMS keys.
//
// *
// Symmetric KMS keys contain a 256-bit symmetric key that never leaves KMS
// unencrypted. To use the KMS key, you must call KMS. You can use a symmetric KMS
// key to encrypt and decrypt small amounts of data, but they are typically used to
// generate data keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#data-keys)
// and data keys pairs
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#data-key-pairs).
// For details, see GenerateDataKey and GenerateDataKeyPair.
//
// * Asymmetric KMS keys
// can contain an RSA key pair or an Elliptic Curve (ECC) key pair. The private key
// in an asymmetric KMS key never leaves KMS unencrypted. However, you can use the
// GetPublicKey operation to download the public key so it can be used outside of
// KMS. KMS keys with RSA key pairs can be used to encrypt or decrypt data or sign
// and verify messages (but not both). KMS keys with ECC key pairs can be used only
// to sign and verify messages.
//
// For information about symmetric and asymmetric KMS
// keys, see Using Symmetric and Asymmetric KMS keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html)
// in the Key Management Service Developer Guide. To create different types of KMS
// keys, use the following guidance: Asymmetric KMS keys To create an asymmetric
// KMS key, use the KeySpec parameter to specify the type of key material in the
// KMS key. Then, use the KeyUsage parameter to determine whether the KMS key will
// be used to encrypt and decrypt or sign and verify. You can't change these
// properties after the KMS key is created. Symmetric KMS keys When creating a
// symmetric KMS key, you don't need to specify the KeySpec or KeyUsage parameters.
// The default value for KeySpec, SYMMETRIC_DEFAULT, and the default value for
// KeyUsage, ENCRYPT_DECRYPT, are the only valid values for symmetric KMS keys.
// Multi-Region primary keys Imported key material To create a multi-Region primary
// key in the local Amazon Web Services Region, use the MultiRegion parameter with
// a value of True. To create a multi-Region replica key, that is, a KMS key with
// the same key ID and key material as a primary key, but in a different Amazon Web
// Services Region, use the ReplicateKey operation. To change a replica key to a
// primary key, and its primary key to a replica key, use the UpdatePrimaryRegion
// operation. This operation supports multi-Region keys, an KMS feature that lets
// you create multiple interoperable KMS keys in different Amazon Web Services
// Regions. Because these KMS keys have the same key ID, key material, and other
// metadata, you can use them interchangeably to encrypt data in one Amazon Web
// Services Region and decrypt it in a different Amazon Web Services Region without
// re-encrypting the data or making a cross-Region call. For more information about
// multi-Region keys, see Using multi-Region keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html)
// in the Key Management Service Developer Guide. You can create symmetric and
// asymmetric multi-Region keys and multi-Region keys with imported key material.
// You cannot create multi-Region keys in a custom key store. To import your own
// key material, begin by creating a symmetric KMS key with no key material. To do
// this, use the Origin parameter of CreateKey with a value of EXTERNAL. Next, use
// GetParametersForImport operation to get a public key and import token, and use
// the public key to encrypt your key material. Then, use ImportKeyMaterial with
// your import token to import the key material. For step-by-step instructions, see
// Importing Key Material
// (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html) in
// the Key Management Service Developer Guide . You cannot import the key material
// into an asymmetric KMS key. To create a multi-Region primary key with imported
// key material, use the Origin parameter of CreateKey with a value of EXTERNAL and
// the MultiRegion parameter with a value of True. To create replicas of the
// multi-Region primary key, use the ReplicateKey operation. For more information
// about multi-Region keys, see Using multi-Region keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html)
// in the Key Management Service Developer Guide. Custom key store To create a
// symmetric KMS key in a custom key store
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html),
// use the CustomKeyStoreId parameter to specify the custom key store. You must
// also use the Origin parameter with a value of AWS_CLOUDHSM. The CloudHSM cluster
// that is associated with the custom key store must have at least two active HSMs
// in different Availability Zones in the Amazon Web Services Region. You cannot
// create an asymmetric KMS key in a custom key store. For information about custom
// key stores in KMS see Using Custom Key Stores
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
// in the Key Management Service Developer Guide . Cross-account use: No. You
// cannot use this operation to create a KMS key in a different Amazon Web Services
// account. Required permissions: kms:CreateKey
// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// (IAM policy). To use the Tags parameter, kms:TagResource
// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// (IAM policy). For examples and information about related permissions, see Allow
// a user to create KMS keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/iam-policies.html#iam-policy-example-create-key)
// in the Key Management Service Developer Guide. Related operations:
//
// *
// DescribeKey
//
// * ListKeys
//
// * ScheduleKeyDeletion
func (c *Client) CreateKey(ctx context.Context, params *CreateKeyInput, optFns ...func(*Options)) (*CreateKeyOutput, error) {
	if params == nil {
		params = &CreateKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateKey", params, optFns, c.addOperationCreateKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateKeyInput struct {

	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the KMS key becomes
	// unmanageable. Do not set this value to true indiscriminately. For more
	// information, refer to the scenario in the Default Key Policy
	// (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam)
	// section in the Key Management Service Developer Guide . Use this parameter only
	// when you include a policy in the request and you intend to prevent the principal
	// that is making the request from making a subsequent PutKeyPolicy request on the
	// KMS key. The default value is false.
	BypassPolicyLockoutSafetyCheck bool

	// Creates the KMS key in the specified custom key store
	// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
	// and the key material in its associated CloudHSM cluster. To create a KMS key in
	// a custom key store, you must also specify the Origin parameter with a value of
	// AWS_CLOUDHSM. The CloudHSM cluster that is associated with the custom key store
	// must have at least two active HSMs, each in a different Availability Zone in the
	// Region. This parameter is valid only for symmetric KMS keys and regional KMS
	// keys. You cannot create an asymmetric KMS key or a multi-Region key in a custom
	// key store. To find the ID of a custom key store, use the DescribeCustomKeyStores
	// operation. The response includes the custom key store ID and the ID of the
	// CloudHSM cluster. This operation is part of the Custom Key Store feature
	// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
	// feature in KMS, which combines the convenience and extensive integration of KMS
	// with the isolation and control of a single-tenant key store.
	CustomKeyStoreId *string

	// Instead, use the KeySpec parameter. The KeySpec and CustomerMasterKeySpec
	// parameters work the same way. Only the names differ. We recommend that you use
	// KeySpec parameter in your code. However, to avoid breaking changes, KMS will
	// support both parameters.
	//
	// Deprecated: This parameter has been deprecated. Instead, use the KeySpec
	// parameter.
	CustomerMasterKeySpec types.CustomerMasterKeySpec

	// A description of the KMS key. Use a description that helps you decide whether
	// the KMS key is appropriate for a task. The default value is an empty string (no
	// description). To set or change the description after the key is created, use
	// UpdateKeyDescription.
	Description *string

	// Specifies the type of KMS key to create. The default value, SYMMETRIC_DEFAULT,
	// creates a KMS key with a 256-bit symmetric key for encryption and decryption.
	// For help choosing a key spec for your KMS key, see How to Choose Your KMS key
	// Configuration
	// (https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose.html)
	// in the Key Management Service Developer Guide . The KeySpec determines whether
	// the KMS key contains a symmetric key or an asymmetric key pair. It also
	// determines the encryption algorithms or signing algorithms that the KMS key
	// supports. You can't change the KeySpec after the KMS key is created. To further
	// restrict the algorithms that can be used with the KMS key, use a condition key
	// in its key policy or IAM policy. For more information, see
	// kms:EncryptionAlgorithm
	// (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-encryption-algorithm)
	// or kms:Signing Algorithm
	// (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-signing-algorithm)
	// in the Key Management Service Developer Guide . Amazon Web Services services
	// that are integrated with KMS
	// (http://aws.amazon.com/kms/features/#AWS_Service_Integration) use symmetric KMS
	// keys to protect your data. These services do not support asymmetric KMS keys.
	// For help determining whether a KMS key is symmetric or asymmetric, see
	// Identifying Symmetric and Asymmetric KMS keys
	// (https://docs.aws.amazon.com/kms/latest/developerguide/find-symm-asymm.html) in
	// the Key Management Service Developer Guide. KMS supports the following key specs
	// for KMS keys:
	//
	// * Symmetric key (default)
	//
	// * SYMMETRIC_DEFAULT (AES-256-GCM)
	//
	// *
	// Asymmetric RSA key pairs
	//
	// * RSA_2048
	//
	// * RSA_3072
	//
	// * RSA_4096
	//
	// * Asymmetric
	// NIST-recommended elliptic curve key pairs
	//
	// * ECC_NIST_P256 (secp256r1)
	//
	// *
	// ECC_NIST_P384 (secp384r1)
	//
	// * ECC_NIST_P521 (secp521r1)
	//
	// * Other asymmetric
	// elliptic curve key pairs
	//
	// * ECC_SECG_P256K1 (secp256k1), commonly used for
	// cryptocurrencies.
	KeySpec types.KeySpec

	// Determines the cryptographic operations
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations)
	// for which you can use the KMS key. The default value is ENCRYPT_DECRYPT. This
	// parameter is required only for asymmetric KMS keys. You can't change the
	// KeyUsage value after the KMS key is created. Select only one valid value.
	//
	// * For
	// symmetric KMS keys, omit the parameter or specify ENCRYPT_DECRYPT.
	//
	// * For
	// asymmetric KMS keys with RSA key material, specify ENCRYPT_DECRYPT or
	// SIGN_VERIFY.
	//
	// * For asymmetric KMS keys with ECC key material, specify
	// SIGN_VERIFY.
	KeyUsage types.KeyUsageType

	// Creates a multi-Region primary key that you can replicate into other Amazon Web
	// Services Regions. You cannot change this value after you create the KMS key. For
	// a multi-Region key, set this parameter to True. For a single-Region KMS key,
	// omit this parameter or set it to False. The default value is False. This
	// operation supports multi-Region keys, an KMS feature that lets you create
	// multiple interoperable KMS keys in different Amazon Web Services Regions.
	// Because these KMS keys have the same key ID, key material, and other metadata,
	// you can use them interchangeably to encrypt data in one Amazon Web Services
	// Region and decrypt it in a different Amazon Web Services Region without
	// re-encrypting the data or making a cross-Region call. For more information about
	// multi-Region keys, see Using multi-Region keys
	// (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html)
	// in the Key Management Service Developer Guide. This value creates a primary key,
	// not a replica. To create a replica key, use the ReplicateKey operation. You can
	// create a symmetric or asymmetric multi-Region key, and you can create a
	// multi-Region key with imported key material. However, you cannot create a
	// multi-Region key in a custom key store.
	MultiRegion *bool

	// The source of the key material for the KMS key. You cannot change the origin
	// after you create the KMS key. The default is AWS_KMS, which means that KMS
	// creates the key material. To create a KMS key with no key material (for imported
	// key material), set the value to EXTERNAL. For more information about importing
	// key material into KMS, see Importing Key Material
	// (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html) in
	// the Key Management Service Developer Guide. This value is valid only for
	// symmetric KMS keys. To create a KMS key in an KMS custom key store
	// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
	// and create its key material in the associated CloudHSM cluster, set this value
	// to AWS_CLOUDHSM. You must also use the CustomKeyStoreId parameter to identify
	// the custom key store. This value is valid only for symmetric KMS keys.
	Origin types.OriginType

	// The key policy to attach to the KMS key. If you provide a key policy, it must
	// meet the following criteria:
	//
	// * If you don't set BypassPolicyLockoutSafetyCheck
	// to true, the key policy must allow the principal that is making the CreateKey
	// request to make a subsequent PutKeyPolicy request on the KMS key. This reduces
	// the risk that the KMS key becomes unmanageable. For more information, refer to
	// the scenario in the Default Key Policy
	// (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam)
	// section of the Key Management Service Developer Guide .
	//
	// * Each statement in the
	// key policy must contain one or more principals. The principals in the key policy
	// must exist and be visible to KMS. When you create a new Amazon Web Services
	// principal (for example, an IAM user or role), you might need to enforce a delay
	// before including the new principal in a key policy because the new principal
	// might not be immediately visible to KMS. For more information, see Changes that
	// I make are not always immediately visible
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_eventual-consistency)
	// in the Amazon Web Services Identity and Access Management User Guide.
	//
	// If you do
	// not provide a key policy, KMS attaches a default key policy to the KMS key. For
	// more information, see Default Key Policy
	// (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default)
	// in the Key Management Service Developer Guide. The key policy size quota is 32
	// kilobytes (32768 bytes). For help writing and formatting a JSON policy document,
	// see the IAM JSON Policy Reference
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies.html) in
	// the Identity and Access Management User Guide .
	Policy *string

	// Assigns one or more tags to the KMS key. Use this parameter to tag the KMS key
	// when it is created. To tag an existing KMS key, use the TagResource operation.
	// Tagging or untagging a KMS key can allow or deny permission to the KMS key. For
	// details, see Using ABAC in KMS
	// (https://docs.aws.amazon.com/kms/latest/developerguide/abac.html) in the Key
	// Management Service Developer Guide. To use this parameter, you must have
	// kms:TagResource
	// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
	// permission in an IAM policy. Each tag consists of a tag key and a tag value.
	// Both the tag key and the tag value are required, but the tag value can be an
	// empty (null) string. You cannot have more than one tag on a KMS key with the
	// same tag key. If you specify an existing tag key with a different tag value, KMS
	// replaces the current tag value with the specified one. When you add tags to an
	// Amazon Web Services resource, Amazon Web Services generates a cost allocation
	// report with usage and costs aggregated by tags. Tags can also be used to control
	// access to a KMS key. For details, see Tagging Keys
	// (https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html).
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateKeyOutput struct {

	// Metadata associated with the KMS key.
	KeyMetadata *types.KeyMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateKey{}, middleware.After)
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
	if err = addOpCreateKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateKey(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "CreateKey",
	}
}
