import { ModelsBoard } from "../../../typescript-fetch-client";
import { useDeleteColumn } from "../../../hooks/boards/administration/useDeleteColumn";
import { BoardColumnDelete } from "./BoardColumnDelete";
export const BoardColumnsList = ({ board }: { board: ModelsBoard }) => {
  const { deleteColumn } = useDeleteColumn();
  return (
    <div>
      <h2 className="text-xl font-bold">Columns</h2>
      <ul>
        {board.columns?.map((column) => (
          <li key={column.id} className="flex items-center gap-2">
            <span>{column.name}</span>
            <BoardColumnDelete column={column} />
          </li>
        ))}
      </ul>
    </div>
  );
};
