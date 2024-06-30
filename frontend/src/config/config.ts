export const API_URL = 'http://localhost:1323/v1/api';

export interface DatacenterType {
  ID: number;
  Tag: string;
  Info: string;
}

export interface RowType {
  ID: number;
  Tag: string;
  Info: string;
  DatacenterID: number;
}
