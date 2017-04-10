wercker-build-trigger
===

Lambda job with [Apex – Serverless Infrastructure](http://apex.run/) to trigger wercker build.

This realizes [Scheduled wercker build by scheduled lambda function](http://docs.aws.amazon.com/lambda/latest/dg/tutorial-scheduled-events-schedule-expressions.html).

## Usage
### 1. Initialize
#### 1.1 Install
See [Apex installation description](http://apex.run/#installation).

#### 1.2 AWS configuration
See [Apex AWS credentials description](http://apex.run/#aws-credentials).


#### 1.3 init

```sh
$ apex init
```

And input project settings.

Maybe you don't want default function `hello`, remove it with `$ rm -rf functions/hello`.

### 2. Deploy
```sh
$ apex deploy
```


### 2. Lambda settings
Configure to trigger Lambda function.

#### 2.1 Environment variables
Configure environment variables in `./project.json`.

|Name||
|---|---|
|WERCKER_TOKEN|Required|
|WERCKER_PIPELINE_ID|Required|
|WERCKER_MESSAGE|Optional|

See [Authentication](http://devcenter.wercker.com/docs/api/getting-started/authentication#using-a-token) and [Run](http://devcenter.wercker.com/docs/api/endpoints/runs#trigger-a-run) in wercker docs for more detail.

#### 2.2 Trigger settings
See [Examples of How to Use AWS Lambda - AWS Lambda](http://docs.aws.amazon.com/lambda/latest/dg/use-cases.html) and configure lambda Trigger settings.


## License
MIT
