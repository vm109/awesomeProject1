### Cloudwatch Logs 
1. we can write our application logs to cloudwatch. 
2. In container environment where the logs are ephermal we can store our logs in cloud. 
3. Cloudwatch logs have logs grouped under `logGroupName`
   1. Each logGroup have multiple log streams 
   2. log stream is where array of logs are stored which came from a single session
   3. that single session is identified by `sequenceToken`
   4. When we write the first message to cloudwatch logs there will not be `sequenceToken`
   5. But for the subsequent log writes we need to send `sequenceToken` with the request to write to CW logs. 
4. 