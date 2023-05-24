package cmd

type CmdHandler interface {
    HandleCmd(c Cmdline, id string) bool
}
