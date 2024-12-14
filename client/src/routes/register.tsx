import { createFileRoute } from "@tanstack/react-router";
import AuthForm from "../components/Auth/Authform";

export const Route = createFileRoute("/register")({
  component: RegisterComponent,
});

function RegisterComponent() {
  return <AuthForm mode="register" />;
}
