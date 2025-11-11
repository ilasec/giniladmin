"use client"

import { useState, useEffect } from "react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogFooter } from "@/components/ui/dialog"
import { Pencil, Trash2 } from "lucide-react"
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip"
import { fetchGroupList, createGroup, deleteGroup, updateGroup } from "@/services/api"
import { withAuth } from "@/components/auth-guard"
import type { Group } from '@/types/types'
import { Eye } from "lucide-react"

// 添加导入
import { useDebounce } from "@/hooks/use-debounce"
import { Search } from "lucide-react"

// 添加导入
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

function GroupsPage() {
  const [groups, setGroups] = useState<Group[]>([])
  const [total, setTotal] = useState(0)
  const [page, setPage] = useState(1)
  const [pageSize] = useState(10)
  const [loading, setLoading] = useState(false)
  const [isCreating, setIsCreating] = useState(false)
  const [createForm, setCreateForm] = useState({
    name: '',
    appId: '',
    permission: {}
  })
  const [createError, setCreateError] = useState('')
  const [editingGroup, setEditingGroup] = useState<Group | null>(null)
  const [editForm, setEditForm] = useState({
    name: '',
    permission: {}
  })
  const [deletingGroup, setDeletingGroup] = useState<Group | null>(null)
  const [viewingPermissions, setViewingPermissions] = useState<Group | null>(null)
  const [permissionForm, setPermissionForm] = useState<{[key: string]: string[]}>({})

  const [keyword, setKeyword] = useState("")
  const debouncedKeyword = useDebounce(keyword, 500)

  
  // 修改 loadGroups 函数
  const loadGroups = async (searchKeyword?: string) => {
    try {
      setLoading(true)
      const { result } = await fetchGroupList({
        page,
        pageSize,
        keyword: searchKeyword || ''
      })
      setGroups(result.list)
      setTotal(result.total)
    } catch (err) {
      console.error("加载用户组失败:", err)
    } finally {
      setLoading(false)
    }
  }

  // 添加搜索处理函数
  const handleSearch = () => {
    setPage(1)
    loadGroups(keyword)
  }

  // 修改 useEffect
  useEffect(() => {
    loadGroups(debouncedKeyword)
  }, [page, debouncedKeyword])

  const handleViewPermissions = (group: Group) => {
    setViewingPermissions(group)
  }

  // 修改 handleCreate 函数
  const handleCreate = async () => {
    if (!createForm.name || !createForm.appId) return
    try {
      await createGroup({
        name: createForm.name,
        AppId: createForm.appId,
        permission: permissionForm
      })
      setIsCreating(false)
      setCreateForm({ name: '', appId: '', permission: {} })
      setPermissionForm({})
      setCreateError('')
      loadGroups()
    } catch (err: any) {
      setCreateError(err.message || '创建用户组失败')
    }
  }

  // 修改 handleEdit 和 handleUpdate 函数
  const handleEdit = (group: Group) => {
    setEditingGroup(group)
    setEditForm({
      name: group.name,
      permission: group.permission || {}
    })
    setPermissionForm(group.permission || {})
  }

  const handleUpdate = async () => {
    if (!editingGroup) return
    try {
      await updateGroup(editingGroup.id, {
        name: editForm.name,
        permission: permissionForm
      })
      setEditingGroup(null)
      setPermissionForm({})
      loadGroups()
    } catch (err) {
      console.error("Failed to update group:", err)
    }
  }

  const handleDelete = (group: Group) => {
    setDeletingGroup(group)
  }

  const confirmDelete = async () => {
    if (!deletingGroup) return
    try {
      await deleteGroup(deletingGroup.id)
      setDeletingGroup(null)
      loadGroups()
    } catch (err) {
      console.error("Failed to delete group:", err)
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
        <h2 className="text-2xl font-bold">用户组管理</h2>

        {/* 搜索和操作区域 */}
        <div className="flex flex-col md:flex-row gap-4">
          <Input
            placeholder="搜索用户组..."
            value={keyword}
            onChange={(e) => setKeyword(e.target.value)}
            className="max-w-sm"
          />
          <div className="flex-1 flex justify-end">
            <Button onClick={() => setIsCreating(true)}>创建用户组</Button>
          </div>
        </div>

        {/* 表格卡片 */}
        <Card>
          <CardHeader>
            <CardTitle>用户组列表</CardTitle>
          </CardHeader>
          <CardContent>
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>ID</TableHead>
                  <TableHead>名称</TableHead>
                  <TableHead>应用ID</TableHead>
                  <TableHead>创建时间</TableHead>
                  <TableHead>更新时间</TableHead>
                  <TableHead>操作</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {groups.map((group) => (
                  <TableRow key={group.id}>
                    <TableCell>{group.id}</TableCell>
                    <TableCell>{group.name}</TableCell>
                    <TableCell>{group.appId}</TableCell>
                    <TableCell>{new Date(group.createdAt).toLocaleString()}</TableCell>
                    <TableCell>{new Date(group.updatedAt).toLocaleString()}</TableCell>

                    <TableCell>
                      <div className="flex gap-2">
                        <TooltipProvider>
                          <Tooltip>
                            <TooltipTrigger asChild>
                              <Button 
                                variant="ghost" 
                                size="icon"
                                onClick={() => handleViewPermissions(group)}
                              >
                                <Eye className="h-4 w-4" />
                              </Button>
                            </TooltipTrigger>
                            <TooltipContent>
                              <p>查看权限</p>
                            </TooltipContent>
                          </Tooltip>

                          <Tooltip>
                            <TooltipTrigger asChild>
                              <Button 
                                variant="ghost" 
                                size="icon"
                                onClick={() => handleEdit(group)}
                              >
                                <Pencil className="h-4 w-4" />
                              </Button>
                            </TooltipTrigger>
                            <TooltipContent>
                              <p>编辑用户组</p>
                            </TooltipContent>
                          </Tooltip>

                          <Tooltip>
                            <TooltipTrigger asChild>
                              <Button 
                                variant="ghost" 
                                size="icon"
                                onClick={() => handleDelete(group)}
                                className="text-red-500 hover:text-red-600"
                              >
                                <Trash2 className="h-4 w-4" />
                              </Button>
                            </TooltipTrigger>
                            <TooltipContent>
                              <p>删除用户组</p>
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

      {/* 创建用户组对话框 */}
      <Dialog open={isCreating} onOpenChange={(open) => {
        if (!open) {
          setIsCreating(false)
          setCreateForm({ name: '', appId: '', permission: {} })
          setPermissionForm({})
          setCreateError('')
        }
      }}>
        <DialogContent className="max-w-3xl">
          <DialogHeader>
            <DialogTitle>创建用户组</DialogTitle>
          </DialogHeader>
          <div className="space-y-4 py-4">
            <div className="flex flex-col gap-4">
              <div className="flex flex-col gap-2">
                <label>用户组名称</label>
                <Input
                  value={createForm.name}
                  onChange={(e) => {
                    setCreateForm(prev => ({
                      ...prev,
                      name: e.target.value
                    }))
                    setCreateError('')
                  }}
                />
              </div>
              <div className="flex flex-col gap-2">
                <label>应用ID</label>
                <Input
                  value={createForm.appId}
                  onChange={(e) => {
                    setCreateForm(prev => ({
                      ...prev,
                      appId: e.target.value
                    }))
                    setCreateError('')
                  }}
                />
              </div>
              <div className="flex flex-col gap-2">
                <label>权限设置</label>
                <div className="space-y-2">
                  {Object.entries(permissionForm).map(([key, values], index) => (
                    <div key={index} className="flex gap-2">
                      <Input
                        placeholder="权限键"
                        value={key}
                        onChange={(e) => {
                          const newKey = e.target.value
                          const newPermissionForm = { ...permissionForm }
                          delete newPermissionForm[key]
                          newPermissionForm[newKey] = values
                          setPermissionForm(newPermissionForm)
                        }}
                      />
                      <Input
                        placeholder="权限值（用逗号分隔）"
                        value={values.join(',')}
                        onChange={(e) => {
                          setPermissionForm({
                            ...permissionForm,
                            [key]: e.target.value.split(',').filter(Boolean)
                          })
                        }}
                      />
                      <Button
                        variant="ghost"
                        size="icon"
                        onClick={() => {
                          const newPermissionForm = { ...permissionForm }
                          delete newPermissionForm[key]
                          setPermissionForm(newPermissionForm)
                        }}
                      >
                        <Trash2 className="h-4 w-4" />
                      </Button>
                    </div>
                  ))}
                  <Button
                    variant="outline"
                    onClick={() => {
                      setPermissionForm({
                        ...permissionForm,
                        [`权限${Object.keys(permissionForm).length + 1}`]: []
                      })
                    }}
                  >
                    添加权限
                  </Button>
                </div>
              </div>
              {createError && (
                <p className="text-sm text-red-500 bg-red-50 p-2 rounded">{createError}</p>
              )}
            </div>
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => {
              setIsCreating(false)
              setCreateForm({ name: '', appId: '', permission: {} })
              setPermissionForm({})
              setCreateError('')
            }}>
              取消
            </Button>
            <Button 
              onClick={handleCreate}
              disabled={!createForm.name || !createForm.appId}
            >
              创建
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* 编辑用户组对话框 */}
      <Dialog open={!!editingGroup} onOpenChange={() => setEditingGroup(null)}>
        <DialogContent className="max-w-3xl">
          <DialogHeader>
            <DialogTitle>编辑用户组</DialogTitle>
          </DialogHeader>
          <div className="space-y-4 py-4">
            <div className="flex flex-col gap-4">
              <div className="flex flex-col gap-2">
                <label>用户组名称</label>
                <Input
                  value={editForm.name}
                  onChange={(e) => setEditForm(prev => ({
                    ...prev,
                    name: e.target.value
                  }))}
                />
              </div>
              <div className="flex flex-col gap-2">
                <label>权限设置</label>
                <div className="space-y-2">
                  {Object.entries(permissionForm).map(([key, values], index) => (
                    <div key={index} className="flex gap-2">
                      <Input
                        placeholder="权限键"
                        value={key}
                        onChange={(e) => {
                          const newKey = e.target.value
                          const newPermissionForm = { ...permissionForm }
                          delete newPermissionForm[key]
                          newPermissionForm[newKey] = values
                          setPermissionForm(newPermissionForm)
                        }}
                      />
                      <Input
                        placeholder="权限值（用逗号分隔）"
                        value={values.join(',')}
                        onChange={(e) => {
                          setPermissionForm({
                            ...permissionForm,
                            [key]: e.target.value.split(',').filter(Boolean)
                          })
                        }}
                      />
                      <Button
                        variant="ghost"
                        size="icon"
                        onClick={() => {
                          const newPermissionForm = { ...permissionForm }
                          delete newPermissionForm[key]
                          setPermissionForm(newPermissionForm)
                        }}
                      >
                        <Trash2 className="h-4 w-4" />
                      </Button>
                    </div>
                  ))}
                  <Button
                    variant="outline"
                    onClick={() => {
                      setPermissionForm({
                        ...permissionForm,
                        [`权限${Object.keys(permissionForm).length + 1}`]: []
                      })
                    }}
                  >
                    添加权限
                  </Button>
                </div>
              </div>
            </div>
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => {
              setEditingGroup(null)
              setEditForm({ name: '', permission: {} })
              setPermissionForm({})
            }}>
              取消
            </Button>
            <Button onClick={handleUpdate} disabled={!editForm.name}>
              保存
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* 确认删除对话框 */}
      <Dialog open={!!deletingGroup} onOpenChange={() => setDeletingGroup(null)}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>确认删除</DialogTitle>
          </DialogHeader>
          <div className="py-4">
            <p>确定要删除用户组 "{deletingGroup?.name}" 吗？此操作不可恢复。</p>
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => setDeletingGroup(null)}>
              取消
            </Button>
            <Button variant="destructive" onClick={confirmDelete}>
              删除
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* 添加查看权限对话框 */}
      <Dialog open={!!viewingPermissions} onOpenChange={() => setViewingPermissions(null)}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>查看权限 - {viewingPermissions?.name}</DialogTitle>
          </DialogHeader>
          <div className="space-y-4 py-4">
            {viewingPermissions?.permission && Object.entries(viewingPermissions.permission).map(([key, values]) => (
              <div key={key} className="border p-4 rounded-md">
                <h3 className="font-medium mb-2">{key}</h3>
                <div className="flex flex-wrap gap-2">
                  {values.map((value, index) => (
                    <span key={index} className="bg-gray-100 px-2 py-1 rounded text-sm">
                      {value}
                    </span>
                  ))}
                </div>
              </div>
            ))}
            {viewingPermissions && !viewingPermissions.permission && (
              <div className="text-center text-muted-foreground">
                暂无权限设置
              </div>
            )}
          </div>
          <DialogFooter>
            <Button onClick={() => setViewingPermissions(null)}>
              关闭
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </>
  )
}

export default withAuth(GroupsPage)