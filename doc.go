/*

Biscuit helps you use AWS KMS to manage infrastructure secrets.

Biscuit uses standard symmetric crypto algorithms (NaCl secretbox or
AES-GCM-256) and ephemeral data keys provided by KMS (using "key
wrapping" or "envelope encryption"). The ciphertexts are stored in a
YAML file that is safe to distribute in source control, deployments,
CI systems, etc.

Biscuit provides a key-value store interface (put and get), and tooling for
grant and policy management across AWS regions.

Quickstart:

	biscuit kms get-caller-identity
	biscuit kms init -f stash.yml
	biscuit put -f stash.yml launch_codes 0000
	biscuit get -f stash.yml launch_codes

*/
package main

