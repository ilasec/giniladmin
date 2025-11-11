import "./globals.css"
import { Inter } from "next/font/google"
import { Sidebar } from "@/components/sidebar"
import { AuthProvider } from "@/hooks/use-auth"
import type React from "react"

const inter = Inter({ subsets: ["latin"] })

export const metadata = {
  title: "Account Authentication Service",
  description: "Manage applications, users, and permissions",
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <AuthProvider>
          <div className="flex h-screen">
            <Sidebar />
            <main className="flex-1 overflow-y-auto p-8">{children}</main>
          </div>
        </AuthProvider>
      </body>
    </html>
  )
}

