import { createFileRoute, useParams } from "@tanstack/react-router";
import BoardView from "../../components/Board/Boardview";

export const Route = createFileRoute("/boards/$boardId")({
  component: BoardView,
});
