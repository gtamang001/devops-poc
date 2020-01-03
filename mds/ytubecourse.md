### Notes on devops training Linked In and sample bash scripts 

Continous integration with Jenkins 
Continous integration is reason to generate a shorter release cycles for the clients and also it is way to get automation in the release process. 

Code to continous integration tool which will build and run tests and then give feed back if needed to the deveoper or else release the artifacts. Reducing the release risk with faster delivery of the finished product. 


### Jenkins 
Set up Jenkins 
- install jenkins in the docker container if you are using container 
```bash
sudo apt get openjdk8-jdkxs
sudo apt-get install apache2 

sudo wget -q -o - https://jenkins-ci.org/debian/jenkins-ci.org.key | sudo apt - key add - 

sudo sh -c 'echo deb http://pkg.jenksin-ci.org/debian binary/ > /etc/apt/sources.list.d/jenkins.list

sudo apt-get update 

sudo apt-get install jenkins 

service jenkins status 

```

### Jenkins Master Slave architecture 

create two slaves in aws 

slave1 and slave2 

these servers can be managed from master node 

### CICD pipline 

get project from https://github.com

install docker on both the slaves 

configure Jenkins to build the project and Slave1 if successfull then build on slave2


Finally, trigger the job using git web-hooks : an event which automatically triggers the pipeline 

Jenkis > create a job > in slave1 > link repository 











