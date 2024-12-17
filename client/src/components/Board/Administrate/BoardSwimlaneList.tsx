import { ModelsBoard } from "../../../typescript-fetch-client";
import { BoardSwimlaneDelete } from "./BoardSwimlaneDelete";

export const BoardSwimlaneList = ({ board }: { board: ModelsBoard }) => {
  return (
    <div>
      <h2 className="text-xl font-bold">Swimlanes</h2>
      <ul>
        {board.swimlanes?.map((swimlane) => (
          <li key={swimlane.id} className="flex items-center gap-2">
            <span>{swimlane.name}</span>
            <BoardSwimlaneDelete swimlane={swimlane} />
          </li>
        ))}
      </ul>
    </div>
  );
};
