import { createLazyFileRoute } from "@tanstack/react-router";
import Notifications from "../../components/profile/notifications";

export const Route = createLazyFileRoute("/profile/notifications")({
  component: Notifications,
});
