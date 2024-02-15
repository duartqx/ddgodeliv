import React from "react";
import { Link, useLocation } from "react-router-dom";
import Logout from "./Logout";
import { Paths as P } from "../../path";

function NavBarLi({ path, atPath, icon }) {
  return (
    <li className="nav-link">
      <Link to={path || "#"}>
        <i className={`bi bi-${icon} ${atPath ? "" : "text-white"}`}></i>
      </Link>
    </li>
  );
}

function NavBarLiGroup({ group }) {
  return (
    <>
      {group.map((p, i) => (
        <>
          <NavBarLi path={p.path} atPath={p.atPath} icon={p.icon} />
          {(i + 1) % 2 === 0 && <hr className="mx-2 text-white" />}
        </>
      ))}
    </>
  );
}

export default function NavBar() {
  const location = useLocation();

  const at = (/** @type {string} */ path) => location.pathname === path;

  return (
    <>
      <nav
        className="d-flex flex-column flex-shrink-0"
        style={{ width: "4.5rem", height: "100vh", backgroundColor: "#000" }}
        data-bs-theme="dark"
      >
        <ul className="nav nav-pills nav-flush flex-column mb-auto text-center pt-2">
          <NavBarLiGroup
            group={[
              { path: P.root, atPath: at(P.root), icon: "house" },
              { path: "#", atPath: false, icon: "bell" },
              { path: P.drivers, atPath: at(P.drivers), icon: "person" },
              { path: P.vehicles, atPath: at(P.vehicles), icon: "truck" },
              {
                path: P.delivery.available,
                atPath: at(P.delivery.available),
                icon: "hourglass",
              },
              {
                path: P.delivery.company,
                atPath: at(P.delivery.company),
                icon: "tags-fill",
              },
            ]}
          />
        </ul>

        <hr className="mx-2 text-white" />

        <div className="d-flex align-items-center justify-content-center p-3">
          <Logout />
        </div>
      </nav>
    </>
  );
}
