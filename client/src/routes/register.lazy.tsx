import { createLazyFileRoute } from "@tanstack/react-router";
import AuthForm from "../components/Auth/Authform";

export const Route = createLazyFileRoute("/register")({
  component: RegisterComponent,
});

function RegisterComponent() {
  return <AuthForm mode="register" />;
}
