# Go-Gin-CRUD-and-JWT-Authentication
This is a sample Go using the Gin framework that provides the following:
1. CRUD Functionality
2. Login Functionality (Email only)
3. JWT Authentication

# How to run the application?
1. Install Go
2. Install Postgres
3. Install Postman, or use any application that allows you to make HTTP Requests
4. Download/Clone the repository
5. Create a .env file inside the folder and add the following details:

6. Run the "go mod download" command through the terminal, this will install all the necessary package (such as Gin)
7. Run the "go run migrate/migrate.go" to setup the database tables

8. Run the "CompileDaemon -command="./go-gin-api"" on the terminal, this will start the application

9. Make a GET request at "localhost:<port_number>/" and check if it returns the following:

10. Make a POST request at "localhost:<port_number>/users/register" with the following request body (JSON) to register:

11.  Make a request at "localhost:<port_number>/users/login" with the following request body (JSON) to :

12.  You can now make requests to the different endpoints that the application has. Here are some examples:
