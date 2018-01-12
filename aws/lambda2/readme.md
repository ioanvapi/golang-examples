# Run Go code in aws lambda using python runtime

## Reference
* http://www.itdadao.com/articles/c19a1057329p0.html
* https://github.com/eawsy/aws-lambda-go-shim
* https://github.com/eawsy/aws-lambda-go-core/

## Technical

I like this approach because the handler function we have to implement has a similar signature as the java function in a java runtime.
For below signature, json.RawMessage is just []byte and you can return any type.
```
func Handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
    // todo
}
```

Also, for simple situations you can use:
```
func Handle(evt String, ctx *runtime.Context) (String, error) {
    // todo
}
```

Or, you can define your own types. The framework will unmarshal and marshal the data for you.
```
type Proba struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

type Result struct {
    Signature string `json:"signature"`
}

func Handle2(evt Proba, ctx *runtime.Context) (Result, error) {
    return Result{
        Signature: fmt.Sprintf("Name %s, Age %d", evt.Name, evt.Age),
    }, nil
}
```

I've tested all these in AWS Lambda Dashboard using input data as json.

## Building
The framework provides a make file and a docker image. They are used to build the code and creates a zip archive that can be uploaded in aws.

The zip file can be uploaded in AWS Lambda Dashboard or using the aws cli. The profile flag is used to switch to another user.
```
aws lambda create-function \
  --profile homeuser \
  --region eu-west-1 \
  --role arn:aws:iam::146844705783:role/TestLambdaRole \
  --function-name preview-go2 \
  --zip-file fileb://handler.zip \
  --runtime python2.7 \
  --handler handler.Handle
```

In order to update the existing lambda function you can use:
```
aws lambda update-function-code \
  --function-name preview-go \
  --zip-file fileb://handler.zip
```
