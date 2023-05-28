import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import { BrowserRouter } from "react-router-dom";
import { ThemeProvider } from "@material-tailwind/react";
import { MaterialTailwindControllerProvider } from "@/context";
import { QueryClient, QueryClientProvider, useQuery } from 'react-query'
import "../public/css/tailwind.css";

const queryClient = new QueryClient()

ReactDOM.createRoot(document.getElementById("root")).render(
  <QueryClientProvider client={queryClient}>
    <BrowserRouter>
      <ThemeProvider>
        <MaterialTailwindControllerProvider>
          <App />
        </MaterialTailwindControllerProvider>
      </ThemeProvider>
    </BrowserRouter>
  </QueryClientProvider>
);
