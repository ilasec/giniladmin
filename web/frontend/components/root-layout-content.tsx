'use client'

import { usePathname } from 'next/navigation'
import { Sidebar } from "@/components/sidebar"
import { TopNav } from "@/components/top-nav"
import { useAuth } from "@/hooks/use-auth"

export function RootLayoutContent({ children }: { children: React.ReactNode }) {
  const pathname = usePathname()
  const { isAuthenticated } = useAuth()
  const isLoginPage = pathname === '/'

  if (isLoginPage || !isAuthenticated) {
    return <>{children}</>
  }

  return (
    <div className="min-h-screen flex">
      <div className="fixed inset-y-0 z-50 w-64">
        <Sidebar />
      </div>
      <div className="flex-1 ml-64 flex flex-col">
        <TopNav />
        <main className="flex-1 p-6 bg-gray-100">
          {children}
        </main>
      </div>
    </div>
  )
}