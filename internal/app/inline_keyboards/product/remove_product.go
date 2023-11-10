package product

// RemoveProductInlineButton предоставляем пользователю с помощью inline кнопки получить
// доступ к функционалу заложенму в команду удаления продукта
type RemoveProductInlineButton struct {
	cmd command
}

// NewRemoveProductInlineButton ctor
func NewRemoveProductInlineButton(cmd command) *RemoveProductInlineButton {
	return &RemoveProductInlineButton{cmd: cmd}
}

// Text возвращает название кнопки
func (btn *RemoveProductInlineButton) Text() string {
	return "🗑 Удалить"
}

// Callback возвращает команду которая должна отработать
func (btn *RemoveProductInlineButton) Callback() string {
	return btn.cmd.Alias()
}
