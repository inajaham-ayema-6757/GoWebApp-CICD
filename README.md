# GoWebApp-CICD
Continuous Integration, Continuous Deployment using Docker, Travis CI, Heroku

# File Discription: -

1. The Simple web application is written in Go language, i.e. "myweb.go" file.
2. Simple Unit-test cases for web application(myweb.go) created are written in "myweb_test.go" file.
3. "Dockerfile" is used to build a Docker image of web application.
4. ".travis.yml" file specifies the programming language used, design building and testing environment.

# Process Flow: -

1. Developer commits his/her work on Github.
2. Step 1 invokes Travis CI and build is triggered for the same.
3. All the dependencies and environment is set by Travis CI.
4. Testing is done using test files provided.
5. Success or failing of tests is notified.
6. Docker image for web application is built and upload on Dockerhub.
7. Same Docker image is pushed in the container registry on Heroku.
8. Using that image, the application is deployed on Heroku.

# Heroku Application link: - 

https://myapp6757.herokuapp.com/




