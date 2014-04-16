# SMTP Callback Verification

This is an API to verify email addresses

## Usage

```sh
$ curl -i "http://smtp-callback-verification.herokuapp.com/verify/?email=paul@gmail.com"
```

Which generates the following response.

```json
{
  "email": "paul@gmail.com",
  "isValid": true
}
```
