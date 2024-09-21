package main

import "os"

func commandExit(_ *cfgCommand) error {
    os.Exit(0)
    return nil
}
