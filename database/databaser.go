package database

import (
	"github.com/google/uuid"
	"github.com/vadim-dmitriev/sittme/state"
	"github.com/vadim-dmitriev/sittme/stream"
)

// Databaser интерфейс взаимодействия с хранилищем данных сервиса
type Databaser interface {
	// Insert записывает в хранилище новый объект хранилища.
	// В случае, если запись невозможна, возвращает ошибку
	Insert(stream.Stream) error

	// Select отдает из хранилища объект трансляции с запрашиваемым
	// идентификатором uuid. В случае, если такой трансляции не существует,
	// возвращает пустой объект трансляции и ошибку
	Select(uuid uuid.UUID) (stream.Stream, error)

	// SelectALl отдает из хранилища список всех существуешь трансляций.
	SelectAll() ([]stream.Stream, error)

	// Delete удалаяет объект трансляции с заданным идентификатором uuid
	Delete(uuid.UUID) error

	// Update обнавляет значение состояния трансляции.
	Update(uuid.UUID, state.Stater) error
}
