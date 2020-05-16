package reddit

import (
	"time"

	"github.com/pkg/errors"
	"github.com/thecsw/mira"
)

// The ProtocolZero client for Reddit
type P0Reddit struct {
	Reddit *mira.Reddit
}

// Login to reddit using package credentials
func Login(clientId, clientSecret, username, password string) (*P0Reddit, error) {
	obj := &P0Reddit{}
	var err error
	obj.Reddit, err = mira.Init(mira.Credentials{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Username:     username,
		Password:     password,
		UserAgent:    "p0",
	})
	return obj, errors.Wrap(err, "failed to auth with reddit, check credentials")
}

// ClearPosts deletes all submissions on your account
// You can pass a <-chan string channel so the function
// will send you recently deleted submissions' IDs
// If your channel is fully buffered, the execution will
// be blocked until you free your channel
// Deletion is permanent via Reddit API
func (r *P0Reddit) ClearPosts(c chan<- string) {
	myName := r.Reddit.Creds.Username
	go func() {
		for {
			// Get the maximum of 100 submissions
			posts, err := r.Reddit.Redditor(myName).Submissions("new", "all", 100)
			// If Reddit API throws an error, just sleep for constant
			// time and try next call
			if err != nil {
				time.Sleep(5 * time.Second)
			}
			// Break the stuff if returned fewer than 1 post
			if len(posts) < 1 {
				break
			}
			// Walk through all the submissions and delete them
			for _, post := range posts {
				err = r.Reddit.Submission(post.GetId()).Delete()
				if err != nil {
					// Continue a failed call
					// We will catch the submission in the next call
					continue
				}
				// If channel exists, pipe the deleted ID
				// If passed nil, then don't pipe the result
				// If the channel is fully buffered, then the
				// call will just halt and wait
				if c != nil {
					c <- post.GetId()
				}
			}
			// Sleep for 2 seconds to avoid call freeze
			time.Sleep(2 * time.Second)
		}
		if c != nil {
			c <- "QUIT"
		}
	}()
}

// ClearComments deletes all comments on your account
// You can pass a <-chan string channel so the function
// will send you recently deleted comments' IDs
// If your channel is fully buffered, the execution will
// be blocked until you free your channel
// Deletion is permanent via Reddit API
func (r *P0Reddit) ClearComments(c chan<- string) {
	myName := r.Reddit.Creds.Username
	go func() {
		for {
			// Get the maximum of 100 submissions
			comments, err := r.Reddit.Redditor(myName).Comments("new", "all", 100)
			// If Reddit API throws an error, just sleep for constant
			// time and try next call
			if err != nil {
				time.Sleep(5 * time.Second)
			}
			// Break the stuff if returned fewer than 1 comment
			if len(comments) < 1 {
				break
			}
			// Walk through all the submissions and delete them
			for _, comment := range comments {
				err = r.Reddit.Comment(comment.GetId()).Delete()
				if err != nil {
					// Continue a failed call
					// We will catch the submission in the next call
					continue
				}
				// If channel exists, pipe the deleted ID
				// If passed nil, then don't pipe the result
				// If the channel is fully buffered, then the
				// call will just halt and wait
				if c != nil {
					c <- comment.GetId()
				}
			}
			// Sleep for 2 seconds to avoid call freeze
			time.Sleep(2 * time.Second)
		}
		if c != nil {
			c <- "QUIT"
		}
	}()
}
