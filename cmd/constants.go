package cmd

const (
	SupernodeHostPortDefault       = "localhost:1234" // SupernodeHostPortDefault is the default supernode Host:port that edges started in `edged` should connect to.
	EdgedServerHostPortDefault     = "localhost:1235" // EdgedServerHostPortDefault is the default Host:port of `edged`.
	SupernodeServerHostPortDefault = "localhost:1236" // SupernodeServerHostPortDefault is the default Host:port of `supernoded`.
)

const (
	CouldNotBindFlagsErrorMessage = "Could not bind flags" // CouldNotBindFlagsErrorMessage is the error message to throw if binding the flags has failed.
)
