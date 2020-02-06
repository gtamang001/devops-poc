# Running Jenkins on AWS guide 
## Requirements 
- Linux Admin or basic understanding 
- Jenkins - intermediate understanding 
- AWS account 
- GitHub account 
## Steps 
- Log into AWS account 
- Create Security Group Under a distinct region for AWS 
  ### Create Security Group 
    1. EC2 Dashboard 
    2. Security Groups 
    3. Create Security Group
    4. Outboud (leave as default)
    5. Inboud add rules

    | Types | Protocol | Port Range | Source | Description | 
    | -----| :--------| :----------| :-----| :-------------| 
    | SSH   | TCP      | 22         | My IP  | SSH         | 
    | HTTP  | TCP      | 80         | Anywhere| HTTP       | 
    | HTTPS | TCP      | 443        | Anywhere | HTTPS     | 

    6. Create 

    Once created tag the security group name by clicking the pencil button. 
   ### Create Key Pairs 
   - Key pairs are cryptographic system for encrypting and decrypting login information 
   - The public key is attached to an instance 
   - The private key is stored locally 
   - Keep private key files safe and secure 
     #### Steps 
     - EC2 
     - Key Pairs
     - Create new key pair 
     - Download pem file 
     - move file to .ssh 
     - run chmod 600 on the file 
  ### Create EC2 instance 
  - EC2 Launch instance 
  - ubuntu 18.04 LTS bionic 
  - Free tier image 
  - Keep everything default and use the jenkins masters key and the above security groups for the the requested details 
  - Launch instance 
  ### Elastic IP configuration 
  - EC2 
  - Create Elastic IP by allocating 
  - Associate the new IP to the intance just created 
  ### Set Up the Software base for Jenkins 
  - SSH into the instance 
  ```bash
  ssh -i ~/.ssh/jenkins-master.pem ubuntu@public-dns-of-intance
  # become root 
  sudo su - 
  # get key for jenkins install verification from jenkins.org
  wget http://pkg.jenkins-ci.org/debian-stable/jenkins-ci.org.key
  # add key 
  apt-key add jenkins-ci.org.key
  # add jenkins debian repo to aptitube sources list 
  echo "deb http://pkg.jenkins-ci.org/debian binary/" > /etc/apt/sources.list.d/jenkins.list
  # begin updating 
  apt update && apt upgrade 
  # install all the needed softwares 
  apt install -y openjdk-8-jdk
  # nginx 
  apt install -y nginx 
  # jenkins 
  apt install -y jenkins
  # confirmation test 
  systemctl status nginx | grep Active 
  systemctl status jenkins | grep Active 
  ```
  ### NGINX - Reverse Proxy 
  User   ---> Nginx ----> Jenkins 
  - Security for Jenkins server 
  All the requests to the app server exposed at port 8080 by default will be routed through the proxy server nginx (web server )
  - Logging benefits 
  #### Configure 
  ```bash 
  # unlink the default conif first 
  unlink /etc/nginx/sites-enabled/default
  # add new config 
  nano /etc/nginx/conf.d/jenkins.conf
  # file 
  upstream jenkins {
        server 127.0.0.1:8080;
   }

   server {
        listen 80 default_server;
        listen [::]:80 default_server;
        location / {
                proxy_pass http://jenkins;
                proxy_set_header Host $host;
                proxy_set_header X-Real-IP $remote_addr;
        }
   }
  # test the syntax
  nginx -t 
  # load the configuration 
  systemctl reload nginx
 ```


