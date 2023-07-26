# IPFS Metadata RestAPI

This project is a REST API that interacts with a DynamoDB table to retrieve metadata.

## Features

- **GET Metadata**: Get the metadata of an IPFS CID via GET requests.
- **GET CIDs**: Get all of the CID inside the database via GET requests.

## Prerequisites

- **Go**: This is a Go project, and you would need Go installed to run or build it.
- **AWS CLI & AWS SDK**: The AWS CLI needs to be set up with appropriate permissions. The program also utilizes the AWS SDK for Go v2.
- **DynamoDB**: Configuring the DynamoDB table with an appropriate name is essential. Please make sure to update the table name in the code accordingly.


## Installation

1. Clone the repository to your local machine.
2. Run `go mod download` to download the project dependencies.
3. Set up your AWS credentials by running `aws configure`.
4. Create a DynamoDB table or update the code to support yours.
5. Start the server by running `go run main.go` or building the dockerfile.

## Usage

Once the server is running, you can interact with it using any HTTP client. Here are some example requests:

- `GET /metadata`: Returns a list of all metadata items in the DynamoDB table.
- `GET /metadata/{cid}`: Returns the metadata item with the specified CID.


## Results
![image](https://github.com/QuincyForbes/GoIPFRest/assets/74159902/1886335e-bb93-410e-83d9-aa4f9989209a)

![image](https://github.com/QuincyForbes/GoIPFRest/assets/74159902/061693d6-2dd8-4895-a707-995f5ffdce9c)
## Contributing

If you'd like to contribute to this project, please fork the repository and submit a pull request.

