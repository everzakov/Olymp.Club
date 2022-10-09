package news

import "fmt"

type NewsFilter struct {
	ID    int32
	Table string
	Key   int32
}

func NewNewFilter() NewsFilter {
	return NewsFilter{
		ID:    -1,
		Table: "",
		Key:   -1,
	}
}

func GetQueryNewsOptions(filter NewsFilter) (string, []interface{}) {
	inters := make([]interface{}, 0)
	count := 1
	query := "Select * from \"NewsModel\" "
	baseLen := len(query)
	if filter.Table != "" {
		query += fmt.Sprintf("WHERE \"TableStruct\"=$%d", count)
		count += 1
		inters = append(inters, filter.Table)
	}

	if filter.ID != -1 {
		if len(query) == baseLen {
			query += " WHERE "
		} else {
			query += " AND "
		}
		query += fmt.Sprintf("id=$%d", count)
		count += 1
		inters = append(inters, filter.ID)
	}
	if filter.Key != -1 {
		if len(query) == baseLen {
			query += " WHERE "
		} else {
			query += " AND "
		}
		query += fmt.Sprintf("\"Key\"=$%d", count)
		count += 1
		inters = append(inters, filter.Key)
	}
	fmt.Println(query)
	return query, inters
}
