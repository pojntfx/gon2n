package cmd

const (
	keyPrefix         = "supernode."
	configFileDefault = ""
	serverHostPortKey = keyPrefix + "serverHostPort"
	configFileKey     = keyPrefix + "configFile"
	listenPortKey     = keyPrefix + "listenPort"
	managementPortKey = keyPrefix + "managementPort"
)

const (
	couldNotBindFlagsErrorMessage = "Could not bind flags"
)
