This repository contains a simple Go application that implements user authentication and interacts with a MySQL database. You can use Docker Compose to easily set up the application and its dependencies.

Prerequisites
Before getting started, make sure you have the following installed on your system:
- Docker
- Docker Compose

Getting Started
1. Clone the Repository
Start by cloning this repository to your local machine:
- git clone https://github.com/muhammadarash1997/user.git
- cd user

2. Configure Environment Variables
Open the docker-compose.yml file and update the environment variables for the app service to match your desired configuration:
- DB_NAME: The name of the MySQL database.
- MYSQL_ROOT_PASSWORD: The root password for the MySQL database.
Save your changes.

3. Build and Start the Docker Containers
Build and start the Docker containers using Docker Compose:
- docker-compose up --build
This will build the Docker image for the Go application and start the MySQL database and your Go application container.

4. Access the Application
Your Go application should now be running. You can access it in your web browser or using tools like curl or Postman. By default, the application is accessible at http://localhost:8080.
You can register and log in to test the authentication features.

Clean Up
To stop and remove the Docker containers, you can press Ctrl+C in the terminal where docker-compose up is running. Then, run the following command to remove the containers:
- docker-compose down