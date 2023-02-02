package types

// структура чтобы конфигурировать информацию об отправке писем
type MailInfo struct {
	From     string
	User     string
	Password string
	Addr     string
	Host     string
	FrontURL string
}
