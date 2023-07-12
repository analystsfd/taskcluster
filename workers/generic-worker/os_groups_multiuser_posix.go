//go:build multiuser && (darwin || linux || freebsd)

package main

func (osGroups *OSGroups) Start() *CommandExecutionError {
	groups := osGroups.Task.Payload.OSGroups
	if len(groups) == 0 {
		return nil
	}
	if config.RunTasksAsCurrentUser {
		osGroups.Task.Infof("Not adding task user to group(s) %v since we are running as current user.", groups)
		return nil
	}
	return nil
}

func (osGroups *OSGroups) Stop(err *ExecutionErrors) {
}
