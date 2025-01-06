import { createFileRoute } from "@tanstack/react-router";
import TaskView from "../../components/Task/TaskView";

export const Route = createFileRoute("/task/$id")({
  component: TaskView,
});
