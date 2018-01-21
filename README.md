# aws-lambda-go-proxy

[![Build Status](https://travis-ci.org/mtojek/aws-lambda-go-proxy.svg?branch=master)](https://travis-ci.org/mtojek/aws-lambda-go-proxy)

Status: **Done**

Briefly, run **Lambda** Go function on your machine and forget about redeployments!

Run **proxy application** as Lambda function and **route all events** (TCP traffic) to a different, **even local** Go application. Forget about long hours of blindly debugging the function code, docker images and SAM models. Pass all **incoming, real Lambda requests** to the Go application that you're **currently developing**.

No need to prepare *Test events* or wrap *http.Request* magically. 

## Disclaimer

The idea of passing TCP traffic to remotely running applications has been described and used to simplify debugging process of Go Lambda functions to prevent redeployments. I hope it will attract more users to [AWS Lambda](https://aws.amazon.com/lambda/) based solutions, including Github community.

## License

See [LICENSE](LICENSE) for the details.
