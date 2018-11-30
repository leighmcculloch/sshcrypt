# sshcrypt

Use the keys in the ssh-agent to encrypt data symmetrically, and sign data.

Works with keys that are being forwarded to a remote host using SSH Agent Forwarding. These applications only use the SSH keys for signing. Encryption is performed by signing a randomly generated challenge and using the signature as the encryption key.

Each encryption uses a random nonce. Each encryption with a signature uses a random challenge that is signed to generate the input to scrypt, along with a randomly generated salt, for generating the key. The random nonce, random challenge, and random salt are stored alongside the encrypted data in the clear. Uses [Golang's implementation of NaCl secret box](https://godoc.org/golang.org/x/crypto/nacl/secretbox).

**The status of this project is experimental.** See [LICENSE](LICENSE).

## Install

```
go get 4d63.com/sshcrypt/...
```

## Example

This example encrypts a file using an intermediary key, that is encrypted with an SSH Agent signature.

Prepare a key, that is encrypted with your SSH Agent, saving it in `key.enc`. Use [randstr](@leighmcculloch/randstr) or a similar application to randomly generate the key.

```
$ randstr | sshcrypt agent-encrypt > key.enc
```

Encrypt a file with the key.

```
$ sshcrypt encrypt -p "$(sshcrypt agent-decrypt < key.enc)" < secrets.yaml > secrets.yaml.enc
```

Decrypt a file with the key.

```
$ sshcrypt decrypt -p "$(sshcrypt agent-decrypt < key.enc)" < secrets.yaml.enc > secrets.yaml
```

You can then store the encrypted `key.enc`, and any files encrypted with it in git, and add the non-encrypted filenames to `.gitignore` to ensure you do not accidentally commit them..

To give other people access to the encrypted files you'll need to meet them on a shared machine, decrypt `key.enc` > `key`, and have them add their encrypted version of the key to `key.enc`.

```
you  $ sshcrypt agent-decrypt < key.enc > key
them $ sshcrypt agent-encrypt < key >> key.enc
you  $ rm key
```

## Usage

### Encrypt with Agent

Encrypts the data symmetrically by signing a random challenge and using the resulting signature as the encryption key. The signature is generated using the keys available in the SSH Agent. Outputs one line of cipher text per key in the SSH Agent.

```
[data] | sshcrypt agent-encrypt
```

```
$ echo hello world | sshcrypt agent-encrypt
[jibberish]
```

### Decrypt with Agent

Decrypts the data by signing the challenge stored alongside the encrypted data and using the resulting signature as the decryption key. The signature is generated using the keys available in the SSH Agent. Outputs the clear text. If multiple encrypted text are passed separated by new lines, each will be attempted and the first available that can be decrypted will be output in clear text. The outpt of sshagentencrypt can be passed directly to sshagentdecrypt.

```
[data] | sshcrypt agent-decrypt
```

```
$ echo [jibberish] | sshcrypt agent-decrypt
hello world
```

### Encrypt with Password

Encrypts the data using the password, along with a random nonce, and random salt stored inside the resulting encrypted data, using Golang's implementation of NaCl secret box.

```
[data] | sshcrypt encrypt -p [password]
```

```
$ echo hello world | sshcrypt decrypt -p "1O685Q7I4^3c"
[jibberish]
```

### Decrypt with Password

Decrypts the data using the password, along with the random nonce, and random salt stored inside the encrypted data, using Golang's implementation of NaCl secret box.

```
[data] | sshcrypt agent-decrypt -p [password]
```

```
$ echo [jibberish] | sshcrypt decrypt -p "1O685Q7I4^3c"
hello world
```

### Sign with Agent

Signs stdin using the keys in the SSH Agent. The output is the signature.

```
$ echo hello world | sshcrypt agent-sign
ssh-rsa Zm6uX...MQjJTKG81VWCP24g==
```

Note: More than one signature will be returned if the SSH Agent contains multiple keys. Signatures wll be separated by new-lines.

### Verify with Public Key

Verifies the signature is a valid signature of stdin using the public key. The output is `Success` or `Failed`, with the exit code reflecting same.

```
[data] | sshcrypt verify -s [signature] -k [public-key]
```

```
$ echo hello world | sshcrypt verify -s "ssh-rsa Zm6uX...MQjJTKG81VWCP24g==" -k "$(curl -s https://github.com/leighmcculloch.keys)"
Success
```

Note: More than one signature and key can be provided by separating them by new-lines. If more than one signature is provided, success will be returned if at least one signature can be verified with one public key.

## Thanks

Thanks to [@ejcx](https://github.com/ejcx)'s [talk](https://twitter.com/ejcx_/status/732595370136494080) on Crypto in Golang.
