import { createFileRoute } from "@tanstack/react-router";
import BoardAdministrateIndex from "../../../components/Board/Administrate/BoardAdministrateIndex";

export const Route = createFileRoute("/boards/administrate/$boardId")({
  component: () => <BoardAdministrateIndex />,
});
