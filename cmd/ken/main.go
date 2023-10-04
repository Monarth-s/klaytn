// Modifications Copyright 2018 The klaytn Authors
// Copyright 2016 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.
//
// This file is derived from cmd/geth/main.go (2018/06/04).
// Modified and improved for the klaytn development.

package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/Monarth-s/klaytn/api/debug"
	"github.com/Monarth-s/klaytn/cmd/utils"
	"github.com/Monarth-s/klaytn/cmd/utils/nodecmd"
	"github.com/Monarth-s/klaytn/console"
	"github.com/Monarth-s/klaytn/log"
	"github.com/urfave/cli/v2"
)

var (
	logger = log.NewModuleLogger(log.CMDKEN)

	// The app that holds all commands and flags.
	app = utils.NewApp(nodecmd.GetGitCommit(), "The command line interface for Klaytn Endpoint Node")
)

func init() {
	// Initialize the CLI app and start ken
	app.Action = nodecmd.RunKlaytnNode
	app.HideVersion = true // we have a command to print the version
	app.Copyright = "Copyright 2018-2023 The klaytn Authors"
	app.Commands = []*cli.Command{
		// See utils/nodecmd/chaincmd.go:
		nodecmd.InitCommand,
		nodecmd.DumpGenesisCommand,

		// See utils/nodecmd/accountcmd.go
		nodecmd.AccountCommand,

		// See utils/nodecmd/consolecmd.go:
		nodecmd.GetConsoleCommand(utils.KenNodeFlags(), utils.CommonRPCFlags),
		nodecmd.AttachCommand,

		// See utils/nodecmd/versioncmd.go:
		nodecmd.VersionCommand,

		// See utils/nodecmd/dumpconfigcmd.go:
		nodecmd.GetDumpConfigCommand(utils.KenNodeFlags(), utils.CommonRPCFlags),

		// See utils/nodecmd/db_migration.go:
		nodecmd.MigrationCommand,

		// See utils/nodecmd/util.go:
		nodecmd.UtilCommand,

		// See utils/nodecmd/snapshot.go:
		nodecmd.SnapshotCommand,
	}
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Flags = utils.KenAppFlags()

	app.CommandNotFound = nodecmd.CommandNotExist
	app.OnUsageError = nodecmd.OnUsageError
	app.Before = nodecmd.BeforeRunNode
	app.After = func(ctx *cli.Context) error {
		debug.Exit()
		console.Stdin.Close() // Resets terminal mode.
		return nil
	}
}

func main() {
	// Set NodeTypeFlag to en
	utils.NodeTypeFlag.Value = "en"

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
