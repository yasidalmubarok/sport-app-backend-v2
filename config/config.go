package config

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"os"
	"sport-app-backend/helper"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// RedisClient adalah instance global Redis
var RedisClient *redis.Client
var ctx = context.Background()

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectDB() *gorm.DB {
	LoadEnv()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return db
}

// ConnectRedis menghubungkan aplikasi ke Redis
func ConnectRedis() *redis.Client {
	LoadEnv() // Memuat variabel lingkungan dari .env

	addr := fmt.Sprintf("%s:%s",
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
	)

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: os.Getenv("REDIS_PASSWORD"), // Bisa kosong jika tidak pakai password
		DB:       0,                           // Default database Redis
	})

	// Tes koneksi Redis
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed connect to Redis:", err)
	}

	fmt.Println("Connected to Redis!")
	return RedisClient
}

func SendEmail(to string, subject string, body string) helper.Error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")

	log.Println("Trying to connect to SMTP server:", smtpHost, "on port", smtpPort)

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	// Buat koneksi ke SMTP server
	conn, err := net.Dial("tcp", smtpHost+":"+smtpPort)
	if err != nil {
		log.Println("Failed to connect to SMTP server:", err)
		return helper.NewInternalServerError("failed to send emai: " + err.Error())
	}
	defer conn.Close()

	// Buat client SMTP
	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		log.Println("Failed to create SMTP client:", err)
		return helper.NewInternalServerError("failed to send ema: " + err.Error())
	}
	defer client.Close()

	// Mulai STARTTLS
	tlsConfig := &tls.Config{ServerName: smtpHost}
	if err = client.StartTLS(tlsConfig); err != nil {
		log.Println("Failed to start TLS:", err)
		return helper.NewInternalServerError("failed to send em: " + err.Error())
	}

	// Autentikasi
	if err = client.Auth(auth); err != nil {
		log.Println("SMTP authentication failed:", err)
		return helper.NewInternalServerError("failed to send e: " + err.Error())
	}

	// Set pengirim & penerima
	if err = client.Mail(smtpUser); err != nil {
		log.Println("Failed to set sender:", err)
		return helper.NewInternalServerError("failed to send: " + err.Error())
	}
	if err = client.Rcpt(to); err != nil {
		log.Println("Failed to set recipient:", err)
		return helper.NewInternalServerError("failed to sen: " + err.Error())
	}

	// Kirim data email
	w, err := client.Data()
	if err != nil {
		log.Println("Failed to send email data:", err)
		return helper.NewInternalServerError("failed to se: " + err.Error())
	}

	msg := []byte("From: " + smtpUser + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		"<html><body><h2>Reset Password Request</h2>" +
		"<p>Hi,</p><p>You requested to reset your password. Use the OTP code below:</p>" +
		"<h3 style='color: #007bff;'>" + body + "</h3>" +
		"<p>This OTP will expire in 5 minutes.</p>" +
		"<br><p>If you didn't request this, please ignore this email.</p>" +
		"<p>Best regards,<br><b>Sport App Team</b></p></body></html>\r\n")

	_, err = w.Write(msg)
	if err != nil {
		log.Println("Failed to write email data:", err)
		return helper.NewInternalServerError("failed to s: " + err.Error())
	}

	err = w.Close()
	if err != nil {
		log.Println("Failed to close writer:", err)
		return helper.NewInternalServerError("failed to: " + err.Error())
	}

	client.Quit()
	log.Println("Email successfully sent to", to)
	return nil
}
