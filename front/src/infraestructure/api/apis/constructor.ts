import { IVoteGroupPort } from "@/domain/ports/ports";
import { createVoteGroup } from "./create-group";
import { getGroups } from "./get-group";
import { CreateVoteGroupRequest } from "../../../domain/request/vote-group";
import { CreateVoteGroupResponse, Group } from "../../../domain/responses/vote-group";
import { getRuntimeEnv } from "@/pkg/env/env";

const { apiUrl } = getRuntimeEnv()

export class VoteGroupAPI implements IVoteGroupPort {
    private readonly baseUrl: string = apiUrl ?? "";

    private async handleRequest<T>(request: Promise<T>): Promise<T> {
        try {
            return await request;
        } catch (error) {
            console.error("API Error:", error);
            throw error;
        }
    }

    createVoteGroup = (data: CreateVoteGroupRequest): Promise<CreateVoteGroupResponse> =>
        this.handleRequest(createVoteGroup(this.baseUrl, data));

    getGroups = (): Promise<Group[]> =>
        this.handleRequest(getGroups(this.baseUrl));
}

export const voteGroupApi = new VoteGroupAPI();
