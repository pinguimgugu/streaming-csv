# Streaming data provider

Basic example to provide a large csv reading data from database using streaming to make then response more faster

 - Docker engine
 - Docker compose

Start application
  # build and running all stack on the fist time you need wait a some time because the process will be create all 100.000 user registers
  ./application.sh serve

  # get all dependency application 
  ./application.sh dep

  # run application 
  ./application.sh run

  # build base image with golang and dependencies manager dep
  ./application.sh build

  # to create registers of users
  ./application.sh create_users

# Open this path in your browser to get csv with 100000 lines in little bit time
  http://localhost:7000/users/csv
