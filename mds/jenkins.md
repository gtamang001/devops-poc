# Jenkins
## Topics 
### Why Jenkis ? 
    Ease of use 
    Built in documentation 
    Open source
    Extensibility, plugins available 
### Key Terminologies 
    - Project 
    - Build 
    - Step 
    - Trigger 
    - Plugins
### Installing Jenkins 
 - Windows 
    Download from the website > zip > unarchive > 
    localhost:8080
    secret admin password and then install neccessary to plugins 
- Mac OS 
    Download from website > unarchive>
    localhost:8080
    terminal> sudo cat path
    paste the password in the jenkins homepage 
- Docker 
    install docker 
    ```bash 
    docker pull jenkins/jenkins:lts
    docker run --detach --publish 8080:8080 --volume jenkins_home:/var/jenkins_home --name jenkins jenkins/jenkins:lts
    # get password from exectin into the container 

    ```
- Standalone ubuntu Download 

### Jenkins Dashboard 
- new item - starts the process of creating new project or build step
- people - users 
- build history completed build 
- Manage Jenkins : configure jenkins 

### Jobs in Jenkins 
    - new item 
    - enter name 
    - build > add build step > echo "Hello Jenkins" 
        - Trigger 
        Trigger builds remotely 
        - build after some other projects 
        - build periodically ( on schedule builds )
        - Github hook (using webhooks)
        - pull scm ( activity in the repo as well)
    - Build Environment
        - delete and build 
### Build Step 
This includes various ways to build your project
### Post build
This step allows us to wrap up the build 
### Jobs in detail 
- Jobs with parameters 
- Choice parameters 
- boolean parameters (Shows up as checkbox)
    #### Scheduled jobs
    - build periodically option
    - similar to cron jobs 
    - 00*** every other day of the week at midnight 
    - Jenkins jobs Schedular Aliases example (practice ??)
### View and Folders 
Views
    - Views displays jobs that meet the criteria 
    - Views are like filter 
Folders 
    - Folders group things together 
    - Folders contain jobs, views and other folders 
- Examples 
    New View , + next to all tabs
    job filters (using regular expression)
    for e.g. for anything build is .\*BUILD.\* 

    Folders 
    New items, name and select folder, description 
    get isolated namespace of jenkins server 
### Pipelines 




