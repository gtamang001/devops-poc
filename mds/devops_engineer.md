# DevOps Key Notes
## Key Phrases and terms 

People over process over tools : emphasize priority

Continuous delivery (Release in small batches)

Lean Sofware Management (Focus on Lean)

Change Control : operational success and change controll

Ifrastructure as code : The whole process should be treated as code

###  10 practical practices for devops 

1. Incident command system
2. Developer on call (involvement)
3. Status pages public status pages (banner message)
4. Blameless portmortoms (review)
5. Embedded teams (sharing)
6. The Cloud -- use cloud -- api driven way .. allows you to treat your infrastructue as code 
7. Andon Cords ( Innovation used by toyota, stop ship on production line, quality control system and stop the pipeline)
8. Dependency Injection -- Loosely coupled dependencies pass dependencies at run time ..
9. Blue green deployment ( two identical system, blue and green )
10. Chaos monkey (netflix )

### How to pick devops tools 

i. Programmable

ii. Verifiable
 
iii. Wellbehaved 

## Building blocks of DevOps 
AGILE : 

SDLC to AGILE development precursor to DevOps introduction 

AGILE's missing piece is System Operations 
Hence Devops is extension of AGILE system. 

LEAN : 

Lean Software Development 
Enhanced feedback loop 
Muda, Muri, Mura (Japanese terms to define ?)
Seven Waste of Software Development (google ?)

ITIL, ITSM, SDLC


are being replaced by DEVOPS 

## Infrastructure Automation 
Infrastructure as code 

Configuration Management

CM toosl are meant for following 

Provisioning -- prepare 

Deployment  -- deploy

Orchestration -- manage 

### Approaches  
imperative,
declarative,
idempoten -- repetition, 
self service,
Canary Deployment, 
Blue/ Green Deployment, 
Immutable Deployment(netflix),

Immutable Deployment ?? 
Deb file ?? 

### Continuous Delivery 

Small + Faster = better 

Continous Integration Practices 

Continuous Delivery mean : 
artifacts should be build once and then immutable 
have same built artifact to all the environments 
dev 
stage and prod 
Keeping track of metrics for furthering in making better 

Role of QA
Don't forget testing 
Code hygiene 
Integration testing 
TDD, BDD, Acceptance driven development 

CICD Tools: 
Version Control : GitHub, Gitlab, Bitbucket 

CI: Jenkins, GoCD, Bamboo

BUILD TOOL :
Maven 
Gulp

Unit Test 
Linters, JUnit
Cucumber ...

Infra test 
Kitchen CI 

Reliability: 
Metrics to improve and increase reliability 

Circuit breaker 

APM tools ?? 
Code Profilers ?? 

Metrics and Monitoring 
Service Metrics 
System Component Metrics 
Application Metrics 


TOOLS 

nagios, saas , pingdom, datadog, newrelic, statd, graphana, 
flapack io 
rerun github 


### Additional DevOps Resources : 
DevOps Days, Velocity, Devops enterprise summit 
10 . Visible Ops 9. Continuous Delivery 8. Release It ! 7. Effective DevOps 6. Lean Software Development and Agile toolkit 5. Web Operations 4. Practice of Cloud System administration 3. DevOps Handbook 2. Leading Transformation 1. The Phoenix Project

Internet Sources : 
devopsweekly.com
dzone.com/devops 
devopscafe
theshipshow.com
arresteddevops
theagileadmin

### Established Devops practices 
Public, Private and Hybrid Cloud

Serverless  : AWS Lamda 

Rugged Manifesto: 
gauntlt 


# Devops Foundation CI/CD 

Continuous Integtration 

Continuous Deployment 

Continuous Delivery : Additional change ( deploy to production like environment)

### Benefits of CI/CD 
Long term benefits 
 - Emplowering teams 
 - Lowerwed Cycle time (in hrs instead of weeks and months or days )
 - Better Security (high performers means better security)
 - Rythm of practice means regularity and easy for people 
 - Productivity increase 
### Build pipelines in practice 
Jenkins, Bamboo, Circlec, unit tests , packaging
1. Version Control 
    Use Git, commit oftent, share codebase, branching strategies
    Trunk based approach or master branch approach, linters and formatter, unit test and unit test results in pull request itself in pre-commit hook
    - Practicle 
    ```bash
    git show commit-id # gives you the detail of commit
    # always run 
    git diff to check 
    # check hooks 
    cat ./hooks/git-pre-commit.hook
    make test 2>&1
    if [ $? -ne 0 ]
    then 
            echo "Failed Tests .."
            exit 1
    fi

    go fmt 2>&1 
    if [$? -ne 0 ]
    then 
            echo "go vet has detected potential issues with your project.."
            exit 1
    fi

    fo vet 2>&1
    if [ $? -ne 0 ]
    then 
            echo " go vet has detected potential issues in your project.."
            exit 1
    fi
    

    ```

    ### Options for CI 
     - Open Source Jenkins  
     - Commercial
     - SaaS for CI (Circle CI, Travis CI)
     A proper CI culture is always a default options while implementing CI process. 

# Jenkins 

Free Open Source Tool 

### Packaging and Artifact management
Benefits of Artifacts 
- Reliability
- Composability (common packaging format, debian packages ?? fpm ??)
- Security ()
- Sharability 
Build it test it and package it and deploy it 

### Sample Nexus artifactory 
Nexus

### Testing in DevOps 
TDD, BDD, and other ways of testing
Unit testing, Integration Testing 
- Tract various metrics for the growth and agility of your team. 
- Testing is always key element to add in your development 
### Testing Practices : 
Using go language sample 
- Using go convey 
```bash 
# 
cd workdirectory 
# run go convey 
make convey 
# open ui at 9999 port 
``` 
### Linters 
```bash 
make lint
```

### Next up is self service Deployments 
 insert image here 
 Same artifacts same way deployments and similar environments for deployments. 

 ### Deployment Tools and Options 
  - Source pulls (caputrano ? )
  - CM systems (chef, puppet )
  - Orchestration push (Rundeck, ansible)
  - Build (Docker, rancher)
  - Others
  - build dashboard for deployment 
### Deployment using Chef 
Server Spec tests 

### Integration Testing
- Testing performed to expose defects in the interfaces and in the interactions between integrated components or systems. 
### End to End Testing 
Regression Testing 
Acceptance Testing 
### Security Testing 
- Static :        
        finding out vulnerable js (using retire js)
- Dynamic : 
        Running code, Gauntlt as part of CI pipeline 
        archni - xx attack, veeracode 
        sonatype 

# DevSecOps 
        DevOps is a culture shift
        Normal Security Scan flow
 1. Develop Code 
 2. Send to Security 
 3. Scan 
 4. Send PDF of result 
 5. find false positive 
 6. Fix security defects 

 This process is too slow. Hence new approach is needed, enter DevSecOps 

 | DevSecOps  | DevOps Rugged | 
 | :--------------- | ---------------| 
 | Overall automation | Developer focused | 

Automation meaning in DevSecOps 
Dont automate pushing a button. 
Real-time notification to the developers.
ChatOps, not pdfs 
### Education 
Educate developers on security basics 
### Educate Dev team 
Make security team auditor instead of being the main focus. Make developer team the main focus of development team. 

### Security team duties 
- provide framework of security to devs 
- provide tool and educate them 
- provide needed expertise 
- work closely with dev team 
- shift security toward left of application development lifecycle.
### Technologies 
APIs
Containers 

### Containers and Security automation 
Using security as container and ship it. 
Use artifactory to keep version of the security tool.

### A typical process 
| Process | Tools | 
| : -----------| --------------------| 
| Code         | Bitbucket           | 
| Task Tracking | Jira               | 
| Automation    | Jenkins            | 
| Configuration | Ansible            | 
| Deployment    | Nagios             | 
Make security work this workflow.
? ChatOps 
? IAST tools 
### Automation on Ops Side 
### Free tools 
- Find Security Bugs (java, jenkins)
- OWASP zed Attack Proxy Object (ZAP) ( Jenkins and Docker )
- sqlmap 
- OpenVAS 
- Recon- ng (recon framework) 
- Glue (docker container)

### Exercise is implementing Glue with JIRA 




















