
# Linux Shell Basics 
## Overview 
- Shell Types 
- Variables and Shell Environment 
- Tips and Tricks 
- Configure Shell History 
- Pattern Matching with Globs 
- Pattern Matching with Extended Globs 
## Using Bash 
- Escapes and quotes 
- Brace and path subsitution 
- Command and variable subsitution 
Nested command subsitution ($ and parenthesis combination)
- Named and Unnamed pipes 
- File redirects and tees 
- 
## Linux Processses 
## Job Scheduling 
## System Services  

0 = stdin , 1 => stnout , 2 => stderr
```bash
command > filename # stdout to file 
command 2> filename # stderr to file 
command < filename # input from filename object or 

# if use &> it means redirect both stdout and stderr to the file 
command | command2 # standard pipe command 
command 2>&1 | command2 # both stdout and stderr go into the pipe
# alternative 
command |& command2
# 
command1 >> file # append the stdout to end or to file 
command &>> file # both stout and stderr to end or to file 
# here documents  '<< ' command
 sort << END 
 banana
 apples
 cherry
 berry
 END
```
Another useful command 
```bash
lsof -p $$ # see the file descriptors of processes open 



