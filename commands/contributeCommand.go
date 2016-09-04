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

package commands

import (
	"botexample/client"
	"bytes"
	"regexp"
)

type ContributeCommand struct{}

func (cc ContributeCommand) MatchRegex() (regex *regexp.Regexp) {
	regex = regexp.MustCompile(`(?i)/contribute`)
	return regex
}

func (hc ContributeCommand) Help() (help string) {
	return "<b>/contribute</b> - find out how you can add your own commands"
}

func (cc ContributeCommand) OnMessage(message client.V2Message, client client.BotClient, handlers []CommandHandler) {
	var buffer bytes.Buffer
	buffer.WriteString("<messageML>")
	buffer.WriteString("Go to <a href=\"https://github.com/SymphonyOSF/botexample\"/>")
	buffer.WriteString("</messageML>")
	client.SendMessageMLMessage(message.StreamId, buffer.String())
}
