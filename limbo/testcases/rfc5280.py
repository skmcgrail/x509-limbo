"""
RFC5280 profile tests.
"""

from datetime import datetime

from cryptography import x509
from cryptography.hazmat.primitives.asymmetric import rsa

from limbo.assets import ee_cert, ext
from limbo.testcases._core import Builder, testcase

# TODO: Intentionally mis-matching algorithm fields.


@testcase
def empty_issuer(builder: Builder) -> None:
    """
    Produces the following **invalid** chain:

    ```
    root -> EE
    ```

    This chain is invalid solely because of the EE cert's construction:
    it has an empty issuer name, which isn't allowed under the RFC 5280 profile.
    """
    # Intentionally empty issuer name.
    issuer = x509.Name([])
    subject = x509.Name.from_rfc4514_string("CN=empty-issuer")
    root = builder.root_ca(issuer=issuer, subject=subject)
    leaf = ee_cert(root)

    builder = builder.client_validation()
    builder = builder.trusted_certs(root).peer_certificate(leaf).fails()


@testcase
def unknown_critical_extension_ee(builder: Builder) -> None:
    """
    Produces the following **invalid** chain:

    ```
    root -> EE
    ```

    The EE cert has an extension, 1.3.6.1.4.1.55738.666.1, that no implementation
    should recognize. As this unrecognized extension is marked as critical, a
    chain should not be built with this EE.
    """
    root = builder.root_ca()
    leaf = ee_cert(
        root,
        extra_extension=ext(
            x509.UnrecognizedExtension(x509.ObjectIdentifier("1.3.6.1.4.1.55738.666.1"), b""),
            critical=True,
        ),
    )

    builder = builder.client_validation()
    builder = builder.trusted_certs(root).peer_certificate(leaf).fails()


@testcase
def unknown_critical_extension_root(builder: Builder) -> None:
    """
    Produces the following **invalid** chain:

    ```
    root -> EE
    ```

    The root has an extension, 1.3.6.1.4.1.55738.666.1, that no implementation
    should recognize. As this unrecognized extension is marked as critical, a
    chain should not be built with this root.
    """

    root = builder.root_ca(
        extra_extension=ext(
            x509.UnrecognizedExtension(x509.ObjectIdentifier("1.3.6.1.4.1.55738.666.1"), b""),
            critical=True,
        )
    )
    leaf = ee_cert(root)

    builder = builder.client_validation()
    builder = builder.trusted_certs(root).peer_certificate(leaf).fails()


@testcase
def unknown_critical_extension_intermediate(builder: Builder) -> None:
    """
    Produces the following **invalid** chain:

    ```
    root -> intermediate (pathlen:0) -> EE
    ```

    The intermediate has an extension, 1.3.6.1.4.1.55738.666.1, that no implementation
    should recognize. As this unrecognized extension is marked as critical, a
    chain should not be built with this intermediate.
    """

    root = builder.root_ca()
    intermediate = builder.intermediate_ca(
        root,
        0,
        extra_extension=ext(
            x509.UnrecognizedExtension(x509.ObjectIdentifier("1.3.6.1.4.1.55738.666.1"), b""),
            critical=True,
        ),
    )
    leaf = ee_cert(intermediate)

    builder = builder.client_validation()
    builder = (
        builder.trusted_certs(root)
        .untrusted_intermediates(intermediate)
        .peer_certificate(leaf)
        .fails()
    )


# TODO: Empty serial number, overlength serial number.


@testcase
def critical_aki(builder: Builder) -> None:
    """
    Produces the following **invalid** chain:

    ```
    root -> EE
    ```

    The root cert has an AKI extension marked as critical, which is disallowed
    under the [RFC 5280 profile]:

    > Conforming CAs MUST mark this extension as non-critical.

    [RFC 5280 profile]: https://www.rfc-editor.org/rfc/rfc5280#section-4.2.1.1
    """
    key = rsa.generate_private_key(public_exponent=65537, key_size=4096)
    root = builder.root_ca(
        key=key,
        aki=ext(
            x509.AuthorityKeyIdentifier.from_issuer_public_key(key.public_key()), critical=True
        ),
    )
    leaf = ee_cert(root)

    builder = builder.client_validation()
    builder = builder.trusted_certs(root).peer_certificate(leaf).fails()


def missing_aki(builder: Builder) -> None:
    """
    Produces the following **invalid** chain:

    ```
    root -> EE
    ```

    The root cert is missing the AKI extension, which is disallowed under the
    [RFC 5280 profile].

    > To facilitate certification path construction, this extension MUST
    > appear in all conforming CA certificates, that is, all certificates
    > including the basic constraints extension (Section 4.2.1.9) where the
    > value of cA is TRUE.

    Note: for roots, the SKI should be the same value as the AKI, therefore,
    this extension isn't strictly necessary, although required by the RFC.

    [RFC 5280 profile]: https://www.rfc-editor.org/rfc/rfc5280#section-4.2.1.2
    """
    root = builder.root_ca(aki=None)
    leaf = ee_cert(root)

    builder = builder.client_validation()
    builder = builder.trusted_certs(root).peer_certificate(leaf).fails()


@testcase
def critical_ski(builder: Builder) -> None:
    """
    Produces the following **invalid** chain:

    ```
    root -> EE
    ```

    The root cert has an SKI extension marked as critical, which is disallowed
    under the [RFC 5280 profile].

    > Conforming CAs MUST mark this extension as non-critical.

    [RFC 5280 profile]: https://www.rfc-editor.org/rfc/rfc5280#section-4.2.1.2
    """
    key = rsa.generate_private_key(public_exponent=65537, key_size=4096)
    root = builder.root_ca(
        key=key,
        ski=ext(x509.SubjectKeyIdentifier.from_public_key(key.public_key()), critical=True),
    )
    leaf = ee_cert(root)

    builder = builder.client_validation()
    builder = builder.trusted_certs(root).peer_certificate(leaf).fails()


@testcase
def missing_ski(builder: Builder) -> None:
    """
    Produces the following **invalid** chain:

    ```
    root -> EE
    ```

    The root cert is missing the SKI extension, which is disallowed under the
    [RFC 5280 profile].

    > To facilitate certification path construction, this extension MUST
    > appear in all conforming CA certificates, that is, all certificates
    > including the basic constraints extension (Section 4.2.1.9) where the
    > value of cA is TRUE.

    Note: for roots, the SKI should be the same value as the AKI, therefore,
    this extension isn't strictly necessary, although required by the RFC.

    [RFC 5280 profile]: https://www.rfc-editor.org/rfc/rfc5280#section-4.2.1.2
    """
    root = builder.root_ca(ski=None)
    leaf = ee_cert(root)

    builder = builder.client_validation()
    builder = builder.trusted_certs(root).peer_certificate(leaf).fails()


@testcase
def multiple_chains_expired_intermediate(builder: Builder) -> None:
    """
    Produces the following chain:

    root 2 -> intermediate (expired) -> root -> EE

    Both roots are trusted. A chain should be built successfully, disregarding
    the expired intermediate certificate and the second root. This scenario is
    known as the "chain of pain"; for further reference, see
    https://www.agwa.name/blog/post/fixing_the_addtrust_root_expiration.
    """
    root = builder.root_ca()
    root_two = builder.root_ca(issuer=x509.Name.from_rfc4514_string("CN=x509-limbo-root-2"))
    ski = x509.SubjectKeyIdentifier.from_public_key(root.key.public_key())  # type: ignore[arg-type]
    expired_intermediate = builder.intermediate_ca(
        root_two,
        1,
        subject=root.cert.subject,
        not_after=datetime.fromisoformat("1988-11-25T00:00:00Z"),
        key=root.key,
        ski=ski,
    )
    leaf = ee_cert(root)

    builder = builder.client_validation()
    builder.trusted_certs(root, root_two).untrusted_intermediates(
        expired_intermediate
    ).peer_certificate(leaf).succeeds()


@testcase
def chain_untrusted_root(builder: Builder) -> None:
    """
    Produces the following chain:

    root (untrusted) -> intermediate -> EE

    The root is not in the trusted set, thus no chain should be built.
    """
    root = builder.root_ca()
    intermediate = builder.intermediate_ca(root, 0)
    leaf = ee_cert(intermediate)

    builder = builder.client_validation()
    builder.trusted_certs().untrusted_intermediates(root, intermediate).peer_certificate(
        leaf
    ).fails()