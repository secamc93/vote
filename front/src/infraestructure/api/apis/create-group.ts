import { CreateVoteGroupRequest } from "../../../domain/request/vote-group";
import { CreateVoteGroupResponse, ErrorResponse } from "../../../domain/responses/vote-group";

export async function createVoteGroup(baseUrl: string, data: CreateVoteGroupRequest): Promise<CreateVoteGroupResponse> {
  try {
    const res = await fetch(`${baseUrl}/api/vote-group`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    });

    const json = await res.json();

    if (!res.ok) {
        throw json as ErrorResponse;
    }

    return json as CreateVoteGroupResponse;
  } catch (error) {
    console.error("Error al crear el grupo:", error);
    throw error;
  }
}


