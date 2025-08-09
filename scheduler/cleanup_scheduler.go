package scheduler

import (
	"context"
	"database/sql"
	"log"
	"time"
	"golang_jwt/repository"
)

type CleanupScheduler struct {
	userRepo repository.UserRepository
	db       *sql.DB
	interval time.Duration
}

func NewCleanupScheduler(userRepo repository.UserRepository, db *sql.DB) *CleanupScheduler {
	return &CleanupScheduler{
		userRepo: userRepo,
		db:       db,
		interval: 24 * time.Hour,
	}
}

func (s *CleanupScheduler) SetInterval(duration time.Duration) {
	s.interval = duration
}

func (s *CleanupScheduler) Start() {
	log.Println("Starting cleanup scheduler...")
	
	s.runCleanup()
	
	ticker := time.NewTicker(s.interval)
	
	go func() {
		for {
			select {
			case <-ticker.C:
				s.runCleanup()
			}
		}
	}()
}

func (s *CleanupScheduler) runCleanup() {
	ctx := context.Background()
	
	tx, err := s.db.Begin()
	if err != nil {
		log.Printf("Error starting cleanup transaction: %v", err)
		return
	}
	
	err = s.userRepo.DeleteExpiredSessions(ctx, tx)
	if err != nil {
		tx.Rollback()
		log.Printf("Error cleaning expired sessions: %v", err)
	} else {
		tx.Commit()
		log.Println("Expired sessions cleanup completed")
	}
}

func (s *CleanupScheduler) ManualCleanup() error {
	log.Println("Manual cleanup triggered")
	
	ctx := context.Background()
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	
	err = s.userRepo.DeleteExpiredSessions(ctx, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	
	tx.Commit()
	log.Println("Manual cleanup completed")
	return nil
}