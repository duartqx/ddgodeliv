import React from "react";
import { Link } from "react-router-dom";
import Logout from "./Logout";

export default function NavSideBar({ onClickDrivers, onClickVehicles }) {
  return (
    <nav
      className="d-flex flex-column flex-shrink-0"
      style={{ width: "4.5rem", height: "100vh", backgroundColor: "#000" }}
      data-bs-theme="dark"
    >
      <Link
        to="/dashboard"
        className="d-block p-3 link-body-emphasis text-center"
      >
        <i className="bi bi-house"></i>
      </Link>
      <ul className="nav nav-pills nav-flush flex-column mb-auto text-center">
        <li className="nav-link">
          <Link to="/dashboard">
            <i className="bi bi-bell"></i>
          </Link>
        </li>

        <li className="nav-link">
          <Link to="#" onClick={onClickDrivers}>
            <i className="bi bi-person"></i>
          </Link>
        </li>

        <li className="nav-link">
          <Link to="#" onClick={onClickVehicles}>
            <i className="bi bi-truck"></i>
          </Link>
        </li>
      </ul>
      <div className="d-flex align-items-center justify-content-center p-3">
        <Logout />
      </div>
    </nav>
  );
}
