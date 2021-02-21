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
