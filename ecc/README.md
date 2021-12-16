ROXE Elliptic Curve Cryptography Wrapper
=========================================

This is a simple wrapper for `btcec`, that handles the specificities
of the format for keys in ROXE.

It was crafted in reference to `roxejs-ecc`, `roxejs-keygen` and the C++
codebase of ROXE.IO Software.

This handles the `ROXE` prefix on public keys, manages the version and
checksums in public and private keys.
