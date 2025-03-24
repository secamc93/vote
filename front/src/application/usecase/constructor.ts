// src/application/voteGroup/voteGroup.constructor.ts

import { CreateVoteGroupRequest } from "@/domain/request/vote-group";
import { CreateVoteGroupResponse, Group } from "@/domain/responses/vote-group";
import { IVoteGroupPort } from "@/domain/ports/ports";
import { voteGroupApi } from "@/infraestructure/api/apis/constructor";

// 🔷 Interfaz pública del use case compuesto
export interface IVoteGroupUseCases {
    getGroups(): Promise<Group[]>;
    createVoteGroup(data: CreateVoteGroupRequest): Promise<CreateVoteGroupResponse>;
  }
  
  // 🔧 Implementación concreta
class VoteGroupUseCases implements IVoteGroupUseCases {
    constructor(private readonly adapter: IVoteGroupPort) {}
  
    async getGroups(): Promise<Group[]> {
      return this.adapter.getGroups();
    }
  
    async createVoteGroup(data: CreateVoteGroupRequest): Promise<CreateVoteGroupResponse> {
      return this.adapter.createVoteGroup(data);
    }
}
  
  // 🏗️ Exporta una función que instancie el use case inyectando voteGroupApi
  export const createVoteGroupUseCases = (): IVoteGroupUseCases => 
  new VoteGroupUseCases(voteGroupApi);