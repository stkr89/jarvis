# Jarvis

Jarvis is a cli based workflow automation tool. It takes its input from a `yml` 
file and performs the required tasks.  
Following are some benefits of automating your workflows using Jarvis:
- Shared and executed by any member of your team
- Added to a git repo and evolve over time
- Work with environment variables
- Executed by an automation server (eg, Jenkins)
- Keep the process consistent and free from manual errors

---
Here is simple `config.yml` file that can be processed by Jarvis.
```
tasks:
  - name: Make sample get request
    task_type:
      http:
        method: GET
        url: https://reqres.in/api/users
        auth:
          basic:
            username: username
            password: password
```
Use following command to verify if the `config.yml` is valid.
```
jarvis verify ./config.yml
```
To initiate execution, run following command.
```
jarvis apply ./config.yml
```
Let's break down this file to understand what each component means.
## Tasks
A Task is an independent unit of work that gets executed as part of the workflow.
A config file can contain any number of tasks. These can be a combination
of any type of task that is supported by Jarvis.
## Task type
A task type defines the kind of work that Jarvis need to perform. Jarvis supports
following type of tasks:
### HTTP
```
http:
  method: GET
  url: https://reqres.in/api/users
```
`http` tasks are responsible for performing actions that require interaction with 
a RESTful endpoint. It has following properties:
#### Method
`method` takes any one of the following values:
- `GET`
- `POST`
- `PUT`
- `DELETE`
#### Url
`url` takes the address of the RESTful endpoint where the request need to be made.
#### Auth
`auth` describes the type authentication scheme required by the RESTful endpoint. It
supports following authentication schemes:
##### Basic
```
auth:
  basic:
    username: username
    password: password
```
`basic` authentication adds a header field in the form of 
`Authorization: Basic <credentials>`, where `<credentials>` is the Base64 encoding of 
`username` and `password` joined by a single colon `:`.
##### Bearer Token
```
auth:
  bearer_token:
    token: token
```
`bearer_token` authentication adds a header field in the form of 
`Authorization: Bearer <token>`.