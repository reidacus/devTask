# devTask
Solution to assigned Dev Task

### Overview
This solution has containerised both a MongoDB and the simple GoLang application and is simply run from the docker-compose file. When running this the first time, the docker-compose will build the image on the fly of the application, copying in all the required files to build and run the GoLang application. Once this image is built, this could be distributed and pushed to a Docker Hub repo for others to use, but for this exercise, it has not been. 

The solution also will create a custom bridge network for the containers to communictae with one another and makes use of DNS resolution, since the default docker bridge network does not allow resolving container names via DNS. 

### Assumptions
There was no specification for redundancy of either the application or the DB. Had these been a production application which was being produced, then we would most likely want to deploy both the front and and backend to a swarm. If a container should cease to function in the swarm configuration, Docker has the ability to quickly and automatically replace that bad container with a new instance using the image files already built. The database would be a bit more tricky as we ideally want to preserve the data should the container get into a bad state. For this reason, we would want to think about using volumes with our docker containers instead rather than retaining the data solely inside the container. 

### Running this Solution
The solution has been tested on both Docker for Windows and Linux. 
1. Clone the following Repo : https://github.com/reidacus/devTask.git
2. Run the docker-compose file from the terminal with docker-compose up -d
3. Since we bind local machine ports to those which are specififed by the container, it is possible to use Postman or another tool to execute the commands as though we were querying localhost or the IP of the host machine. 


