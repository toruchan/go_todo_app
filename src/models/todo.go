package models

import(
	"github.com/wcl48/valval"
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
	"database/sql"
	"time"
    "regexp"
	"fmt"
)

type Todo struct {
    Id int64
    Todo string `sql:"size:255"`
	DeletedFlag bool
    CreatedAt time.Time
    UpdatedAT time.Time
    DeletedAt time.Time
}

func TodoValidate(todo Todo)(error){
    Validator := valval.Object(valval.M{
        "Todo": valval.String(
                valval.MaxLength(20),
                valval.Regexp(regexp.MustCompile(`^[a-z ]+$`)),
        ),
		"DeletedFlag": valval.Bool(),
    })

    return Validator.Validate(todo)
}
func InsertTask(todo string){
	db, _ := gorm.Open("mysql", "root:password@/go_tutorial")
	db.Create(&Todo{0,todo,false,time.Now(),time.Now(),time.Now()})
}

func FindAll() *sql.Rows {
	db, _ := gorm.Open("mysql", "root:password@/go_tutorial")
	eventsEx, err := db.Raw("SELECT id FROM todos WHERE id = 16").Rows()
	fmt.Println("%v",eventsEx)
	fmt.Println("%v",err)
	return eventsEx
}

func SelectAll()(map[int]string) {
	todos := map[int]string{}
	flag := 0
	db, err := sql.Open("mysql", "root:password@/go_tutorial")
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Query("SELECT id,todo FROM todos WHERE deleted_flag = ?", flag)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		id, todo := 0, ""
		err = rows.Scan(&id, &todo)
		if err != nil {
			fmt.Println(err)
		}
		todos[id] = todo
	}
	return todos
}
