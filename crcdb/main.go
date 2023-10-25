import (
    "encoding/json"
    "net/http"
    "github.com/jackc/pgx/v4"
    "context"
)

// CreateOrderAssignmentHandler обрабатывает запрос на присвоение заказа курьеру.
func CreateOrderAssignmentHandler(db *pgx.Conn) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Распарсить JSON запроса в структуру данных для операции присвоения заказа
        var assignment OrderAssignment
        err := json.NewDecoder(r.Body).Decode(&assignment)
        if err != nil {
            http.Error(w, "Ошибка разбора JSON: "+err.Error(), http.StatusBadRequest)
            return
        }

        // Выполнить операцию присвоения заказа
        err = AssignOrder(db, assignment.CourierID, assignment.OrderID)
        if err != nil {
            http.Error(w, "Ошибка при присвоении заказа: "+err.Error(), http.StatusInternalServerError)
            return
        }

        // Отправить ответ об успешном присвоении заказа
        response := map[string]string{"message": "Заказ успешно присвоен курьеру"}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    }
}

// AssignOrder выполняет операцию присвоения заказа курьеру в базе данных.
func AssignOrder(db *pgx.Conn, courierID int, orderID int) error {
    // Здесь вы должны выполнить операцию присвоения заказа в базе данных, например, обновить статус заказа и связать его с курьером.
    // Пример:
    query := `
        UPDATE orders
        SET status = 'assigned',
            courier_id = $1
        WHERE id = $2;
    `
    _, err := db.Exec(context.Background(), query, courierID, orderID)
    return err
}
