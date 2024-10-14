# gRPC-Go Project

This project is a basic demonstration of gRPC in Go, featuring four types of gRPC communication:
- Unary RPC
- Server-side streaming RPC
- Client-side streaming RPC
- Bidirectional streaming RPC

The project illustrates how gRPC can be used to communicate between a client and a server using Protocol Buffers for data serialization. Each communication type is implemented as a separate function, demonstrating the versatility of gRPC for various use cases.

## Project Explanation

### 1. Unary RPC
Unary RPC is the simplest form of gRPC communication, where the client sends a single request to the server, and the server returns a single response. This is similar to a traditional function call.

### 2. Server-side Streaming RPC
In server-side streaming, the client sends a request to the server and receives a stream of responses. The server can send multiple responses before the connection is closed by the client.

### 3. Client-side Streaming RPC
With client-side streaming, the client sends a stream of requests to the server. After the client has finished sending the requests, the server processes them and sends back a single response.

### 4. Bidirectional Streaming RPC
In bidirectional streaming, both the client and server send streams of messages to each other. The two streams operate independently, meaning the client and server can read and write in any order.

### Proto File (`greet.proto`)

The project uses a Protocol Buffers definition file (`greet.proto`) to define the gRPC service and message types. This file is used to generate the Go code for the client and server.


## Setup Instructions

### 1. Clone the Repository

Clone the project repository to your local machine:

\`\`\`bash
git clone https://github.com/yourusername/gRPC-GO.git
cd gRPC-GO
\`\`\`

### 2. Install Dependencies

Run the following command to download the required dependencies:

\`\`\`bash
go mod tidy
\`\`\`

### 3. Compiling the Proto File

To generate the gRPC code from the \`greet.proto\` file, use the following command:

\`\`\`bash
protoc --go_out=. --go-grpc_out=. proto/greet.proto
\`\`\`

Make sure the \`protoc\` compiler and the plugins \`protoc-gen-go\` and \`protoc-gen-go-grpc\` are installed and available in your system's \`$PATH\`.

## Running the Server

1. **Navigate to the Server Directory:**

   \`\`\`bash
   cd server
   \`\`\`

2. **Start the gRPC Server:**

   \`\`\`bash
   go run *.go
   \`\`\`

   This command will start the server, listening for incoming gRPC requests.

## Running the Client

1. **Navigate to the Client Directory:**

   Open a new terminal window and navigate to the client directory:

   \`\`\`bash
   cd client
   \`\`\`

2. **Run the Client:**

   Use the following commands to run the client with different communication types:

   - **Unary RPC:**

     \`\`\`bash
     go run *.go 
     \`\`\`

   - **Server-side Streaming RPC:**

     \`\`\`bash
     go run *.go 
     \`\`\`

   - **Client-side Streaming RPC:**

     \`\`\`bash
     go run *.go 
     \`\`\`

   - **Bidirectional Streaming RPC:**

     \`\`\`bash
     go run *.go 
     \`\`\`

