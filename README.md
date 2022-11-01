# microservice_for_naris_app

## How to request data
- By default, the app runs on port 8000. You can request a specific user's details as follows:
  - GET http://localhost:8000/api/users/<user's email address>
- To get all users, you can use the following GET request:
  - GET http://localhost:8000/api/users"

- To create a new user:
  - POST http://localhost:8000/api/users
and include the following details in the request:
{
    "firstname":"New",
    "lastname":"Guy",
    "id": 1,
    "email": "useremail@email.com"
}


## How to receive data

![image](https://user-images.githubusercontent.com/6415751/199136668-2baab98c-03b0-458d-aa82-7dfe61a0c38b.png)
