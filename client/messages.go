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


package client

type V2BaseMessage struct {
	Id            string
	Timestamp     string
	StreamId      string
	V2messageType string
}

type Attachment struct {
	// TODO
}

type V2Message struct {
	Message     string
	FromUserId  int64
	Attachments []Attachment
	StreamId    string
}

type UserJoinedRoomMessage struct {
	AddedByUserId     int64
	MemberAddedUserId int64
	Id                string
	Timestamp         string
	StreamId          string
}

type UserLeftRoomMessage struct {
	RemovedByUserId  int64
	MemberLeftUserId int64
	Id               string
	Timestamp        string
	StreamId         string
}

type RoomMemberPromotedToOwnerMessage struct {
	PromotedByUserId  int64
	PromotedUserId int64
	Id               string
	Timestamp        string
	StreamId         string
}

type RoomMemberDemotedFromOwnerMessage struct {
	DemotedByUserId  int64
	DemotedUserId int64
	Id               string
	Timestamp        string
	StreamId         string
}

