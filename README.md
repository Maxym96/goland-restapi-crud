            ## CRUD REST-API ##

            ## GET EVENT BY ID ##

            ## GET http://localhost:8081/v1/event/

            ##Headers:

            ## key: id
            ## value: Event id from database
            ## ID field IS REQUIRED!
            ## return Event from db

            ## GET ALL EVENTS ##

            ##GET http://localhost:8081/v1/event/all

            ##return list all Events from db

            ## ADD NEW EVENT ##

            ##POST http://localhost:8081/v1/event/add

            ##Body:
                ##x-www-form-urlencoded:
                    ##name
                    ##description
                    ##date_and_time

            ## All fields ARE REQUIRED! ##

            ##return added Event (id, name, description and date_and_time)

            ## UPDATE EVENT ##

            ##PUT http://localhost:8081/v1/event/update

            ##Body:
            ##x-www-form-urlencoded:
            ##id
            ##name
            ##description
            ##date_and_time

            ##ONLY ID field IS REQUIRED!

            ##return updated Event with new fields (name, description and date_and_time)

            ## DELETE EVENT ##

            ##DELETE http://localhost:8081/v1/event/delete

            ##Headers:
            ##key: id
            ##value: Event id from database

            ##ID field IS REQUIRED!

            ##return true if Event is deleted from db