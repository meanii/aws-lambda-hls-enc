# aws-lambda-hls-enc
> Go-powered AWS Lambda tool for swift and efficient HLS file encryption, empowering direct encryption of HLS files within Lambda. Boosts flexibility and scalability for managing media streaming workflows on AWS.

## Deployment

1. Create AWS Lambda function with API Gateway
   - Runtime: `Amazon Linux 2`
   - Architecture: `arm64`

2. Deploy the compiled and zipped file to Lambda, using the artifacts from this [release](https://github.com/meanii/aws-lambda-hls-enc/releases).

3. Now, you need to configure a few environment variables.

| Name                          | Value Type      | Description                          |
|-------------------------------| ----------------| -------------------------------------|
| CLOUDFRONT_ACCESS_KEY_ID      | String          | Cloudfront Public KeyID              |
| CLOUDFRONT_ORIGIN             | String          | Cloudfront CDN Domain                |
| CLOUDFRONT_PRIVATE_KEY_BASE64 | String          | Cloudfront Generated Private Key     |
| EXPIRE_TIME                   | String          | Signature URL Expiry Time, e.g. (2h) |

