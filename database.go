package main

var db = []Post{
	{
		Id:       1,
		Upvotes:  1,
		Title:    "My cat is the cutest!",
		Link:     "https://i.imgur.com/jseZqNK.jpg",
		Username: "alicia98",
		Comments: []Comment{
			{
				Id:       1,
				Upvotes:  2,
				Text:     "She's such a cutie! :3",
				Username: "raahi014",
			},
		},
	},
	{
		Id:       2,
		Upvotes:  -432,
		Title:    "Thomas Jefferson circa 2015",
		Link:     "https://i.redd.it/xn9auq3xdoa51.jpg",
		Username: "ZzturtleszZ",
		Comments: []Comment{},
	},
}
