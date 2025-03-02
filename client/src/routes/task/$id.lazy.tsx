import { createLazyFileRoute } from "@tanstack/react-router";
import TaskView from "../../components/Task/TaskView";

export const Route = createLazyFileRoute("/task/$id")({
  component: TaskView,
});
