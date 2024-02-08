# oapi-codegen-cultibio

## Sensor Data API

This repository contains the OpenAPI specification for the Sensor Data API. The API allows you to add sensor data to the database.

## Overview

- **Version:** 1.0.0
- **License:** MIT

## API Documentation

### Add Sensor Data

Endpoint to add sensor data to the database.

- **URL:** `http://localhost/api/v1.0/sensor-data`
- **Method:** POST
- **Request Body:**
  - Description: Sensor data to add to the database
  - Content Type: `application/json`
  - Schema: [SensorData](#sensordata)
- **Responses:**
  - `200 OK`: Sensor data added successfully
  - `400 Bad Request`: Invalid input data
  - `500 Internal Server Error`

## Schemas

### SensorData

Object representing sensor data.

- **Properties:**
  - `SensorID` (integer)
  - `Value` (number)
  - `Timestamp` (string, format: date-time)

## Usage

To use this API, make HTTP requests to the specified endpoints with the required parameters.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
