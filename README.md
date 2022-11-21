# Simple container

## Run
To run it from source, you need to have go installed.
```shell
$sudo go run cmd/mcontainer/main.go
```

Then you will be permitted to a containerized shell. For instance, if type **inside** container:
```shell
$echo $$
1
```
*`$$` stands for self PID.