/*
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package jira

// Jira REST API resources

type ErrorResult struct {
	ErrorMessages []string `url:"errorMessages"`
	Errors        []string `url:"errors"`
}

type GetMyPermissionsQueryOptions struct {
	ProjectKey  string   `url:"projectKey,omitempty"`
	Permissions []string `url:"permissions,comma,omitempty"`
}

type GetIssueQueryOptions struct {
	Fields     []string `url:"fields,comma,omitempty"`
	Expand     []string `url:"expand,comma,omitempty"`
	Properties []string `url:"properties,comma,omitempty"`
}

type GetIssueCommentQueryOptions struct {
	StartAt    int      `url:"startAt,omitempty"`
	MaxResults int      `url:"maxResults,omitempty"`
	OrderBy    string   `url:"orderBy,omitempty"`
	Expand     []string `url:"expand,comma,omitempty"`
}

type Permission struct {
	ID             string `json:"id"`
	Key            string `json:"key"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Description    string `json:"description"`
	HavePermission bool   `json:"havePermission"`
}

type Permissions struct {
	Permissions map[string]Permission `json:"permissions"`
}

type Project struct {
	Expand      string `json:"expand,omitempty"`
	Self        string `json:"self,omitempty"`
	ID          string `json:"id,omitempty"`
	Key         string `json:"key,omitempty"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
	Email       string `json:"email,omitempty"`
	Name        string `json:"name,omitempty"`
}

type Issue struct {
	Expand      string                 `json:"expand"`
	Self        string                 `json:"self"`
	ID          string                 `json:"id"`
	Key         string                 `json:"key"`
	Fields      IssueFields            `json:"fields"`
	Changelog   PageOfChangelogs       `json:"changelog"`
	Properties  map[string]interface{} `json:"properties"`
	Transitions []IssueTransition      `json:"transitions"`
}

type IssueFields struct {
	Status      StatusDetails  `json:"status"`
	Comment     PageOfComments `json:"comment"`
	Project     Project        `json:"project"`
	Type        IssueType      `json:"issuetype"`
	Summary     string         `json:"summary,omitempty"`
	Description string         `json:"description,omitempty"`
}

type IssueTransition struct {
	ID   string        `json:"id,omitempty"`
	Name string        `json:"name,omitempty"`
	To   StatusDetails `json:"to,omitempty"`
}

type IssueType struct {
	Self        string `json:"self,omitempty"`
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	IconURL     string `json:"iconUrl,omitempty"`
	Name        string `json:"name,omitempty"`
}

type IssueFieldsInput struct {
	Type        *IssueType `json:"issuetype,omitempty"`
	Project     *Project   `json:"project,omitempty"`
	Summary     string     `json:"summary,omitempty"`
	Description string     `json:"description,omitempty"`
}

type IssueInput struct {
	Fields     IssueFieldsInput `json:"fields,omitempty"`
	Properties []EntityProperty `json:"properties,omitempty"`
}

type IssueTransitionInput struct {
	Transition IssueTransition `json:"transition"`
}

type CreatedIssue struct {
	ID   string `json:"id"`
	Key  string `json:"key"`
	Self string `json:"self"`
}

type EntityProperty struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type StatusDetails struct {
	Self        string `json:"self"`
	Description string `json:"description"`
	IconURL     string `json:"iconUrl"`
	Name        string `json:"name"`
	ID          string `json:"id"`
}

type UserDetails struct {
	Self         string `json:"self"`
	AccountID    string `json:"accountId"`
	EmailAddress string `json:"emailAddress"`
	DisplayName  string `json:"displayName"`
	Active       bool   `json:"active"`
	TimeZone     string `json:"timeZone"`
	AccountType  string `json:"accountType"`
}

type Changelog struct {
	ID      string          `json:"id"`
	Author  UserDetails     `json:"author"`
	Created string          `json:"created"`
	Items   []ChangeDetails `json:"items"`
}

type ChangeDetails struct {
	Field      string `json:"field"`
	FieldType  string `json:"fieldtype"`
	FieldID    string `json:"fieldId"`
	From       string `json:"from"`
	FromString string `json:"fromString"`
	To         string `json:"to"`
	ToString   string `json:"toString"`
}

type Comment struct {
	Self    string      `json:"self"`
	ID      string      `json:"id"`
	Author  UserDetails `json:"author"`
	Body    string      `json:"body"`
	Created string      `json:"created"`
}

type CommentInput struct {
	Body       string           `json:"body,omitempty"`
	Properties []EntityProperty `json:"properties,omitempty"`
}

type PageOfChangelogs struct {
	StartAt    int         `json:"startAt"`
	MaxResults int         `json:"maxResults"`
	Total      int         `json:"total"`
	Histories  []Changelog `json:"histories"`
}

type PageOfComments struct {
	StartAt    int       `json:"startAt"`
	MaxResults int       `json:"maxResults"`
	Total      int       `json:"total"`
	Comments   []Comment `json:"comments"`
}
