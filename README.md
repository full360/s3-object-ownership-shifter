# S3 File Copier Lambda

The S3 File Copier Lambda provides an AWS lambda function which copy objects between s3 2 S3 buckets, As soon as an object arrives at the source bucket, the lambda will be trigger and copy the object to the target bucket, even if in another aws account the lambda will transfer the ownership of the object in case is needed.

## Installation

All resources are set in a terraform module so all you need to do is add the module to your terraform script and apply the changes.

### Step 1 User permissions

Make sure your user has the right permissions to apply the infrastructure, attach this IAM policy to your user

```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "Stmt1540566775041",
      "Action": [
        "lambda:AddPermission",
        "lambda:CreateFunction",
        "lambda:TagResource",
        "lambda:UpdateFunctionCode"
      ],
      "Effect": "Allow",
      "Resource": "*"
    },
    {
      "Sid": "Stmt1540566865566",
      "Action": [
        "s3:PutBucketNotification"
      ],
      "Effect": "Allow",
      "Resource": "*"
    },
    {
      "Sid": "Stmt1540566973814",
      "Action": [
        "iam:AttachRolePolicy",
        "iam:CreatePolicy",
        "iam:CreateRole"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
```

### Step 2 Add the terraform module
The module implementation should follow this structure

```
module "s3-file-copier" {
  source = "git::ssh://git@gitlab.full360.com/experiments/s3-copier.git?ref=master//terraform/modules/s3-file-copier"

  region                    = "${var.region}"
  aws_s3_source_bucket_name = "${var.aws_s3_source_bucket_name}"
  aws_s3_target_bucket_name = "${var.aws_s3_target_bucket_name}"
  ownership_full_control    = "${var.ownership_full_control}"
}
```

And the values of the variables should be like this

```
variable "region" {}
variable "aws_s3_source_bucket_name" {}
variable "aws_s3_target_bucket_name" {}
variable "ownership_full_control" {}
variable "env" {}

region                    = "us-west-2"
env                       = "dev"
aws_s3_source_bucket_name = "bucket-source-test-name"
aws_s3_target_bucket_name = "bucket-target-test-name"
ownership_full_control    = "emailaddress=exacmplet@example.com"
```

We recommend to follow the directory structure on `example-region-name`

```    ├── example-region-name
    │   └── _shared
    │       └── s3-file-copier
    │           ├── env
    │           │   └── prod.tfvars
    │           ├── main.tf
    │           │   └── prod
    │           └── variables.tf
```

- Create the module following the schema in `example-region-name/_shared/s3-file-copier/main.tf.sample`

- Create a directory inside `terraform` folder, use the name of the region where module will run, example `us_west`

- Inside `us_west` create `_shared/s3-file-copier` directories and copy the main and the variables

```
cp terraform/example-region-name/_shared//main.tf.sample terraform/us_west/_shared/s3-file-copier/main.tf && cp terraform/example-region-name/_shared/s3-file-copier/env/prod.tfvars.sample terraform/us_west/_shared/s3-file-copier/env/prod.tfvars && cp terraform/example-region-name/_shared/s3-file-copier/variables.tf.sample terraform/us_west/_shared/s3-file-copier/variables.tf
```

**NOTE: In case the transference will be in the same aws account you can leave the ownership full control empty ""**

### Step 3 Set up the terraform module infrastructure

Use the correct credentials **please**. Export the `AWS_PROFILE` or append it to
every command:

    export AWS_PROFILE=aws_profile

Once terraform is in place and ready to use we can `init` our directory:

Initializing:

    terraform init

Now that we have our terraform initialized we are ready to continue an plan our infrastructure.

    terraform plan

The output of this should say that there are **Plan: 8 to add, 0 to change, 0 to destroy.** to apply. If that's
the case you are done.

Now you are ready to apply them (remember to always plan fist) use the following command:

    terraform apply

In case you follow the example inside `example-region-name` you will need to add the parameter `-var-file=env/prod.tfvars` to the plan and apply commands, i.e `terraform plan -var-file=env/prod.tfvars`