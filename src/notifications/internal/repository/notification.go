package repository

import (
	"github.com/gocql/gocql"
	"noty-notifications/internal/database"
	"noty-notifications/internal/dto"
	"noty-notifications/internal/model"
)

type NotificationRepository interface {
	NotificationCreator
	NotificationUpdater
	NotificationFinder
}

type NotificationCreator interface {
	Create(notification model.Notification) error
}

type NotificationUpdater interface {
	Update(notification model.Notification) error
}

type NotificationFinder interface {
	FindByID(notification *dto.NotificationInfo, id string) error
}

type CassandraNotificationRepository struct {
	DB *gocql.Session
}

func NewNotificationRepository() *CassandraNotificationRepository {
	sess := database.InitCassandra()
	return &CassandraNotificationRepository{DB: sess}
}

func (rp *CassandraNotificationRepository) Create(notification model.Notification) error {
	return rp.DB.Query(
		"INSERT INTO notifications(id, name, description, strategy, notify_at, user_id) VALUES(?,?,?,?,?,?)",
		notification.ID,
		notification.Name,
		notification.Description,
		notification.Strategy,
		notification.NotifyAt.UnixNano(),
		notification.UserID,
	).Exec()
}

func (rp *CassandraNotificationRepository) Update(notification model.Notification) error {
	return rp.DB.Query(
		"UPDATE notifications SET name = ?, description = ?, strategy = ?, notify_at = ? WHERE id = ? IF EXISTS",
		notification.Name,
		notification.Description,
		notification.Strategy,
		notification.NotifyAt.UnixNano(),
		notification.ID,
	).Exec()
}

func (rp *CassandraNotificationRepository) FindByID(notification *dto.NotificationInfo, id string) error {
	if err := rp.DB.Query("SELECT * FROM notifications WHERE id=?", id).
		Consistency(gocql.One).
		Scan(
			&notification.ID,
			&notification.Name,
			&notification.Description,
			&notification.NotifyAt,
			&notification.Strategy,
			&notification.UserID,
		); err != nil {
		return err
	}

	return nil
}
