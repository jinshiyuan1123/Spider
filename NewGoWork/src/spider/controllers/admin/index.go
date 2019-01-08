package admin

type IndexController struct{
	AdminController
}

func (self *IndexController) Index(){
	self.display()
}