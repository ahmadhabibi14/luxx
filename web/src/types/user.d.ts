interface avatar {
  String: string;
  Valid: boolean;
}
export interface User {
  username: string;
  fullname: string;
  email: string;
  avatar: Avatar;
  join_at: Date;
}