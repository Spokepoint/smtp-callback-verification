# SMTP Callback Verification

A simple REST API service to verify email addresses. Using [SMTP Callback Verification](http://en.wikipedia.org/wiki/Callback_verification).

## Usage

Make a request.

```sh
curl -i "http://localhost/verify/?email=paul@gmail.com"
```

Which generates the following response.

```json
{
  "email": "paul@gmail.com",
  "isValid": true
}
```

### Running it locally

1. Build project with Go. `go build`
2. Spin up server locally. `PORT=8080 ./smtp-callback-verification` Note that server port is read from the environment.
3. Send a request. `curl -i "http://localhost:8080/verify/?email=paul@gmail.com"`

### Deploy to Heroku

1. Create a Heroku app with Go buildpack. `heroku create -b https://github.com/kr/heroku-buildpack-go.git`
2. Deploy this app to Heroku. `git push heroku master`
3. Send a request. `curl -i "http://<YOUR_APP>.herokuap.com/verify/?email=paul@gmail.com"`