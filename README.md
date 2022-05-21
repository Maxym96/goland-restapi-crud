            ## CRUD REST-API ##

            ## Get Event (by ID)

            ## GET http://localhost:8081/v1/event/

            ##Headers:

            ## key: id
            ## value: Event id from database
            ## ID field IS REQUIRED!
            ## return Event from db

            ##Get All Events

            ##GET http://localhost:8081/v1/event/all

            ##return list all Events from db

            ##Add new Event

            ##POST http://localhost:8081/v1/event/add

            ##Body:
            ##x-www-form-urlencoded:
            ##name
            ##description
            ##date_and_time

            ##All fields ARE REQUIRED!

            ##return added Event (id, name, description and date_and_time)

            ##Update Event

            ##PUT http://localhost:8081/v1/event/update

            ##Body:
            ##x-www-form-urlencoded:
            ##id
            ##name
            ##description
            ##date_and_time

            ##ONLY ID field IS REQUIRED!

            ##return updated Event with new fields (name, description and date_and_time)

            ##Delete Event

            ##DELETE http://localhost:8081/v1/event/delete

            ##Headers:
            ##key: id
            ##value: Event id from database

            ##ID field IS REQUIRED!

            ##return true if Event is deleted from db