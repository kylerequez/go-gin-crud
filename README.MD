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
![Alt text](./README-Images/image.png)
6. Run the "go mod download" command through the terminal, this will install all the necessary package (such as Gin)
7. Run the "go run migrate/migrate.go" to setup the database tables
![Alt text](./README-Images/image-1.png)
8. Run the "CompileDaemon -command="./go-gin-api"" on the terminal, this will start the application
![Alt text](./README-Images/image-2.png)
9. Make a GET request at "localhost:<port_number>/" and check if it returns the following:
![Alt text](./README-Images/image-6.png)
10. Make a POST request at "localhost:<port_number>/users/register" with the following request body (JSON) to register:
![Alt text](./README-Images/image-4.png)
11.  Make a request at "localhost:<port_number>/users/login" with the following request body (JSON) to :
![Alt text](./README-Images/image-5.png)
12.  You can now make requests to the different endpoints that the application has. Here are some examples:
![Alt text](./README-Images/image-7.png)
![Alt text](./README-Images/image-8.png)
![Alt text](./README-Images/image-9.png)
![Alt text](./README-Images/image-10.png)
![Alt text](./README-Images/image-11.png)
![Alt text](./README-Images/image-12.png)