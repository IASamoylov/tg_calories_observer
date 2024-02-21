package product

// EditProductInlineButton предоставляем пользователю с помощью inline кнопки получить
// доступ к функционалу заложенму в команду редактирования продукта
type EditProductInlineButton struct {
	cmd command
}

// NewEditProductInlineButton ctor
func NewEditProductInlineButton(cmd command) *EditProductInlineButton {
	return &EditProductInlineButton{cmd: cmd}
}

// Text возвращает название кнопки
func (btn *EditProductInlineButton) Text() string {
	return "✏️ Изменить"
}

// Callback возвращает команду которая должна отработать
func (btn *EditProductInlineButton) Callback() string {
	return btn.cmd.Alias()
}
