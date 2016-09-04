/*
 *
 *
 * Copyright 2016 Symphony Communication Services, LLC
 *
 * Licensed to Symphony Communication Services, LLC under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 *
 */


package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	AgentUrl          string
	SessionAuthUrl    string
	KeyManagerAuthUrl string
	PodUrl            string
	KeyFilePath       string
	CertFilePath      string
	StreamId          string
}

type ConfigurationLoader struct {
	ConfigurationFileName string
}

func (configurationLoader ConfigurationLoader) Load(configurationFile string) (configuration Configuration) {
	file, _ := os.Open(configurationFile)
	decoder := json.NewDecoder(file)
	configuration = Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("Configuration=%+v\n", configuration)
	return
}
