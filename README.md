# AWS Alexa Go

![Amazon Alexa](https://m.media-amazon.com/images/G/01/mobile-apps/dex/avs/docs/ux/branding/mark2._TTH_.png)

## How it works

See `example/main.go`

### How to upload your code to AWS using Lambda

Build a binary that runs on Linux and zip it up into a deployment package.

```
$ GOOS=linux go build -o lambda_handler
$ zip deployment.zip main
```

Assuming you have already installed `aws-cli`:

`NOTE:` \
function-name: name which handles the requests (in our example it is alexaDispatchIntentHandler) \ 
handler: name of the binary file (here it is lambda_handler) \
region: in order to run Alexa Skills with AWS Lambda you need to choose eu-west-1, us-east-1 or eu-west-1

```
aws lambda create-function \
  --region eu-west-1 \
  --function-name alexaDispatchIntentHandler \
  --memory 128 \
  --role arn:aws:iam::<account-id>:role/<role> \
  --runtime go1.x \
  --zip-file fileb://deployment.zip \
  --handler lambda_handler
```

Verify if your function was uploaded:

![Lambda](/images/lambda.png)
