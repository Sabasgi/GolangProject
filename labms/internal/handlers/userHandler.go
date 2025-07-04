package handlers

// type UserHandler struct {
// 	service services.UserServiceStruct
// }

//	func NewUserHandler(s services.UserServiceStruct) *UserHandler {
//		return &UserHandler{
//			service: s,
//		}
//	}
/*
func (uh *MainHandlers) CreateUserRoute(c *gin.Context) {
	var user models.UserModel
	err := c.Bind(&user)
	if err != nil {
		fmt.Println("ERROR : CreateUserRoute", err)
		c.JSON(http.StatusExpectationFailed, map[string]string{"error": err.Error()})
		return
	}
	Err := uh.Us.CreateService(user)
	if Err != nil {
		c.JSON(http.StatusExpectationFailed, map[string]string{"error": Err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "user is created successfully"})
	return
}
func (uh *MainHandlers) UpdateUserRoute(c *gin.Context) {
	var user models.UserModel
	err := c.Bind(&user)
	if err != nil {
		// fmt.Println("ERROR : CreateUserRoute", err)
		c.JSON(http.StatusExpectationFailed, map[string]string{"error": err.Error()})
		return
	}
	Err := uh.Us.UpdateService(user)
	if Err != nil {
		c.JSON(http.StatusExpectationFailed, map[string]string{"error": Err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "user updated successfully"})
	return
}
func (uh *MainHandlers) GetAllUsersRoute(c *gin.Context) {
	users, err := uh.Us.GetAll()
	if err != nil {
		fmt.Println("ERROR : GetAllProductsRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
	return
}
func (uh *MainHandlers) GetOneUsersRoute(c echo.Context) {
	var u models.UserModel
	err := c.Bind(&u)
	if err != nil {
		// fmt.Println("ERROR : CreateUserRoute", err)
		c.JSON(http.StatusExpectationFailed, map[string]string{"error": err.Error()})
		return
	}
	user, Err := uh.Us.GetOneService(u)
	if err != nil {
		fmt.Println("ERROR : GetAllProductsRoute", Err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
	return
}
*/
