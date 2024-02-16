import React from "react";
import {
    createBrowserRouter,
    Navigate,
    RouterProvider,
} from "react-router-dom";
import Login from "./pages/Login";
import Register from "./pages/Register";
import Dashboard from "./pages/Dashboard";
import Layout from "./pages/Layout";
import NotSignedInRouter from "./middlewares/NotSignedInRouter";
import PrivateRouter from "./middlewares/PrivateRouter";
import AuthProvider from "./middlewares/AuthProvider";
import DriverList from "./components/SideBarList/Driver/DriverList";
import VehiclesList from "./components/SideBarList/Vehicle/VehiclesList";
import { Paths } from "./path";

function App() {
    const router = createBrowserRouter([
        {
            path: "/",
            element: (
                <PrivateRouter>
                    <Layout />
                </PrivateRouter>
            ),
            children: [
                { index: true, element: <Dashboard /> },
                { path: Paths.drivers, element: <DriverList /> },
                { path: Paths.vehicles, element: <VehiclesList /> },
                { path: Paths.delivery.available, element: <></> },
                { path: Paths.delivery.company, element: <></> },
            ],
        },
        {
            path: Paths.login,
            element: (
                <NotSignedInRouter>
                    <Login />
                </NotSignedInRouter>
            ),
        },
        {
            path: Paths.register,
            element: (
                <NotSignedInRouter>
                    <Register />
                </NotSignedInRouter>
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
