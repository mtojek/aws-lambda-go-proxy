# aws-lambda-go-proxy

[![Build Status](https://travis-ci.org/mtojek/aws-lambda-go-proxy.svg?branch=master)](https://travis-ci.org/mtojek/aws-lambda-go-proxy)

Status: **Done**

Run a **proxy application** as Lambda function and **route all events** (TCP traffic) to a different, **even local** Go application. Forget about long hours of blindly debugging the function code, docker images and SAM models. Pass all **incoming, real Lambda requests** to the Go application that you're **currently developing**.
No need to prepare *Test events* or wrap *http.Request* magically.

## License

See [LICENSE](LICENSE) for the details.
