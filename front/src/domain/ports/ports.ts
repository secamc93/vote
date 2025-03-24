// src/core/ports/voteGroup.port.ts

import { CreateVoteGroupRequest } from "@/domain/request/vote-group";
import { CreateVoteGroupResponse, Group } from "@/domain/responses/vote-group";

export interface IVoteGroupPort {
  createVoteGroup(data: CreateVoteGroupRequest): Promise<CreateVoteGroupResponse>;
  getGroups(): Promise<Group[]>;
}
