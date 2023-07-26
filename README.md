# IPFS Metadata RestAPI

This project is a REST API that interacts with a DynamoDB table to retrieve metadata about content-addressed data.

## Getting Started

To get started with this project, you'll need to have the following installed on your machine:

- **Go**
- **AWS CLI**

You'll also need to have an AWS account set up with the appropriate permissions to create and manage DynamoDB tables.

## Installation

1. Clone the repository to your local machine.
2. Run `go mod download` to download the project dependencies.
3. Set up your AWS credentials by running `aws configure`.
4. Create a DynamoDB table by running:
   `aws dynamodb create-table --table-name MetadataTable --attribute-definitions AttributeName=Cid,AttributeType=S --key-schema AttributeName=Cid,KeyType=HASH --billing-mode PAY_PER_REQUEST --region us-east-1`

sql
Copy code 5. Start the server by running `go run main.go`.

## Usage

Once the server is running, you can interact with it using any HTTP client. Here are some example requests:

- `GET /metadata`: Returns a list of all metadata items in the DynamoDB table.
- `GET /metadata/{cid}`: Returns the metadata item with the specified CID.


## Contributing

If you'd like to contribute to this project, please fork the repository and submit a pull request.

