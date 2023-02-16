package help

import (
	"net/http"
	"severyanochka/internals/app/handler/static"
	"severyanochka/internals/utils/md"
)

// HelpController Контроллер для отправки документации по работе с API
type HelpController struct {
	md *md.MdWorkerImpl
}

// New Создание нового контроллера
func New(md *md.MdWorkerImpl) *HelpController {
	controller := new(HelpController)
	controller.md = md
	return controller
}

// ShowMainPageDoc Отправка документации с API главной страницы
func (controller *HelpController) ShowMainPageDoc(w http.ResponseWriter, r *http.Request) {
	html := controller.md.GetMainPageDocApi()
	static.WrapHtmlPage(w, r, html)
}
