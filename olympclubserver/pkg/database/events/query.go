package events

import "fmt"

type EventFilter struct {
	ID       int32  `json:"id"`
	Short    string `json:"short"`
	HolderID int32  `json:"holder_id"`
}

func NewEventFilter() EventFilter {
	return EventFilter{
		ID:       -1,
		Short:    "",
		HolderID: -1,
	}
}

func GetQueryEventOptions(filter EventFilter) (string, []interface{}) {
	inters := make([]interface{}, 0)
	count := 1
	query := "Select * from \"EventModel\" "
	baseLen := len(query)
	if filter.ID != -1 {
		query += fmt.Sprintf("WHERE id=$%d", count)
		count += 1
		inters = append(inters, filter.ID)
	}
	if filter.Short != "" {
		if len(query) == baseLen {
			query += " WHERE "
		} else {
			query += " AND "
		}
		query += fmt.Sprintf("\"Short\"=$%d", count)
		count += 1
		inters = append(inters, filter.Short)
	}
	if filter.HolderID != -1 {
		if len(query) == baseLen {
			query += " WHERE "
		} else {
			query += " AND "
		}

		query += fmt.Sprintf("holderId=$%d", count)
		count += 1
		inters = append(inters, filter.HolderID)
	}
	fmt.Println(query)
	return query, inters
}
