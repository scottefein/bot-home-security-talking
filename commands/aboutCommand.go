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
	"os"
	"regexp"
)

type AboutCommand struct{}

func (aboutCommand AboutCommand) MatchRegex() (regex *regexp.Regexp) {
	regex = regexp.MustCompile(`(?i)/about`)
	return regex
}

func (aboutCommand AboutCommand) Help() (help string) {
	return "<b>/about</b> - what is this?"
}

var currentHost, _ = os.Hostname()

func (aboutCommand AboutCommand) OnMessage(message client.V2Message, client client.BotClient, handlers []CommandHandler) {
	var buffer bytes.Buffer
	buffer.WriteString("<messageML>")
	buffer.WriteString("I live on " + currentHost + "<br/>")
	buffer.WriteString("I was written in Go. It's a language written by Google... <a href=\"https://golang.org\"/><br/>")
	buffer.WriteString("I use the API's from <a href=\"https://developers.symphony.com\"/> to authenticate, listen for messages, and reply.<br/>")
	buffer.WriteString("</messageML>")

	client.SendMessageMLMessage(message.StreamId, buffer.String())
}
