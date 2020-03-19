package cmd

const (
	SupernodeHostPortDefault       = ":1234"                           // SupernodeHostPortDefault is the default supernode Host:port that edges started in `edged` should connect to.
	EdgedServerHostPortDefault     = ":1060"                           // EdgedServerHostPortDefault is the default Host:port of `edged`.
	SupernodeServerHostPortDefault = ":1050"                           // SupernodeServerHostPortDefault is the default Host:port of `supernoded`.
	HostPortDocs                   = "Host:port of the server to use." // HostPortDocs is the documentation for the host:port flag.
	ConfigurationFileDocs          = "Configuration file to use."      // ConfigurationFileDocs is the documentation for the configuration file flag.)
)

const (
	CouldNotBindFlagsErrorMessage        = "Could not bind flags"         // CouldNotBindFlagsErrorMessage is the error message to throw if binding the flags has failed.
	CouldNotStartRootCommandErrorMessage = "Could not start root command" // CouldNotStartRootCommandErrorMessage is the error message to throw if starting the root command has failed.
)
