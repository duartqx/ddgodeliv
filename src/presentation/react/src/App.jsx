import React from "react";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Login from "./pages/Login";
import Register from "./pages/Register";
import Dashboard from "./pages/Dashboard";
import LoggedInRouter from "./middlewares/LoggedInRouter";
import PrivateRouter from "./middlewares/PrivateRouter";

function App() {
  const router = createBrowserRouter([
    {
      path: "/login",
      element: (
        <LoggedInRouter>
          <Login />
        </LoggedInRouter>
      ),
    },
    {
      path: "/register",
      element: (
        <LoggedInRouter>
          <Register />
        </LoggedInRouter>
      ),
    },
    {
      path: "/dashboard",
      element: (
        <PrivateRouter>
          <Dashboard />
        </PrivateRouter>
      ),
    },
  ]);

  return <RouterProvider router={router} />;
}

export default App;
