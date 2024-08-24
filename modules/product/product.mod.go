package product

type ProductModule struct {
	Ctl *ProductController
	Svc *ProductService
}

func New() *ProductModule {
	svc := newService()
	return &ProductModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
