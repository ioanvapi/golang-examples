# Run Go code in aws lambda using Nodejs runtime

This approach assumes you can build go code for amazon linux machine (GOOS=linux GOARCH=amd64).
Also, you need the nodejs handler file that will spawn a new process for go built app.
The arguments passed to nodejs handler will be passed to the go code as program arguments.
We can read them from os.Args.

```javascript
var child_process = require('child_process');

exports.handler = function(event, context) {
    var proc = child_process.spawn('./main', [ JSON.stringify(event) ], { stdio: 'inherit' });
    proc.on('close', function(code) {
        if(code !== 0) {
            return context.done(new Error("Process exited with non-zero status code"));
        }
        context.succeed("Completed with success!!");
    });
}
```

You have to zip the nodejs file and the go binary ('main' in our example) in the same zip file that will be uploaded to aws lambda later.

```
I don't know how to pass the lambda context to the go code and how to return something from go code to the nodejs code.
```






