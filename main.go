package main

import "fmt"

/* Создать приложение, которое сначала выдает меню:
	1. Посмотреть закладки
	2. Добавить закладку
	3. Удалить закладку
	4. Выход
При 1 - Выводит закладки
При 2 - 2 поля ввода названия и адреса и после добавление
При 3 - Ввод названия и удаления по нему
При 4 - Завершение
*/

func main() {

	bookmarks := map[string]string{} //создание мапы
	fmt.Println("Приложение для закладок")
Menu:		//labels с помощью его мы выходим из цикла, а не из кейса
	for {			//кейсы с функциями
		variant := getMenu()
		switch variant {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			addBookmark(bookmarks)
		case 3:
			deleteBookMark(bookmarks)
		case 4:
			break Menu
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант")
	fmt.Println("1. Посмотреть закладку")
	fmt.Println("2. добавить закладку")
	fmt.Println("3. удалить закладку")
	fmt.Println("4. Выйти")
	fmt.Scan(&variant)
	return variant
}


// функция просмотра закладок
func printBookmarks(bookmarks map[string]string) {
	if len(bookmarks) == 0 {
		fmt.Println("Пока нету закладок")
	}
	for key, value := range bookmarks {
		fmt.Println(key, ": ", value)
	}
}

// функция добавления закладок
func addBookmark(bookmarks map[string]string) {
	var newBookmarkKey string
	var newBookmarkValue string
	fmt.Print("Введите название: ")
	fmt.Scan(&newBookmarkKey)
	fmt.Print("Введите ссылку: ")
	fmt.Scan(&newBookmarkValue)
	bookmarks[newBookmarkKey] = newBookmarkValue
}

// функция удаления закладок
func deleteBookMark(bookmarks map[string]string) {
	var bookmarkKeyToDelete string
	fmt.Print("Введите название: ")
	fmt.Scan(&bookmarkKeyToDelete)
	delete(bookmarks, bookmarkKeyToDelete)
}
