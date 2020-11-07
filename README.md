# go-DES

Implementation of DES algorithm in golang.

## How to use ?
**Example:**
```
$ go-DES -e "I am a message to encrypt": Will encrypt the message in binary
$ go-DES -e -b "I am a message to encrypt": Will encrypt the message in binary
$ go-DES -e -h "I am a message to encrypt": Will encrypt the message in hexadecimal
$ go-DES -e -b64 "I am a message to encrypt": Will encrypt the message in base64
$ go-DES -d -h "762336c2b5360cc38751": Will decrypt the hexadecimal message
```

**Arguments:**
--help: Print this helper
-e: Encrypt a message
-d: Decrypt a message
-h: Use hexadecimal
-b64: Use base64
-b: Use binary

### sources
The DES algorithm in-use here is describe here: [french](https://www.commentcamarche.net/contents/204-introduction-au-chiffrement-avec-des), [english](https://ccm.net/contents/134-introduction-to-encryption-with-des)
