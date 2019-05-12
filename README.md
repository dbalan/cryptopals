# Cryptopals solutions

[![builds.sr.ht status](https://builds.sr.ht/~dbalan/cryptopals.svg)](https://builds.sr.ht/~dbalan/cryptopals?)



Solutions to [cryptopals](https://cryptopals.com) challenges in [Go](https://golang.org).

## Running
Solutions are implemented as a testcases.

```bash
cd cryptopals
go test -v ./...
```

## Solutions
### Set 1

1. Convert hex to base64: https://github.com/dbalan/cryptopals/blob/master/set1/ch1_test.go#L16
1. Fixed XOR: https://github.com/dbalan/cryptopals/blob/master/set1/ch2.go
1. Single-byte XOR cipher: https://github.com/dbalan/cryptopals/blob/master/set1/ch3.go#L62
1. Detect single-character XOR: https://github.com/dbalan/cryptopals/blob/master/set1/ch4_test.go#L9
1. Implement repeating-key XOR: https://github.com/dbalan/cryptopals/blob/master/set1/ch5_test.go#L19
1. Break repeating-key XOR: https://github.com/dbalan/cryptopals/blob/master/set1/ch6_test.go#L41
1. AES in ECB mode: https://github.com/dbalan/cryptopals/blob/master/set1/ch7_test.go#L11
1. Detect AES in ECB mode: https://github.com/dbalan/cryptopals/blob/master/set1/ch8_test.go#L11

### Set 2

1. Implement PKCS#7 padding: https://github.com/dbalan/cryptopals/blob/master/set2/pkcs.go
1. Implement CBC mode: https://github.com/dbalan/cryptopals/blob/master/set2/aescbc.go
1. An ECB/CBC detection oracle: https://github.com/dbalan/cryptopals/blob/master/set2/aes_detect_mode.go#L7
1. Byte-at-a-time ECB decryption (Simple): https://github.com/dbalan/cryptopals/blob/master/set2/aes_ecb_attack1.go
1. ECB cut-and-paste: https://github.com/dbalan/cryptopals/blob/master/set2/ecbcutpaste_test.go
1. Byte-at-a-time ECB decryption (Harder): https://github.com/dbalan/cryptopals/blob/master/set2/aes_ecb_attack2.go
1. PKCS#7 padding validation: https://github.com/dbalan/cryptopals/blob/master/set2/pkcs.go
1. CBC bitflipping attacks: https://github.com/dbalan/cryptopals/blob/master/set2/cbc_bitflipping_test.go

### Set 3

1. The CBC padding oracle: https://github.com/dbalan/cryptopals/blob/master/set3/padding_oracle.go
1. Implement CTR, the stream cipher mode: https://github.com/dbalan/cryptopals/blob/master/set3/aesctr.go
1. Break fixed-nonce CTR mode using substitutions: TODO
1. Break fixed-nonce CTR statistically: https://github.com/dbalan/cryptopals/blob/master/set3/aesctr_stat.go
1. Implement the MT19937 Mersenne Twister RNG: https://github.com/dbalan/cryptopals/blob/master/set3/mt19937.go
1. Crack an MT19937 seed: https://github.com/dbalan/cryptopals/blob/master/set3/mt_stream_cipher.go
1. Clone an MT19937 RNG from its output: https://github.com/dbalan/cryptopals/blob/master/set3/mt_clone.go
1. Create the MT19937 stream cipher and break it: https://github.com/dbalan/cryptopals/blob/master/set3/mt_stream_cipher_test.go

### Set 4

1. Break "random access read/write" AES CTR: https://github.com/dbalan/cryptopals/blob/master/set4/aesctr_attack.go
1. CTR bitflipping: https://github.com/dbalan/cryptopals/blob/master/set4/ctrbitflipping.go
1. Recover the key from CBC with IV=Key: https://github.com/dbalan/cryptopals/blob/master/set4/ivattack.go
1. Implement a SHA-1 keyed MAC: https://github.com/dbalan/cryptopals/blob/master/sha/sha.go
1. Break a SHA-1 keyed MAC using length extension: https://github.com/dbalan/cryptopals/blob/master/set4/sha_len_ext.go
1. Break an MD4 keyed MAC using length extension: TODO (MD construction similar to sha1)
1. Implement and break HMAC-SHA1 with an artificial timing leak: TODO
1. Break HMAC-SHA1 with a slightly less artificial timing leak: TODO

### Set 5

1. Implement Diffie-Hellman: https://github.com/dbalan/cryptopals/blob/master/set5/dh.go
1. Implement a MITM key-fixing attack on Diffie-Hellman with parameter injection: https://github.com/dbalan/cryptopals/blob/master/set5/dh_mitm_test.go
1. Implement DH with negotiated groups, and break with malicious "g" parameters: https://github.com/dbalan/cryptopals/blob/master/set5/dh_mitm_primes.go
1. Implement Secure Remote Password (SRP)
: https://github.com/dbalan/cryptopals/blob/master/set5/simple_srp.go
1. Break SRP with a zero key: https://github.com/dbalan/cryptopals/blob/master/set5/srp_test.go
1. Offline dictionary attack on simplified SRP: https://github.com/dbalan/cryptopals/blob/master/set5/evil_ssrp.go
1. Implement RSA: https://github.com/dbalan/cryptopals/blob/master/rsa/rsa.go
1. Implement an E=3 RSA Broadcast attack: https://github.com/dbalan/cryptopals/blob/master/set5/broadcast_rsa_attack_test.go

### Set 6
1. Implement unpadded message recovery oracle: https://github.com/dbalan/cryptopals/blob/master/set6/rsa_recovery_test.go
1. Bleichenbacher's e=3 RSA Attack : https://github.com/dbalan/cryptopals/blob/master/set6/rsa_sign_test.go
1. DSA key recovery from nonce: https://github.com/dbalan/cryptopals/blob/master/set6/dsa_key_recovery.go
1. DSA nonce recovery from repeated nonce: https://github.com/dbalan/cryptopals/blob/master/set6/dsa_repeated_nonce.go
1. DSA parameter tampering: https://github.com/dbalan/cryptopals/blob/master/set6/dsa_parameter_tampering.go
1. RSA parity oracle: https://github.com/dbalan/cryptopals/blob/master/set6/rsa_parity_oracle.go
1. Bleichenbacher's PKCS 1.5 Padding Oracle (Simple Case)
1. Bleichenbacher's PKCS 1.5 Padding Oracle (Complete Case
