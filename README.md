# CreateOrderAssignmentHandler
CreateOrderAssignmentHandler handler, which waits for a POST request with a JSON body containing courierID and orderID for the order assignment operation. We then call the AssignOrder function, which performs the appropriate update in the database. If the order assignment is successful, a successful JSON response is returned.
