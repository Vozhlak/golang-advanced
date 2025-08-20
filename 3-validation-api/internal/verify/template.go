package verify

import (
	"bytes"
	"html/template"
)

const verifyEmailTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>Verify Your Email</title>
    <style>
        body { font-family: Arial, sans-serif; text-align: center; margin-top: 50px; }
        .container { max-width: 600px; margin: 0 auto; padding: 20px; }
        .btn { 
            display: inline-block; 
            padding: 12px 24px; 
            background-color: #007BFF; 
            color: white; 
            text-decoration: none; 
            border-radius: 5px; 
            font-size: 16px; 
            margin: 20px 0;
        }
        .footer { margin-top: 40px; color: #777; font-size: 14px; }
    </style>
</head>
<body>
    <div class="container">
        <h1>📧 Подтвердите ваш email</h1>
        <p>Чтобы завершить регистрацию, нажмите кнопку ниже:</p>
        <a href="{{.URL}}" class="btn">Подтвердить email</a>
        <p><small>Ссылка действует 24 часа.</small></p>
        <div class="footer">
            Если вы не регистрировались у нас — просто проигнорируйте это письмо.
        </div>
    </div>
</body>
</html>
`

var verifyEmailTmpl = template.Must(template.New("email").Parse(verifyEmailTemplate))

func renderVerifyEmail(url string) (string, error) {
	var buf bytes.Buffer
	err := verifyEmailTmpl.Execute(&buf, struct{ URL string }{URL: url})
	return buf.String(), err
}
