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


package nytimes

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty"
)

var newsClient = resty.New()

type News struct {
	APIKey string
}

type Result struct {
	Title    string
	Url      string
	Abstract string
}

type NewsResult struct {
	Results []Result
}

func (news *News) TopStories() (result []Result) {
	resp, err := newsClient.R().
		Get("http://developer.nytimes.com/proxy/https/api.nytimes.com/svc/topstories/v2/home.json?api-key=" + news.APIKey)
	if err != nil {
		fmt.Println(err)
		fmt.Println(resp)
		return
	} else {
		var newsResult NewsResult
		if err := json.Unmarshal(resp.Body(), &newsResult); err != nil {
			fmt.Println(err)
			return
		}
		result = newsResult.Results
		fmt.Println(resp.String())
	}
	return
}
