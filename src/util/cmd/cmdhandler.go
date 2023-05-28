package cmd

type CmdHandler interface {
    HandleCmd(c Cmdline) bool
}
