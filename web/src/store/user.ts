import { Store } from "@tanstack/solid-store";
import type { User } from "~/types";

export const UserStore = new Store({
  user: null as User | null,
});
