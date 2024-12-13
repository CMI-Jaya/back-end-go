package utils

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

// Fungsi untuk mengirimkan notifikasi WhatsApp
func SendWhatsAppNotification(to string, body string) error {
	// Validasi input
	from := os.Getenv("TWILIO_WHATSAPP_FROM")
	if from == "" || to == "" {
		return fmt.Errorf("invalid 'from' or 'to' number")
	}

	// Buat Twilio client
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})

	// Siapkan parameter pesan
	messageParams := &openapi.CreateMessageParams{}
	messageParams.SetTo(to)
	messageParams.SetFrom(from)
	messageParams.SetBody(body)

	// Kirim pesan
	resp, err := client.Api.CreateMessage(messageParams)
	if err != nil {
		return fmt.Errorf("failed to send WhatsApp message: %w", err)
	}

	// Log hasil pengiriman
	fmt.Printf("Message sent successfully: SID %s\n", *resp.Sid)
	return nil
}
