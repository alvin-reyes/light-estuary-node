# Light Estuary Node

*THIS IS A WIP*

Bare minimum node to run Estuary specific functions such as:

- Gateway to serve files
- Upload/Store retrieve from Filecoin
- Pinning and Gateway Service
- Deal Making (Storage and Retrieval)
- Customizable Bucket/Staging Area for CIDs
- UI to upload, download and monitor network and deals

## Deal-Making
One of the core value proposition of Estuary is it can manage the storage and retrieval deals. This light node uses a simplistic approach in making deals.

- Reactively process contents on buckets and make deals.
- Buckets runs on worker groups to concurrently process CIDs.
- Creates a CAR file for each set of bucket, meaning a CAR file has links to other files. 
- The CAR file is then sent to the filecoin network via filcient.
- Uses FIFO to select content without deals. 
- Replication default count of 3 for each content.


## Dashboard / Gateway 
This node comes with it's own light gateway to serve directories and files.

View the gateway using:
- https://localhost:1313
- https://localhost:1313/dashboard
- https://localhost:1313/gw/ipfs/:cid
- https://localhost:1313/ipfs/:cid

# Installation
To install the node, clone this repo and build.

## `make`
```
make all
```

## `go build`
```
go build -tags netgo -ldflags '-s -w' -o lnode
```

## `docker`
```
docker build .
```

# Running the daemon
Running the daemon will initialize the node configuration and the gateway at port 1313
```
./lnode daemon
```

# CLI commands
The following commands will store the file, dir, car or cid into the local blockstore.


## Store / Upload
```
./lnode pin <path>
./lnode pin-file <path>
./lnode pin-car <path>
./lnode pin-dir <path>
./lnode pin-cid <path>
```

## Retrieval 
```
./lnode retrieve <cid> <miners>
./lnode retrieve <cid>
```

## Jobs
```
./lnode jobs deal-run <cid>
./lnode jobs deal-run-bucket <bucket-uuid>
```

# API endpoints
```
GET /api/v1/content/upload
```



