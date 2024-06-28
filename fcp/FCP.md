1. Compile the program:

```bash
    go build -o smolfed fcp.go
```

2. Open multiple terminals

3. In each terminal, set the environment variables and run the program with different instances IDs:

```bash
    export INSTANCE_ID=1
    export INSTANCE_IP=127.0.0.1
    export INSTANCE_PORT=8001
    export INSTANCE_METADATA="Instance 1"
    ./smolfed
```

```bash
    export INSTANCE_ID=2
    export INSTANCE_IP=127.0.0.2
    export INSTANCE_PORT=8002
    export INSTANCE_METADATA="Instance 2"
    ./smolfed
```

```bash
    export INSTANCE_ID=3
    export INSTANCE_IP=127.0.0.3
    export INSTANCE_PORT=8003
    export INSTANCE_METADATA="Instance 3"
    ./smolfed
```