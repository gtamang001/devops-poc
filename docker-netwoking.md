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

### What if you wanted no part of nc ? then do following 
```bash
# run the server container as is 
docker run --rm -ti -p 45678:45678 -p 45679:45679 --name echo-server ubuntu:14.04 bash /
&& nc -lp 45678 | nc -lp 45679

# on the instance of terminal access the docker host as follows 
# fist instance 

docker run --rm -ti ubuntu:14.04 bash /
&& nc host.docker.internal 45678 
# second instance 

docker run --rm -ti ubuntu:14.04 bash /
&& nc host.docker.internal 45679
# start typing messages from one container to another container
```

# Container Networking within localhost examples 

We are using the same ubuntu image for the demo here 
```bash 
# create network you want to work with 
# * by default docker provides you some networks if you don't want to create new network 
# run
docker network ls # to view the networks 

# create your own network 
docker network create test-network

# run two containers or machines on this network using the ubuntu image 
docker run --rm -ti --net test-network --name apple-server ubuntu:14.04 bash
#

# create another machine 
docker run --rm -ti --net test-network --name orange-server ubuntu:14.04 bash
#
# try pinging each machine from the other using ping command 
```
### Add more Networks  

Create another network and make the network discoverable from apple-server 
```bash
# 
docker network create applesonly
# 
# now connect the applesonly with apple-server
docker network connect applesonly apple-server

```
Now create another machine in the applesonly server 
```bash
docker run --rm -ti --net applesonly --name red-apple ubuntu:14.04 bash
# try pinging apple-server or nc to apple server 
# nc apple-server  or 

ping apple-server ## should see results 
ping orange-server ## shouldn't be able to 

```
### Container networking key terms and ideas 
- Ethernet - Moves "frames" on wire or Wifi 
- Ip-Layer - moves packets on a local network 
- Routing - forwards packets between local networks
- Ports - adddress particular program on a computer 

Docker uses bridges to accomplish most of the networking groundworks To view the bridges use following test and explore more 
```bash 
# try running a container with --net flag with host
docker run --rm -ti --net=host ununtu:16.04 bash
# in the container run update and installl bridge-utils tool 
apt-get update && apt-get install bridge-utils
# now run bridge-utils to see bridge in your environment
brctl show

```
### Bridges 
Docker use bridges to create virtual networks in your computer

Function as software switches 

They control the ehternet layer 

Turn this off using 
```bash
# docker run --net=host options image-name command 
```

### Routing 
Uses firewall rules to move packets between networks 

Uses NAT ( Network Address Translation )

Change the source address on the way out 

Change the destination address on the way back in 

```bash 
# too to use 
# sudo iptables -n -L -t nat
# 
```
Routing Test
```bash 
docker run --rm -ti --net=host --priviledged=true ubuntu bash

# 
apt-get update
apt-get install iptables 
# visualize iptables 
iptables -n -L -t nat 
# Run a separate container with port exposed 
# run above command and visualize the port forwarding 
```
At the core the port mapping in docker is accomplished using port forwarding

### Namespaces 
Linus feature 

Allows  processes to be attached  to private network  segments 

Private networks are shared into bridged network with  rest of the containers 

Containers have virtual network cards 

Containers get their own copy of  networking stack




