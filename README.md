# grpc-test

Sample for a Go language gRPC client and server using Bazel as the build system.

# In 2 terminals

Run the server

```bazel run counter/server:server```

And the client

```bazel run counter/client:client```

# Notes

This does not do any king of gRPC authentication (*yet).