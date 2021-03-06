# aws-lambda-go-proxy

[![Build Status](https://travis-ci.org/mtojek/aws-lambda-go-proxy.svg?branch=master)](https://travis-ci.org/mtojek/aws-lambda-go-proxy)

Status: **Done**

![Gopher and Lambda](https://cdn-images-1.medium.com/max/400/1*SncdHqDPypbypdx2QmP6iQ.jpeg)

*Source: medium.com*


Briefly, run **Lambda** Go function on your machine and forget about redeployments!

Run **proxy application** as Lambda function and **route all events** (TCP traffic) to a different, **even local** Go application. Forget about long hours of blindly debugging the function code, docker images and SAM models. Pass all **incoming, real Lambda requests** to the Go application that you're **currently developing**.

No need to prepare *Test events* or wrap *http.Request* magically. 

## Quickstart

*In this case I will show you how to process requests sent to API Gateway.*

Build **aws-lambda-go-proxy** package:

```bash
$ go get github.com/mtojek/aws-lambda-go-proxy
```

Go to the source directory and run *make* to create a deployment artifact:

```bash
$ make
go get -v ./...
GOOS=linux go build -o main
zip deployment.zip main
updating: main (deflated 65%)
rm main
```

In the mean time, let's build and run a sample Go application (for educational reasons) - [Hello World](https://github.com/aws-samples/lambda-go-samples/blob/master/main.go):

```bash
$ go get github.com/aws-samples/lambda-go-samples
$ _LAMBDA_SERVER_PORT=9999 lambda-go-samples
```

This will expose the *lambda-go-samples* behind the port 9999 (TCP). If you're curious about this sample, feel free to read the [blog post](https://aws.amazon.com/blogs/compute/announcing-go-support-for-aws-lambda/).

Now, it's time to deploy **aws-lambda-go-proxy** on AWS Lambda. I hope you can figure this out on your own (if not, please read first the mentioned blog post). Few hints:

* *deployment.zip* will be created if you run *make* command, it contains the binary of *aws-lambda-go-proxy*
* set *LAMBDA_DEBUG_PROXY* environment variable to the remote address of the target application, e.g.: *0.tcp.ngrok.io:12345*
* configure *API Gateway* trigger (API name: *lambda-go*, Deployment stage: *prod*, Security: *open* remember, this is not secure and only for educational reasons)

As you may see I set the *LAMBDA_DEBUG_PROXY* variable to *ngrok* address. It is due a NAT that's in front of me. With [ngrok](https://ngrok.com/) you can open a direct network tunnel to a specified exposed local port:

```bash
$ ngrok tcp 9999
```

Here you can find more details regarding [TCP tunnels](https://ngrok.com/docs#tcp) in ngrok.

Once you have API Gateway API deployed and you a URL similar to the following:

https://hefunef32.execute-api.us-west-2.amazonaws.com/prod/main

Feel free to execute *curl* command:

```bash
$ curl -XPOST -d "horse" "https://hefunef32.execute-api.us-west-2.amazonaws.com/prod/main"
Hello horse
```

Easter Egg: https://twitter.com/bbctwo/status/549296800709234688 :)

*Notice: It may be necessary to hit curl few times to let Lambda establish a connection with your machine. Internal server error should appear.*

You should also notice some log entries appearing in your local console:

```bash
2018/01/21 04:12:01 Processing Lambda request ae976f75-e0af-478c-ad06-3d460fe1881d
2018/01/21 04:12:05 Processing Lambda request be6fd715-5cfd-4ee7-8f67-2a20c0a98af2
2018/01/21 04:14:14 Processing Lambda request a7288c78-596e-4772-a4c4-2d1dc9622410
```

## Behind the scenes

**aws-lambda-go-proxy** runs a TCP proxy to the target host defined via environment variable *LAMBDA_DEBUG_PROXY*. Whole Lambda traffic is passed to that address, what actually allows for processing Lambda events on a different machine, e.g. a developer host (debugging reasons).

As the Lambda Go runtime uses *net/rpc* internals to communicate, you can preview whole communication with *tcpdump*. Here are some sniffs:

```bash
InvokeRequest........Payload..........
ClientContext.tId.....XAmznTraceId.....Deadline......InvokedFunctionArn.....CognitoIdentityId.....CognitoIdentityPoolId....
...;......InvokeRequest_Timestamp........Seconds.....Nanos.......[......{"resource":"/main","path":"/main","httpMethod":"POST","headers":{"Accept":"*/*","CloudFront-Forwarded-Proto":"https","CloudFront-Is-Desktop-Viewer":"true","CloudFront-Is-Mobile-Viewer":"false","CloudFront-Is-SmartTV-Viewer":"false","CloudFront-Is-Tablet-Viewer":"false","CloudFront-Viewer-Country":"PL","Content-Type":"application/x-www-form-urlencoded","Host":"<redacted>.execute-api.us-west-2.amazonaws.com","User-Agent":"curl/7.30.0","Via":"1.1 <redacted>.cloudfront.net (CloudFront)","X-Amz-Cf-Id":"<redacted>","X-Amzn-Trace-Id":"<redacted>","X-Forwarded-For":"<redacted>","X-Forwarded-Port":"443","X-Forwarded-Proto":"https"},"queryStringParameters":null,"pathParameters":null,"stageVariables":null,"requestContext":{"requestTime":"<redacted>","path":"/prod/main","accountId":"<redacted>","protocol":"HTTP/1.1","resourceId":"<redacted>","stage":"prod","requestTimeEpoch":<redacted>,"requestId":"<redacted>","identity":{"cognitoIdentityPoolId":null,"accountId":null,"cognitoIdentityId":null,"caller":null,"sourceIp":"<redacted>","accessKey":null,"cognitoAuthenticationType":null,"cognitoAuthenticationProvider":null,"userArn":null,"userAgent":"curl/7.30.0","user":null},"resourcePath":"/main","httpMethod":"POST","apiId":"<redacted>"},"body":"horse","isBase64Encoded":false}.$<redacted>.JRoot=<redacted>;Sampled=0.....@>..?]....3arn:aws:lambda:us-west-2:<redacted>:function:main.
```

```bash
E.....@.@...........'...-\h...!..."..{.....
J.|.J.|......Function.Invoke...;...6{"statusCode":200,"headers":null,"body":"Hello horse"}
```

## Disclaimer

The idea of passing TCP traffic to remotely running applications has been described and used to simplify debugging process of Go Lambda functions to prevent redeployments. I hope it will attract more users to [AWS Lambda](https://aws.amazon.com/lambda/) based solutions, including Github community.

## License

See [LICENSE](LICENSE) for the details.
