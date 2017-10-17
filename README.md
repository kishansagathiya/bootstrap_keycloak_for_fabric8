# fabric8-bootstrap-keycloak
This repository has been created for the issue 92 for the https://github.com/fabric8-services/fabric8-auth repository
The intentions of this project is to
  - Version the Configuration
  - Instead of changing configuration manually through Keycloak UI, this let's us make changes in the json and then the code would apply the configuration to our Keycloak 
    
bootstrap package contains methods for 
        -   parsing json file (bootstrap/configuration.json)
        -   getting keycloak admin token ( using form-urlencoded post request by passing username and password)
        -   creating a new realm ( using post request by passing a json string)

How to run this?
```
[user@localhost fabric8-bootstrap-keycloak]$ go install
[user@localhost fabric8-bootstrap-keycloak]$ $GOPATH/bin/fabric8-bootstrap-keycloak

```
Running this would print contents of configuration.json, admin token message on creating a new realm, and errors
