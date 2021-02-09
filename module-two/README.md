# Importing a module

```
# make the module directories
mkdir moduleone
mkdir moduletwo
```

```
# Create the go.mod file
1 module <MODULE_NAME>
2
3 go <GO_VERSION>
4
5 replace <MODULE_TO_IMPORT> => <PATH_TO_OTHER_MODULE>
```

```
# Run the top level module
go run moduletwo.go
```