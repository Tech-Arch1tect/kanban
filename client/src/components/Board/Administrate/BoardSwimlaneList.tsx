import { ModelsBoard } from "../../../typescript-fetch-client";
import { BoardSwimlaneDelete } from "./BoardSwimlaneDelete";
import { useDraggableSwimlanes } from "../../../hooks/boards/administration/useDraggableSwimlanes";

export const BoardSwimlaneList = ({ board }: { board: ModelsBoard }) => {
  const { swimlanes, onDragStart, onDragOver, onDrop } = useDraggableSwimlanes(
    board.swimlanes ?? []
  );

  return (
    <div>
      <h2 className="text-xl font-bold">Swimlanes</h2>
      <ul>
        {swimlanes.map((swimlane) => (
          <li
            key={swimlane.id}
            className="flex items-center gap-2"
            draggable
            onDragStart={(e) => onDragStart(e, swimlane.id!)}
            onDragOver={onDragOver}
            onDrop={(e) => onDrop(e, swimlane.id!)}
          >
            <span>{swimlane.name}</span>
            <BoardSwimlaneDelete swimlane={swimlane} />
          </li>
        ))}
      </ul>
    </div>
  );
};
