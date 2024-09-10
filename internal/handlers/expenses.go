package handlers

import (
	"clear-cut/internal/models"
	"clear-cut/internal/storage"
	"encoding/json"
	"net/http"
)

// AddExpenseHandler handles adding a new expense
func AddExpenseHandler(w http.ResponseWriter, r *http.Request) {
	var expense models.Expense

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Insert expense into the database
	db := storage.GetDB()
	result, err := db.Exec("INSERT INTO expenses (description, amount, paid_by, group_id) VALUES (?, ?, ?, ?)", 
		expense.Description, expense.Amount, expense.PaidBy, expense.GroupID)
	if err != nil {
		http.Error(w, "Failed to add expense", http.StatusInternalServerError)
		return
	}

	// Retrieve the last inserted ID
	expenseID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to retrieve expense ID", http.StatusInternalServerError)
		return
	}
	expense.ID = int(expenseID)

	// Send the response
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(expense); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// GetExpensesHandler returns all expenses for a specific group
func GetExpensesHandler(w http.ResponseWriter, r *http.Request) {
	groupID := r.URL.Query().Get("group_id")
	if groupID == "" {
		http.Error(w, "Group ID is required", http.StatusBadRequest)
		return
	}

	var expenses []models.Expense
	db := storage.GetDB()
	rows, err := db.Query("SELECT id, description, amount, paid_by, group_id FROM expenses WHERE group_id = ?", groupID)
	if err != nil {
		http.Error(w, "Failed to fetch expenses", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Read expenses from the rows
	for rows.Next() {
		var expense models.Expense
		if err := rows.Scan(&expense.ID, &expense.Description, &expense.Amount, &expense.PaidBy, &expense.GroupID); err != nil {
			http.Error(w, "Failed to scan expense", http.StatusInternalServerError)
			return
		}
		expenses = append(expenses, expense)
	}

	// Handle any error encountered while iterating over rows
	if err := rows.Err(); err != nil {
		http.Error(w, "Failed to fetch expenses", http.StatusInternalServerError)
		return
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(expenses); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
