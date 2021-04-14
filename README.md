# devTask
Solution to assigned Dev Task

### Overview
This solution has containerised both a MongoDB and the simple GoLang application and is simply run from the docker-compose file. When running this the first time, the docker-compose will build the image on the fly of the application, copying in all the required files to build and run the GoLang application. Once this image is built, this could be distributed and pushed to a Docker Hub repo for others to use, but for this exercise, it has not been. 

The solution also will create a custom bridge network for the containers to communictae with one another and makes use of DNS resolution, since the default docker bridge network does not allow resolving container names via DNS. 

### Running this Solution
The solution has been tested on both Docker for Windows and Linux. 
1. Clone the following Repo : https://github.com/reidacus/devTask.git
2. run the docker-compose file from the terminal with docker-compose up -d
3. Since we bind local machine ports to those which are specififed by the container, it is possible to use Postman or another tool to execute the commands as though we were querying localhost.


