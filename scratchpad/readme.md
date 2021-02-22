# Design Notes 

Messages between the two sides of the system will be json formatted, this makes it much easier to  parse (good parsers available in both languages ) and is plenty efficient given we are not going to be pushing that much data.  

Defs: 
    **CLI app** is the go application that the user can call.  
    **Service app **  This is the G/LabVIEW app that receives the messages 


## Messags from CLI to G service applicaiton 
 

These messages contain the data input into the CLI by the user, using the [go kong parser library](https://github.com/alecthomas/kong)  

In order to make the first level parser on the service app side generic we will embed the class in the message header and use the factory pattern to provide a message specific parser.  

```json
{
    "type":"CommandType.lvclass",
    "data":{
    ///data for this message type goes here 
    }
}
```
Order of operations
*    CLI app sends msg to service app 
*    service app parses message and sends an ack.  
*    Command specific process takes place. Each process is responsible for returning the exit command to the CLI app at the appropriate time.  

The above process cuts out persistent terminals or prompting the user for input but this will be a future feature if needed. 


## Messages from Service App to CLI app 

**Print Line ** Print output to CLI 
** Exit With Error ** Exit the CLI app with an error code 