# CWL - A CloudWatch Logs CLI

CWL helps you monitor CloudWatch logs on the command line.

The default AWS CLI displays logs in JSON format, and while that can
be processed with another tool like `jq`, it's a bit of a pain.

`cwl` tries to make working with AWS CloudWatch logs as easy as
possible by simplifying the parameters and choosing sane defaults.

## Configuring

CWL expects that you have your AWS credentials be configured in one of the standard ways - either with an AWS credentials file or with environment variables. You can read more about standard credential formats for AWS in the  [SDK documentation](http://docs.aws.amazon.com/sdk-for-go/latest/v1/developerguide/configuring-sdk.title.html).

### Credentials File

Create an INI file at `~/.aws/credentials`. Each section of the INI file represents a different credentials profile. The default
profile will be used by default. See the -profile parameter to choose a different profile.

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

## Parameters

## Contributing

## License
