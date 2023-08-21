package task

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"gorm.io/gorm"
)

type (
	Services interface {
		GetAll()
		Create()
		Delete()
		Complete()
	}

	ser struct {
		db *gorm.DB
	}
)

func NewSer(db *gorm.DB) Services {
	return &ser{
		db: db,
	}
}
func tables(task []Task) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Complete", "ID", "Name")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, task := range task {
		status := "[ ]"

		if task.Complete == "true" {
			status = "[✓]"
		}

		tbl.AddRow(status, task.Id, task.Name)
	}

	tbl.Print()
}
func (s ser) GetAll() {

	task, err := repo.GetAll(repo{db: s.db})
	if err != nil {
		log.Fatal(err)
	}

	tables(task)

}
func (s ser) Create() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Name:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	err := repo.Create(repo{db: s.db}, name)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("La tarea '%s' fue agregada ", name)
}

func (s ser) Delete() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("id:")
	stringid, _ := reader.ReadString('\n')
	stringid = strings.TrimSpace(stringid)

	id, err := strconv.Atoi(stringid)

	err = repo.Delete(repo{db: s.db}, id)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("La tarea con el id: '%d' fue eliminada ", id)
}

func (s ser) Complete() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("id:")
	stringid, _ := reader.ReadString('\n')
	stringid = strings.TrimSpace(stringid)

	id, err := strconv.Atoi(stringid)

	err = repo.Complete(repo{db: s.db}, id)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Se ha completado la tarea con el id: %d ", id)
}
