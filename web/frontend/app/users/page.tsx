"use client"

import { useState, useEffect } from "react"
// 在导入部分添加 createUser
import { fetchUserList, deleteUser, updateUser, updateUserPassword, createUser } from "@/services/api"
import { withAuth } from "@/components/auth-guard"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { useDebounce } from "@/hooks/use-debounce"
import type { User } from '@/types/types'
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogFooter } from "@/components/ui/dialog"
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"

import { Pencil, Key, Trash2 } from "lucide-react"
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip"

import { Search } from "lucide-react"


// 在组件顶部添加状态
// 添加新的状态
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"

// 添加导入
import {
  Pagination,
  PaginationContent,
  PaginationEllipsis,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "@/components/ui/pagination"

function UsersPage() {
  const [users, setUsers] = useState<User[]>([])
  const [total, setTotal] = useState(0)
  const [page, setPage] = useState(1)
  const [pageSize] = useState(10)
  const [keyword, setKeyword] = useState("")
  const [loading, setLoading] = useState(false)
  const [editingUser, setEditingUser] = useState<User | null>(null)
  const [passwordUser, setPasswordUser] = useState<User | null>(null)
  const [deletingUser, setDeletingUser] = useState<User | null>(null)
  const [newPassword, setNewPassword] = useState('')
  const [confirmPassword, setConfirmPassword] = useState('')
  const [passwordError, setPasswordError] = useState('')
  const [isCreating, setIsCreating] = useState(false)
  const [createForm, setCreateForm] = useState({
    userName: '',
    password: '',
    confirmPassword: ''
  })
  const [createError, setCreateError] = useState('')
  const [editForm, setEditForm] = useState({
    userName: '',
    enable: 1,
    system: 0,
  })
  const debouncedKeyword = useDebounce(keyword, 500)

  const loadUsers = async (searchKeyword?: string) => {
    try {
      setLoading(true)
      const { result } = await fetchUserList({
        page,
        pageSize,
        keyword: searchKeyword || ''
      })
      setUsers(result.list)
      setTotal(result.total)
    } catch (err) {
      console.error("Failed to load users:", err)
    } finally {
      setLoading(false)
    }
  }

  const handleSearch = () => {
    setPage(1)
    loadUsers(keyword)
  }

  // 修改 useEffect 依赖
  useEffect(() => {
    loadUsers(debouncedKeyword)
  }, [page, debouncedKeyword])

  const handleEdit = (user: User) => {
    setEditingUser(user)
    setEditForm({
      userName: user.userName,
      enable: user.enable,
      system: user.system,
    })
  }

  const handleUpdate = async () => {
    if (!editingUser) return
    try {
      await updateUser(editingUser.id, editForm)
      setEditingUser(null)
      loadUsers()
    } catch (err) {
      console.error("Failed to update user:", err)
    }
  }

  const handlePasswordUpdate = async () => {
    if (!passwordUser || !newPassword) return
    if (newPassword !== confirmPassword) {
      setPasswordError('两次输入的密码不一致')
      return
    }
    try {
      await updateUserPassword(passwordUser.id, newPassword)
      setPasswordUser(null)
      setNewPassword('')
      setConfirmPassword('')
      setPasswordError('')
      loadUsers()
    } catch (err) {
      console.error("Failed to update password:", err)
    }
  }

  const handleDelete = (user: User) => {
    setDeletingUser(user)
  }

  const confirmDelete = async () => {
    if (!deletingUser) return
    try {
      await deleteUser(deletingUser.id)
      setDeletingUser(null)
      loadUsers()
    } catch (err) {
      console.error("Failed to delete user:", err)
    }
  }

  // 添加创建用户处理函数
  // 修改 handleCreate 函数
  const handleCreate = async () => {
    if (!createForm.userName || !createForm.password) return
    if (createForm.password !== createForm.confirmPassword) {
      setCreateError('两次输入的密码不一致')
      return
    }
    try {
      await createUser({
        userName: createForm.userName,
        password: createForm.password
      })
      setIsCreating(false)
      setCreateForm({ userName: '', password: '', confirmPassword: '' })
      setCreateError('')
      loadUsers()
    } catch (err: any) {
      // 设置后端返回的错误信息
      setCreateError(err.message || '创建用户失败')
    }
  }

  // 添加分页计算
  const totalPages = Math.ceil(total / pageSize)
  
  const getPageRange = () => {
    const range: (number | 'ellipsis')[] = []
    const maxVisiblePages = 5
    
    if (totalPages <= maxVisiblePages) {
      return Array.from({ length: totalPages }, (_, i) => i + 1)
    }

    range.push(1)
    
    let start = Math.max(2, page - 1)
    let end = Math.min(totalPages - 1, page + 1)
    
    if (page - 1 > 2) {
      range.push('ellipsis')
    }
    
    for (let i = start; i <= end; i++) {
      range.push(i)
    }
    
    if (totalPages - page > 2) {
      range.push('ellipsis')
    }
    
    if (end < totalPages) {
      range.push(totalPages)
    }
    
    return range
  }

  return (
    <>
      <div className="p-6 space-y-6">
        <h2 className="text-2xl font-bold">用户管理</h2>

        {/* 搜索和操作区域 */}
        <div className="flex flex-col md:flex-row gap-4">
    <div className="flex gap-4">
      <Input
        placeholder="搜索用户..."
        value={keyword}
        onChange={(e) => setKeyword(e.target.value)}
        className="w-64"
      />
    </div>
    <div className="flex-1 flex justify-end">
      <Button onClick={() => setIsCreating(true)}>创建用户</Button>
    </div>
  </div>

        {/* 表格卡片 */}
        <Card>
          <CardHeader>
            <CardTitle>用户列表</CardTitle>
          </CardHeader>
          <CardContent>
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>ID</TableHead>
                  <TableHead>用户名</TableHead>
                  <TableHead>昵称</TableHead>
                  <TableHead>邮箱</TableHead>
                  <TableHead>电话</TableHead>
                  <TableHead>状态</TableHead>
                  <TableHead>系统</TableHead>
                  <TableHead>创建时间</TableHead>
                  <TableHead>更新时间</TableHead>
                  <TableHead>操作</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {users.map((user) => (
                  <TableRow key={user.id}>
                    <TableCell>{user.id}</TableCell>
                    <TableCell>{user.userName}</TableCell>
                    <TableCell>{user.nickName || '-'}</TableCell>
                    <TableCell>{user.email || '-'}</TableCell>
                    <TableCell>{user.phone || '-'}</TableCell>
                    <TableCell>{user.enable ? '启用' : '禁用'}</TableCell>
                    <TableCell>{user.system ? '是' : '否'}</TableCell>
                    <TableCell>{new Date(user.createdAt).toLocaleString()}</TableCell>
                    <TableCell>{new Date(user.updatedAt).toLocaleString()}</TableCell>

                    <TableCell>
                      <div className="flex gap-2">
                        <TooltipProvider>
                          <Tooltip>
                            <TooltipTrigger asChild>
                              <Button 
                                variant="ghost" 
                                size="icon"
                                onClick={() => handleEdit(user)}
                              >
                                <Pencil className="h-4 w-4" />
                              </Button>
                            </TooltipTrigger>
                            <TooltipContent>
                              <p>编辑用户</p>
                            </TooltipContent>
                          </Tooltip>
                        
                          <Tooltip>
                            <TooltipTrigger asChild>
                              <Button 
                                variant="ghost" 
                                size="icon"
                                onClick={() => setPasswordUser(user)}
                              >
                                <Key className="h-4 w-4" />
                              </Button>
                            </TooltipTrigger>
                            <TooltipContent>
                              <p>修改密码</p>
                            </TooltipContent>
                          </Tooltip>
                        
                          <Tooltip>
                            <TooltipTrigger asChild>
                              <Button 
                                variant="ghost" 
                                size="icon"
                                onClick={() => handleDelete(user)}
                                className="text-red-500 hover:text-red-600"
                              >
                                <Trash2 className="h-4 w-4" />
                              </Button>
                            </TooltipTrigger>
                            <TooltipContent>
                              <p>删除用户</p>
                            </TooltipContent>
                          </Tooltip>
                        </TooltipProvider>
                      </div>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
            {total > 0 ? (
            <div className="mt-4 flex flex-col items-center justify-center gap-4">
              <Pagination>
                <PaginationContent>
                  <PaginationItem>
                    <PaginationPrevious 
                      onClick={() => page > 1 && setPage(page - 1)}
                      className={page === 1 ? 'pointer-events-none opacity-50' : ''}
                    />
                  </PaginationItem>
                  {getPageRange().map((pageNum, index) => (
                    <PaginationItem key={index}>
                      {pageNum === 'ellipsis' ? (
                        <PaginationEllipsis />
                      ) : (
                        <PaginationLink 
                          onClick={() => setPage(pageNum as number)}
                          isActive={page === pageNum}
                        >
                          {pageNum}
                        </PaginationLink>
                      )}
                    </PaginationItem>
                  ))}
                  <PaginationItem>
                    <PaginationNext 
                      onClick={() => page < totalPages && setPage(page + 1)}
                      className={page === totalPages ? 'pointer-events-none opacity-50' : ''}
                    />
                  </PaginationItem>
                </PaginationContent>
              </Pagination>
              <p className="text-sm text-muted-foreground text-center">
                显示 {total > 0 ? ((page - 1) * pageSize) + 1 : 0} - {Math.min(page * pageSize, total)} 项，共 {total} 项
              </p>
            </div>
          ) : (
            <div className="mt-4 text-center text-muted-foreground">
              暂无数据
            </div>
          )}
          </CardContent>
        </Card>
      </div>

      <Dialog open={!!passwordUser} onOpenChange={() => {
          setPasswordUser(null)
          setNewPassword('')
          setConfirmPassword('')
          setPasswordError('')
        }}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>修改密码</DialogTitle>
          </DialogHeader>
          <div className="space-y-4 py-4">
            <div className="flex flex-col gap-4">
              <div className="flex flex-col gap-2">
                <label>新密码</label>
                <Input
                  type="password"
                  value={newPassword}
                  onChange={(e) => {
                    setNewPassword(e.target.value)
                    setPasswordError('')
                  }}
                />
              </div>
              <div className="flex flex-col gap-2">
                <label>确认密码</label>
                <Input
                  type="password"
                  value={confirmPassword}
                  onChange={(e) => {
                    setConfirmPassword(e.target.value)
                    setPasswordError('')
                  }}
                />
              </div>
              {passwordError && (
                <p className="text-sm text-red-500">{passwordError}</p>
              )}
            </div>
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => {
              setPasswordUser(null)
              setNewPassword('')
              setConfirmPassword('')
              setPasswordError('')
            }}>
              取消
            </Button>
            <Button 
              onClick={handlePasswordUpdate}
              disabled={!newPassword || !confirmPassword || newPassword !== confirmPassword}
            >
              保存
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      <Dialog open={!!editingUser} onOpenChange={() => setEditingUser(null)}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>编辑用户</DialogTitle>
          </DialogHeader>
          <div className="space-y-4 py-4">
            <div className="grid gap-4">
              <div className="flex flex-col gap-2">
                <label>用户名</label>
                <Input
                  value={editForm.userName}
                  onChange={(e) => setEditForm(prev => ({
                    ...prev,
                    userName: e.target.value
                  }))}
                />
              </div>
              <div className="flex flex-col gap-2">
                <label>状态</label>
                <select
                  value={editForm.enable}
                  onChange={(e) => setEditForm(prev => ({
                    ...prev,
                    enable: Number(e.target.value)
                  }))}
                  className="p-2 border rounded"
                >
                  <option value={1}>启用</option>
                  <option value={0}>禁用</option>
                </select>
              </div>
              <div className="flex flex-col gap-2">
                <label>系统账号</label>
                <select
                  value={editForm.system}
                  onChange={(e) => setEditForm(prev => ({
                    ...prev,
                    system: Number(e.target.value)
                  }))}
                  className="p-2 border rounded"
                >
                  <option value={1}>是</option>
                  <option value={0}>否</option>
                </select>
              </div>
            </div>
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => setEditingUser(null)}>
              取消
            </Button>
            <Button onClick={handleUpdate}>
              保存
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* 添加确认删除对话框 */}
      <Dialog open={!!deletingUser} onOpenChange={() => setDeletingUser(null)}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>确认删除</DialogTitle>
          </DialogHeader>
          <div className="py-4">
            <p>确定要删除用户 "{deletingUser?.userName}" 吗？此操作不可恢复。</p>
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => setDeletingUser(null)}>
              取消
            </Button>
            <Button variant="destructive" onClick={confirmDelete}>
              删除
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* 修改创建用户对话框 */}
      <Dialog open={isCreating} onOpenChange={(open) => {
        if (!open) {
          setIsCreating(false)
          setCreateForm({ userName: '', password: '', confirmPassword: '' })
          setCreateError('')
        }
      }}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>创建用户</DialogTitle>
          </DialogHeader>
          <div className="space-y-4 py-4">
            <div className="flex flex-col gap-4">
              <div className="flex flex-col gap-2">
                <label>用户名</label>
                <Input
                  value={createForm.userName}
                  onChange={(e) => {
                    setCreateForm(prev => ({
                      ...prev,
                      userName: e.target.value
                    }))
                    setCreateError('')  // 清除错误信息
                  }}
                />
              </div>
              <div className="flex flex-col gap-2">
                <label>密码</label>
                <Input
                  type="password"
                  value={createForm.password}
                  onChange={(e) => {
                    setCreateForm(prev => ({
                      ...prev,
                      password: e.target.value
                    }))
                    setCreateError('')
                  }}
                />
              </div>
              <div className="flex flex-col gap-2">
                <label>确认密码</label>
                <Input
                  type="password"
                  value={createForm.confirmPassword}
                  onChange={(e) => {
                    setCreateForm(prev => ({
                      ...prev,
                      confirmPassword: e.target.value
                    }))
                    setCreateError('')
                  }}
                />
              </div>
              {createError && (
                <p className="text-sm text-red-500 bg-red-50 p-2 rounded">{createError}</p>
              )}
            </div>
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => {
              setIsCreating(false)
              setCreateForm({ userName: '', password: '', confirmPassword: '' })
              setCreateError('')
            }}>
              取消
            </Button>
            <Button 
              onClick={handleCreate}
              disabled={!createForm.userName || !createForm.password || !createForm.confirmPassword || createForm.password !== createForm.confirmPassword}
            >
              创建
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </>
  )
}

export default withAuth(UsersPage)