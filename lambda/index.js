var AWS = require('aws-sdk');

console.log('Loading function');

exports.handler = function(event, context, callback) {
    console.log(JSON.stringify(event, null, 2));
    event.Records.forEach(function(record) {
        console.log(record.eventID);
        console.log(record.eventName);
        console.log('DynamoDB Record: %j', record.dynamodb);

        if (record['dynamodb']['NewImage']['Id'] % 2 == 1) {
            var s3 = new AWS.S3();
            var params = {
                Bucket : "mathias-bucket",
                Key : "/lambda/new_update.txt",
                Body : record['dynamodb']['NewImage']['Id']
            }
            s3.putObject(params, function(err, data) {
                if (err) console.log(err, err.stack); 
                else     console.log(data);
            });
        } else {
            console.log("Do nothing..")
        }

    });
    callback(null, "message");
};