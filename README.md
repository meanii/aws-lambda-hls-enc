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
| CLOUDFRONT_ORIGIN             | String          | Cloudfront CDN Domain                |
| CLOUDFRONT_ACCESS_KEY_ID      | String          | Cloudfront Public KeyID              |
| CLOUDFRONT_PRIVATE_KEY_BASE64 | String          | Cloudfront Generated Private Key     |
| EXPIRE_TIME                   | String          | Signature URL Expiry Time, e.g. (2h) |

4. IMPORTANT: To ensure the entire process functions correctly, it is necessary to pre-sign your master.m3u8 file with the corresponding CDN origin URL. Subsequently, update the domain to the deployed Lambda function.

Example:

1. Pre-sign your master.m3u8 file with the CDN origin URL:
- Original URL: `https://cdn-origin.com/master.m3u8`
- Pre-signed URL: `https://cdn-origin.com/master.m3u8?Expires=1610000000&Signature=xxxx&Key-Pair-Id=xxxx`

2. After deploying the Lambda function, update the domain to the Lambda function URL:
- Updated URL: `https://your-lambda-function-url/master.m3u8?Expires=1610000000&Signature=xxxx&Key-Pair-Id=xxxx`

Ensure to follow these steps for seamless integration with the Lambda function.
