# x509
A simple program to print certificate information.

# Usage
The `x509` tool is takes only a hostname / url as an argument
and will make a dial that hostname with tls.Dial


```
e % ./x509 ejj.io
{
     "IsCA": false,
     "Version": 3,
     "SerialNumber": 15054454955724096237524256656952086492,
     "Issuer": "USCloudflare, Inc.  Cloudflare Inc ECC CA-3",
     "Subject": "USCloudflare, Inc.  ejj.io",
     "NotBefore": "2022-05-02T00:00:00Z",
     "NotAfter": "2023-05-02T23:59:59Z",
     "KeyUsage": 1,
     "DNSNames": [
         "ejj.io",
         "*.ejj.io"
     ],
     "EmailAddresses": null,
     "IPAddresses": null,
     "URIs": null,
     "PermittedDNSDomainsCritical": false,
     "PermittedDNSDomains": null,
     "ExcludedDNSDomains": null,
     "PermittedIPRanges": null,
     "ExcludedIPRanges": null,
     "PermittedEmailAddresses": null,
     "ExcludedEmailAddresses": null,
     "PermittedURIDomains": null,
     "ExcludedURIDomains": null
 }
```
