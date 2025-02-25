# SCA README

## Overview

This project utilizes Docker to manage and deploy services efficiently. Before starting, ensure that the Docker engine is running on your machine.

## Getting Started

To start the project, navigate to the directory containing the `Makefile` and run the following command:

```bash
make up
```

This command will start the Docker containers using `docker-compose`. All environment variables are sourced from the `deployments/.env` file.

## API Endpoints

### Working with Cats

You can interact with the cat management system using the following API endpoints:

- **Create a new cat**  
    `POST /newCat/:CatName/:Salary/:Expireins/:Bread`  
    Creates a new cat in the system with the specified name, salary, expiration date, and bread type.
    
- **List all cats**  
    `GET /cats`  
    Retrieves a list of all cats in the system.
    
- **Get a cat by ID**  
    `GET /cat/:CatID`  
    Retrieves details of a cat specified by its ID.
    
- **Delete a cat**  
    `DELETE /deleteCat/:CatID`  
    Deletes a cat from the system by its ID.
    
- **Update cat's salary**  
    `POST /updateSalary/:CatID/:Salary`  
    Changes the salary for a cat specified by its ID.


### Working with Missions

You can manage missions using the following API endpoints:

- **Create a new mission**  
    `POST /newMission?names=Name,Target&coutrys=Countrys,Target`  
    Creates a new mission with the specified names and targets.
    
- **List all missions**  
    `GET /missions`  
    Retrieves a list of all missions.
    
- **Delete a mission**  
    `DELETE /deleteMission/:ID`  
    Deletes a mission by its ID.
    
- **Assign a cat to a mission**  
    `POST /assign/:CatID/:MissionID`  
    Assigns a cat to the specified mission.
    
- **Update mission status**  
    `POST /updateMission/id`  
    Changes the status of a mission to complete.

### Working with Targets

You can manage targets associated with missions using the following API endpoints:

- **Add target to mission**  
    `POST /newTarget/:MissionID/:Name/:Country`  
    Adds a target to the specified mission.
    
- **Delete a target**  
    `DELETE /deleteTarget/:TargetName`  
    Deletes a target by its name.
    
- **Write a Note**  
    `POST /updateNote/:TargetName/:Note`  
    Writes a note associated with the specified target.
    
- **Change Status of Target**  
    `POST /updateTarget/:TargetName`  
    Updates the status of the specified target.    
## Stopping the Docker Compose

To stop the Docker containers and remove them, use the following command:

```bash
make down
```

This will clean up the Docker environment used for the project.