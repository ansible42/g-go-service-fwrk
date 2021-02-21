#  G - Go Service Framework 

Design concept is to provide a framework from which to via scripting create a CLI interface for G applications that can sustain an asynchronous interface via loopback tcp. This was inspired by the amazing [g-cli](https://github.com/JamesMc86/G-CLI) package created by @jamesmc86.  


## Requirements 

*  Provide command line access to G applications that are running as a service or are called at runtime.  
*  Provide robust supplort for common command line features such as 
  *  commands 
  *  arguments  
  *  flags 
  *  sub-commands  
  *  robust help text 
  *  async processes query 
  *  Faster exeuction of commands 
  *  bash completion (strech goal) 

## Design Concept  

Go was selected as the CLI interface language because it is easy to compile across platforms and I know it works on cRIO's.  Also I hate Python. IT is pretty easy to script the creation of go and create the command cases from G type definitions.  Each command will have a type def and when the G application is compiled a go project will be created or updated to contain the updated type defs and corresponding commands, flags, etc. Then the go application will be generated from the scripted code. Ideally this should get to the point where the developer need not know go in order to use this framework.  

Then there will be a reciever template that you the developer must include within the target application.  Ideally I want to make this plugin able to handle all archetectures from simple queue driven state machines, Actor Framework and DQMH.  
