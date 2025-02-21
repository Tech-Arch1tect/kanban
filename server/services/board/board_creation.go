package board

import (
	"server/database/repository"
	"server/models"
)

func (bs *BoardService) CreateBoard(name, slug string, swimlaneNames, columnNames []string) (models.Board, error) {
	swimlanes := make([]models.Swimlane, len(swimlaneNames))
	for i, name := range swimlaneNames {
		swimlanes[i] = models.Swimlane{Name: name, Order: i}
	}

	columns := make([]models.Column, len(columnNames))
	for i, name := range columnNames {
		columns[i] = models.Column{Name: name, Order: i}
	}

	board := models.Board{
		Name:      name,
		Slug:      slug,
		Swimlanes: swimlanes,
		Columns:   columns,
	}

	err := bs.db.BoardRepository.Create(&board)
	if err != nil {
		return models.Board{}, err
	}

	return bs.db.BoardRepository.GetByID(board.ID)
}

func (bs *BoardService) DeleteBoard(id uint) error {
	// delete swimlanes
	swimlanes, err := bs.db.SwimlaneRepository.GetAll(repository.WithWhere("board_id = ?", id))
	if err != nil {
		return err
	}

	for _, swimlane := range swimlanes {
		err = bs.ss.DeleteSwimlane(swimlane.ID)
		if err != nil {
			return err
		}
	}

	// delete columns
	columns, err := bs.db.ColumnRepository.GetAll(repository.WithWhere("board_id = ?", id))
	if err != nil {
		return err
	}

	for _, column := range columns {
		err = bs.cs.DeleteColumn(column.ID)
		if err != nil {
			return err
		}
	}

	// delete tasks
	tasks, err := bs.db.TaskRepository.GetAll(repository.WithWhere("board_id = ?", id))
	if err != nil {
		return err
	}

	for _, task := range tasks {
		err = bs.ts.DeleteTask(task.ID)
		if err != nil {
			return err
		}
	}

	// delete board
	return bs.db.BoardRepository.Delete(id)
}
