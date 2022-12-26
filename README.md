# Light Estuary Node

Bare minimum node to run Estuary specific functions such as:

- Gateway to serve files
- Upload/Store retrieve from Filecoin
- Pinning and Gateway Service
- Deal-Making
- Customizable Bucket/Staging Area for CIDs
- UI to upload, download and monitor network and deals

## Deal-Making
One of the core value proposition of Estuary is it can manage the storage and retrieval deals. This light node uses a simplistic approach in making deals.

- Reactively creates new buckets when it reaches a threshold of number of items per bucket.
- Creates a CAR file for each set of bucket - meaning a CAR file has links to other files. The CAR file is then sent to the filecoin network via filcient

## Gateway 
This node comes with it's own light gateway to serve directories and files.


# Installation
To install the node, clone this repo and build.

## `go build`
```
go build -o lnode
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
```
./lnode store-file <path>
./lnode store-car <path>
./lnode store-dir <path>
./lnode store-cid <path>
```

# API endpoints
```

```



