import React, { useContext, useEffect } from "react";
import { TitleContext } from "../middlewares/TitleContext";

/**
 * @typedef {{
 *  selected: "drivers"|"vehicles"|"pending"|"taken"
 *  isSelected: (field: string) => boolean
 * }} SideBar
 */

export default function Dashboard() {
  const { setTitle } = useContext(TitleContext)

  useEffect(() => setTitle("Home"), [])

  return <main className="d-flex"> </main>;
}
