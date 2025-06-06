package main

import "fmt"

type ClothingItem struct {
	ID           int
	Name         string
	Color        string
	Category     string
	numCategory  int // numeric representation of category
	Formality    string
	numFormality int // numberic representation of formality
	Weather      string
	numWeather   int // numberic representation of weathers
	LastWorn     string
}

const NMAX int = 99

type wardrobe [NMAX]ClothingItem

func main() {
	var item wardrobe
	var total, choice int

	welcome()
	setupWardrobe(&item, &total, 1)

	for {
		menu()
		fmt.Scan(&choice)
		menuChoiceValidity(choice)

		fmt.Println()
		switch choice {
		case 1:
			viewAllClothingItems(item, total)
		case 2:
			editWardrobe(&item, &total)
		case 3:
			sorting(&item, total)
		case 4:
			searchClothingItem(item, total)
		case 5:
			recommendOutfit(item, total)
		}
		if choice == 6 {
			break
		}
	}

	fmt.Println("Thank you for using WellDressed! Goodbye!")
	fmt.Println("=========================================================================")
}

func welcome() {
	var userName string

	fmt.Print("Enter your name (one word, e.g. Terry): ")
	fmt.Scan(&userName)

	fmt.Println("=========================================================================")
	fmt.Printf(" /\\_/\\   Hello %s, welcome to WellDressed!\n", userName)
	fmt.Println("( o.o )  This is your digital wardrobe.")
	fmt.Println(" > ^ <   Let's get started!")
	fmt.Println("=========================================================================")
	fmt.Println()
	fmt.Println("First, let's set up your wardrobe.")
}

func setupWardrobe(item *wardrobe, total *int, NextID int) {
	addClothingItem(item, total, NextID)
	fmt.Println("Wardrobe setup complete!")
	fmt.Println("You can now edit, sort, search, and get outfit recommendations.")
	fmt.Println()
}

func addClothingItem(item *wardrobe, total *int, NextID int) {
	var n, i, spaceLeft int

	spaceLeft = NMAX - *total
	if spaceLeft > 0 {
		fmt.Printf("You can add up to 99 clothing items. You have %d spaces left.\n", spaceLeft)
		fmt.Print("Enter the number of clothing items in your wardrobe: ")
		fmt.Scan(&n)

		// ensures validity
		for n < 1 || n > spaceLeft {
			fmt.Printf("Invalid number of clothing items. Please enter a positive integer less than %d.\n", spaceLeft+1)
			fmt.Print("Enter the number of clothing items in your wardrobe: ")
			fmt.Scan(&n)
		}

		*total += n

		fmt.Println()
		for i = NextID - 1; i < *total; i++ {
			item[i].ID = NextID
			fmt.Printf("Adding clothing item (ID: %d)\n", NextID)

			enterClothingName(item, i)
			chooseClothingCategory(item, i)
			chooseClothingWeather(item, i)
			chooseClothingFormality(item, i)
			enterClothingColor(item, i)
			enterClothingLastWorn(item, i)

			NextID++
			fmt.Println()
		}

		viewAllClothingItems(*item, *total)
		fmt.Println("Clothing item added succesfully!")
	} else {
		fmt.Println("Cannot add item, data has reached maximum.")
	}

}

func enterClothingName(item *wardrobe, index int) {
	var name string

	fmt.Print("Enter clothing name (one word and use capital for initial, e.g. Sundress): ")
	fmt.Scan(&name)

	// ensures validity
	for name[0] < 'A' || name[0] > 'Z' {
		fmt.Println("Invalid name, please try again.")
		fmt.Print("Enter clothing name (one word and use capital for initial, e.g. sundress): ")
		fmt.Scan(&name)
	}

	(*item)[index].Name = name
}

func chooseClothingCategory(item *wardrobe, index int) {
	var categoryChoice int

	fmt.Println("Choose clothing category:")
	fmt.Println("1. Top")
	fmt.Println("2. Bottoms")
	fmt.Println("3. Dress")
	fmt.Println("4. Shoes")
	fmt.Println("5. Accessories")
	fmt.Print("Enter your choice (1-5): ")
	fmt.Scan(&categoryChoice)

	// ensures validity
	for categoryChoice < 1 || categoryChoice > 5 {
		fmt.Println("Invalid choice, please try again.")
		fmt.Println("Choose clothing category:")
		fmt.Println("1. Top")
		fmt.Println("2. Bottoms")
		fmt.Println("3. Dress")
		fmt.Println("4. Shoes")
		fmt.Println("5. Accessories")
		fmt.Print("Enter your choice (1-5): ")
		fmt.Scan(&categoryChoice)
	}

	(*item)[index].numCategory = categoryChoice
	(*item)[index].Category = category(categoryChoice)
}

func category(num int) string {
	var result string

	switch num {
	case 1:
		result = "Tops"
	case 2:
		result = "Bottoms"
	case 3:
		result = "Dress"
	case 4:
		result = "Shoes"
	case 5:
		result = "Accesories"
	}

	return result
}

func chooseClothingWeather(item *wardrobe, index int) {
	var weatherChoice int

	fmt.Println("Choose the suitable weather category")
	fmt.Println("1. Hot")
	fmt.Println("2. Temperate")
	fmt.Println("3. Cold")
	fmt.Print("Enter your choice (1-3): ")
	fmt.Scan(&weatherChoice)

	// ensures validity
	for weatherChoice < 1 || weatherChoice > 3 {
		fmt.Println("Invalid choice, please try again.")
		fmt.Println("Choose the suitable weather category")
		fmt.Println("1. Hot")
		fmt.Println("2. Temperate")
		fmt.Println("3. Cold")
		fmt.Print("Enter your choice (1-3): ")
		fmt.Scan(&weatherChoice)
	}

	(*item)[index].numWeather = weatherChoice
	(*item)[index].Weather = weather(weatherChoice)
}

func weather(num int) string {
	var result string

	switch num {
	case 1:
		result = "Hot"
	case 2:
		result = "Temperate"
	case 3:
		result = "Cold"
	}

	return result
}

func chooseClothingFormality(item *wardrobe, index int) {
	var formalityChoice int

	fmt.Println("Choose clothing formality")
	fmt.Println("1. Casual")
	fmt.Println("2. Semi-Formal")
	fmt.Println("3. Formal")
	fmt.Print("Enter your choice (1-3): ")
	fmt.Scan(&formalityChoice)

	// ensures validity
	for formalityChoice < 1 || formalityChoice > 3 {
		fmt.Println("Invalid formality, please try again.")
		fmt.Println("Choose clothing formality")
		fmt.Println("1. Casual")
		fmt.Println("2. Semi-Formal")
		fmt.Println("3. Formal")
		fmt.Print("Enter your choice (1-3): ")
		fmt.Scan(&formalityChoice)
	}

	(*item)[index].numFormality = formalityChoice
	(*item)[index].Formality = formality(formalityChoice)
}

func formality(num int) string {
	var result string

	switch num {
	case 1:
		result = "Casual"
	case 2:
		result = "Semi-Formal"
	case 3:
		result = "Formal"
	}

	return result
}

func enterClothingColor(item *wardrobe, index int) {
	var color string

	fmt.Print("Enter clothing color (one word and use capital for initial, e.g. Cyan): ")
	fmt.Scan(&color)

	// ensures validity
	for color[0] < 'A' || color[0] > 'Z' {
		fmt.Println("Invalid color, please try again.")
		fmt.Print("Enter clothing color (one word and use capital for initial, e.g. Cyan): ")
		fmt.Scan(&color)
	}

	(*item)[index].Color = color
}

func enterClothingLastWorn(item *wardrobe, index int) {
	var lastWorn string

	fmt.Print("Enter last worn date (YYYY-MM-DD): ")
	fmt.Scan(&lastWorn)

	// ensures validity
	for len(lastWorn) < 10 || lastWorn[4] != '-' || lastWorn[7] != '-' {
		fmt.Println("Invalid date format, please try again.")
		fmt.Print("Enter last worn date (YYYY-MM-DD): ")
		fmt.Scan(&lastWorn)
	}

	(*item)[index].LastWorn = lastWorn
}

func viewAllClothingItems(item wardrobe, total int) {
	var i int

	fmt.Println("===============================================VIEW ALL CLOTHING ITEMS===============================================")
	fmt.Printf("|| %-3s | %-15s | %-15s | %-15s | %-15s | %-15s | %-15s ||\n", "ID", "Last Worn", "Name", "Category", "Weather", "Formality", "Color")
	fmt.Println("---------------------------------------------------------------------------------------------------------------------")
	for i = 0; i < total; i++ {
		fmt.Printf("|| %-3d | %-15s | %-15s | %-15s | %-15s | %-15s | %-15s ||\n", item[i].ID, item[i].LastWorn, item[i].Name, item[i].Category, item[i].Weather, item[i].Formality, item[i].Color)
	}
	fmt.Println("=====================================================================================================================")
}

func menu() {
	fmt.Println("╔═════════════════════════════════════════╗")
	fmt.Println("║               WellDressed               ║")
	fmt.Println("║        Digital Wardrobe Assistant       ║")
	fmt.Println("╠═════════════════════════════════════════╣")
	fmt.Println("║ 1. View My Wardrobe                     ║")
	fmt.Println("║ 2. Edit My Wardrobe                     ║")
	fmt.Println("║ 3. Sort My Wardrobe                     ║")
	fmt.Println("║ 4. Search Clothing Item                 ║")
	fmt.Println("║ 5. Recommend Me an Outfit!              ║")
	fmt.Println("║ 6. Exit                                 ║")
	fmt.Println("╚═════════════════════════════════════════╝")
	fmt.Print("Choose an option (1-6): ")
}

func menuChoiceValidity(choice int) int {
	for choice < 1 || choice > 6 {
		fmt.Println("Invalid choice, please try again.")
		fmt.Print("Choose an option (1-6): ")
		fmt.Scan(&choice)
	}

	return choice
}

func editWardrobe(item *wardrobe, total *int) {
	var choice int

	fmt.Println("Edit My Wardrobe")
	fmt.Println("1. Modify Clothing Item")
	fmt.Println("2. Remove Clothing Item")
	fmt.Println("3. Add Clothing Item")
	fmt.Print("Choose an option (1-3): ")
	fmt.Scan(&choice)

	// ensures validity
	for choice < 1 || choice > 3 {
		fmt.Println("Invalid choice, please try again.")
		fmt.Print("Edit My Wardrobe")
		fmt.Println("1. Modify Clothing Item")
		fmt.Println("2. Remove Clothing Item")
		fmt.Println("3. Add Clothing Item")
		fmt.Print("Choose an option (1-3): ")
		fmt.Scan(&choice)
	}

	switch choice {
	case 1:
		modifyClothingItem(item, *total)
	case 2:
		removeClothingItem(item, total)
	case 3:
		addClothingItem(item, total, *total+1)
	}
}

func modifyClothingItem(item *wardrobe, total int) {
	var i, id int
	var found bool = false

	viewAllClothingItems(*item, total)

	fmt.Print("Enter the ID of the clothing item to modify: ")
	fmt.Scan(&id)

	i = 0
	for i < total && !found {
		if (*item)[i].ID == id {
			found = true
			modifyByField(item, i)

			viewAllClothingItems(*item, total)
			fmt.Println("Clothing item modified successfully!")
		}
		i++
	}

	if !found {
		fmt.Println("Item not found.")
	}
}

func modifyByField(item *wardrobe, index int) {
	var field int

	fmt.Print("Enter which field to modify (1-Last Worn, 2-Name, 3-Category, 4-Weather, 5-Formality, 6-Color): ")
	fmt.Scan(&field)

	// ensures validity
	for field < 1 || field > 6 {
		fmt.Println("Invalid choice, please try again.")
		fmt.Print("Enter which field to modify (1-Last Worn, 2-Name, 3-Category, 4-Weather, 5-Formality, 6-Color): ")
		fmt.Scan(&field)
	}

	switch field {
	case 1:
		enterClothingLastWorn(item, index)
	case 2:
		enterClothingName(item, index)
	case 3:
		chooseClothingCategory(item, index)
	case 4:
		chooseClothingWeather(item, index)
	case 5:
		chooseClothingFormality(item, index)
	case 6:
		enterClothingColor(item, index)
	}
}

func removeClothingItem(item *wardrobe, total *int) {
	var i, j, id int
	var found bool = false

	viewAllClothingItems(*item, *total)
	fmt.Print("Enter the ID of the clothing item to remove: ")
	fmt.Scan(&id)

	i = 0
	for i < *total && !found {
		if (*item)[i].ID == id {
			found = true
			for j = i; j < *total-1; j++ {
				(*item)[j] = (*item)[j+1]
				(*item)[j].ID--
			}
			*total--
			viewAllClothingItems(*item, *total)
			fmt.Println("Clothing item removed successfully!")
		}
		i++
	}

	if !found {
		fmt.Println("Item not found.")
	}
}

func sorting(sortThis *wardrobe, total int) {
	fmt.Println("Sort by Most Formal: ")
	selectionSortDescendingFormal(sortThis, total)
	printSorted(*sortThis, total)

	fmt.Println("Sort by Least Formal: ")
	selectionSortAscendingFormal(sortThis, total)
	printSorted(*sortThis, total)

	fmt.Println("Sort by Last Worn: ")
	insertionSortByLastWorn(sortThis, total)
	printSorted(*sortThis, total)
}

func selectionSortDescendingFormal(arr *wardrobe, total int) {
	var i, j, maxIndex int
	var temp ClothingItem

	for i = 0; i < total-1; i++ {
		maxIndex = i
		for j = i + 1; j < total; j++ {
			if arr[j].numFormality > arr[maxIndex].numFormality {
				maxIndex = j
			}
		}

		temp = arr[maxIndex]
		arr[maxIndex] = arr[i]
		arr[i] = temp
	}
}

func selectionSortAscendingFormal(arr *wardrobe, total int) {
	var i, j, minIndex int
	var temp ClothingItem

	for i = 0; i < total-1; i++ {
		minIndex = i
		for j = i + 1; j < total; j++ {
			if arr[j].numFormality < arr[minIndex].numFormality {
				minIndex = j
			}
		}

		temp = arr[minIndex]
		arr[minIndex] = arr[i]
		arr[i] = temp

	}
}

func insertionSortByLastWorn(arr *wardrobe, total int) {
	var i, j int
	var key ClothingItem

	for i = 1; i < total; i++ {
		key = arr[i]
		j = i - 1
		for j >= 0 && arr[j].LastWorn < key.LastWorn {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func printSorted(arr wardrobe, total int) {
	var i int

	fmt.Println("====================================================SORTED  ITEMS====================================================")
	fmt.Printf("|| %-3s | %-15s | %-15s | %-15s | %-15s | %-15s | %-15s ||\n", "ID", "Last Worn", "Name", "Category", "Weather", "Formality", "Color")
	fmt.Println("---------------------------------------------------------------------------------------------------------------------")
	for i = 0; i < total; i++ {
		fmt.Printf("|| %-3d | %-15s | %-15s | %-15s | %-15s | %-15s | %-15s ||\n", arr[i].ID, arr[i].LastWorn, arr[i].Name, arr[i].Category, arr[i].Weather, arr[i].Formality, arr[i].Color)
	}
	fmt.Println("=====================================================================================================================")
}

func searchClothingItem(item wardrobe, total int) {
	var searchChoice int

	fmt.Println("Choose a search option:")
	fmt.Println("1. Search all by color")
	fmt.Println("2. Search an item by category")
	fmt.Print("Choose an option (1-2): ")
	fmt.Scan(&searchChoice)

	// ensures validity
	for searchChoice < 1 || searchChoice > 2 {
		fmt.Println("Invalid choice, please try again.")
		fmt.Println("Choose a search option:")
		fmt.Println("1. Search all by color")
		fmt.Println("2. Search an item by category")
		fmt.Print("Choose an option (1-2): ")
		fmt.Scan(&searchChoice)
	}

	switch searchChoice {
	case 1:
		sequentialSearchColor(item, total)
	case 2:
		binarySearchCategory(item, total)
	}
}

func sequentialSearchColor(item wardrobe, total int) {
	var i int
	var searchedColor string
	var found bool = false

	fmt.Print("Enter color to search (one word and use capital for initial, e.g. Cyan): ")
	fmt.Scan(&searchedColor)

	// ensures validity
	if searchedColor[0] < 'A' || searchedColor[0] > 'Z' {
		fmt.Println("Invalid input, please try again.")
		fmt.Print("Enter color to search (one word and use capital for initial, e.g. Cyan): ")
		fmt.Scan(&searchedColor)
	}

	fmt.Println("====================================================SEARCH RESULT====================================================")
	fmt.Printf("|| %-3s | %-15s | %-15s | %-15s | %-15s | %-15s | %-15s ||\n", "ID", "Last Worn", "Name", "Category", "Weather", "Formality", "Color")
	fmt.Println("---------------------------------------------------------------------------------------------------------------------")
	for i = 0; i < total; i++ {
		if item[i].Color == searchedColor {
			found = true
			fmt.Printf("|| %-3d | %-15s | %-15s | %-15s | %-15s | %-15s | %-15s ||\n", item[i].ID, item[i].LastWorn, item[i].Name, item[i].Category, item[i].Weather, item[i].Formality, item[i].Color)
		}
	}

	if !found {
		fmt.Println("Item not found.")
	}
	fmt.Println("=====================================================================================================================")
}

func selectionSortAscendingCategory(arr *wardrobe, total int) {
	var i, j, minIndex int
	var temp ClothingItem

	for i = 0; i < total-1; i++ {
		minIndex = i
		for j = i + 1; j < total; j++ {
			if arr[j].numCategory < arr[minIndex].numCategory {
				minIndex = j
			}
		}

		temp = arr[minIndex]
		arr[minIndex] = arr[i]
		arr[i] = temp
	}
}

func binarySearchCategory(item wardrobe, total int) {
	var choice, left, right, mid int
	var found bool = false

	fmt.Print("Choose category to search (1-Tops, 2-Bottoms, 3-Dress, 4-Shoes, 5-Accesories): ")
	fmt.Scan(&choice)

	// ensures validity
	for choice < 1 || choice > 5 {
		fmt.Println("Invalid category, please try again.")
		fmt.Println("Choose category to search (1-Tops, 2-Bottoms, 3-Dress, 4-Shoes, 5-Accesories): ")
		fmt.Scan(&choice)
	}

	selectionSortAscendingCategory(&item, total)
	fmt.Println("====================================================SEARCH RESULT====================================================")
	fmt.Printf("|| %-3s | %-15s | %-15s | %-15s | %-15s | %-15s | %-15s ||\n", "ID", "Last Worn", "Name", "Category", "Weather", "Formality", "Color")
	fmt.Println("---------------------------------------------------------------------------------------------------------------------")
	left = 0
	right = total - 1
	for left <= right && !found {
		mid = (left + right) / 2
		if item[mid].numCategory < choice {
			left = mid + 1
		} else if item[mid].numCategory > choice {
			right = mid - 1
		} else {
			found = true
			fmt.Printf("|| %-3d | %-15s | %-15s | %-15s | %-15s | %-15s | %-15s ||\n", item[mid].ID, item[mid].LastWorn, item[mid].Name, item[mid].Category, item[mid].Weather, item[mid].Formality, item[mid].Color)
		}
	}

	if !found {
		fmt.Println("Item not found.")
	}
	fmt.Println("=====================================================================================================================")
}

func recommendOutfit(item wardrobe, total int) {
	var preferWeather, preferFormality, preferCombo int

	userPreferences(&preferCombo, &preferWeather, &preferFormality)

	fmt.Println()
	fmt.Println("Your recommended outfit combo:")
	findRecommendation(item, total, preferCombo, preferWeather, preferFormality)
}

func userPreferences(preferCombo, preferWeather, preferFormality *int) {
	fmt.Print("Enter 1 for top-bottom combo or 2 for dress outfit recommendation: ")
	fmt.Scan(preferCombo)

	// ensures validity
	for *preferCombo < 1 || *preferCombo > 2 {
		fmt.Println("Invalid choice, please try again.")
		fmt.Print("Enter 1 for top-bottom combo or 2 for dress outfit recommendation: ")
		fmt.Scan(preferCombo)
	}

	fmt.Print("Enter weather to dress for (1-Hot, 2-Temperate, 3-Cold): ")
	fmt.Scan(preferWeather)

	// ensures validity
	for *preferWeather < 1 || *preferWeather > 3 {
		fmt.Println("Invalid choice, please try again.")
		fmt.Print("Enter weather to dress for (1-Hot, 2-Temperate, 3-Cold): ")
		fmt.Scan(preferWeather)
	}

	fmt.Print("Enter occasion formality to dress for (1-Casual, 2-Semi-Formal, 3-Formal): ")
	fmt.Scan(preferFormality)

	// ensures validity
	for *preferFormality < 1 || *preferFormality > 3 {
		fmt.Println("Invalid choice, please try again.")
		fmt.Print("Enter occasion formality to dress for (1-Casual, 2-Semi-Formal, 3-Formal): ")
		fmt.Scan(preferFormality)
	}
}

func findRecommendation(item wardrobe, total int, preferCombo, preferWeather, preferFormality int) {
	var iTop, iBottoms, iDress int
	var iShoes, iAccessories int

	if preferCombo == 1 {
		iTop = sequentialSearchRecommendation(item, total, preferWeather, preferFormality, "Tops")
		if iTop < 0 {
			fmt.Println("Matched top not found")
		} else {
			fmt.Printf("Top ID: %d, Top Name: %s\n", item[iTop].ID, item[iTop].Name)
		}

		iBottoms = sequentialSearchRecommendation(item, total, preferWeather, preferFormality, "Bottoms")
		if iBottoms < 0 {
			fmt.Println("Matched bottoms not found")
		} else {
			fmt.Printf("Bottom ID: %d, Bottom Name: %s\n", item[iBottoms].ID, item[iBottoms].Name)
		}

	} else {
		iDress = sequentialSearchRecommendation(item, total, preferWeather, preferFormality, "Dress")
		if iDress < 0 {
			fmt.Println("Matched dress not found")
		} else {
			fmt.Printf("Dress ID: %d, Dress Name: %s\n", item[iDress].ID, item[iDress].Name)
		}
	}

	iShoes = sequentialSearchRecommendation(item, total, preferWeather, preferFormality, "Shoes")
	if iShoes < 0 {
		fmt.Println("Matched shoes not found")
	} else {
		fmt.Printf("Shoes ID: %d, Shoes Name: %s\n", item[iShoes].ID, item[iShoes].Name)
	}

	iAccessories = sequentialSearchRecommendation(item, total, preferWeather, preferFormality, "Accesories")
	if iAccessories < 0 {
		fmt.Println("Matched accessories not found")
	} else {
		fmt.Printf("Accessories ID: %d, Accessories Name: %s\n", item[iAccessories].ID, item[iAccessories].Name)
	}
}

func sequentialSearchRecommendation(item wardrobe, total int, preferWeather, preferFormality int, keyCategory string) int {
	var i int
	var found bool = false

	i = total - 1
	for i >= 0 && !found {
		found = item[i].numWeather == preferWeather && item[i].numFormality == preferFormality && item[i].Category == keyCategory
		i--
	}

	if found {
		i++
	}

	return i
}

