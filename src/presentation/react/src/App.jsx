import React from "react";
import {
  createBrowserRouter,
  Navigate,
  RouterProvider,
} from "react-router-dom";
import Login from "./pages/Login";
import Register from "./pages/Register";
import Dashboard from "./pages/Dashboard";
import NotSignedInRouter from "./middlewares/NotSignedInRouter";
import PrivateRouter from "./middlewares/PrivateRouter";
import AuthProvider from "./middlewares/AuthProvider";

function App() {
  const router = createBrowserRouter([
    {
      path: "/",
      element: <Navigate to="/dashboard" replace />,
    },
    {
      path: "/login",
      element: (
        <NotSignedInRouter>
          <Login />
        </NotSignedInRouter>
      ),
    },
    {
      path: "/register",
      element: (
        <NotSignedInRouter>
          <Register />
        </NotSignedInRouter>
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

  return (
    <AuthProvider>
      <RouterProvider router={router} />
    </AuthProvider>
  );
}

export default App;
