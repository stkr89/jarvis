# Jarvis

Jarvis is a cli based workflow automation tool. It takes its input from a `yml` 
file and performs the required tasks.  
Following are some benefits of automating your workflows using Jarvis:
- Can be shared and executed by any member of your team
- Can be added to a git repo and evolve over time
- Can work with environment variables
- Can be executed by an automation server (eg, Jenkins)
- Can keep the process consistent and free from manual errors

---
Here is sample `config.yml` file that can be processed by Jarvis
```
tasks:
  - name: Make sample get request
    task_type:
      http:
        method: GET
        url: https://reqres.in/api/users?page=2
        auth:
          basic:
            username: ${username}
            password: ${password}
  - name: Make sample post request
    task_type:
      http:
        method: POST
        url: https://reqbin.com/echo/post/json
        auth:
          basic:
            username: ${username}
            password: ${password}
```
Use following command to verify if the `config.yml` is valid
```
jarvis verify ./config.yml
```
To initiate execution, run following command
```
jarvis apply ./config.yml
```