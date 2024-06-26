// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package githubql

import (
	"context"
	"time"

	"github.com/Khan/genqlient/graphql"
)

// __contributionsViewInput is used internally by genqlient
type __contributionsViewInput struct {
	Username string    `json:"username"`
	From     time.Time `json:"from"`
	To       time.Time `json:"to"`
}

// GetUsername returns __contributionsViewInput.Username, and is useful for accessing the field via an interface.
func (v *__contributionsViewInput) GetUsername() string { return v.Username }

// GetFrom returns __contributionsViewInput.From, and is useful for accessing the field via an interface.
func (v *__contributionsViewInput) GetFrom() time.Time { return v.From }

// GetTo returns __contributionsViewInput.To, and is useful for accessing the field via an interface.
func (v *__contributionsViewInput) GetTo() time.Time { return v.To }

// contributionsViewResponse is returned by contributionsView on success.
type contributionsViewResponse struct {
	// Lookup a user by login.
	User *contributionsViewUser `json:"user"`
}

// GetUser returns contributionsViewResponse.User, and is useful for accessing the field via an interface.
func (v *contributionsViewResponse) GetUser() *contributionsViewUser { return v.User }

// contributionsViewUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A user is an individual's account on GitHub that owns repositories and can make new content.
type contributionsViewUser struct {
	// The collection of contributions this user has made to different repositories.
	ContributionsCollection contributionsViewUserContributionsCollection `json:"contributionsCollection"`
}

// GetContributionsCollection returns contributionsViewUser.ContributionsCollection, and is useful for accessing the field via an interface.
func (v *contributionsViewUser) GetContributionsCollection() contributionsViewUserContributionsCollection {
	return v.ContributionsCollection
}

// contributionsViewUserContributionsCollection includes the requested fields of the GraphQL type ContributionsCollection.
// The GraphQL type's documentation follows.
//
// A contributions collection aggregates contributions such as opened issues and commits created by a user.
type contributionsViewUserContributionsCollection struct {
	// A calendar of this user's contributions on GitHub.
	ContributionCalendar contributionsViewUserContributionsCollectionContributionCalendar `json:"contributionCalendar"`
}

// GetContributionCalendar returns contributionsViewUserContributionsCollection.ContributionCalendar, and is useful for accessing the field via an interface.
func (v *contributionsViewUserContributionsCollection) GetContributionCalendar() contributionsViewUserContributionsCollectionContributionCalendar {
	return v.ContributionCalendar
}

// contributionsViewUserContributionsCollectionContributionCalendar includes the requested fields of the GraphQL type ContributionCalendar.
// The GraphQL type's documentation follows.
//
// A calendar of contributions made on GitHub by a user.
type contributionsViewUserContributionsCollectionContributionCalendar struct {
	// Determine if the color set was chosen because it's currently Halloween.
	IsHalloween bool `json:"isHalloween"`
	// A list of the weeks of contributions in this calendar.
	Weeks []contributionsViewUserContributionsCollectionContributionCalendarWeeksContributionCalendarWeek `json:"weeks"`
}

// GetIsHalloween returns contributionsViewUserContributionsCollectionContributionCalendar.IsHalloween, and is useful for accessing the field via an interface.
func (v *contributionsViewUserContributionsCollectionContributionCalendar) GetIsHalloween() bool {
	return v.IsHalloween
}

// GetWeeks returns contributionsViewUserContributionsCollectionContributionCalendar.Weeks, and is useful for accessing the field via an interface.
func (v *contributionsViewUserContributionsCollectionContributionCalendar) GetWeeks() []contributionsViewUserContributionsCollectionContributionCalendarWeeksContributionCalendarWeek {
	return v.Weeks
}

// contributionsViewUserContributionsCollectionContributionCalendarWeeksContributionCalendarWeek includes the requested fields of the GraphQL type ContributionCalendarWeek.
// The GraphQL type's documentation follows.
//
// A week of contributions in a user's contribution graph.
type contributionsViewUserContributionsCollectionContributionCalendarWeeksContributionCalendarWeek struct {
	// The days of contributions in this week.
	ContributionDays []contributionsViewUserContributionsCollectionContributionCalendarWeeksContributionCalendarWeekContributionDaysContributionCalendarDay `json:"contributionDays"`
}

// GetContributionDays returns contributionsViewUserContributionsCollectionContributionCalendarWeeksContributionCalendarWeek.ContributionDays, and is useful for accessing the field via an interface.
func (v *contributionsViewUserContributionsCollectionContributionCalendarWeeksContributionCalendarWeek) GetContributionDays() []contributionsViewUserContributionsCollectionContributionCalendarWeeksContributionCalendarWeekContributionDaysContributionCalendarDay {
	return v.ContributionDays
}

// contributionsViewUserContributionsCollectionContributionCalendarWeeksContributionCalendarWeekContributionDaysContributionCalendarDay includes the requested fields of the GraphQL type ContributionCalendarDay.
// The GraphQL type's documentation follows.
//
// Represents a single day of contributions on GitHub by a user.
type contributionsViewUserContributionsCollectionContributionCalendarWeeksContributionCalendarWeekContributionDaysContributionCalendarDay struct {
	// How many contributions were made by the user on this day.
	ContributionCount int `json:"contributionCount"`
	// The hex color code that represents how many contributions were made on this day compared to others in the calendar.
	Color string `json:"color"`
}

// GetContributionCount returns contributionsViewUserContributionsCollectionContributionCalendarWeeksContributionCalendarWeekContributionDaysContributionCalendarDay.ContributionCount, and is useful for accessing the field via an interface.
func (v *contributionsViewUserContributionsCollectionContributionCalendarWeeksContributionCalendarWeekContributionDaysContributionCalendarDay) GetContributionCount() int {
	return v.ContributionCount
}

// GetColor returns contributionsViewUserContributionsCollectionContributionCalendarWeeksContributionCalendarWeekContributionDaysContributionCalendarDay.Color, and is useful for accessing the field via an interface.
func (v *contributionsViewUserContributionsCollectionContributionCalendarWeeksContributionCalendarWeekContributionDaysContributionCalendarDay) GetColor() string {
	return v.Color
}

// The query or mutation executed by contributionsView.
const contributionsView_Operation = `
query contributionsView ($username: String!, $from: DateTime!, $to: DateTime!) {
	user(login: $username) {
		contributionsCollection(from: $from, to: $to) {
			contributionCalendar {
				isHalloween
				weeks {
					contributionDays {
						contributionCount
						color
					}
				}
			}
		}
	}
}
`

func contributionsView(
	ctx context.Context,
	client graphql.Client,
	username string,
	from time.Time,
	to time.Time,
) (*contributionsViewResponse, error) {
	req := &graphql.Request{
		OpName: "contributionsView",
		Query:  contributionsView_Operation,
		Variables: &__contributionsViewInput{
			Username: username,
			From:     from,
			To:       to,
		},
	}
	var err error

	var data contributionsViewResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
