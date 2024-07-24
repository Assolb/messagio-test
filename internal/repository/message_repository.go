package repository

type MessageRepository struct {
	Database *Database
}

func NewMessageRepository(db *Database) *MessageRepository {
	return &MessageRepository{Database: db}
}

func (r *MessageRepository) CreateTable() error {
	query := `
    CREATE TABLE IF NOT EXISTS messages (
        id SERIAL PRIMARY KEY,
        text TEXT NOT NULL,
        status VARCHAR(50) NOT NULL
    );`

	_, err := r.Database.Db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepository) InsertMessage(text string) (int, error) {
	var id int
	query := `INSERT INTO messages (text, status) VALUES ($1, 'pending') RETURNING id`
	err := r.Database.Db.QueryRow(query, text).Scan(&id)
	return id, err
}

func (r *MessageRepository) UpdateMessageStatus(id int, status string) error {
	query := `UPDATE messages SET status = $1 WHERE id = $2`
	_, err := r.Database.Db.Exec(query, status, id)
	return err
}

func (r *MessageRepository) GetMessageStats() (map[string]int, error) {
	stats := make(map[string]int)
	query := `SELECT status, COUNT(*) FROM messages GROUP BY status`
	rows, err := r.Database.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var status string
		var count int
		if err := rows.Scan(&status, &count); err != nil {
			return nil, err
		}
		stats[status] = count
	}

	return stats, nil
}
