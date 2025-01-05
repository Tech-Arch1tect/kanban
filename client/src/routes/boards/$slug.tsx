import { createFileRoute } from "@tanstack/react-router";
import BoardView from "../../components/Board/Boardview";

export const Route = createFileRoute("/boards/$slug")({
  component: BoardView,
});
