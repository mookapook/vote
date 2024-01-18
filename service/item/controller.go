package vote

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/go-playground/validator/v10"
	echo "github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	userMap   = make(map[string]User)
	UserLogin = make(map[string]DataUser)
	userVote  = make(map[string][]string)
)

func NewController() *Controller {

	m := model()
	c := Controller{
		model: m,
	}
	//go LoadAdmin()
	go genTokenuser()

	UserLogin["Alice"] = DataUser{
		"user1",
		"123456",
		"Alice",
	}

	UserLogin["Lamda"] = DataUser{
		"user2",
		"123456x",
		"Lamda",
	}
	m.createIndex()
	//go c.autoBatchData("")
	return &c
}

//Create itemVote
func (c *Controller) CreateitemVote(ctx echo.Context) error {
	// parse request body and extract parameters

	var data Action
	if err := ctx.Bind(&data); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request body1",
		})
	}

	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request body",
		})
	}
	user := ctx.Get("userId").(string)
	u, _ := mapUser(user)
	if u == "" {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "Unauthorized",
		})
	}
	//log.Println(u.Name)
	data.CreateTime = time.Now()
	data.UpDateTime = time.Now()
	data.UserName = u
	data.Vote = 0
	data.Status = "open"

	datareturn, err := c.model.CreateItem(&data)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":  "Failed to create ",
			"status": 0,
		})
	}

	// return response as JSON
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": datareturn,
	})
}

//Update itemVote
func (c *Controller) UpdateitemVote(ctx echo.Context) error {
	// parse request body and extract parameters

	var data Action
	if err := ctx.Bind(&data); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request body1",
		})
	}

	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request body",
		})
	}

	//log.Println(u.Name)

	ID, err := Converthex(ctx.Param("id"))
	log.Println(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusNoContent, map[string]interface{}{
			"error": "No Content",
		})
	}
	CanEdit, err := c.model.CheckVoteMoreZero(ID)
	if err != nil || CanEdit == false {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"error": "Cannot Edit / Remove",
		})
	}

	user := ctx.Get("userId").(string)
	u, _ := mapUser(user)
	if u == "" {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "Unauthorized",
		})
	}

	data.ID = ID

	datareturn, err := c.model.UpdateItem(&data)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to Update ",
		})
	}

	// return response as JSON
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": datareturn,
	})
}

//Update itemVote
func (c *Controller) Removeitem(ctx echo.Context) error {
	ID, err := Converthex(ctx.Param("id"))
	log.Println(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusNoContent, map[string]interface{}{
			"error": "No Content",
		})
	}
	CanEdit, err := c.model.CheckVoteMoreZero(ID)
	if err != nil || CanEdit == false {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"error": "Cannot Edit / Remove",
		})
	}

	user := ctx.Get("userId").(string)
	u, _ := mapUser(user)
	if u == "" {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "Unauthorized",
		})
	}

	//data.ID = ID

	datareturn, err := c.model.DeleteItem(ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to Remove ",
		})
	}

	// return response as JSON
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": datareturn,
	})
}

func (c *Controller) itemVoteByID(ctx echo.Context) error {
	ID, err := Converthex(ctx.Param("id"))
	log.Println(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusNoContent, map[string]interface{}{
			"error": "No Content",
		})
	}

	user := ctx.Get("userId").(string)
	log.Println(user)
	u, _ := mapUser(user)
	if u == "" {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "Unauthorized",
		})
	}

	if val, ok := userVote[ctx.Param("id")]; ok {
		//val = append(val, userid)
		if Contains(val, user) {
			return ctx.JSON(http.StatusOK, map[string]interface{}{
				"error": "You have voted",
				"data":  "",
			})
		}

	}

	open, err := c.model.GetItemVoteByID(ID)
	if err != nil || open.Status == "close" {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"error": "Cannot Vote or Item Close",
			"data":  "",
		})
	}

	q := primitive.M{}
	q["itemid"] = ID
	q["userid"] = user
	CanNotVote := c.model.CheckVote(q)
	if CanNotVote == true {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"error": "You have voted",
			"data":  "",
		})
	}

	v := VoteUser{}
	v.CreateTime = time.Now()
	v.UpDateTime = time.Now()
	v.Itemid = ID
	v.UserID = user
	cx, err := c.model.VoteItemByUser(&v)
	log.Println(err)
	if err != nil || cx == false {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Cannot Vote",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error": "",
		"data":  cx,
	})
}

func (c *Controller) OpenCloseItem(ctx echo.Context) error {
	// parse request body and extract parameters

	var data OpenClose
	if err := ctx.Bind(&data); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request body1",
		})
	}

	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request body",
		})
	}

	//log.Println(u.Name)

	ID, err := Converthex(ctx.Param("id"))
	log.Println(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusNoContent, map[string]interface{}{
			"error": "No Content",
		})
	}

	if data.Status != Open {
		if data.Status != Close {
			return ctx.JSON(http.StatusOK, map[string]interface{}{
				"error": "No Have Status Setup",
			})
		}
	}
	// CanEdit, err := c.model.CheckVoteMoreZero(ID)
	// if err != nil || CanEdit == false {
	// 	return ctx.JSON(http.StatusNoContent, map[string]interface{}{
	// 		"error": "Cannot Edit / Remove",
	// 	})
	// }

	// user := ctx.Get("userId").(string)
	// u, _ := mapUser(user)
	// if u == "" {
	// 	return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
	// 		"error": "Unauthorized",
	// 	})
	// }

	datareturn, err := c.model.OpenCloseItem(ID, data.Status)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to OpenCloseItem ",
		})
	}

	// return response as JSON
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error": "",
		"data":  datareturn,
	})
}

func (c *Controller) GetAllItem(ctx echo.Context) error {
	params := fmt.Sprintf("?1=%s", "1")
	page := 1

	if tm := ctx.QueryParam("page"); !IsEmpty(tm) {
		page = StringToInt(tm)

		page := page + 1
		params += fmt.Sprintf("&page=%d", page)
	}

	// if tm := ctx.FormValue("id"); !helper.IsEmpty(tm) {
	// 	productId = helper.StringToInt(tm)
	// 	params += fmt.Sprintf("&id=%s", tm)
	// }

	if tm := ctx.QueryParam("sort"); !IsEmpty(tm) {
		//sortz := helper.StringToInt(tm)

		//params += fmt.Sprintf("&sort=%s", "vote")
	}
	status := Open
	if tm := ctx.QueryParam("status"); !IsEmpty(tm) {
		params += fmt.Sprintf("&status=%s", tm)
		if tm == Close {
			status = Close
		} else if tm == All {
			status = All
		}
	}
	limit := 10
	skip := (page - 1) * limit
	if page == 1 {
		page := page + 1
		params += fmt.Sprintf("&page=%d", page)
	}
	user := ctx.Get("userId").(string)
	datareturn := c.model.GetAllItem(skip, limit+1, "vote", user, status)
	next := ""
	if len(datareturn) > limit {

		//temp := datainfo.Content
		datareturn = datareturn[:len(datareturn)-1]

		next = params
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error": "",
		"data":  datareturn,
		"next":  next,
	})
}

func (c *Controller) LoginUser(ctx echo.Context) error {
	var data DataUser
	if err := ctx.Bind(&data); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request body1",
		})
	}

	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request body",
		})
	}
	token := ""
	var val DataUser
	var ok bool
	if val, ok = UserLogin[data.UserName]; !ok {
		//val = append(val, userid)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "No UserName",
		})

	}

	if strings.ToLower(val.Password) == strings.ToLower(strings.TrimSpace(data.Password)) && strings.ToLower(val.UserName) == strings.ToLower(strings.TrimSpace(data.UserName)) {
		token, err = GenerateJWT(val.UserID, data.UserName)
		if err != nil {
			return ctx.JSON(http.StatusOK, map[string]interface{}{
				"error": "Cannot Gen Token",
			})
		}
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error": "",
		"token": token,
	})
}

func (c *Controller) ClearALL(ctx echo.Context) error {

	datareturn, err := c.model.ClearItemAndVoteALL()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to OpenCloseItem ",
		})
	}

	// return response as JSON
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error": "",
		"data":  datareturn,
	})
}

func (c *Controller) ClearbyItem(ctx echo.Context) error {
	// parse request body and extract parameters

	//log.Println(u.Name)

	ID, err := Converthex(ctx.Param("id"))
	log.Println(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusNoContent, map[string]interface{}{
			"error": "No Content",
		})
	}

	datareturn, err := c.model.ClearItemAndVoteByID(ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to ClearItem ",
		})
	}

	// return response as JSON
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error": "",
		"data":  datareturn,
	})
}

func (c *Controller) ExportItem(ctx echo.Context) error {

	// Create Excel file
	f := excelize.NewFile()
	f.NewSheet("Sheet1")

	// Write headers
	//headers := []string{"Name", "Description", "Vote", "User", "CreateDate", "Status"}
	row := 1

	f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), "Name")
	f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), "Description")
	f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), "Vote")
	f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), "User")
	f.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), "CreateDate")
	f.SetCellValue("Sheet1", fmt.Sprintf("F%d", row), "Status")

	row++
	datareturn := c.model.ReportItem("", "", "")
	// Write data
	for _, doc := range datareturn {

		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), doc.Name)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), doc.Description)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), doc.Vote)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), doc.UserName)
		datetimeString := time.Date(doc.CreateTime.Year(), doc.CreateTime.Month(), doc.CreateTime.Day(), doc.CreateTime.Hour(), doc.CreateTime.Minute(), doc.CreateTime.Second(), 0, time.UTC).Format("2006-01-02 15:04:05")
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), datetimeString)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", row), doc.Status)
		row++
	}

	// Save Excel file
	if err := f.SaveAs("exportitem.xlsx"); err != nil {
		log.Fatal(err)
	}
	ctx.Response().Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Response().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=exportitem.xlsx"))

	return f.Write(ctx.Response())
}

func (c *Controller) ExportVoteByItem(ctx echo.Context) error {
	ID, err := Converthex(ctx.Param("id"))
	log.Println(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusNoContent, map[string]interface{}{
			"error": "No Content",
		})
	}
	// Create Excel file
	f := excelize.NewFile()
	f.NewSheet("Sheet1")

	// Write headers
	//headers := []string{"Name", "Description", "Vote", "User", "CreateDate", "Status"}
	row := 1

	f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), "Name")
	f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), "CreateDate")
	row++
	datareturn := c.model.ReportVoteItemById(ID, "", "")
	// Write data
	for _, doc := range datareturn {
		u, _ := mapUser(doc.UserID)
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), u)
		datetimeString := time.Date(doc.CreateTime.Year(), doc.CreateTime.Month(), doc.CreateTime.Day(), doc.CreateTime.Hour(), doc.CreateTime.Minute(), doc.CreateTime.Second(), 0, time.UTC).Format("2006-01-02 15:04:05")
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), datetimeString)
		row++
	}

	// Save Excel file
	if err := f.SaveAs("voteitem.xlsx"); err != nil {
		log.Fatal(err)
	}
	ctx.Response().Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Response().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=voteitem.xlsx"))

	return f.Write(ctx.Response())
}
