import React from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider, createRouter } from "@tanstack/react-router";
import { routeTree } from "./routeTree.gen";

import "./index.css";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { LocalSettingsProvider } from "./context/LocalSettingsContext";
import { ServerSettingsProvider } from "./context/ServerSettingsContext";

// Set up a Router instance
const router = createRouter({
  routeTree,
  defaultPreload: "intent",
});

const queryClient = new QueryClient();

// Register things for typesafety
declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

const rootElement = document.getElementById("app")!;

if (!rootElement.innerHTML) {
  const root = ReactDOM.createRoot(rootElement);
  root.render(
    <LocalSettingsProvider>
      <QueryClientProvider client={queryClient}>
        <ServerSettingsProvider>
          <RouterProvider router={router} />
        </ServerSettingsProvider>
      </QueryClientProvider>
    </LocalSettingsProvider>
  );
}
