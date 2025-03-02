import { createLazyFileRoute } from "@tanstack/react-router";
import BoardAdministrateIndex from "../../../components/Board/Administrate/BoardAdministrateIndex";

export const Route = createLazyFileRoute("/boards/administrate/$boardId")({
  component: () => <BoardAdministrateIndex />,
});
