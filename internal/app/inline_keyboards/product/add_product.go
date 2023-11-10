package product

// AddProductInlineButton предоставляем пользователю с помощью inline кнопки получить
// доступ к функционалу заложенму в команду добавления продукта
type AddProductInlineButton struct {
	cmd command
}

// NewAddProductInlineButton ctor
func NewAddProductInlineButton(cmd command) *AddProductInlineButton {
	return &AddProductInlineButton{cmd: cmd}
}

// Text возвращает название кнопки
func (btn *AddProductInlineButton) Text() string {
	return "➕ Добавить"
}

// Callback возвращает команду которая должна отработать
func (btn *AddProductInlineButton) Callback() string {
	return btn.cmd.Alias()
}
