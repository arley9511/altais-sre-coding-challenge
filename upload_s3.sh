#!/bin/bash

files=(
"./files/test.txt"
)

BUCKET_NAME="s3-altais-test-4"
AWS_PROFILE="arley_tests"

for file in "${files[@]}"
do

  aws s3 cp "$file" s3://$BUCKET_NAME --profile "$AWS_PROFILE"

done
