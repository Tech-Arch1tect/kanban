import { createFileRoute } from "@tanstack/react-router";
import AuthForm from "../components/Authform";

export const Route = createFileRoute("/login")({
  component: LoginComponent,
});

function LoginComponent() {
  return <AuthForm mode="login" />;
}
