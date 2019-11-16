# Streming data provider

Basic example to provide a large csv reading data from database using streaming to make then response more faster

 - Docker engine
 - Docker compose

Start application
  # build and runnig all stack on the fist time you need wait a lot time because the process will be create all user registers
  ./application serve

  # get all dependency applicationu 
  ./application dep

  # run application 
  ./application run

  # build base image with golang and dependencies manager dep
  ./application build

  # to create registers of users
  ./application.sh create_users

# Just make this requisition to get csv using streaming on terminal
