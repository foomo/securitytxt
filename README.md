# SecurityTXT

Go package that provides a http.Handler and Middleware to serve a security txt file.

The contents of the file are controlled via environment variables.

```
EXPIRES=2025-03-26T11:00:00.000Z
COMMENT=this is a comment
CONTACT=mailto:security@org.com
ACKNOWLEDGMENT=https://example.com/halloffame
CANONICAL=https://example.com/canonical
ENCRYPTION=https://example.com/pgpkey.txt
HIRING=https://example.com/hiring
PREFERRED_LANGUAGES=en, de
POLICY=https://example.com/policy
CSAF=https://example.com/csaf
```