<div align="center">
  <img src="https://www.seekpng.com/png/full/399-3990193_building-a-go-web-app-from-scratch-to.png" width="200px">
  <span> &nbsp &nbsp &nbsp &nbsp</span>
  <img src="https://pngimg.com/uploads/mysql/mysql_PNG6.png" width="200px"> 
  <span> &nbsp &nbsp &nbsp &nbsp</span>
  <img src="https://i0.wp.com/www.docker.com/blog/wp-content/uploads/2013/11/homepage-docker-logo.png?fit=400%2C331&ssl=1" width="200px">
</div>
<br />

<div align="center">
  <h3> REST API build with Golang, MySQL, Docker </h3>
  <p> Understanding the basics of these technologies! </p>
</div>

## Requirements
 - Docker

## Clonning and Running
- To run this app, you need to create an ".env" file in the root folder.
- This file should contain the following environment variables
```
MYSQL_ROOT_PASSWORD=root_password
MYSQL_DATABASE=development
MYSQL_USER=user_password
MYSQL_PASSWORD=user_password
MYSQL_PORT=3306
API_HOST=http://localhost
API_PORT=5000
```

- Use the following command to build images and start the application
```
$ docker-compose up
```
