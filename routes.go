package main

func initializeRoutes() {

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	//router.Use(setUserStatus())

	// Handle the index route
	router.GET("/",main_index )

	// Group user related routes together
	userRoutes := router.Group("/u")
	{
		// Handle the GET requests at /u/ping
		// Show the {"pong":"Hello"}
		userRoutes.GET("/ping", ping)

		// Handle the GET requests at /main/index , /posts/index , /users/index
		// html_multi_rendering
		userRoutes.GET("/main/index",main_index)
		userRoutes.GET("/posts/index",posts_index)
		userRoutes.GET("/users/index",users_index)

		// Handle the GET requests at /u/hello/:name
		// Show the hello name
		userRoutes.GET("/hello/:name", get_url_one_param)

		// Handle the GET requests at /u/hello/name/action
		// Show the hello name and action
		userRoutes.GET("/hello/:name/*action", get_url_multi_param)

		// Handle the GET requests at /u/query/name/action
		// Show the hello name , action , firstname , lastname 
		userRoutes.GET("/query/:name/*action", get_url_by_query)

		// Handle POST requests at /u/post
		// Show the post data , { "user": "Pgluffy" }
		userRoutes.POST("/post", get_post_json_format)

		// Handle POST requests at /u/form_post
		// Show the post data
		userRoutes.POST("/form_post", get_post_form)

		// Handle POST requests at /u/cookie
		// Show the cookie
		userRoutes.POST("/cookie", getcookie)

		// Handle POST requests at /u/bind_json
		// Bind struct and verify Req json value
		userRoutes.POST("/bind_json", bind_json)

		// Handle POST requests at /u/bindPostData
		// Bind struct Person
		userRoutes.POST("/bindPostData", bindPostData)

		// Handle the GET requests at /u/getb , /u/getc , /u/getd
		// Show the StructA --> StructB --> StructC --> StructD
		userRoutes.GET("/getb", GetDataB)
		userRoutes.GET("/getc", GetDataC)
		userRoutes.GET("/getd", GetDataD)

		// Handle the GET requests at /u/middleware
		// Show the middleware1 -> middleware2 start --> handler --> middleware2 end
		userRoutes.GET("/middleware", middleware1, middleware2, handler)
	}
}
