import { Store } from "@tanstack/solid-store";
import type { Bangumi } from "~/types";

export const BangumiStore = new Store({
  bangumi: null as Bangumi | null,
});
