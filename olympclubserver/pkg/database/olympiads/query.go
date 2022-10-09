package olympiads

import "fmt"

type OlympiadFilter struct {
	Subject       string
	Level         string
	Grade         string
	BigOlympiadID int32
	OlympiadID    int32
	OlympiadShort string
}

func NewOlympiadFilter() OlympiadFilter {
	return OlympiadFilter{
		Subject:       "",
		Level:         "",
		Grade:         "",
		BigOlympiadID: -1,
		OlympiadID:    -1,
		OlympiadShort: "",
	}
}

func GetQueryOlympiadOptions(filter OlympiadFilter) (string, []interface{}) {
	inters := make([]interface{}, 0)
	count := 1
	query := "Select * from \"OlympiadModel\" "
	baseLen := len(query)
	if filter.Subject != "" {
		query += fmt.Sprintf("WHERE Subject = $%d", count)
		count += 1
		inters = append(inters, filter.Subject)
	}
	if filter.Level != "" {
		if len(query) == baseLen {
			query += " WHERE "
		} else {
			query += " AND "
		}
		query += fmt.Sprintf("Level = $%d", count)
		count += 1
		inters = append(inters, filter.Level)
	}
	if filter.Grade != "" {
		if len(query) == baseLen {
			query += " WHERE "
		} else {
			query += " AND "
		}
		query += fmt.Sprintf("Grade LIKE $%d", count)
		count += 1
		inters = append(inters, fmt.Sprintf("%%%s%%", filter.Grade))
	}
	if filter.BigOlympiadID != -1 {
		if len(query) == baseLen {
			query += " WHERE "
		} else {
			query += " AND "
		}
		query += fmt.Sprintf("big_olympiad_id = $%d", count)
		count += 1
		inters = append(inters, filter.BigOlympiadID)
	}
	if filter.OlympiadID != -1 {
		if len(query) == baseLen {
			query += " WHERE "
		} else {
			query += " AND "
		}
		query += fmt.Sprintf("id = $%d", count)
		count += 1
		inters = append(inters, filter.OlympiadID)
	}
	if filter.OlympiadShort != "" {
		if len(query) == baseLen {
			query += " WHERE "
		} else {
			query += " AND "
		}
		query += fmt.Sprintf("short = $%d", count)
		count += 1
		inters = append(inters, filter.OlympiadShort)
	}
	fmt.Println(query)
	return query, inters
}

type BigOlympiadFilter struct {
	ID    int32
	Name  string
	Short string
}

func NewBigOlympiadFilter() BigOlympiadFilter {
	return BigOlympiadFilter{
		ID:    -1,
		Name:  "",
		Short: "",
	}
}

func GetQueryBigOlympiadOptions(filter BigOlympiadFilter) (string, []interface{}) {
	inters := make([]interface{}, 0)
	count := 1
	query := "Select * from \"BigOlympiadModel\" "
	baseLen := len(query)
	if filter.Name != "" {
		query += fmt.Sprintf("WHERE name = $%d", count)
		count += 1
		inters = append(inters, filter.Name)
	}
	if filter.Short != "" {
		if len(query) == baseLen {
			query += " WHERE "
		} else {
			query += " AND "
		}
		query += fmt.Sprintf("Short = $%d", count)
		count += 1
		inters = append(inters, filter.Short)
	}
	if filter.ID != -1 {
		if len(query) == baseLen {
			query += " WHERE "
		} else {
			query += " AND "
		}
		query += fmt.Sprintf("big_olympiad_id = $%d", count)
		count += 1
		inters = append(inters, filter.ID)
	}
	fmt.Println(query)
	return query, inters
}
