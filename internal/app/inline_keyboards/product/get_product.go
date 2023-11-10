package product

// GetProductInlineButton предоставляем пользователю с помощью inline кнопки получить
// доступ к функционалу заложенму в команду поиска продукта
type GetProductInlineButton struct {
	cmd command
}

// NewGetProductInlineButton ctor
func NewGetProductInlineButton(cmd command) *GetProductInlineButton {
	return &GetProductInlineButton{cmd: cmd}
}

// Text возвращает название кнопки
func (btn *GetProductInlineButton) Text() string {
	return "🔎 Найти"
}

// Callback возвращает команду которая должна отработать
func (btn *GetProductInlineButton) Callback() string {
	return btn.cmd.Alias()
}
