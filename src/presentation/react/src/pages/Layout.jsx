import React from "react";
import NavBar from "../components/NavBar/NavBar";
import { Outlet } from "react-router-dom";

export default function Layout() {
  return (
    <>
      <main className="d-flex" style={{ width: "100vw"}}>
        <NavBar />
        <Outlet />
      </main>
    </>
  );
}
