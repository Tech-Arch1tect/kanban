import * as React from "react";
import { Link, Outlet, createRootRoute } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/router-devtools";
import Navbar from "../components/navbar";

export const Route = createRootRoute({
  component: RootComponent,
  notFoundComponent: () => <div>404 Not found</div>,
});

function RootComponent() {
  return (
    <>
      <Navbar />
      <hr />
      <Outlet />
    </>
  );
}
