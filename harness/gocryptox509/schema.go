// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package main

import "encoding/json"
import "fmt"
import "reflect"
import "time"

type ExpectedResult string

const ExpectedResultFAILURE ExpectedResult = "FAILURE"
const ExpectedResultSUCCESS ExpectedResult = "SUCCESS"

type Feature string

const FeatureHasCertPolicies Feature = "has-cert-policies"
const FeatureMaxChainDepth Feature = "max-chain-depth"
const FeatureNameConstraintDn Feature = "name-constraint-dn"
const FeatureNoCertPolicies Feature = "no-cert-policies"
const FeaturePedanticPublicSuffixWildcard Feature = "pedantic-public-suffix-wildcard"
const FeaturePedanticRfc5280 Feature = "pedantic-rfc5280"
const FeaturePedanticSerialNumber Feature = "pedantic-serial-number"
const FeaturePedanticWebpki Feature = "pedantic-webpki"
const FeaturePedanticWebpkiEku Feature = "pedantic-webpki-eku"
const FeatureRfc5280IncompatibleWithWebpki Feature = "rfc5280-incompatible-with-webpki"

type KeyUsage string

const KeyUsageCRLSign KeyUsage = "cRLSign"
const KeyUsageContentCommitment KeyUsage = "contentCommitment"
const KeyUsageDataEncipherment KeyUsage = "dataEncipherment"
const KeyUsageDecipherOnly KeyUsage = "decipher_only"
const KeyUsageDigitalSignature KeyUsage = "digitalSignature"
const KeyUsageEncipherOnly KeyUsage = "encipher_only"
const KeyUsageKeyAgreement KeyUsage = "keyAgreement"
const KeyUsageKeyCertSign KeyUsage = "keyCertSign"
const KeyUsageKeyEncipherment KeyUsage = "keyEncipherment"

type KnownEKUs string

const KnownEKUsAnyExtendedKeyUsage KnownEKUs = "anyExtendedKeyUsage"
const KnownEKUsClientAuth KnownEKUs = "clientAuth"
const KnownEKUsCodeSigning KnownEKUs = "codeSigning"
const KnownEKUsEmailProtection KnownEKUs = "emailProtection"
const KnownEKUsOCSPSigning KnownEKUs = "OCSPSigning"
const KnownEKUsServerAuth KnownEKUs = "serverAuth"
const KnownEKUsTimeStamping KnownEKUs = "timeStamping"

// The top-level testcase container.
type Limbo struct {
	// One or more testcases in this testsuite
	Testcases []Testcase `json:"testcases" yaml:"testcases" mapstructure:"testcases"`

	// The limbo schema version; this must currently always be 1
	Version LimboVersion `json:"version" yaml:"version" mapstructure:"version"`
}

type LimboVersion int

type PeerKind string

const PeerKindDNS PeerKind = "DNS"
const PeerKindIP PeerKind = "IP"
const PeerKindRFC822 PeerKind = "RFC822"

// Represents a peer (i.e., end entity) certificate's name (Subject or SAN).
type PeerName struct {
	// The kind of peer name
	Kind interface{} `json:"kind" yaml:"kind" mapstructure:"kind"`

	// The peer's name
	Value string `json:"value" yaml:"value" mapstructure:"value"`
}

type SignatureAlgorithm string

const SignatureAlgorithmDSAWITHSHA1 SignatureAlgorithm = "DSA_WITH_SHA1"
const SignatureAlgorithmDSAWITHSHA224 SignatureAlgorithm = "DSA_WITH_SHA224"
const SignatureAlgorithmDSAWITHSHA256 SignatureAlgorithm = "DSA_WITH_SHA256"
const SignatureAlgorithmDSAWITHSHA384 SignatureAlgorithm = "DSA_WITH_SHA384"
const SignatureAlgorithmDSAWITHSHA512 SignatureAlgorithm = "DSA_WITH_SHA512"
const SignatureAlgorithmECDSAWITHSHA1 SignatureAlgorithm = "ECDSA_WITH_SHA1"
const SignatureAlgorithmECDSAWITHSHA224 SignatureAlgorithm = "ECDSA_WITH_SHA224"
const SignatureAlgorithmECDSAWITHSHA256 SignatureAlgorithm = "ECDSA_WITH_SHA256"
const SignatureAlgorithmECDSAWITHSHA3224 SignatureAlgorithm = "ECDSA_WITH_SHA3_224"
const SignatureAlgorithmECDSAWITHSHA3256 SignatureAlgorithm = "ECDSA_WITH_SHA3_256"
const SignatureAlgorithmECDSAWITHSHA3384 SignatureAlgorithm = "ECDSA_WITH_SHA3_384"
const SignatureAlgorithmECDSAWITHSHA3512 SignatureAlgorithm = "ECDSA_WITH_SHA3_512"
const SignatureAlgorithmECDSAWITHSHA384 SignatureAlgorithm = "ECDSA_WITH_SHA384"
const SignatureAlgorithmECDSAWITHSHA512 SignatureAlgorithm = "ECDSA_WITH_SHA512"
const SignatureAlgorithmED25519 SignatureAlgorithm = "ED25519"
const SignatureAlgorithmED448 SignatureAlgorithm = "ED448"
const SignatureAlgorithmGOSTR34102012WITH34112012256 SignatureAlgorithm = "GOSTR3410_2012_WITH_3411_2012_256"
const SignatureAlgorithmGOSTR34102012WITH34112012512 SignatureAlgorithm = "GOSTR3410_2012_WITH_3411_2012_512"
const SignatureAlgorithmGOSTR341194WITH34102001 SignatureAlgorithm = "GOSTR3411_94_WITH_3410_2001"
const SignatureAlgorithmRSASSAPSS SignatureAlgorithm = "RSASSA_PSS"
const SignatureAlgorithmRSAWITHMD5 SignatureAlgorithm = "RSA_WITH_MD5"
const SignatureAlgorithmRSAWITHSHA1 SignatureAlgorithm = "RSA_WITH_SHA1"
const SignatureAlgorithmRSAWITHSHA224 SignatureAlgorithm = "RSA_WITH_SHA224"
const SignatureAlgorithmRSAWITHSHA256 SignatureAlgorithm = "RSA_WITH_SHA256"
const SignatureAlgorithmRSAWITHSHA3224 SignatureAlgorithm = "RSA_WITH_SHA3_224"
const SignatureAlgorithmRSAWITHSHA3256 SignatureAlgorithm = "RSA_WITH_SHA3_256"
const SignatureAlgorithmRSAWITHSHA3384 SignatureAlgorithm = "RSA_WITH_SHA3_384"
const SignatureAlgorithmRSAWITHSHA3512 SignatureAlgorithm = "RSA_WITH_SHA3_512"
const SignatureAlgorithmRSAWITHSHA384 SignatureAlgorithm = "RSA_WITH_SHA384"
const SignatureAlgorithmRSAWITHSHA512 SignatureAlgorithm = "RSA_WITH_SHA512"

// Represents an individual Limbo testcase.
type Testcase struct {
	// A list of testcase IDs that this testcase is mutually incompatible with
	ConflictsWith []string `json:"conflicts_with,omitempty" yaml:"conflicts_with,omitempty" mapstructure:"conflicts_with,omitempty"`

	// A short, Markdown-formatted description
	Description string `json:"description" yaml:"description" mapstructure:"description"`

	// For client-side validation: the expected peer name, if any
	ExpectedPeerName interface{} `json:"expected_peer_name,omitempty" yaml:"expected_peer_name,omitempty" mapstructure:"expected_peer_name,omitempty"`

	// For server-side validation: the expected peer names
	ExpectedPeerNames []PeerName `json:"expected_peer_names" yaml:"expected_peer_names" mapstructure:"expected_peer_names"`

	// The expected validation result
	ExpectedResult interface{} `json:"expected_result" yaml:"expected_result" mapstructure:"expected_result"`

	// A constraining list of extended key usages, either in well-known form or as
	// OIDs
	ExtendedKeyUsage []KnownEKUs `json:"extended_key_usage" yaml:"extended_key_usage" mapstructure:"extended_key_usage"`

	// Zero or more human-readable tags that describe OPTIONAL functionality described
	// by this testcase. Implementers should use this to specify testcases for
	// non-mandatory X.509 behavior (like certificate policy validation) or for
	// 'pedantic' cases. Consumers that don't understand a given feature should skip
	// tests that are marked with it.
	Features []Feature `json:"features,omitempty" yaml:"features,omitempty" mapstructure:"features,omitempty"`

	// A short, unique identifier for this testcase
	Id string `json:"id" yaml:"id" mapstructure:"id"`

	// A constraining list of key usages
	KeyUsage []KeyUsage `json:"key_usage" yaml:"key_usage" mapstructure:"key_usage"`

	// The maximum chain-building depth
	MaxChainDepth *int `json:"max_chain_depth,omitempty" yaml:"max_chain_depth,omitempty" mapstructure:"max_chain_depth,omitempty"`

	// The PEM-encoded peer (EE) certificate
	PeerCertificate string `json:"peer_certificate" yaml:"peer_certificate" mapstructure:"peer_certificate"`

	// A list of acceptable signature algorithms to constrain against
	SignatureAlgorithms []SignatureAlgorithm `json:"signature_algorithms" yaml:"signature_algorithms" mapstructure:"signature_algorithms"`

	// A list of PEM-encoded CA certificates to consider trusted
	TrustedCerts []string `json:"trusted_certs" yaml:"trusted_certs" mapstructure:"trusted_certs"`

	// A list of PEM-encoded untrusted intermediates to use during path building
	UntrustedIntermediates []string `json:"untrusted_intermediates" yaml:"untrusted_intermediates" mapstructure:"untrusted_intermediates"`

	// The kind of validation to perform
	ValidationKind interface{} `json:"validation_kind" yaml:"validation_kind" mapstructure:"validation_kind"`

	// The time at which to perform the validation
	ValidationTime *time.Time `json:"validation_time,omitempty" yaml:"validation_time,omitempty" mapstructure:"validation_time,omitempty"`
}

type ValidationKind string

const ValidationKindCLIENT ValidationKind = "CLIENT"
const ValidationKindSERVER ValidationKind = "SERVER"

var enumValues_ExpectedResult = []interface{}{
	"SUCCESS",
	"FAILURE",
}
var enumValues_Feature = []interface{}{
	"has-cert-policies",
	"no-cert-policies",
	"pedantic-public-suffix-wildcard",
	"name-constraint-dn",
	"pedantic-webpki",
	"pedantic-webpki-eku",
	"pedantic-serial-number",
	"max-chain-depth",
	"pedantic-rfc5280",
	"rfc5280-incompatible-with-webpki",
}
var enumValues_KeyUsage = []interface{}{
	"digitalSignature",
	"contentCommitment",
	"keyEncipherment",
	"dataEncipherment",
	"keyAgreement",
	"keyCertSign",
	"cRLSign",
	"encipher_only",
	"decipher_only",
}
var enumValues_KnownEKUs = []interface{}{
	"anyExtendedKeyUsage",
	"serverAuth",
	"clientAuth",
	"codeSigning",
	"emailProtection",
	"timeStamping",
	"OCSPSigning",
}
var enumValues_LimboVersion = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *LimboVersion) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_LimboVersion {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_LimboVersion, v)
	}
	*j = LimboVersion(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *KeyUsage) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_KeyUsage {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_KeyUsage, v)
	}
	*j = KeyUsage(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Limbo) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["testcases"]; !ok || v == nil {
		return fmt.Errorf("field testcases in Limbo: required")
	}
	if v, ok := raw["version"]; !ok || v == nil {
		return fmt.Errorf("field version in Limbo: required")
	}
	type Plain Limbo
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Limbo(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Feature) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_Feature {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_Feature, v)
	}
	*j = Feature(v)
	return nil
}

var enumValues_PeerKind = []interface{}{
	"RFC822",
	"DNS",
	"IP",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PeerKind) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_PeerKind {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_PeerKind, v)
	}
	*j = PeerKind(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Testcase) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["description"]; !ok || v == nil {
		return fmt.Errorf("field description in Testcase: required")
	}
	if v, ok := raw["expected_peer_names"]; !ok || v == nil {
		return fmt.Errorf("field expected_peer_names in Testcase: required")
	}
	if v, ok := raw["expected_result"]; !ok || v == nil {
		return fmt.Errorf("field expected_result in Testcase: required")
	}
	if v, ok := raw["extended_key_usage"]; !ok || v == nil {
		return fmt.Errorf("field extended_key_usage in Testcase: required")
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id in Testcase: required")
	}
	if v, ok := raw["key_usage"]; !ok || v == nil {
		return fmt.Errorf("field key_usage in Testcase: required")
	}
	if v, ok := raw["peer_certificate"]; !ok || v == nil {
		return fmt.Errorf("field peer_certificate in Testcase: required")
	}
	if v, ok := raw["signature_algorithms"]; !ok || v == nil {
		return fmt.Errorf("field signature_algorithms in Testcase: required")
	}
	if v, ok := raw["trusted_certs"]; !ok || v == nil {
		return fmt.Errorf("field trusted_certs in Testcase: required")
	}
	if v, ok := raw["untrusted_intermediates"]; !ok || v == nil {
		return fmt.Errorf("field untrusted_intermediates in Testcase: required")
	}
	if v, ok := raw["validation_kind"]; !ok || v == nil {
		return fmt.Errorf("field validation_kind in Testcase: required")
	}
	type Plain Testcase
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["conflicts_with"]; !ok || v == nil {
		plain.ConflictsWith = []string{}
	}
	if v, ok := raw["features"]; !ok || v == nil {
		plain.Features = []Feature{}
	}
	*j = Testcase(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ExpectedResult) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ExpectedResult {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ExpectedResult, v)
	}
	*j = ExpectedResult(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *KnownEKUs) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_KnownEKUs {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_KnownEKUs, v)
	}
	*j = KnownEKUs(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PeerName) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["kind"]; !ok || v == nil {
		return fmt.Errorf("field kind in PeerName: required")
	}
	if v, ok := raw["value"]; !ok || v == nil {
		return fmt.Errorf("field value in PeerName: required")
	}
	type Plain PeerName
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = PeerName(plain)
	return nil
}

var enumValues_ValidationKind = []interface{}{
	"CLIENT",
	"SERVER",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ValidationKind) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ValidationKind {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ValidationKind, v)
	}
	*j = ValidationKind(v)
	return nil
}

var enumValues_SignatureAlgorithm = []interface{}{
	"RSA_WITH_MD5",
	"RSA_WITH_SHA1",
	"RSA_WITH_SHA224",
	"RSA_WITH_SHA256",
	"RSA_WITH_SHA384",
	"RSA_WITH_SHA512",
	"RSA_WITH_SHA3_224",
	"RSA_WITH_SHA3_256",
	"RSA_WITH_SHA3_384",
	"RSA_WITH_SHA3_512",
	"RSASSA_PSS",
	"ECDSA_WITH_SHA1",
	"ECDSA_WITH_SHA224",
	"ECDSA_WITH_SHA256",
	"ECDSA_WITH_SHA384",
	"ECDSA_WITH_SHA512",
	"ECDSA_WITH_SHA3_224",
	"ECDSA_WITH_SHA3_256",
	"ECDSA_WITH_SHA3_384",
	"ECDSA_WITH_SHA3_512",
	"DSA_WITH_SHA1",
	"DSA_WITH_SHA224",
	"DSA_WITH_SHA256",
	"DSA_WITH_SHA384",
	"DSA_WITH_SHA512",
	"ED25519",
	"ED448",
	"GOSTR3411_94_WITH_3410_2001",
	"GOSTR3410_2012_WITH_3411_2012_256",
	"GOSTR3410_2012_WITH_3411_2012_512",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SignatureAlgorithm) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SignatureAlgorithm {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SignatureAlgorithm, v)
	}
	*j = SignatureAlgorithm(v)
	return nil
}
