package main

import "fmt"

type Category struct {
	id int
	name int
	path string
}

type Product struct {
	id int
	categoryId int
	name string
}

type Categorys []Category

func (cats Categorys) GetNameByID (id int) int {
	var  tmpName int
	for _, cat := range cats{
		if(cat.id == id){
			tmpName = cat.name
			break
		}
	}
	return tmpName
}

func (prod Product) PrintAllData(){
	fmt.Println(prod);
}

func main(){
	var product = Product{1,10,"Test product"}
	var cat1 = Category{1,11,"testPathCat1"}
	var cat2 = Category{2,22,"testPathCat2"}
	var cat3 = Category{3,33,"testPathCat3"}
	var cat4 = Category{4,44,"testPathCat4"}

	var cats = Categorys{cat1,cat2,cat3,cat4}

	fmt.Println(cats.GetNameByID(2))
	fmt.Println(cats.GetNameByID(4))
	fmt.Println(cats.GetNameByID(1))
	fmt.Println(cats.GetNameByID(3))
	product.PrintAllData()

}
