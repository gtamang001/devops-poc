## Docker Certified Exam Guide
Following will be tested
- Demonstrate the ability to upgrade the Docker engine 
- Complete setup of repo, select a storage driver, and complete installation of Docker engine on multiple platforms
i. Linux
ii.  Amazon web services
iii.  Windows server
iv. Windows 10 and 
v. Mac OS 

- Configure logging drivers (splunk,journald,etc)
- Setup swarm, configure managers, add nodes, and setup backup schedule
- Create and manage user and teams 
- Interpret errors to troubleshoot installation issues without assistance
- outline the sizing requirement prior to installation
- Understand namespaces, cgroups, and  configuration of certificates
- use certificate-based client-server authentication to ensure a Docker daemon has the rights to access images on a registry
- Consistently repeat steps to deply Docker engine, UCP, and DTR on AWS and on premises in an HA config
- Complete configuration of backups of backups for UCP and DTR 
- Configure the Docker daemon to start on boot. 
15 % of the exam total comes from this section 

Setup 

For various docket tutorial setups, we need to Install a virtual machine client. In this case  we are using vmware fusion. There are others tools as well out there for virtualization. 
You can install vmware fusion from online for free. 
Once downloaded open the Vmware with ubuntu system. [This feature can be used to set up any software system or program i.e. you can run windows on mac if you have vmware] 
.For this tutorial choose ubuntu server [This can be downloaded as well from online ]
Once the ubuntu is running, run 
```bash
sudo apt-get install openssh-server # it  will install ssh in the client
# then run 
ip addr # take ip address and then run 
ssh username@ip # enter credentials and log in
```

## Overview Docker edition 

- CE Container engine with orchestration, networking, and security 
  Desktop

- EE Cotainer engine with orchestration, networking, and security.
     Certified insfrastructure, plugins, and ISV containers 
     Image management [Docker trusted registry]
     Container app amangement [Universal Control plane]
     Secure image scanning [Image scanning]
    
  Servers 
  ### When selecting a platform chose from following 
  - Production or Dev 
  - Type of application 
  - Cloud or on prem
  - Scalability requirement 
  - Pay as you go or fixed 
  - Linux or Windows 
  * Find information at docker compatibility matrix.

  ### Docker sizing
  UCP - 8 GB minimum for manager 
      - 4 GB worker node and 
      - 3 GB of disk space 
  
  Production grade application setup may require following on average or in general 
   - 16 GB Ram for manager nodes running DTR 
   - 4 vCPUs for manager nodes or nodes running DTR 
   - 25 - 100 GB of free disk space 
   Note: Generally windows images require more disk space than linux images.

   ### Docker Engine 
   You may already be familiar with what the Docker Engine is and how the typical Docker architecture lays out but if not, here's a quick review for you before we jump into the universal control plane. So, the Docker Engine is the heart of Docker and when people think of Docker, they're actually thinking of the Docker Engine. The Docker Engine is designed as a client server application and it's really made up of three different things. It starts off with dockerd or the Docker daemon which is installed when you install Docker and that's the server, that's the Docker server itself. Along with the installation of the Docker Engine, you receive a RESTful API which is important because that defines the interface that all other programs use to talk to the daemon and there's so many different pieces that make up the typical Docker ecosystem, both tools from Docker as well as third-party tools. And then finally, you have the Docker client. So, this is the actual Docker command that you run as a client to talk to the Docker server, to pull down images, build images and instantiate containers. No matter what version of Docker you're using whether it's the Community Edition or the Enterprise Edition the Docker Engine is the required foundation that makes it all possible. Now let's review the typical Docker architecture. The Docker daemon is installed on the Docker host. That Docker host could be your desktop or laptop computer, it could be a server in the data center or it could be a virtual machine running up in the cloud. From there, the Docker host is used to execute or instantiate your containers and images. It's administered through the Docker client which could be on the same host as the Docker daemon or it could be remote. That's the beauty of the Docker client server architecture. Using the Docker client you can pull down images from a registry and then execute those images as containers running on the Docker host.

   ### Namespaces 
   he Docker engine utilizes something called 'Namespaces' to isolate what's happening in the running containers from the operating system that those containers are running on. With Namespaces the kernal resources such as the process ID, user IDs, network storage, and inner process communications or IPC, can all be virtualized and shared between the host operating system and the containers running on top. Namespaces weren't created by Docker. Linux Namespaces are a core feature of the Linux kernal and have been around since 2002. Since that time there's been a lot of enhancement around Namespaces and Docker has capitalized on those enhancements in the Docker engine. Docker utilizes process, mount, IPC, network, and user Namespaces to isolate what's happening on the Docker host from what's happening in the Docker containers. Thankfully, Microsoft has even added the equivalent of Namespace isolation in Windows so that Docker for Windows could provide the same functionality. Namespaces are similar in concept to what a hypervisor does to provide the virtual resources like virtual CPU, virtual memory and virtual storage to a virtual machine. Namespaces keep containers isolated until Docker administrators, for example, allow containers to communicate over the Docker virtual networks on the same host. With the Namespace isolation in Docker operating systems and applications running in containers feel like they have their own process trees, file systems, network connections, and more. It's even possible in Docker to map a user account in a container to a user account in the host operating system. Here's a simple example of how Namespace isolation works. For example, here I am in an Ubuntu Docker host, and if I do a ps minus ef you can see there are roughly a hundred different processes running on this host. If I do the ip adder command or ip address it'll list out the ip addresses. You can there's roughly some 70 different network interfaces on this host. However if I do a docker run minus it and let's run alpine Linux and go into the shell, we pull down an alpine Linux image, we're running that as a container, and now if I perform the same commands here, for example ps minus ef, we have exactly two processes. So what's happening in this operating system in the container is isolated off the process isolation. Using Namespaces is isolating us off from what's happening in the Docker host and vice versa. Another example here is if I run the ip adder command, you can see we have exactly two interfaces, with two different ip addresses on those interfaces. Again, another example of how Namespaces work to isolate off network resources from the Docker host. So that's how Namespaces work in Docker.
   commands
   ```bash
    pf -ef 
    ip addr
    ```
### CGROUPS 
 In addition to name spaces the Docker engine also utilizes cgroups or control groups for isolation the difference is that control groups provide resource isolation, Linux C groups were introduced in 2006 and are available in most Linux distributions today. The Docker engine allows you to leverage these cgroups to prioritize the resources for your running containers. In general, cgroups can be used to control a variety of resources but in Docker, the most common uses are for container CPU limits, CPU reservations, memory limits, and memory reservations. You should know that to use cgroups in Docker, there are strict kernel requirements, so if you plan to use cgroups, make sure that your kernel and your version of Linux is going to be compatible. Most importantly cgroups prevent runaway containers from bringing down an entire Docker host. A common example would be cgroups preventing a (mumbles) service attack or preventing noisy neighbor syndrome in a multi-tenant environment. For example when you have multiple companies or multiple users on a single host all sharing resources if one particular application from one company is maximizing that server's resources you don't want that to slow down all of the other companies or all of the other applications running on the same host. In Docker for Windows, control groups are called job objects but they're just as useful as Linux cgroups. You can use Docker control groups with the Docker run command, for example minus C or --cpu-shares will allow you to allocate a specific number of shares to that container that you're bringing up. Similarly the -memory bytes command will allow you to set a memory limit or --memory-reservation will allow you to set a memory soft limit. For more information on cgroups in Docker, I recommend you check out this Docker documentation entitled limiting a container's resources. For your Docker certified associate exam, you should be aware of what cgroups are and what command syntax you might use to control container resources in a production environment.

In general CGROUPS can be used to control  variety of reources, but in Docker, the most common uses are container 
- CPU limits 
- CPU reservations 
- Memory limits 
- Memory reservations 

For exam, we need to know how to manage containers using cgroup commands 

Use Limit container's resources page read in docker documentation.

## Installing on Linux distros 
Pre Requisits
- Proper version 
- Network connected 
- Uninstall docker
```bash
sudo apt -get remove docker docker-engine docker-ce docker.io
sudo apt-get update
```
- Make modification to the Linux package installer  (apt) to add the Docker repository 
```bash
sudo apt-get install apt-transport-https ca-certificates curl software-properties-common
```
Add the Docker Official GPG Key to Apt 
```bash
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
```
Verify that you now have the docker GPG key 
```bash
sudo apt-key fingerprint 0EBFCD88
```
Add the Docker repository to Apt 
```bash
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
```
- Update packages 
```bash
sudo apt-get update
```
- Install Docker 
```bash
sudo apt-get install docker-ce=17.12.0~ce-0~ubuntu
```
or 
```bash
sudo apt-get install docker-ce
```
Before verify Add Groups and Users 
```bash
sudo groupadd docker 
sudo usermd -aG docker $USER
```
Log out and log back in for this to take effect 
- Verify 
```bash
docker -version 
docker run hello-worl
```
### Installing Docker on windows 
Compatible after windows 2016

### Installing Docker on AWS 
Use centOS amazon instance 
- create instance 
- Create key pair 
- Launch instance 
- connect to new instance using ssh 
```bash
ssh -i "linkedin-pem" centos@ec2-18-221-131-68.us.east-2.compute.amazonaws.com
```
Deploy CE on centos 
- uninstall docker 
```bash 
sudo yum remove docker docker-common docker-selinux docker-engine
```
- setup the repositories that docker will need 
```bash
sudo yum install -y yum-utils device-mapper-persistent-data lvm2
```
- add docker-ce repo to yum package manager 
```bash
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
```
- Yum install docker ce 
```bash
sudo yum install docker-ce 
```
- Start docker and add user 
Use systemctl 
```bash
sudo systemctl start docker 
sudo groupadd docker 
sudo usermod -aG docker $USER
```
- Verify 
```bash
docker ren hello-world
```

### Docker Universal Control Plane  
Intall UCP 
From docker doc 
```bash
docker container run --rm -it --name ucp -v /var/run/docker.sock:/var/run/docker.sock docker/ucp:2.2.5 install --interactive

```
Grab the url from the log and login 

### Upgrading docker 
Find out the version of docker 
```bash
# check docker version 
docker version 
# check docker configuration folder 
cd /var/lib/docker && ls -l 
cd swarm
tar swarm swarm-bak.tar # back it up incase you need
apt-get update # update host packages
# follow install instructions to intall docker 
```
## Configuring Docker 
 ### Storage Drivers 
Storing data needed for the containers, recommended to use volumes. Or can be stored within the containers, storage driver. Checkout, Docker storage drivers doc in docker documentation.
```bash
docker info | more # storage driver 
```
 ### Docker repositories and Docker Hub
 A Docker repository is a collection of related Docker images with the same name but different tags. A registry, on the other hand, is a place where you store your images. A registry can be public, it can be private, and it can have multiple repositories. The most common and default registry for Docker is Docker's own public registry, Docker Hub. As part of a new Docker installation it's common to set up a repository and link it to Docker Hub
 ### Login to docker hub 
 ```bash
 docker login # enter docker hub credentials, create one if you don't have one
 # pull down an images 
 docker pull <imagename>
 docker pull alpine
 # Create various tags of the images 
 docker tag <image-id> docker-id/alpine:tag1 
 docker tag <image-id> docker-id/alpine:tag2 
 docker tag <image-id> docker-id/alpine:tag3
 # push these images to docker hub 
 docker push docker-id/<repository>
 # check and validate by going into the docker hub

 # Delete images locally 
 docker rm <image hash>

 # Pull back the image from docker hub 
 docker pull docker-id/repository:tag1
 # this should pull down the image 
 ```
 ### Docker Swarm 
 Docker Swarm is the cluster management and orchestration feature that's built into the Docker Engine. Docker Swarm is available whether or not you have Docker Community Edition or Enterprise Edition, and it's available on all platforms.
 
 ### Set up to simulate 
 - Use vmware fusion to spin off 3 machines (linux nodes)
 - Node 1 : Manager Node - install docker engine and intall ucp 
 - Node 2 and Node 3 : Worker nodes 
 ```bash

 ```
 ### Creating and Managing users 
  Docker's Universal Control Plane offers multi-user management and role-based access control. You can go in here to User Management and you can create new organizations and teams inside those organizations and then users inside the teams and then finally, roles for those users
 ### Set up 
 - Organization 
  - Teams [A,B,C]
   - Add users to these teams 
 - Create user 
 - Set up roles 

 ### Docker daemon start manually 
 ```bash
 systemctl status docker 
 # check if service is enabled
 # disable on boot  
 systemctl disable docker
 # enable 
 systemctl enable docker 
 ```
 ### Backing up docker 
 - Docker Swarm Cluster 
 ```bash
 # Go to manager 
 systemctl stop docker 
 cd /var/lib/docker 
 # Copy and archive the swarm directory 
 systemctl start docker 
 ```
 - Universal Control Plane
  ```bash
  docker container run \
  --log-driver non --rm \ 
  --name cup \ 
  -v /var/run/docker.sock:/var/run/docker.sock \
  docker/ucp:2.2.6 backup \
  --id <ucp-instance-id> \
  --passphrase "secret" > /tmp/backup.tar 
  ```
 - Docker trusted Registry
 ```bash
 # images 
 sudo tar -cf backup-images.tar \ $(dirname $(docker volume inspect --format '{{.Mountpoint}}' dtr-registry-<replica-id>))
  # Dtr metadata 
  read -sp 'ucp password: ' UCP_PASSWORD; \
  docker run --log-driver non -i --rm \
  --env UCP_PASSWORD=$UCP_PASSWORD \
  docker/dtr:2.3.11 backup \
  --ucp-url <ucp-url> \
  --ucp-instance-tls \
  --ucp-username <ucp-username> \
  --existing-replica-id <replica-id> > backup-metadata.tar 
 ``` 
 - Container volume Data 
 ```bash
 
 ```
