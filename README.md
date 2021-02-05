# bZapp

bZapp is a slack app to facilitate a standup meeting in a work from home or remote workplace. It simulates a whiteboard with today and tomorrow's meetings or events and lists teams goals. 

![bZapp whiteboard](docs/bZapp%20board.jpg?raw=true "bZapp whiteboard") 


## Quick Start

This project uses :

- [AWS CLI](https://aws.amazon.com/cli/)
- [AWS SAM CLI](https://github.com/awslabs/aws-sam-cli)
- [Docker CE](https://www.docker.com/community-edition)
- [Slack-go](https://github.com/slack-go/slack)
- [Go 1.15](https://golang.org/)
- [watchexec](https://github.com/mattgreen/watchexec)

Install the CLI tools and Docker CE

```console
$ brew install awscli go node python@2 watchexec
$ pip2 install aws-sam-cli
$ open https://store.docker.com/search?type=edition&offering=community
```

<details>
<summary>We may want to upgrade existing tools...</summary>
&nbsp;

```console
$ brew upgrade awscli go node python@2 watchexec
$ pip2 install --upgrade aws-sam-cli
```
</details>

<details>
<summary>We may also want to configure the AWS CLI with IAM keys to develop and deploy our application...</summary>
&nbsp;

Follow the [Creating an IAM User in Your AWS Account](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_create.html) doc to create a IAM user with programmatic access. Call the user `admin` and attach the "Administrator Access" policy for now.

Then configure the CLI. Here we are creating a new profile that we can switch to with `export AWS_PROFILE=bZapp`. This will help us isolate our experiments from other AWS work.

Configure an AWS profile with keys and switch to the profile:

```console
$ aws configure --profile bZapp
AWS Access Key ID [None]: AKIA................
AWS Secret Access Key [None]: PQN4CWZXXbJEgnrom2fP0Z+z................
Default region name [None]: us-east-1
Default output format [None]: json

$ export AWS_PROFILE=bZapp
$ aws iam get-user
{
    "User": {
        "Path": "/",
        "UserName": "Admin",
        "UserId": "**************",
        "Arn": "arn:aws:iam::*************:user/Admin",
        "CreateDate": "20**-**-88T23:11:39+00:00"
    }
}
```
</details>

### Get the App

We start by getting and testing the `github.com/dctid/bZapp`.

```console
$ git clone https://github.com/dctid/bZapp.git
$ cd ~/bZapp

$ make test
go test -v ./...
go: finding github.com/aws/aws-xray-sdk-go v1.0.0-rc.8
go: finding github.com/aws/aws-lambda-go v1.6.0
go: finding github.com/aws/aws-sdk-go v1.15.49
...
=== RUN   TestInteraction/open_edit_events
--- PASS: TestInteraction/open_edit_events (0.00s)
...
ok     github.com/dctid/bZapp      0.014s
PASS
```

This gives us confidence in our Go environment.

### Develop the App

We can then build the app and start a development server:

```console
$ make app
cd ./handlers/dashboard && GOOS=linux go build...
2018/02/25 08:03:12 Connected to Docker 1.35
2018/02/16 07:40:32 Fetching lambci/lambda:go1.x image for go1.x runtime...

Mounting handler (go1.x) at http://127.0.0.1:3000/users/{id} [DELETE]
Mounting handler (go1.x) at http://127.0.0.1:3000/users/{id} [PUT]
Mounting handler (go1.x) at http://127.0.0.1:3000/users/{id} [GET]
Mounting handler (go1.x) at http://127.0.0.1:3000/ [GET]
Mounting handler (go1.x) at http://127.0.0.1:3000/users [POST]
```

Now we can access our HTTP functions on port 3000:

```console
$ curl http://localhost:3000
<html><body><h1>bZapp dashboard</h1></body></html>
```

We can also invoke a function directly:

```
$ echo '{}' | sam local invoke DashboardFunction
...
START RequestId: b69240d1-66a6-1dfd-89d1-296bdee4081c Version: $LATEST
END RequestId: b69240d1-66a6-1dfd-89d1-296bdee4081c
REPORT RequestId: b69240d1-66a6-1dfd-89d1-296bdee4081c  Init Duration: 433.77 ms        Duration: 4.00 ms       Billed Duration: 100 ms Memory Size: 128 MB     Max Memory Used: 30 MB  

{"statusCode":200,"headers":{"Content-Type":"text/html"},"body":"\u003chtml\u003e\u003cbody\u003e\u003ch1\u003ebZapp dashboard\u003c/h1\u003e\u003c/body\u003e\u003c/html\u003e\n"}
```

Note: if you see `No AWS credentials found. Missing credentials may lead to slow startup...`, review `aws configure list` and your `AWS_PROFILE` env var.

This gives us confidence in our development environment.

The local environment also starts a [docker image](docker-compose.yml) that creates a local [DynamDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Introduction.html) instance.

To shutdown this DynamoDB instance when done developing locally:

```console
$ make dynamo-stop

docker-compose down
Removing dynamodb ... done
Removing network bzapp_default
```  


### Deploy the App

Now we can package and deploy the app:

```console
$ make deploy
make_bucket: pkgs-*************-us-east-1
Uploading to ********************  3018471 / 3018471.0  (100.00%)
Waiting for changeset to be created
Waiting for stack create/update to complete
Successfully created/updated stack - bZapp

ApiUrl	https://********.execute-api.us-east-1.amazonaws.com/Prod
```

Now we can access our HTTP functions on AWS:

```console
$ curl https://********.execute-api.us-east-1.amazonaws.com/Prod
<html><body><h1>bZapp dashboard</h1></body></html>
```

Look at that speedy 11 ms duration! Go is faster than the minimum billing duration of 100 ms.

This gives us confidence in our production environment.

## Configuring Slack

TBD

## Contributing

Find a bug or see a way to improve the project? [Open an issue](https://github.com/dctid/bZapp/issues).


