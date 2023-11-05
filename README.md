# Bank Holiday Application

Fetch bank holidays data from a given url and process the data
as per requirement and serve the data in different endpoints.


## To build
From `cmd` directory execute `go build -o bank-holiday`

## To run
`./bank-holiday`, if built already following above.  
To build and run directly from `cmd` directory execute `go run .`

## To call the APIs
go to any web browser and hit respective urls.  
homepage: http://localhost:8082/  
To get all bank holidays that fall within 2023 from all regions: http://localhost:8082/getHolidays2023  
To get all bank holidays for the region “England and Wales" where “bunting” is false: http://localhost:8082/getHolidaysEnglandAndWalesNoBunting  
To get all bank holidays title and date that fall within 2024 from all regions.: http://localhost:8082/getHolidays2024TitleAndDate  

Also tools like curl, postman etc. can be used.  

## To run tests
While the application is running, from the root directory of the project, run `go test ./test -v`

## Future Action Items:
- Build, run and test through build tools like using Makefile
- Containerization of the application using tools like Docker
- More unit tests