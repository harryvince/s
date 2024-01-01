# S

## What is S and why
S is a secrets manager built on top of AWS Parameter Store. 

You might be asking why? First thing is I just got sick of creating a .env
everytime I clone a repo etc.. also configuring variables in ci/cd ain't fun.
Having to keep track and so on.

I just wanted something that worked and sat on top of my own accounts. There
are other free tools online that acomplish a similar thing this is just my
attempt at an easy maintainable way of doing it.

[CLI](cli)

## Packages (WIP)
[Node](packages/node)

## Required IAM Permissions
Sample policy:
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "Permissions",
            "Effect": "Allow",
            "Action": [
                "ssm:PutParameter",
                "ssm:DeleteParameter",
                "ssm:GetParametersByPath",
                "ssm:GetParameter"
            ],
            "Resource": "arn:aws:ssm:*:ACCOUNT_ID:parameter/*"
        }
    ]
}
```
