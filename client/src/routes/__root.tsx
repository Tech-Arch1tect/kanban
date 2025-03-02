import { Outlet, createRootRoute } from "@tanstack/react-router";
import React from "react";
const Navbar = React.lazy(() => import("../components/layout/navbar"));

export const Route = createRootRoute({
  component: RootComponent,
  notFoundComponent: () => <div>404 Not found</div>,
});

function RootComponent() {
  return (
    <>
      <Navbar />
      <Outlet />
    </>
  );
}
