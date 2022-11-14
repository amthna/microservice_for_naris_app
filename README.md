# microservice_for_naris_app

## How to request data
- first, run the service by entering the service directory and running the command "go run main.go"
- By default, the app runs on port 8000. You can request a specific user's details as follows:
  - POST http://localhost:8000/api/check
and include the following details in the request:
{
    "first":"userfirstname",
    "email": "useremail@userdomain.com"
}

- Your "database" file (employees.json) will be read from. If the requested user doesn't exist, they'll be added to the database file.
- That's it! all extraneous endpoints (delete user, update user, etc.) have been removed with this update.

## How to receive data

![image](https://user-images.githubusercontent.com/6415751/199136668-2baab98c-03b0-458d-aa82-7dfe61a0c38b.png)
