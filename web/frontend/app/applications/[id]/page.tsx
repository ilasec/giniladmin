"use client"

import { useState, useEffect, useCallback } from "react"
import { useParams } from "next/navigation"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogFooter } from "@/components/ui/dialog"
import { fetchAppDetail, fetchAppUsers, addAppUser, removeAppUser, updateAppUserGroup } from "@/services/api"
import { withAuth } from "@/components/auth-guard"
import type { App, User, Group, AppUser } from '@/types/types'

function ApplicationDetailPage() {
  const params = useParams()
  const appId = params.id as string
  
  const [app, setApp] = useState<App | null>(null)
  const [users, setUsers] = useState<AppUser[]>([])
  const [groups, setGroups] = useState<Group[]>([])
  const [isAddingUser, setIsAddingUser] = useState(false)
  const [selectedUser, setSelectedUser] = useState("")
  const [selectedGroup, setSelectedGroup] = useState("")
  const [availableUsers, setAvailableUsers] = useState<User[]>([])

  // 加载应用详情
  const loadAppDetail = useCallback(async () => {
    try {
      const { result } = await fetchAppDetail(appId)
      setApp(result.app)
      setGroups(result.groups)
    } catch (error) {
      console.error("加载应用详情失败:", error)
    }
  }, [appId])

  // 加载应用用户
  const loadAppUsers = useCallback(async () => {
    try {
      const { result } = await fetchAppUsers(appId)
      setUsers(result.list)
    } catch (error) {
      console.error("加载用户列表失败:", error)
    }
  }, [appId])

  useEffect(() => {
    loadAppDetail()
    loadAppUsers()
  }, [loadAppDetail, loadAppUsers])

  // 处理添加用户
  const handleAddUser = async () => {
    if (!selectedUser || !selectedGroup) return
    try {
      await addAppUser(appId, {
        userId: Number(selectedUser),
        groupId: Number(selectedGroup)
      })
      setIsAddingUser(false)
      setSelectedUser("")
      setSelectedGroup("")
      loadAppUsers()
    } catch (error) {
      console.error("添加用户失败:", error)
    }
  }

  // 处理移除用户
  const handleRemoveUser = async (userId: number) => {
    try {
      await removeAppUser(appId, userId)
      loadAppUsers()
    } catch (error) {
      console.error("移除用户失败:", error)
    }
  }

  // 处理更新用户组
  const handleUpdateGroup = async (userId: number, groupId: number) => {
    try {
      await updateAppUserGroup(appId, userId, groupId)
      loadAppUsers()
    } catch (error) {
      console.error("更新用户组失败:", error)
    }
  }

  return (
    <div className="p-6 space-y-6">
      {/* 应用信息卡片 */}
      <Card>
        <CardHeader>
          <CardTitle>应用信息</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="grid grid-cols-2 gap-4">
            <div>
              <label className="font-medium">应用名称</label>
              <p>{app?.name}</p>
            </div>
            <div>
              <label className="font-medium">应用ID</label>
              <p className="font-mono">{app?.appId}</p>
            </div>
            <div>
              <label className="font-medium">应用密钥</label>
              <p className="font-mono">{app?.secret}</p>
            </div>
            <div>
              <label className="font-medium">创建时间</label>
              <p>{app?.createdAt && new Date(app.createdAt).toLocaleString()}</p>
            </div>
          </div>
        </CardContent>
      </Card>

      {/* 用户管理卡片 */}
      <Card>
        <CardHeader>
          <div className="flex justify-between items-center">
            <CardTitle>用户管理</CardTitle>
            <Button onClick={() => setIsAddingUser(true)}>添加用户</Button>
          </div>
        </CardHeader>
        <CardContent>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>用户名</TableHead>
                <TableHead>昵称</TableHead>
                <TableHead>用户组</TableHead>
                <TableHead>操作</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {users.map((user) => (
                <TableRow key={user.id}>
                  <TableCell>{user.userName}</TableCell>
                  <TableCell>{user.nickName}</TableCell>
                  <TableCell>
                    <select
                      value={user.groupId}
                      onChange={(e) => handleUpdateGroup(user.id, Number(e.target.value))}
                      className="p-2 border rounded"
                    >
                      {groups.map(group => (
                        <option key={group.id} value={group.id}>
                          {group.name}
                        </option>
                      ))}
                    </select>
                  </TableCell>
                  <TableCell>
                    <Button
                      variant="ghost"
                      size="sm"
                      onClick={() => handleRemoveUser(user.id)}
                      className="text-red-500 hover:text-red-600"
                    >
                      移除
                    </Button>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </CardContent>
      </Card>

      {/* 添加用户对话框 */}
      <Dialog open={isAddingUser} onOpenChange={setIsAddingUser}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>添加用户</DialogTitle>
          </DialogHeader>
          <div className="space-y-4 py-4">
            <div className="flex flex-col gap-4">
              <div>
                <label className="block mb-2">选择用户</label>
                <select
                  value={selectedUser}
                  onChange={(e) => setSelectedUser(e.target.value)}
                  className="w-full p-2 border rounded"
                >
                  <option value="">请选择用户</option>
                  {availableUsers.map(user => (
                    <option key={user.id} value={user.id}>
                      {user.userName}
                    </option>
                  ))}
                </select>
              </div>
              <div>
                <label className="block mb-2">选择用户组</label>
                <select
                  value={selectedGroup}
                  onChange={(e) => setSelectedGroup(e.target.value)}
                  className="w-full p-2 border rounded"
                >
                  <option value="">请选择用户组</option>
                  {groups.map(group => (
                    <option key={group.id} value={group.id}>
                      {group.name}
                    </option>
                  ))}
                </select>
              </div>
            </div>
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => setIsAddingUser(false)}>
              取消
            </Button>
            <Button
              onClick={handleAddUser}
              disabled={!selectedUser || !selectedGroup}
            >
              添加
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  )
}

export default withAuth(ApplicationDetailPage)