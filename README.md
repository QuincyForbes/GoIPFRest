# IPFS Metadata RestAPI

This project is a REST API that interacts with a DynamoDB table to retrieve metadata.

## Getting Started

To get started with this project, you'll need to have the following installed on your machine:

- **Go**
- **AWS CLI**

You'll also need to have an AWS account set up with the appropriate permissions to create and manage DynamoDB tables.

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

