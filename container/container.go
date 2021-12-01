package container

import (
	"go-micro/database"
	"go-micro/presentation"
	"go-micro/repository"
	"go-micro/service"
)

type Container struct {
	WebServer presentation.WebServer
	DataBase  *database.Database
}

func NewContainer() Container {
	database := database.Database{}

	return Container{
		DataBase: &database,

		WebServer: presentation.WebServer{
			Router: presentation.WebRouter{
				ArticleController: presentation.ArticleController{
					Service: service.Article{
						IArticle: repository.Article{
							DataBase: &database},
					},
				},
				ArticleTypeController: presentation.ArticleTypeController{
					Service: service.ArticleType{
						IArticleType: repository.ArticleType{
							DataBase: &database},
					},
				},
				InvoiceController: presentation.InvoiceController{
					InvoiceService: service.Invoice{
						IInvoice: repository.Invoice{
							DataBase: &database},
					},
				},
			},
		},
	}
}
