export enum Role {
    Admin,
    General,
    Guest,
}

export interface User {
    id: number;
    username: string;
    email: string;
    role: Role;
    avatar: string;
    isActive: boolean;
}

// 给我一些判断User Role的方法
export const UserUtils = {
    isAdmin: (user: User) => user.role === Role.Admin,
    isGeneral: (user: User) => user.role === Role.General,
    isGuest: (user: User) => user.role === Role.Guest,
};