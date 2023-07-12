package main

func (osGroups *OSGroups) UpdateCommands() {
	taskContext.pd.RefreshLoginSession(taskContext.User.Name, taskContext.User.Password)
	for _, command := range osGroups.Task.Commands {
		command.SysProcAttr.Token = taskContext.pd.LoginInfo.AccessToken()
	}
}
