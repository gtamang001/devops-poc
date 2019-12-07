## A sample demonstration of docker networking using netcat 

### Get NetCat using homebrew 
```bash 
brew install netcat
```
### Create Docker Container Server 
```bash
# We need to create a docker container which will act as the server for # us to try to connect to 
# we need a container which will have netcat available in it
# the image I use is ubuntu:14.04 
# create container with docker run 
docker run --rm -ti -p 45678:45678 -p 45679:45679 --name echo-server ubuntu:14.04 bash 
# --rm : for removing the container after it is stoped
# -ti : interactive mode 
# -p expose ports<host-port mapped to : containerport> 
# keep in mind two ports are exposed
# --name : name of our container server 
# bash for running bash shell in interactive mode
### 
### In the interactive bash shell promted after running the command
### above 
### run 
nc -lp 45678 | nc -lp 45679
# -lp : listening on this port 
# | pipe the output of this command into another nc command listening 
# on second port 
```

After the commands above are executed open two instances of terminal and run following command in respective terminals 
```bash 
# first instance run 
nc localhost 45678
## type anything to send message to the listening port of echo server 
# this message should show up in the following terminal window and vice
# versa
# second instance run 
nc localhost 45679

```
The process above explains the basics of networking at work in the docker environment. This base can be expanded further to enrich the concept and use docker networking to our advantage.