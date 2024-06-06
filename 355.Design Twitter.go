package main

/*
import (
	"fmt"
	"sort"
	"time"
)


type tweet struct {
	id     int
	userId int
	time   int64
}

type orderedQueue struct {
	data []*tweet
}

func (o *orderedQueue) init(allTweets map[int][]*tweet) map[int][]*tweet {
	o.data = []*tweet{}

	for index, pt := range allTweets {
		if len(pt) == 0 {
			continue
		}
		o.insert(pt[len(pt)-1])
		allTweets[index] = pt[:len(pt)-1]
	}
	return allTweets
}

func (o *orderedQueue) getLast() *tweet {
	length := len(o.data)
	if length == 0 {
		return &tweet{id: 0}
	}

	answer := o.data[length-1]
	o.data = o.data[:length-1]
	return answer
}

func (o *orderedQueue) insert(t *tweet) {
	o.data = append(o.data, t)
	sort.Slice(o.data, func(i, j int) bool {
		return o.data[i].time < o.data[j].time
	})
}

type twitterUser struct {
	id          int
	followeeIds map[int]struct{}
	tweets      []*tweet
}

func (t *twitterUser) init(id int) {
	t.id = id
	t.followeeIds = map[int]struct{}{id: {}}
	t.tweets = []*tweet{}
}

func (t *twitterUser) follow(id int) {
	_, ok := t.followeeIds[id]
	if !ok {
		t.followeeIds[id] = struct{}{}
	}
}

func (t *twitterUser) unfollow(id int) {
	_, ok := t.followeeIds[id]
	if ok && id != t.id {
		delete(t.followeeIds, id)
	}
}

func (t *twitterUser) postTweet(tweetId int) {
	t.tweets = append(t.tweets, &tweet{id: tweetId, userId: t.id, time: time.Now().UnixNano()})
}

type Twitter struct {
	users map[int]*twitterUser
}

func Constructor() Twitter {
	return Twitter{
		users: map[int]*twitterUser{},
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	user, ok := this.users[userId]
	if ok {
		user.postTweet(tweetId)
	} else {
		newUser := &twitterUser{id: userId, followeeIds: map[int]struct{}{userId: {}}, tweets: []*tweet{}}
		newUser.postTweet(tweetId)
		this.users[userId] = newUser
	}
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	f, ok := this.users[followerId]
	if ok {
		f.follow(followeeId)
	} else {
		newf := twitterUser{id: followerId, followeeIds: map[int]struct{}{}, tweets: []*tweet{}}
		newf.follow(followeeId)
		this.users[followerId] = &newf
	}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	f, ok := this.users[followerId]
	if !ok {
		return
	}
	f.unfollow(followeeId)
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	user, ok := this.users[userId]
	if !ok {
		return []int{}
	}

	allTweets := map[int][]*tweet{userId: user.tweets}
	otherIds := user.followeeIds
	for id := range otherIds {
		otherUser, ok := this.users[id]
		if !ok {
			continue
		}
		allTweets[id] = otherUser.tweets
	}

	answer := getFirst10Tweet(allTweets)
	return answer
}

func getFirst10Tweet(allTweets map[int][]*tweet) []int {
	orderedQueue := orderedQueue{}
	orderedQueue.init(allTweets)
	fmt.Println(orderedQueue.data)

	// get tweets ordered by time
	answer := []int{}
	for len(answer) < 10 && len(allTweets) > 0 {
		currentTweet := orderedQueue.getLast()
		if currentTweet.id == 0 {
			return answer
		}
		answer = append(answer, currentTweet.id)

		userid := currentTweet.userId
		ut, ok := allTweets[userid]
		if !ok {
			continue
		} else if len(ut) == 0 {
			delete(allTweets, userid)
			continue
		}

		orderedQueue.insert(ut[len(ut)-1])
		allTweets[userid] = ut[:len(ut)-1]
	}
	return answer
}
*/
