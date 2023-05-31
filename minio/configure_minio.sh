#!/bin/sh

# Wait for the MinIO server to start
#while ! curl --output /dev/null --silent --head --fail http://localhost:9000; do
#    sleep 1
#done

# Configure the bucket policy with custom access rules
#mc config host add myminio http://localhost:9000 "hello" "world"
#mc policy set download myminio/test-bucket --recursive
#mc admin policy set myminio custom-policy.json
#mc admin policy set myminio read-only-policy.json
#mc admin user add myminio custom-user custom-password
#mc admin policy set myminio custom-policy custom-user