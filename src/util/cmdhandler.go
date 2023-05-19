package util

type CmdHandler interface {
    HandleCmd(cmd [] string, id string) bool
}
