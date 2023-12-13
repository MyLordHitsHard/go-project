# Gofr Project

This is a simple Go project that uses Gofr framework and provides a RESTful API for managing a car garage.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

You need to have Go installed on your machine. You can download it from the [official website](https://golang.org/dl/).
You need to have Docker Desktop installed on your machine.

### Installing

To install the project, follow these steps:

1. Clone the repository: git clone https://github.com/MyLordHitsHard/go-project.git
2. Navigate to the project directory: cd go-project
3. Download the dependencies: go get .
4. Run the following commands in a terminal to create a table in MySQL docker image
   - `docker run --name gofr-project -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=car_garage -p 3306:3306 -d mysql:8.0.30`
   - `docker exec -it gofr-project mysql -uroot -proot123 car_garage -e "CREATE TABLE carGarage (id INT AUTO_INCREMENT PRIMARY KEY, license_plate    VARCHAR(255) UNIQUE, make VARCHAR(255), model VARCHAR(255), color VARCHAR(255), entry_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP, repair_status VARCHAR(255));"`


## Running the tests

To run the tests, use the following command: `go test .`


## Running the application

To start the application, use the following command: `go run .`


The application will start and listen on port 9000.

## API Endpoints
- `GET/car-garage`: Get a list of all the cars in the garage.
- `GET/car-garage/{license-plate}`: Get the details of a specific car.
- `POST/car-garage`: Create a new car in the garage.
- `PUT/car-garage/{license-plate}`: Update the details of a car from the garage
- `DELETE/car-garage/{license_plate}`: Delete a car from the garage.

## Built With

- [Go](https://golang.org/) - The programming language used.
- [Gofr](https://gofr.dev/) - The framework used.

## Authors

- Himanshu Dutt - Programmer - [MyLordHitsHard](https://github.com/MyLordHitsHard)



