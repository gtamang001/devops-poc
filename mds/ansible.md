# Ansible Notes 

Ansible is task execution Engine

It target the local computer and also all the computers attached to that computers via network

Very Scalable 

Execute list of tasks on various machines 

Lightweight Fleet Management System

Communication uses SSH : 
Widely trusted and used communication protocol 

If Ansible Operator can ssh into a target system, he/she can perfrom actions on it using Ansible.

Can be installed using python pip command 

Inventory 
Collection of Hosts IP addresses 
Sometimes categorized into groups and Subgroups 

Inventory Variables : 
key : Value pairs 

Tasks : Scripts to run on the hosts 

Existing Modules and Custom Modules 

### Playbook 
Ansible Executables 
created to execute on the inventories 

Different types of plays in the playbook

Infrastructure Provisioning 

On Demand Infrastructure (Cloud)

Orchestration Automation

### Ad Hock executions 

