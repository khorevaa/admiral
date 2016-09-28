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
	"bytes"
	"encoding/json"
	"fmt"

	"admiral/credentials"
	"admiral/functions"
	"admiral/hosts"
	"admiral/resourcePools"

	"github.com/spf13/cobra"
)

func init() {
	initCustomPropertiesList()
	initCustomPropertiesSet()
	initCustomPropertiesRemove()
}

var customPropertiesListCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists current properties of given entity.",
	Long:  "Lists current properties of given entity.",

	Run: func(cmd *cobra.Command, args []string) {
		output := RunCustomPropertiesList(args)
		processOutput(output, nil)
	},
}

func initCustomPropertiesList() {
	customPropertiesListCmd.Flags().StringVar(&cpHostIP, "host", "", "IP of the host that you want to manage custom properties.")
	customPropertiesListCmd.Flags().StringVar(&cpCredID, "credentials", "", "ID of the credentials that you want to manage custom properties.")
	customPropertiesListCmd.Flags().StringVar(&cpResPoolID, "resource-pool", "", "ID of the resource pool that you want to manage custom properties.")
	CustomPropertiesRootCmd.AddCommand(customPropertiesListCmd)
}

func RunCustomPropertiesList(args []string) string {
	var buffer bytes.Buffer

	if cpHostIP != "" {
		buffer.WriteString(hostCpString())
		buffer.WriteString("\n")
	}

	if cpCredID != "" {
		buffer.WriteString(credCpString())
		buffer.WriteString("\n")
	}

	if cpResPoolID != "" {
		buffer.WriteString(rpCpString())
	}
	return buffer.String()
}

func hostCpString() string {
	cpHost := hosts.GetPublicCustomProperties(cpHostIP)
	if cpHost == nil {
		return "Host with this IP not found."
	}
	cpJson, err := json.MarshalIndent(cpHost, "", "    ")
	functions.CheckJson(err)
	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("Custom Properties of Host: %s\n", cpHostIP))
	buffer.WriteString(fmt.Sprint(string(cpJson)))
	return buffer.String()
}

func credCpString() string {
	cpCred := credentials.GetPublicCustomProperties(cpCredID)
	if cpCred == nil {
		return "Credentials with this ID not found."
	}
	cpJson, err := json.MarshalIndent(cpCred, "", "    ")
	functions.CheckJson(err)
	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("Custom Properties of Credentials: %s\n", cpCredID))
	buffer.WriteString(fmt.Sprint(string(cpJson)))
	return buffer.String()
}

func rpCpString() string {
	cpRp := resourcePools.GetPublicCustomProperties(cpResPoolID)
	if cpRp == nil {
		return "Resource pool with this ID not found."
	}
	cpJson, err := json.MarshalIndent(cpRp, "", "    ")
	functions.CheckJson(err)
	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("Custom Properties of Resource pool: %s\n", cpCredID))
	buffer.WriteString(fmt.Sprint(string(cpJson)))
	return buffer.String()
}

var customPropertiesSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set custom property to given entity.",
	Long:  "Set custom property to given entity.",

	Run: func(cmd *cobra.Command, args []string) {
		output := RunCustomPropertiesSet(args)
		processOutput(output, nil)
	},
}

func initCustomPropertiesSet() {
	customPropertiesSetCmd.Flags().StringSliceVarP(&cpKeys, "key", "k", []string{}, "(Required) Keys of custom property.")
	customPropertiesSetCmd.Flags().StringSliceVarP(&cpVals, "value", "v", []string{}, "(Required) Values of custom property.")
	customPropertiesSetCmd.Flags().StringVar(&cpHostIP, "host", "", "IP of the host that you want to manage custom properties.")
	customPropertiesSetCmd.Flags().StringVar(&cpCredID, "credentials", "", "ID of the credentials that you want to manage custom properties.")
	customPropertiesSetCmd.Flags().StringVar(&cpResPoolID, "resource-pool", "", "ID of the resource pool that you want to manage custom properties.")
	CustomPropertiesRootCmd.AddCommand(customPropertiesSetCmd)
}

func RunCustomPropertiesSet(args []string) string {
	var buffer bytes.Buffer
	if cpHostIP != "" {
		if hosts.AddCustomProperties(cpHostIP, cpKeys, cpVals) {
			buffer.WriteString("Host's custom properties are set.")
			buffer.WriteString("\n")
		}
	}

	if cpResPoolID != "" {
		if resourcePools.AddCustomProperties(cpResPoolID, cpKeys, cpVals) {
			buffer.WriteString("Resource pool's custom properties are set.")
			buffer.WriteString("\n")
		}
	}

	if cpCredID != "" {
		if credentials.AddCustomProperties(cpCredID, cpKeys, cpVals) {
			buffer.WriteString("Credentials's custom properties are set.")
		}
	}
	return buffer.String()
}

var customPropertiesRemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove custom property to given entity.",

	Run: func(cmd *cobra.Command, args []string) {
		output := RunCustomPropertiesRemove(args)
		processOutput(output, nil)
	},
}

func initCustomPropertiesRemove() {
	customPropertiesRemoveCmd.Flags().StringSliceVarP(&cpKeys, "key", "k", []string{}, "(Required) Keys of custom property.")
	customPropertiesRemoveCmd.Flags().StringVar(&cpHostIP, "host", "", "IP of the host that you want to manage custom properties.")
	customPropertiesRemoveCmd.Flags().StringVar(&cpCredID, "credentials", "", "ID of the credentials that you want to manage custom properties.")
	customPropertiesRemoveCmd.Flags().StringVar(&cpResPoolID, "resource-pool", "", "ID of the resource pool that you want to manage custom properties.")
	CustomPropertiesRootCmd.AddCommand(customPropertiesRemoveCmd)
}

func RunCustomPropertiesRemove(args []string) string {
	var buffer bytes.Buffer
	if cpHostIP != "" {
		if hosts.RemoveCustomProperties(cpHostIP, cpKeys) {
			buffer.WriteString("Host's custom properties are removed.")
			buffer.WriteString("\n")
		}
	}

	if cpResPoolID != "" {
		if resourcePools.RemoveCustomProperties(cpResPoolID, cpKeys) {
			buffer.WriteString("Resource pool's custom properties are removed.")
			buffer.WriteString("\n")
		}
	}

	if cpCredID != "" {
		if credentials.RemoveCustomProperties(cpCredID, cpKeys) {
			buffer.WriteString("Credentials's custom properties are removed.")
		}
	}
	return buffer.String()
}