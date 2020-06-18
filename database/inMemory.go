package database

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/vadim-dmitriev/sittme/state"
	"github.com/vadim-dmitriev/sittme/stream"
)

type InMemory struct {
	streams    []*stream.Stream
	streamsMap map[uuid.UUID]*stream.Stream
	sync.RWMutex
}

func NewInMemory() Databaser {
	inMemory := InMemory{
		streams:    make([]*stream.Stream, 0),
		streamsMap: make(map[uuid.UUID]*stream.Stream, 0),
	}

	return &inMemory
}

func (im *InMemory) Insert(newStream *stream.Stream) error {
	im.Lock()
	defer im.Unlock()

	im.streams = append(im.streams, newStream)
	im.streamsMap[newStream.UUID] = newStream

	return nil
}

func (im *InMemory) Select(uuid uuid.UUID) (*stream.Stream, error) {
	stream, ok := im.streamsMap[uuid]
	if !ok {
		return nil, fmt.Errorf("stream %s not found", uuid.String())
	}

	return stream, nil
}

func (im *InMemory) SelectAll() []*stream.Stream {
	im.RLock()
	defer im.RUnlock()

	return im.streams
}

func (im *InMemory) Delete(uuid uuid.UUID) error {
	im.Lock()
	defer im.Unlock()

	if _, ok := im.streamsMap[uuid]; !ok {
		// элемент не был найден
		return fmt.Errorf("stream %s not found", uuid.String())
	}

	for i, stream := range im.streams {
		if stream.UUID == uuid {
			im.streams[i] = im.streams[len(im.streams)-1]
			im.streams = im.streams[:len(im.streams)-1]
			break
		}
	}

	return nil
}

func (im *InMemory) Update(uuid uuid.UUID, newStreamState state.Stater) error {
	im.Lock()
	defer im.Unlock()

	// проверка на существование трансляции с заданым uuid

	// изменение состояния этой трансляции

	return nil
}