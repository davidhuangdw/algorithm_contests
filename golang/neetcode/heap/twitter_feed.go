package heap

import "container/heap"

type TimedTweet struct {
	t, id int
}

type Twitter struct {
	time       int
	userTweets map[int][]TimedTweet
	userFollow map[int]map[int]bool
}

func ConstructorTw() Twitter {
	return Twitter{0, make(map[int][]TimedTweet), make(map[int]map[int]bool)}
}

func (tw *Twitter) PostTweet(userId int, tweetId int) {
	tw.userTweets[userId] = append(tw.userTweets[userId], TimedTweet{tw.time, tweetId})
	tw.time++
}

func (tw *Twitter) GetNewsFeed(userId int) (twIds []int) {
	type UserTw struct {
		tw      TimedTweet
		user, i int
	}
	maxHeap := NewHeap(func(a, b UserTw) bool { return a.tw.t > b.tw.t })
	if _, ok := tw.userFollow[userId]; !ok {
		tw.userFollow[userId] = make(map[int]bool)
	}
	tw.userFollow[userId][userId] = true
	for followee, _ := range tw.userFollow[userId] {
		n := len(tw.userTweets[followee])
		if n > 0 {
			heap.Push(maxHeap, UserTw{tw.userTweets[followee][n-1], followee, n - 1})
		}
	}
	for k := 10; k > 0 && maxHeap.Len() > 0; k-- {
		ut := heap.Pop(maxHeap).(UserTw)
		twIds = append(twIds, ut.tw.id)
		if ut.i > 0 {
			heap.Push(maxHeap, UserTw{tw.userTweets[ut.user][ut.i-1], ut.user, ut.i - 1})
		}
	}
	return twIds
}

func (tw *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := tw.userFollow[followerId]; !ok {
		tw.userFollow[followerId] = make(map[int]bool)
	}
	tw.userFollow[followerId][followeeId] = true
}

func (tw *Twitter) Unfollow(followerId int, followeeId int) {
	delete(tw.userFollow[followerId], followeeId)
}
