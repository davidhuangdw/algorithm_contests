package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwitter(t *testing.T) {
	tests := []struct {
		name     string
		actions  []string
		params   [][]int
		expected [][]int
	}{
		{
			name:     "basic post and get feed",
			actions:  []string{"Twitter", "postTweet", "getNewsFeed"},
			params:   [][]int{{}, {1, 5}, {1}},
			expected: [][]int{{}, {}, {5}},
		},
		{
			name:     "multiple posts same user",
			actions:  []string{"Twitter", "postTweet", "postTweet", "getNewsFeed"},
			params:   [][]int{{}, {1, 5}, {1, 3}, {1}},
			expected: [][]int{{}, {}, {}, {3, 5}},
		},
		{
			name:     "follow and get feed",
			actions:  []string{"Twitter", "postTweet", "follow", "getNewsFeed"},
			params:   [][]int{{}, {2, 6}, {1, 2}, {1}},
			expected: [][]int{{}, {}, {}, {6}},
		},
		{
			name:     "multiple users follow",
			actions:  []string{"Twitter", "postTweet", "postTweet", "follow", "follow", "getNewsFeed"},
			params:   [][]int{{}, {1, 5}, {2, 6}, {1, 2}, {1, 3}, {1}},
			expected: [][]int{{}, {}, {}, {}, {}, {6, 5}},
		},
		{
			name:     "unfollow removes tweets",
			actions:  []string{"Twitter", "postTweet", "follow", "unfollow", "getNewsFeed"},
			params:   [][]int{{}, {2, 6}, {1, 2}, {1, 2}, {1}},
			expected: [][]int{{}, {}, {}, {}, nil},
		},
		{
			name: "max 10 tweets in feed",
			actions: []string{"Twitter", "postTweet", "postTweet", "postTweet", "postTweet",
				"postTweet", "postTweet", "postTweet", "postTweet", "postTweet",
				"postTweet", "postTweet", "getNewsFeed"},
			params: [][]int{{}, {1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6},
				{1, 7}, {1, 8}, {1, 9}, {1, 10}, {1, 11}, {1}},
			expected: [][]int{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
				{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}},
		},
		{
			name:     "empty feed for non-existent user",
			actions:  []string{"Twitter", "getNewsFeed"},
			params:   [][]int{{}, {999}},
			expected: [][]int{{}, nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var twitter *Twitter
			var results [][]int

			for i, action := range tt.actions {
				switch action {
				case "Twitter":
					tw := ConstructorTw()
					twitter = &tw
					results = append(results, []int{})
				case "postTweet":
					twitter.PostTweet(tt.params[i][0], tt.params[i][1])
					results = append(results, []int{})
				case "getNewsFeed":
					feed := twitter.GetNewsFeed(tt.params[i][0])
					results = append(results, feed)
				case "follow":
					twitter.Follow(tt.params[i][0], tt.params[i][1])
					results = append(results, []int{})
				case "unfollow":
					twitter.Unfollow(tt.params[i][0], tt.params[i][1])
					results = append(results, []int{})
				}
			}

			assert.Equal(t, tt.expected, results, "Test case: %s", tt.name)
		})
	}
}
