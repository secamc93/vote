import { Group } from "../../../domain/responses/vote-group";
import { ErrorResponse } from "../../../domain/responses/vote-group";

export async function getGroups(baseUrl: string): Promise<Group[]> {
  try {
    const res = await fetch(`${baseUrl}/api/get-groups`, {
      method: "GET",
      headers: { "Content-Type": "application/json" },
    });

    const json = await res.json();
    if (!res.ok) throw json as ErrorResponse;
    return json as Group[];
  } catch (error) {
    console.error("Error al obtener los grupos:", error);
    throw error;
  }
}