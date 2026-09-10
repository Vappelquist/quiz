package quiz

func Seed() []Question {
	return []Question{{
		Id:   "1",
		Text: "Summery cocktail with a base of rum, lime juice, and sugar, served over ice in a highball glass.",
		Options: []Option{

			{Id: "a", Text: "Mojito"},
			{Id: "b", Text: "Piña Colada"},
			{Id: "c", Text: "Mai Tai"},
			{Id: "d", Text: "Dark 'n' Stormy"}},
		CorrectId: "a",
	}, {
		Id:   "2",
		Text: "A classic cocktail made with gin, vermouth, and Campari from Italy.",
		Options: []Option{
			{Id: "a", Text: "Negroni"},
			{Id: "b", Text: "Americano"},
			{Id: "c", Text: "Martini"},
			{Id: "d", Text: "Manhattan"}},
		CorrectId: "a",
	}, {
		Id:   "3",
		Text: "A tropical cocktail made with rum, coconut cream, and pineapple juice.",
		Options: []Option{
			{Id: "a", Text: "Piña Colada"},
			{Id: "b", Text: "Mai Tai"},
			{Id: "c", Text: "Daiquiri"},
			{Id: "d", Text: "Zombie"}},
		CorrectId: "a",
	}, {
		Id:   "4",
		Text: "A classic cocktail made with whiskey, sweet vermouth, and bitters.",
		Options: []Option{
			{Id: "a", Text: "Manhattan"},
			{Id: "b", Text: "Old Fashioned"},
			{Id: "c", Text: "Sazerac"},
			{Id: "d", Text: "Boulevardier"}},
		CorrectId: "a",
	}, {
		Id:   "5",
		Text: "A refreshing cocktail made with vodka, tomato juice, and various spices and flavorings.",
		Options: []Option{
			{Id: "a", Text: "Bloody Mary"},
			{Id: "b", Text: "Moscow Mule"},
			{Id: "c", Text: "Caipirinha"},
			{Id: "d", Text: "Tom Collins"}},
		CorrectId: "a",
	}}
}
