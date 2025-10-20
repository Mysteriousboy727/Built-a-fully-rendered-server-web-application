package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"gofr.dev/pkg/gofr"
)

type SystemMetrics struct {
	Timestamp    time.Time `json:"timestamp"`
	MemoryUsage  string    `json:"memory_usage"`
	ActiveJobs   int       `json:"active_jobs"`
	JobsExecuted int       `json:"jobs_executed"`
}

type EmailNotification struct {
	To      string    `json:"to"`
	Subject string    `json:"subject"`
	Body    string    `json:"body"`
	SentAt  time.Time `json:"sent_at"`
}

var (
	jobCounter     int
	metricsHistory []SystemMetrics
	emailQueue     []EmailNotification
)

func main() {
	app := gofr.New()

	// System monitoring: every 30 seconds
	app.AddCronJob("*/30 * * * * *", "system-monitor", monitorSystemHealth)

	// Data cleanup: every hour
	app.AddCronJob("0 * * * *", "data-cleanup", cleanupOldLogs)

	// Daily report generation: 9 AM every day
	app.AddCronJob("0 9 * * *", "daily-report", generateDailyReport)

	// Email queue processor: every 2 minutes
	app.AddCronJob("*/2 * * * *", "email-processor", processEmailQueue)

	// Database backup simulation: every 6 hours
	app.AddCronJob("0 */6 * * *", "db-backup", performDatabaseBackup)

	// API health check: every 15 seconds
	app.AddCronJob("*/15 * * * * *", "api-health-check", checkAPIHealth)

	// Add REST endpoints for monitoring
	app.GET("/metrics", getMetrics)
	app.GET("/email-queue", getEmailQueue)
	app.POST("/schedule-email", scheduleEmail)

	app.Run()
}

// Monitor system health and collect metrics
func monitorSystemHealth(ctx *gofr.Context) {
	metrics := SystemMetrics{
		Timestamp:    time.Now(),
		MemoryUsage:  "125MB", // In production, use runtime.ReadMemStats()
		ActiveJobs:   6,
		JobsExecuted: jobCounter,
	}

	metricsHistory = append(metricsHistory, metrics)

	// Keep only last 100 metrics
	if len(metricsHistory) > 100 {
		metricsHistory = metricsHistory[1:]
	}

	ctx.Logger.Infof("[SystemMonitor] Health Check - Memory: %s, Jobs Executed: %d",
		metrics.MemoryUsage, metrics.JobsExecuted)

	jobCounter++
}

// Cleanup old log files and temporary data
func cleanupOldLogs(ctx *gofr.Context) {
	cutoffTime := time.Now().Add(-24 * time.Hour)
	cleaned := 0

	// Simulate cleanup of old metrics
	var newHistory []SystemMetrics
	for _, metric := range metricsHistory {
		if metric.Timestamp.After(cutoffTime) {
			newHistory = append(newHistory, metric)
		} else {
			cleaned++
		}
	}
	metricsHistory = newHistory

	ctx.Logger.Infof("[DataCleanup] Cleaned %d old records older than 24 hours", cleaned)
}

// Generate daily performance report
func generateDailyReport(ctx *gofr.Context) {
	report := fmt.Sprintf(`
	=== Daily System Report ===
	Generated: %s
	Total Jobs Executed: %d
	Metrics Collected: %d
	Emails Queued: %d
	System Status: Healthy
	========================
	`, time.Now().Format("2006-01-02 15:04:05"),
		jobCounter, len(metricsHistory), len(emailQueue))

	// In production, save to file or send via email
	ctx.Logger.Infof("[DailyReport] %s", report)

	// Queue email notification
	emailQueue = append(emailQueue, EmailNotification{
		To:      "admin@example.com",
		Subject: "Daily System Report",
		Body:    report,
		SentAt:  time.Now(),
	})
}

// Process queued emails
func processEmailQueue(ctx *gofr.Context) {
	if len(emailQueue) == 0 {
		ctx.Logger.Infof("[EmailProcessor] No emails in queue")
		return
	}

	processed := 0
	for _, email := range emailQueue {
		// Simulate email sending
		ctx.Logger.Infof("[EmailProcessor] Sending email to %s - Subject: %s",
			email.To, email.Subject)
		processed++
	}

	// Clear queue after processing
	emailQueue = []EmailNotification{}
	ctx.Logger.Infof("[EmailProcessor] Processed %d emails", processed)
}

// Simulate database backup
func performDatabaseBackup(ctx *gofr.Context) {
	backupFile := fmt.Sprintf("backup_%s.sql", time.Now().Format("20060102_150405"))

	// Simulate backup process
	ctx.Logger.Infof("[DatabaseBackup] Starting backup to %s", backupFile)

	// In production, perform actual backup using pg_dump, mysqldump, etc.
	time.Sleep(100 * time.Millisecond) // Simulate backup time

	ctx.Logger.Infof("[DatabaseBackup] ✅ Backup completed: %s (Size: 2.4GB)", backupFile)
}

// Check external API health
func checkAPIHealth(ctx *gofr.Context) {
	// Simulate API health check
	endpoints := []string{
		"https://api.service1.com/health",
		"https://api.service2.com/status",
		"https://api.service3.com/ping",
	}

	for _, endpoint := range endpoints {
		// In production, make actual HTTP requests
		status := "UP"
		responseTime := "45ms"

		ctx.Logger.Infof("[APIHealthCheck] %s - Status: %s, Response: %s",
			endpoint, status, responseTime)
	}
}

// REST endpoint: Get metrics
func getMetrics(ctx *gofr.Context) (interface{}, error) {
	return map[string]interface{}{
		"total_metrics":   len(metricsHistory),
		"jobs_executed":   jobCounter,
		"last_10_metrics": getLastNMetrics(10),
	}, nil
}

// REST endpoint: Get email queue
func getEmailQueue(ctx *gofr.Context) (interface{}, error) {
	return map[string]interface{}{
		"queue_size": len(emailQueue),
		"emails":     emailQueue,
	}, nil
}

// REST endpoint: Schedule new email
func scheduleEmail(ctx *gofr.Context) (interface{}, error) {
	var email EmailNotification
	if err := ctx.Bind(&email); err != nil {
		return nil, err
	}

	email.SentAt = time.Now()
	emailQueue = append(emailQueue, email)

	ctx.Logger.Infof("[API] Email scheduled for %s", email.To)

	return map[string]interface{}{
		"status":  "queued",
		"message": "Email added to queue",
	}, nil
}

// Helper: Get last N metrics
func getLastNMetrics(n int) []SystemMetrics {
	if len(metricsHistory) < n {
		return metricsHistory
	}
	return metricsHistory[len(metricsHistory)-n:]
}

// Save metrics to file (optional)
func saveMetricsToFile() error {
	data, err := json.MarshalIndent(metricsHistory, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("metrics.json", data, 0644)
}
