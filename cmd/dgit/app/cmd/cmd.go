/*
Copyright 2018 COMPANY

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"io"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"matt-rickard.com/docker-merge/pkg/dgit/constants"
	"matt-rickard.com/docker-merge/pkg/dgit/version"
)

var v string

var rootCmd = &cobra.Command{
	Use:   "dgit",
	Short: "",
}

func NewRootCommand(out, err io.Writer) *cobra.Command {
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if err := SetUpLogs(err, v); err != nil {
			return err
		}
		logrus.Infof("%+v", version.Get())
		return nil
	}

	rootCmd.AddCommand(NewCmdVersion(out))
	rootCmd.AddCommand(NewCmdMerge(out))

	rootCmd.PersistentFlags().StringVarP(&v, "verbosity", "v", constants.DefaultLogLevel.String(), "Log level (debug, info, warn, error, fatal, panic")
	return rootCmd
}

func SetUpLogs(out io.Writer, level string) error {
	logrus.SetOutput(out)
	lvl, err := logrus.ParseLevel(v)
	if err != nil {
		return errors.Wrap(err, "parsing log level")
	}
	logrus.SetLevel(lvl)
	return nil
}
