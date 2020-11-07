package main

import "fmt"

type product struct {
	pID  string
	name string
}

type cart struct {
	products []product
	getCart  iCart
}

type iCart interface {
	getAll() ([]product, int)
	getOne(id string) (product, int)
	update(id int, newID string, entry string)
	delete(id string)
	add(item product)
}

func newProduct(id string, pName string) product {
	return product{
		pID:  id,
		name: pName,
	}
}

func newCart(id string, pName string) product {
	return newProduct(id, pName)
}

func (cartItems *cart) getOne(id string) (product, int) {
	if len(cartItems.products) == 0 {
		return product{}, -1
	}
	for index, item := range cartItems.products {
		if item.pID == id {
			return item, index
		}
	}
	return product{}, -1
}

func (cartItems *cart) getAll() ([]product, int) {
	if len(cartItems.products) == 0 {
		return cartItems.products, -1
	}
	return cartItems.products, 1
}

func (cartItems *cart) add(item product) {
	cartItems.products = append(cartItems.products, item)
	fmt.Println("Item Added Successfully !")
}

func (cartItems *cart) delete(id int) {
	if len(cartItems.products) == 0 {
		fmt.Println("Cart is empty !")
	}
	(*cartItems).products = append((*cartItems).products[:id], (*cartItems).products[id+1:]...)
	fmt.Println("Item Removed....")
}

func (cartItems *cart) update(id int, newID string, entry string) {
	if len(cartItems.products) == 0 {
		fmt.Println("Cart is empty !")
	}
	(*cartItems).products[id].pID = newID
	(*cartItems).products[id].name = entry
	fmt.Println("Item Updated....!")
}

func main() {

	myCart := new(cart)

	var input int
	for {
		fmt.Println("Choose your choice (1 : Add, 2 : Update, 3 : Delete, 4 : All Details, 5 : Get One Details")
		fmt.Scan(&input)

		if input == 0 {
			fmt.Println("Terminating Process !")
			break
		}

		switch input {
		case 1:
			var id string
			var name string
			fmt.Println("Please enter product id and name,seperating by a space..")
			fmt.Scan(&id, &name)
			newEntry := newCart(id, name)
			myCart.add(newEntry)
		case 2:
			var id string
			var name string
			fmt.Print("Enter product id : ")
			fmt.Scan(&id)
			item, isPresent := myCart.getOne(id)
			if isPresent == -1 {
				fmt.Print("Product Not Found !")
				continue
			}
			fmt.Println("Product Found,Kindly enter 9 to change only pID,8 to change product name and 6 to change boths...")
			var nQ int
			var newID string
			fmt.Scan(&nQ)
			if nQ == 8 {
				fmt.Print("Enter product name : ")
				fmt.Scan(&name)
				id = item.pID
			} else if nQ == 9 {
				fmt.Print("Enter new product id : ")
				fmt.Scan(&newID)
				id = newID
				name = item.name
			} else if nQ == 6 {
				fmt.Print("Enter new product id and name : ")
				fmt.Scan(&newID, &name)
				id = newID
			}

			myCart.update(isPresent, id, name)
		case 3:
			var id string
			fmt.Print("Enter product id : ")
			fmt.Scan(&id)
			_, isPresent := myCart.getOne(id)
			if isPresent == -1 {
				fmt.Print("Product Not Found !")
				continue
			}
			myCart.delete(isPresent)

		case 4:
			fmt.Println("Showing all product details...")
			product, isPresent := myCart.getAll()
			if isPresent == -1 {
				fmt.Println("No Product are there !")
				continue
			}
			for _, item := range product {
				fmt.Printf("Product ID - %s , Product name - %s\n", item.pID, item.name)
			}

		case 5:
			var id string
			fmt.Print("Please enter the product id to view the details:- ")
			fmt.Scan(&id)
			item, index := myCart.getOne(id)
			if index == -1 {
				fmt.Println("Details Not Found !")
				continue
			}
			fmt.Printf("Product ID - %s , Product name - %s", item.pID, item.name)

		default:
			fmt.Println("No match")
		}
	}
}
