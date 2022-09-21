import json
from model import predict

def lambda_handler(event, context):

    data=event["queryStringParameters"]['data']

    
    return {
        "statusCode": 200,
        "body": json.dumps(
            {
                "test": predict(event["queryStringParameters"]['data']),
            }
        ),
    }
