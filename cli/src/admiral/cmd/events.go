/*
 * Copyright (c) 2016 VMware, Inc. All Rights Reserved.
 *
 * This product is licensed to you under the Apache License, Version 2.0 (the "License").
 * You may not use this product except in compliance with the License.
 *
 * This product may include a number of subcomponents with separate copyright notices
 * and license terms. Your use of these subcomponents is subject to the terms and
 * conditions of the subcomponent's license, as noted in the LICENSE file.
 */

package cmd

import (
	"fmt"

	"admiral/events"

	"github.com/spf13/cobra"
)

func init() {
	eventCmd.Flags().BoolVar(&clearAll, "clear", false, "Clear all logged requests.")
	RootCmd.AddCommand(eventCmd)
}

var eventCmd = &cobra.Command{
	Use:   "events",
	Short: "Prints events log.",
	Long:  "Prints events log.",

	Run: func(cmd *cobra.Command, args []string) {
		RunEvents()
	},
}

func RunEvents() {
	el := events.EventList{}
	count := el.FetchEvents()
	if clearAll {
		el.ClearAllEvent()
		return
	}
	if count < 1 {
		fmt.Println("n/a")
	}
	el.Print()
}