# S3 Object Ownership Shifter

The S3 Object Ownership Shifter provides an AWS lambda function which copy objects between S3 buckets, As soon as an object arrives at the source bucket, the lambda will be triggered and copy the object to the target bucket, even if in another aws account the lambda will transfer the ownership of the object in case is needed.

## Installation

The zip file is under `build/s3copier/main.zip` all you need to do is upload the lambda function to the AWS account and setting the environment using the env vars

* `TARGET_S3_BUCKET` specifies the bucket name where the objects will arrive, ex: `my.target.bucket.test`

* `OWNERSHIP_FULL_CONTROL` use this var in case you want to transfer the ownership of the file ex: `emailaddress=exacmplet@example.com`

* `FILE_FILTER` use this in case you want to add a filter for only transfer objects with certain name, if you want to transfer objects which contains the word `pineapple` just set `pineapple` in the `FILE_FILTER` var.

The trigger of the lambda will be the source bucket where the objects you want to shift are.

## Recommendation

We always recommend use [terraform](http://terraform.io) for set the infrastructure, you can import this module to apply it https://github.com/full360/s3-object-ownership-shifter-tf

## Updates

In case you want to modify you can find the codebase under `cmd/s3copier/main.go` for build the zip you can use the make rule `make build` this will generate the `zip` file to upload it in aws.







