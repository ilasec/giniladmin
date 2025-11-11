import type { ApiResponse, LoginRequest, LoginResult, CaptchaResult, UserListParams,
    UserListResult,
    UpdateUserParams,
    CreateUserParams,
    CreateAppParams, 
    AppListParams,
     AppListResult, AppDetailResult} from '@/types/types'
import type { CreateGroupParams, GroupListParams, GroupListResult, UpdateGroupParams } from '@/types/types'
import type { 
  AppUserListResult, 
  AddAppUserParams, 
  UpdateAppUserGroupParams 
} from '@/types/types'

// 添加通用的请求处理函数
async function fetchWithAuth(url: string, options: RequestInit = {}): Promise<Response> {
  const token = localStorage.getItem('token');
  if (!token && !url.includes('/auth/')) {
    throw new Error('未登录或登录已过期');
  }

  const headers = {
    'accept': 'application/json',
    ...(token && !url.includes('/auth/') ? { 'x-token': token } : {}),
    ...options.headers,
  };

  return fetch(url, { ...options, headers });
}

async function handleResponse<T>(response: Response): Promise<ApiResponse<T>> {
  const data: ApiResponse<T> = await response.json();
  if (data.status !== 200) {
    throw new Error(data.message || '请求失败');
  }
  return data;
}

export async function fetchCaptcha(): Promise<ApiResponse<CaptchaResult>> {
  const response = await fetchWithAuth('/api/v1/auth/captcha');
  return handleResponse<CaptchaResult>(response);
}

export async function login(data: LoginRequest): Promise<ApiResponse<LoginResult>> {
  const response = await fetchWithAuth('/api/v1/auth/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(data)
  });
  
  return handleResponse<LoginResult>(response);
}



export async function fetchUserList(params: UserListParams = {}): Promise<ApiResponse<UserListResult>> {
  const queryParams = new URLSearchParams();
  if (params.page) queryParams.append('page', params.page.toString());
  if (params.pageSize) queryParams.append('pageSize', params.pageSize.toString());
  if (params.keyword) {
    queryParams.append('keyword', params.keyword);
  } else {
    queryParams.append('keyword', '');
  }

  const url = `/api/v1/system/user${queryParams.toString() ? `?${queryParams.toString()}` : ''}`;
  const response = await fetchWithAuth(url);
  return handleResponse<UserListResult>(response);
}

export async function deleteUser(id: number): Promise<ApiResponse<null>> {
  const response = await fetchWithAuth(`/api/v1/system/user/${id}`, {
    method: 'DELETE'
  });
  return handleResponse<null>(response);
}

export async function updateUser(id: number, data: UpdateUserParams): Promise<ApiResponse<null>> {
  const response = await fetchWithAuth(`/api/v1/system/user/${id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(data)
  });
  return handleResponse<null>(response);
}

export async function updateUserPassword(userId: number, password: string): Promise<ApiResponse<null>> {
  const response = await fetchWithAuth(`/api/v1/system/user/${userId}/password`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ password })
  });
  return handleResponse<null>(response);
}

export async function createUser(data: CreateUserParams): Promise<ApiResponse<null>> {
  const response = await fetchWithAuth('/api/v1/system/user', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(data)
  });
  return handleResponse<null>(response);
}


export async function createApp(data: CreateAppParams): Promise<ApiResponse<null>> {
  const response = await fetchWithAuth('/api/v1/system/app', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(data)
  });
  return handleResponse<null>(response);
}



export async function fetchAppList(params: AppListParams = {}): Promise<ApiResponse<AppListResult>> {
  const queryParams = new URLSearchParams();
  if (params.page) queryParams.append('page', params.page.toString());
  if (params.pageSize) queryParams.append('pageSize', params.pageSize.toString());
  if (params.keyword) {
    queryParams.append('keyword', params.keyword);
  } else {
    queryParams.append('keyword', '');
  }

  const url = `/api/v1/system/app${queryParams.toString() ? `?${queryParams.toString()}` : ''}`;
  const response = await fetchWithAuth(url);
  return handleResponse<AppListResult>(response);
}

export async function deleteApp(id: number): Promise<ApiResponse<null>> {
  const response = await fetchWithAuth(`/api/v1/system/app/${id}`, {
    method: 'DELETE'
  });
  return handleResponse<null>(response);
}



export async function createGroup(data: CreateGroupParams): Promise<ApiResponse<null>> {
  const response = await fetchWithAuth('/api/v1/system/group', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(data)
  });
  return handleResponse<null>(response);
}

export async function fetchGroupList(params: GroupListParams = {}): Promise<ApiResponse<GroupListResult>> {
  const queryParams = new URLSearchParams();
  if (params.page) queryParams.append('page', params.page.toString());
  if (params.pageSize) queryParams.append('pageSize', params.pageSize.toString());
  if (params.keyword) {
    queryParams.append('keyword', params.keyword);
  } else {
    queryParams.append('keyword', '');
  }

  const url = `/api/v1/system/group${queryParams.toString() ? `?${queryParams.toString()}` : ''}`;
  const response = await fetchWithAuth(url);
  return handleResponse<GroupListResult>(response);
}

export async function updateGroup(id: number, data: UpdateGroupParams): Promise<ApiResponse<null>> {
  const response = await fetchWithAuth(`/api/v1/system/group/${id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(data)
  });
  return handleResponse<null>(response);
}

export async function deleteGroup(id: number): Promise<ApiResponse<null>> {
  const response = await fetchWithAuth(`/api/v1/system/group/${id}`, {
    method: 'DELETE'
  });
  return handleResponse<null>(response);
}





// 获取应用详情
export async function fetchAppDetail(id: string): Promise<ApiResponse<AppDetailResult>> {
  // 模拟数据
  const mockApp = {
    id: 1,
    name: "测试应用",
    appId: "test_app_001",
    secret: "app_secret_xxx",
    createdAt: "2024-01-20T10:00:00Z",
    enable: true,
    updatedAt: "2024-01-20T10:00:00Z"
  }

  const mockGroups =  [
      { 
        id: 1, 
        name: "管理员",
        appId: "AAAA",
        createdAt: "2024-01-20T10:00:00Z",
        updatedAt: "2024-01-20T10:00:00Z"
      },
      { 
        id: 2, 
        name: "普通用户",
        appId: "AAAA",
        createdAt: "2024-01-20T10:00:00Z",
        updatedAt: "2024-01-20T10:00:00Z"
      },
      { 
        id: 3, 
        name: "访客",
        appId: "AAAA",
        createdAt: "2024-01-20T10:00:00Z",
        updatedAt: "2024-01-20T10:00:00Z"
      }
    ]

  return {
    status: 200,
    message: "success",
    result: {
      app: mockApp,
      groups: mockGroups
    }
  }
}

// 获取应用用户列表
export async function fetchAppUsers(appId: string): Promise<ApiResponse<AppUserListResult>> {
  // 模拟数据
  const mockUsers = [
    {
      id: 1,
      uuid: "sdfsdf",
      headerImg:"aaaa",
      authorityId: 0,
      Gid: 0,
      Uid: 0,
      Pid: 0,
      userName: "admin",
      nickName: "管理员",
      email: "admin@example.com",
      phone: "13800138000",
      enable: 1,
      system: 1,
      createdAt: "2024-01-20T10:00:00Z",
      updatedAt: "2024-01-20T10:00:00Z",
      groupId: 1,
      groupName: "管理员"
    },
    {
      id: 2,
      uuid: "sdfsdf",
      headerImg:"aaaa",
      authorityId: 0,
      Gid: 0,
      Uid: 0,
      Pid: 0,
      userName: "user1",
      nickName: "用户1",
      email: "user1@example.com",
      phone: "13800138001",
      enable: 1,
      system: 1,
      createdAt: "2024-01-20T10:00:00Z",
      updatedAt: "2024-01-20T10:00:00Z",
      groupId: 2,
      groupName: "普通用户"
    }
  ]

  return {
    status: 200,
    message: "success",
    result: {
      list: mockUsers,
      total: mockUsers.length
    }
  }
}

// 添加应用用户
export async function addAppUser(appId: string, data: AddAppUserParams): Promise<ApiResponse<null>> {
  // 模拟成功响应
  return {
    status: 200,
    message: "success",
    result: null
  }
}

// 移除应用用户
export async function removeAppUser(appId: string, userId: number): Promise<ApiResponse<null>> {
  // 模拟成功响应
  return {
    status: 200,
    message: "success",
    result: null
  }
}

// 更新应用用户组
export async function updateAppUserGroup(
  appId: string, 
  userId: number, 
  groupId: number
): Promise<ApiResponse<null>> {
  // 模拟成功响应
  return {
    status: 200,
    message: "success",
    result: null
  }
}