export interface User {
  id: number
  uuid: string
  userName: string
  nickName: string
  headerImg: string
  authorityId: number
  phone: string
  email: string
  enable: number
  system: number
  createdAt: string
  updatedAt: string
  Gid: number
  Uid: number
  Pid: number
}

export interface ApiResponse<T> {
  message: string
  status: number
  result: T
}
  
export  interface CaptchaResult {
    CaptchaId: string
    CaptchaLength: number
    OpenCaptcha: boolean
    PicPath: string
  }
  
  export interface LoginRequest {
    username: string
    password: string
    captcha: string
    captchaId: string
  }
  
  export interface LoginResult {
    token: string
    user: {
      id: string
      username: string
      // 根据实际返回数据添加其他字段
    }
  }

  export interface UserListResult {
    list: User[]
    total: number
  }
  
  export interface UserListParams {
    page?: number
    pageSize?: number
    keyword?: string
  }
  
  export interface CreateUserParams {
      userName: string
      password: string
    }
    
    export interface UpdateUserParams {
      userName?: string
      enable?: number
      // ... 其他可选字段
    }


    export interface CreateAppParams {
      name: string
    }

export interface AppListParams {
  page?: number
  pageSize?: number
  keyword?: string
}

export interface App {
  id: number
  name: string
  appId: string
  secret: string
  createdAt: string
  updatedAt: string
}

export interface AppListResult {
  list: App[]
  total: number
}

export interface Permission {
  [key: string]: string[]
}

export interface CreateGroupParams {
  AppId: string
  name: string
  permission: Permission
}

export interface UpdateGroupParams {
  name: string
  permission: Permission
}

export interface Group {
  id: number
  name: string
  appId: string
  permission?: Permission
  createdAt: string
  updatedAt: string
}

export interface GroupListParams {
  page?: number
  pageSize?: number
  keyword?: string
}

export interface GroupListResult {
  list: Group[]
  total: number
}

export interface AppDetailResult {
  app: App;
  groups: Group[];
}

export interface AppUserListResult {
  list: AppUser[];
  total: number;
}

export interface AppUser extends User {
  groupId: number;
  groupName: string;
}

export interface AddAppUserParams {
  userId: number;
  groupId: number;
}

export interface UpdateAppUserGroupParams {
  groupId: number;
}