package budget

import (
	"github.com/magleff/gobro/features/expensefixed"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewBudget(t *testing.T) {
	budget := NewBudget("32.52")
	assert.Equal(t, float32(32.52), budget.InitialBalance, "Should init the budget with an initial balance.")
	assert.Equal(t, true, budget.Active, "Should make the budget active by default.")
	assert.Equal(t, time.Now().Format("2006-01-02"),
		budget.StartDate.Format("2006-01-02"),
		"Should set the start date to the current date.")
	assert.Equal(t, 0, len(budget.Expenses), "Should init an empty slice for expenses.")
}

func TestNewBudgetWithExpensesFixed(t *testing.T) {
	expensesFixed := []expensefixed.ExpenseFixed{*expensefixed.NewExpenseFixed("60.50", "test"),
		*expensefixed.NewExpenseFixed("50.25", "test2")}
	budget := NewBudgetWithExpensesFixed(expensesFixed, "32.52")
	assert.Equal(t, 2, len(budget.Expenses), "Should init the budget with fixed expenses")
}
