"use client"

import { useAuth } from "@/hooks/use-auth"
import Link from "next/link"
import { usePathname } from "next/navigation"
import { LayoutDashboard, Users, UserCircle, Shield, Key, Lock, FileCheck, PiIcon as Api, LogOut } from "lucide-react"
import { cn } from "@/lib/utils"

const navItems = [
  { name: "仪表盘", href: "/", icon: LayoutDashboard },
  { name: "应用管理", href: "/applications", icon: Shield },
  { name: "用户管理", href: "/users", icon: Users },
  { name: "用户组", href: "/groups", icon: UserCircle },
  // { name: "角色管理", href: "/roles", icon: Key },
  // { name: "权限管理", href: "/permissions", icon: Lock },
  // { name: "认证管理", href: "/authentication", icon: FileCheck },
  // { name: "授权管理", href: "/authorization", icon: Shield },
  // { name: "API管理", href: "/api", icon: Api },
]

export function Sidebar() {
  const { isAuthenticated, logout } = useAuth()
  const pathname = usePathname()

  if (!isAuthenticated) {
    return null
  }

  return (
    <div className="w-64 bg-gradient-to-b from-gray-900 to-gray-800 text-white p-6 min-h-screen flex flex-col">
      <div className="mb-8">
        <h1 className="text-2xl font-bold bg-gradient-to-r from-blue-400 to-purple-500 bg-clip-text text-transparent">
          管理系统
        </h1>
        <p className="text-gray-400 text-sm mt-1">欢迎使用管理后台</p>
      </div>
      
      <nav className="flex-1 space-y-1">
        {navItems.map((item) => {
          const isActive = pathname === item.href
          return (
            <Link
              key={item.name}
              href={item.href}
              className={cn(
                "flex items-center gap-3 rounded-lg px-4 py-3 text-sm font-medium transition-all duration-200",
                "hover:bg-gray-700/50 hover:text-white",
                isActive 
                  ? "bg-blue-500/20 text-blue-400 border-l-4 border-blue-500" 
                  : "text-gray-300"
              )}
            >
              <item.icon className={cn(
                "h-5 w-5 transition-colors",
                isActive ? "text-blue-400" : "text-gray-400"
              )} />
              <span>{item.name}</span>
            </Link>
          )
        })}
      </nav>

      <div className="pt-4 border-t border-gray-700">
        <button
          onClick={logout}
          className="w-full flex items-center gap-3 px-4 py-3 text-sm font-medium text-gray-300 rounded-lg hover:bg-red-500/20 hover:text-red-400 transition-all duration-200"
        >
          <LogOut className="h-5 w-5" />
          <span>退出登录</span>
        </button>
      </div>
    </div>
  )
}

