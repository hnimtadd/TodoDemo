# Todos app with mongodb backend and es search

```How to run locally```:
Install elasticsearch
Install mongodb
Add elasticsearch and mongodb url into ```.env``` file
Generate tls keys and put in app folder.
Example:
```
# Key considerations for algorithm "RSA" ≥ 2048-bit
openssl genrsa -out localhost.key 2048

# Key considerations for algorithm "ECDSA" (X25519 || ≥ secp384r1)
# https://safecurves.cr.yp.to/
# List ECDSA the supported curves (openssl ecparam -list_curves)
openssl ecparam -genkey -name secp384r1 -out localhost.key
openssl req -new -x509 -sha256 -key localhost.key -out localhost.crt -days 3650
```
Run ```go run main.go```, server will serve with port ```PORT```  from .env file
