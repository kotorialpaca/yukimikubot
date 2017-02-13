package services

import (
	"sync"

	"github.com/rylio/ytdl"
)

type Queue struct {
	Queue []ytdl.VideoInfo
	mutex sync.RWMutex
}

func NewQueue() *Queue {

	return &Queue{
		Queue: make([]ytdl.VideoInfo, 0),
	}

}

func (q *Queue) Length() int {

	q.mutex.RLock()
	l := len(q.Queue)
	q.mutex.RUnlock()

	return l
}
