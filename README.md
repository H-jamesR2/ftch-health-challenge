# HTTP Endpoint Health Check Monitor

This Go program monitors the health of multiple HTTP endpoints defined in a YAML configuration file. It sends periodic HTTP requests to each endpoint, checking if they are up or down, based on the response status code and latency. Then, the program logs the cumulative availability percentage of each endpoint's domain to the console after each monitoring cycle.

## Table of Contents
- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Configuration File](#configuration-file)
- [Running the Program](#running-the-program)
- [Logging](#logging)
- [Testing](#testing)
- [Error Handling](#error-handling)
- [License](#license)

## Features
- Monitors HTTP endpoints defined in a YAML configuration file.
- Supports GET and POST requests with headers and body.
- Logs the availability percentage of each domain every 15 seconds.
- Flags endpoints as "UP" if the response code is in the 2xx range and latency is under 500ms.
- Graceful error handling with detailed logging (function name, timestamp, and issue).

## Requirements
- Go 1.16 or later
- Internet access (for sending HTTP requests)
- A valid YAML configuration file defining the endpoints to monitor

## Installation
1. Ensure Go is installed on your system. You can verify it by running:
   ```bash
   go version
   ```
   If Go is not installed, follow the [official installation guide](https://golang.org/doc/install).

2. Clone this repository to your local machine:
   ```bash
   git clone https://github.com/your-username/http-health-checker.git
   cd http-health-checker
   ```

3. Build the project:
   ```bash
   go build -o http-checker
   ```

   This command will create an executable binary named `http-checker` in your current directory. On Windows, the binary will be `http-checker.exe`.

## Configuration File

The program accepts a single input argument, which is the file path to a YAML configuration file that defines the HTTP endpoints to monitor.

### Example `config.yaml`:

```yaml
- name: Fetch Index Page
  url: https://fetch.com/
  method: GET
  headers:
    user-agent: fetch-synthetic-monitor
- name: Fetch Careers Page
  url: https://fetch.com/careers
  method: GET
  headers:
    user-agent: fetch-synthetic-monitor
- name: Fetch Some Post Endpoint
  url: https://fetch.com/some/post/endpoint
  method: POST
  headers:
    content-type: application/json
    user-agent: fetch-synthetic-monitor
  body: '{"foo":"bar"}'
- name: Fetch Rewards Index Page
  url: https://www.fetchrewards.com/
```

### Configuration Fields:
- `name` (required): Descriptive name for the HTTP endpoint.
- `url` (required): The URL of the HTTP endpoint.
- `method` (optional): HTTP method (e.g., GET, POST). Defaults to `GET` if not provided.
- `headers` (optional): A set of headers to include in the request.
- `body` (optional): The request body (for `POST` requests).

## Running the Program

Once the binary is built, you can run the program by passing the YAML configuration file as an argument.

### 1. Run the Program

Execute the binary along with the configuration file:

On Unix-based systems (Linux/macOS):

```bash
./http-checker <path_to_config_file.yaml>
```

On Windows:

```bash
http-checker.exe <path_to_config_file.yaml>
```

Example:
```bash
./http-checker ./config.yaml
```

OR you can either directly run the yaml file:
```bash
go run main.go <path_to_config_file.yaml>
```

```bash
go run main.go config.yaml
```

This will start monitoring the endpoints listed in `config.yaml`. The program will:
- Check each endpoint every 15 seconds.
- Log the cumulative availability percentage for each domain at the end of every 15-second cycle.

### Sample Output:

```
2024-09-17T12:00:00Z INFO fetch.com has 33% availability percentage
2024-09-17T12:00:00Z INFO www.fetchrewards.com has 100% availability percentage
2024-09-17T12:00:15Z INFO fetch.com has 67% availability percentage
2024-09-17T12:00:15Z INFO www.fetchrewards.com has 50% availability percentage
```

## Logging

By default, the program logs all errors and information messages to the console. Errors include:
- Issues with loading the configuration file
- HTTP request failures or unexpected status codes
- Response latencies exceeding 500ms

Each log entry includes:
- Timestamp
- Function where the error occurred
- Detailed error message

You can optionally log output to a file. To do this, modify the `init()` function in `main.go` to point to a log file, as demonstrated below:

```go
logFile, err := os.OpenFile("healthcheck.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
if err != nil {
	log.Fatalf("Could not open log file: %v", err)
}
log.SetOutput(logFile)
```

## Testing

To replicate testing of the program:
1. Create a valid YAML configuration file with the endpoints you want to test. A sample `config.yaml` is provided.
   
2. Run the program using the command:
   ```bash
   ./http-checker ./config.yaml
   ```

3. Monitor the console (or log file) for outputs related to endpoint availability.

### Test Case Examples:
- **Test a correct GET endpoint**: Ensure a `200 OK` response is returned within 500ms.
- **Test a delayed endpoint**: Set up an endpoint with a delayed response (>500ms) to confirm it gets flagged as `DOWN`.
- **Test incorrect URL**: Add a non-existent domain to test the error handling and logging.

## Error Handling

The program includes comprehensive error handling and logging. If the YAML file is incorrectly formatted or if fields like `name` or `url` are missing, the program will:
1. Log the error with a timestamp and the function where it occurred.
2. Stop execution and notify the user.

Example of an error log:
```
2024-09-17T12:05:00Z ERROR in LoadConfig.Validate: missing 'url' for endpoint: Fetch Index Page
2024-09-17T12:05:00Z ERROR in CheckEndpoint.ResponseCheck: Get "https://invalid-url.com/": dial tcp: lookup invalid-url.com: no such host
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.