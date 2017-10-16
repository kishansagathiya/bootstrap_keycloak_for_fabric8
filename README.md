# bootstrap_keycloak_for_fabric8
This repository has been created for the issue 92 for the https://github.com/fabric8-services/fabric8-auth repository

    bootstrap package contains methods for 
        -   parsing json file (bootstrap/configuration.json)
        -   getting keycloak admin token ( using form-urlencoded post request by passing username and password)
        -   creating a new realm ( using post request by passing a json string)
    
    Running this would print contents of configuration.json, admin token message on creating a new realm, and errors