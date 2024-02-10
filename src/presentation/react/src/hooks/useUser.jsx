import { useState } from "react";
import getUser from "../services/user";

export default async function useUser() {
  const [user, setUser] = useState(await getUser());
  return { user, setUser };
}
