package routes

import (
	"context"
	"fmt"
	"log"

	"github.com/lerryjay/email-service/pkg/pb"
	mail "github.com/xhit/go-simple-mail/v2"
)

func (config Smtp) Send(recipient string, from string, subject string, text string) bool {

	server := mail.NewSMTPClient()
	server.Host = config.Host
	server.Port = int(config.Port)
	server.Username = config.Username
	server.Password = config.Password
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Create email
	email := mail.NewMSG()
	email.SetFrom(from)
	email.AddTo(recipient)
	email.AddCc(from)
	email.SetSubject(subject)

	email.SetBody(mail.TextHTML, text)
	// email.AddAttachment("super_cool_file.png")

	// Send email
	err = email.Send(smtpClient)
	if err != nil {
		log.Panicln("Send Mail Erro:", err)
		return false
	}

	return true
}

func (smtp Smtp) SendEmail(ctx context.Context, req *pb.SendEmailRequest) (*pb.SendEmailResponse, error) {
	sent := smtp.Send(req.Recipient, req.SenderEmail, req.Subject, req.Template)
	if sent {
		return &pb.SendEmailResponse{Send: 1, Fails: 0}, nil
	}
	fmt.Println("Email Sent Successfully!")

	return &pb.SendEmailResponse{Send: 0, Fails: 1}, nil
}

func (smtp Smtp) SendBulkEmail(ctx context.Context, req *pb.SendBulkEmailRequest) (*pb.SendBulkEmailResponse, error) {
	fails := 0
	for x := 0; x < len(req.Recipients); x++ {
		sent := smtp.Send(req.Recipients[x], req.SenderEmail, req.Subject, req.Template)
		if !sent {
			fails++
		}
	}

	return &pb.SendBulkEmailResponse{Send: int32(len(req.Recipients) - fails), Fails: int32(fails)}, nil

}
