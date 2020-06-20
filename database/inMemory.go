package database

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/vadim-dmitriev/sittme/state"
	"github.com/vadim-dmitriev/sittme/stream"
)

// InMemory инплементирует интерфейс Databaser.
// В качестве хранилища выступает оперативная память
type InMemory struct {
	streams []stream.Stream
	sync.RWMutex
}

// NewInMemory создает новый экземпляр InMemory и
// возвращает указатель
func NewInMemory() Databaser {
	inMemory := InMemory{
		streams: make([]stream.Stream, 0),
	}

	return &inMemory
}

// Insert сохраняет новый объект трансляции
func (im *InMemory) Insert(newStream stream.Stream) error {
	im.Lock()
	defer im.Unlock()

	im.streams = append(im.streams, newStream)

	return nil
}

// Select производит поиск трансляции с заданным идентификатором uuid.
// Возвращает error в случае, если трансляция не найдена
func (im *InMemory) Select(uuid uuid.UUID) (stream.Stream, error) {
	im.RLock()
	defer im.RUnlock()

	for _, stream := range im.streams {
		if stream.UUID == uuid {
			return stream, nil
		}
	}
	return stream.New(), fmt.Errorf("stream '%s' not found", uuid.String())

}

// SelectAll возвращает список объектов всех существующих трансляций
func (im *InMemory) SelectAll() ([]stream.Stream, error) {
	im.RLock()
	defer im.RUnlock()

	return im.streams, nil
}

// Delete удаляет объект трансляции с заданным идентификатором uuid
func (im *InMemory) Delete(uuid uuid.UUID) error {
	im.Lock()
	defer im.Unlock()

	for i, stream := range im.streams {
		if stream.UUID == uuid {
			close(im.streams[i].StateChan)
			im.streams[i] = im.streams[len(im.streams)-1]
			im.streams = im.streams[:len(im.streams)-1]
			break
		}
	}

	return nil
}

// Update изменяет состояние трансляции с идентификатором uuid на новое newState
func (im *InMemory) Update(uuid uuid.UUID, newState state.Stater) error {
	im.Lock()
	defer im.Unlock()

	for i, stream := range im.streams {
		if stream.UUID == uuid {
			im.streams[i].SetState(newState)
			break
		}
	}

	return nil
}
