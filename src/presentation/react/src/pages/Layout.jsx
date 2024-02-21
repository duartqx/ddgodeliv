import React, { useContext, useState } from "react";
import NavBar from "../components/NavBar/NavBar";
import { TitleContext } from "../middlewares/TitleContext";
import { Outlet } from "react-router-dom";

export default function Layout() {
  const [title, setTitle] = useState("");
  return (
    <>
      <main className="d-flex" style={{ width: "100vw" }}>
        <NavBar />
        <TitleContext.Provider value={{ setTitle }}>
          <div className="d-flex flex-column">
            <div
              className="d-flex align-items-center justify-content-center fw-light"
              style={{
                width: "calc(100vw - 4.5rem)",
                height: "4rem",
                backgroundColor: "#f0f2f7",
              }}
            >
              <span>{title}</span>
            </div>
            <Outlet />
          </div>
        </TitleContext.Provider>
      </main>
    </>
  );
}
