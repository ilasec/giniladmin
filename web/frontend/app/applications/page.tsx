"use client"

import { useState, useEffect } from "react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogFooter } from "@/components/ui/dialog"
import { Pencil, Trash2, Eye } from "lucide-react"
import { useRouter } from "next/navigation"

import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip"
import { fetchAppList, createApp, deleteApp } from "@/services/api"
import { withAuth } from "@/components/auth-guard"
import type { App } from '@/types/types'
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

function ApplicationsPage() {
  const router = useRouter()
  
  const [apps, setApps] = useState<App[]>([])
  const [total, setTotal] = useState(0)
  const [page, setPage] = useState(1)
  const [pageSize] = useState(10)
  const [loading, setLoading] = useState(false)
  const [isCreating, setIsCreating] = useState(false)
  const [createForm, setCreateForm] = useState({ name: '' })
  const [createError, setCreateError] = useState('')
  const [deletingApp, setDeletingApp] = useState<App | null>(null)
  const [keyword, setKeyword] = useState("")
  const debouncedKeyword = useDebounce(keyword, 500)

  const loadApps = async (searchKeyword?: string) => {
    try {
      setLoading(true)
      const { result } = await fetchAppList({
        page,
        pageSize,
        keyword: searchKeyword || ''
      })
      setApps(result.list)
      setTotal(result.total)
    } catch (err) {
      console.error("加载应用失败:", err)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    loadApps(debouncedKeyword)
  }, [page, debouncedKeyword])

  const handleSearch = () => {
    setPage(1)
    loadApps(keyword)
  }

  const handleCreate = async () => {
    if (!createForm.name) return
    try {
      await createApp({
        name: createForm.name
      })
      setIsCreating(false)
      setCreateForm({ name: '' })
      setCreateError('')
      loadApps()
    } catch (err: any) {
      setCreateError(err.message || '创建应用失败')
    }
  }

  const handleDelete = (app: App) => {
    setDeletingApp(app)
  }

  const confirmDelete = async () => {
    if (!deletingApp) return
    try {
      await deleteApp(deletingApp.id)
      setDeletingApp(null)
      loadApps()
    } catch (err) {
      console.error("Failed to delete app:", err)
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
        <h2 className="text-2xl font-bold">应用管理</h2>
        
        {/* 搜索和操作区域 */}
        <div className="flex flex-col md:flex-row gap-4">
          <Input
            placeholder="搜索应用..."
            value={keyword}
            onChange={(e) => setKeyword(e.target.value)}
            className="max-w-sm"
          />
          <div className="flex-1 flex justify-end">
            <Button onClick={() => setIsCreating(true)}>创建应用</Button>
          </div>
        </div>
    
        {/* 表格卡片 */}
        <Card>
          <CardHeader>
            <CardTitle>应用列表</CardTitle>
          </CardHeader>
          <CardContent>
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>ID</TableHead>
                  <TableHead>名称</TableHead>
                  <TableHead>应用ID</TableHead>
                  <TableHead>应用密钥</TableHead>
                  <TableHead>创建时间</TableHead>
                  <TableHead>操作</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {apps.map((app) => (
                  <TableRow key={app.id}>
                    <TableCell>{app.id}</TableCell>
                    <TableCell>{app.name}</TableCell>
                    <TableCell className="font-mono">{app.appId}</TableCell>
                    <TableCell className="font-mono">{app.secret}</TableCell>
                    <TableCell>{new Date(app.createdAt).toLocaleString()}</TableCell>
                    <TableCell>
                      <div className="flex gap-2">
                        <TooltipProvider>
                          <Tooltip>
                            <TooltipTrigger asChild>
                              <Button 
                                variant="ghost" 
                                size="icon"
                                onClick={() => router.push(`/applications/${app.id}`)}
                              >
                                <Eye className="h-4 w-4" />
                              </Button>
                            </TooltipTrigger>
                            <TooltipContent>
                              <p>查看详情</p>
                            </TooltipContent>
                          </Tooltip>
                        </TooltipProvider>

                        <TooltipProvider>
                          <Tooltip>
                            <TooltipTrigger asChild>
                              <Button 
                                variant="ghost" 
                                size="icon"
                                onClick={() => handleDelete(app)}
                                className="text-red-500 hover:text-red-600"
                              >
                                <Trash2 className="h-4 w-4" />
                              </Button>
                            </TooltipTrigger>
                            <TooltipContent>
                              <p>删除应用</p>
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
      </div> {/* 添加缺少的闭合标签 */}

      {/* 创建应用对话框 */}
      <Dialog open={isCreating} onOpenChange={(open) => {
        if (!open) {
          setIsCreating(false)
          setCreateForm({ name: '' })
          setCreateError('')
        }
      }}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>创建应用</DialogTitle>
          </DialogHeader>
          <div className="space-y-4 py-4">
            <div className="flex flex-col gap-4">
              <div className="flex flex-col gap-2">
                <label>应用名称</label>
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
              {createError && (
                <p className="text-sm text-red-500 bg-red-50 p-2 rounded">{createError}</p>
              )}
            </div>
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => {
              setIsCreating(false)
              setCreateForm({ name: '' })
              setCreateError('')
            }}>
              取消
            </Button>
            <Button 
              onClick={handleCreate}
              disabled={!createForm.name}
            >
              创建
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    
      {/* 确认删除对话框 */}
      <Dialog open={!!deletingApp} onOpenChange={() => setDeletingApp(null)}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>确认删除</DialogTitle>
          </DialogHeader>
          <div className="py-4">
            <p>确定要删除应用 "{deletingApp?.name}" 吗？此操作不可恢复。</p>
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => setDeletingApp(null)}>
              取消
            </Button>
            <Button variant="destructive" onClick={confirmDelete}>
              删除
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </>
  )
}

export default withAuth(ApplicationsPage)