import { createLazyFileRoute } from "@tanstack/react-router";
import BoardView from "../../components/Board/Boardview";

export const Route = createLazyFileRoute("/boards/$slug")({
  component: BoardView,
});
