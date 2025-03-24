export interface Group {
    id: number;
    name: string;
    created_at: string;
    houses: House[];
  }
  
export interface House {
    id: number;
    name: string;
    created_at: string;
    vote_group_id: number;
  }

export interface CreateVoteGroupResponse {
    group_id: number;
    message: string;
  }
  
export interface ErrorResponse {
    status_code: number;
    message: string;
  }
  