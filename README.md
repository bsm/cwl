# CWL - A CloudWatch Logs CLI

CWL helps you monitor CloudWatch logs on the command line.

The default AWS CLI displays logs in JSON format, and while that can
be processed with another tool like `jq`, it's a bit of a pain.

`cwl` tries to make working with AWS CloudWatch logs as easy as
possible by simplifying the parameters and choosing sane defaults.

## Installing

For now, the best way to install the latest stable version is with `go get`:

```
go get http://gopkg.in/commondream/cwl.v1
```

You'll need a fully setup Golang toolchain to do so. See the [Golang Getting Started](https://golang.org/doc/install) guide for more information.

## Configuring

CWL expects that you have your AWS credentials be configured in one of the standard ways - either with an AWS credentials file or with environment variables. You can read more about standard credential formats for AWS in the  [SDK documentation](http://docs.aws.amazon.com/sdk-for-go/latest/v1/developerguide/configuring-sdk.title.html).

### Credentials File

Create an INI file at `~/.aws/credentials`. Each section of the INI file represents a different credentials profile. The default
profile will be used by default. See the [-profile parameter](#-profile-AWS_PROFILE) to choose a different profile.

Example credentials file:

```
[default]
aws_access_key_id = your_id
aws_secret_access_key = your_key

[another_profile]
aws_access_key_id = your_other_id
aws_secret_access_key = your_other_key
```

### Environment Variables

You can configure your credentials with the `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables.

```
AWS_ACCESS_KEY_ID=[your_id] AWS_SECRET_ACCESS_KEY=[your_key] cwl [parameters]
```

## Command-Line Parameters

### -profile AWS_PROFILE

The AWS credential profile to use from your credentials file. Defaults to the AWS default credentials either in the credentials file or in the environment.

### -region AWS_REGION

Required. Specifies the AWS region to use, for example `us-east-1`.

### -group GROUP_NAME

Required. Specifies the Log Group name to use.

### -streams STREAM1,STREAM2

Comma separated list of stream names. Defaults to loading all streams.

### -filter STRING

Filters the log events with the given filter. See [Filter and Pattern Syntax](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/FilterAndPatternSyntax.html) for more information about filter syntax.

### -start TIME

Specifies start time. Defaults to "1 minute ago". See [specifying times](#specifying-times) for more information about time parameter formats.

### -end TIME

Specifies end time. Defaults to the current time. See [specifying times](#specifying-times) for more information about time parameter formats.

### -fullStreamNames

By default CWL only displays the first 10 characters of each stream name. This parameter tells CWL to display the full stream name.

### Specifying Times

Times can be specified in one of three ways:

1. `now` specifies the current time.
2. `5 minutes ago` specifies a fuzzy time based relative to the current time. You can specify the interval in `seconds`, `minutes`, `hours`, or `days`. Both singular and plural forms of the interval work, so both `minute` and `minutes` are valid. The numeric portion of the date must be a numeral - so `5` works but `five` does not.
3. `2016-01-01T07:45:00Z` specifies RFC 3339 format.

## Questions

Have a question about how to use CWL? Submit an issue and I'll do my best to answer it in a timely manner.

## Contributing

To contribute to CWL, please fork the repository, make your change, and submit a pull request.

## License

Copyright (C) 2016 by Alan Johnson

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
