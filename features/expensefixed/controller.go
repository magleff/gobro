package expensefixed

import (
	"github.com/magleff/gobro/database"
	"log"
	"strconv"
)

type ExpenseFixedController struct {
	Datastore *ExpenseFixedDataStore
}

func NewController(DB *database.Database) *ExpenseFixedController {
	instance := new(ExpenseFixedController)
	instance.Datastore = NewDatastore(DB)
	return instance
}

func (self ExpenseFixedController) CreateExpenseFixed(amount string, description string) {
	self.Datastore.CreateExpenseFixed(*NewExpenseFixed(amount, description))
}

func (self ExpenseFixedController) ListExpensesFixed() []ExpenseFixed {
	return self.Datastore.ListExpensesFixed()
}

func (self ExpenseFixedController) RemoveExpenseFixed(index string) {
	self.Datastore.RemoveExpenseFixed(parseIndex(index))
}

func parseIndex(index string) int32 {
	parsedIndex, err := strconv.ParseInt(index, 0, 32)
	if err != nil {
		log.Fatal(err)
	}
	return int32(parsedIndex)
}