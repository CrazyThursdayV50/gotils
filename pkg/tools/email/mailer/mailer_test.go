package mailer

import (
	"testing"

	"github.com/CrazyThursdayV50/gotils/pkg/tools/email/mail"
)

func TestMailer(t *testing.T) {
	var username = "mailer tester"
	var sender = "xx@xx.com"
	var password = "xx"
	var receiver = "yy@xx.com"
	var endpoint = "smtp.xx.com"
	var port = 465
	var mailer, err = New(
		WithUsername(username),
		WithSender(sender, password),
		WithSMTP(endpoint, port),
	)
	if err != nil {
		t.Fatalf("new mailer failed: %v", err)
	}

	var to = []string{
		sender,
		receiver,
	}
	var cc = []string{
		sender,
		receiver,
	}

	var mail = mail.New(
		mail.WithFrom(sender),
		mail.WithTo(to...),
		mail.WithCc(cc...),
		mail.WithSubject("MailerTest"),
		mail.WithBody(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Reset Your Password</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f5f5f5;
            margin: 0;
            padding: 0;
        }
        
        .container {
            max-width: 600px;
            margin: 50px auto;
            background-color: white;
            padding: 40px;
            border-radius: 8px;
            box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
        }
        
        h1 {
            color: #333;
            text-align: center;
            margin-bottom: 30px;
        }
        
        .reset-link {
            display: block;
            background-color: #007bff;
            color: white;
            text-decoration: none;
            padding: 16px 32px;
            border-radius: 6px;
            text-align: center;
            font-size: 18px;
            transition: background-color 0.3s ease;
        }
        
        .reset-link:hover {
            background-color: #0056b3;
        }
        
        p {
            color: #666;
            line-height: 1.5;
            margin-top: 20px;
            text-align: center;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Reset Your Password</h1>
        <a href="https://example.com/reset-password" class="reset-link" target="_blank">Reset Password</a>
        <p>Click the button above to reset your password. This link will open in a new window.</p>
    </div>
</body>
</html>
`),
		mail.WithHTML(),
	)

	t.Logf("mail body: %s", mail.Message())
	err = mailer.Send(mail)
	if err != nil {
		t.Fatalf("send mail failed: %v", err)
	}
}
